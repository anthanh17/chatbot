package api

import (
	"chatbot-ai/embeddings"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pgvector/pgvector-go"
)

type promptRequest struct {
	Question string `json:"question" binding:"required,min=1"`
}

type promptResponse struct {
	Answer string `json:"answer"`
}

/*
* The server receives the query
* B1: convert question query to vector embeddings
* B2: search vector above in database, calculate vector via cosin, euclid find similar values
* B3: returns the most relevant results
 */
func (server *Server) prompt(ctx *gin.Context) {
	var req promptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// B1: convert question query to vector embeddings
	apiKey := server.config.OpenaiApiKey
	if apiKey == "" {
		log.Fatal("ApiKey empty!")
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("ApiKey empty!")))
		return
	}

	vectorEmbeddings, err := embeddings.FetchEmbeddings([]string{req.Question}, apiKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// B2: search vector above in database
	query := `
		SELECT content, (embedding <=> $1) as cosine_distance
		FROM documents
		ORDER BY cosine_distance
		LIMIT 1;
	`

	// Execute the query
	rows, err := server.store.Query(context.Background(), query, pgvector.NewVector(vectorEmbeddings[0]))
	if err != nil {
		fmt.Println("Query error:", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer rows.Close()

	var content string
	var cosineDistance float64
	for rows.Next() {
		err = rows.Scan(&content, &cosineDistance)
		if err != nil {
			fmt.Println("Error retrieving data:", err)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if rows.Err() != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// B3: returns the most relevant results
	ctx.JSON(http.StatusOK, promptResponse{
		Answer: content,
	})
}
