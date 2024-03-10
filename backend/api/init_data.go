package api

import (
	"chatbot-ai/embeddings"
	"chatbot-ai/util"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pgvector/pgvector-go"
)

type dataResponse struct {
	Message string `json:"message"`
}

/*
* B1: Get all text from file
* B2: Split text into chunks
* B3: Convert each chunks to vector embedding
* B4: Store vector embeddings to pgvector database
 */
func (server *Server) initializeData(ctx *gin.Context) {
	// B1: Get all text from file
	textData, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading PDF:", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Config split
	chunkSize := server.config.ChunkSize
	overlapPct := float64(server.config.OverlapPct) / float64(100)

	// B2: Split text into chunks
	chunks, err := util.ChunkText(string(textData), chunkSize, overlapPct)
	if err != nil {
		fmt.Println("Error splitting text:", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// B3: Convert each chunks to vector embedding
	apiKey := server.config.OpenaiApiKey
	if apiKey == "" {
		log.Fatal("ApiKey empty!")
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("ApiKey empty!")))
		return
	}

	// Call api open ai generate embedding, vectorEmbeddings
	vectorEmbeddings, err := embeddings.FetchEmbeddings(chunks, apiKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// B4: Store vector embeddings to pgvector database
	// 4.1: Drop database if exists
	_, err = server.store.Exec(ctx, "DROP TABLE IF EXISTS documents")
	if err != nil {
		panic(err)
	}

	// 4.2: Create table
	_, err = server.store.Exec(ctx, "CREATE TABLE documents (id bigserial PRIMARY KEY, content text, embedding vector(1536))")
	if err != nil {
		panic(err)
	}

	// 4.3: Insert to database
	for i, content := range chunks {
		_, err := server.store.Exec(ctx, "INSERT INTO documents (content, embedding) VALUES ($1, $2)", content, pgvector.NewVector(vectorEmbeddings[i]))
		if err != nil {
			panic(err)
		}
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Message: "initialize data from file success!",
	})

}
