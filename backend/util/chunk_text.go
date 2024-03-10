package util

import "fmt"

// ChunkText divides text into chunks with specified size and overlap
func ChunkText(text string, chunkSize int, overlapPct float64) ([]string, error) {
	if chunkSize <= 0 || overlapPct < 0 || overlapPct >= 1 {
		return nil, fmt.Errorf("invalid arguments: chunkSize must be positive, overlap must be between 0 and 99")
	}

	overlap := int(float64(chunkSize) * overlapPct)
	chunks := []string{}
	start := 0

	for len(text[start:]) > chunkSize {
		end := start + chunkSize
		if start > 0 {
			start -= overlap // Apply overlap for subsequent chunks
		}
		chunks = append(chunks, text[start:end])
		start = end
	}
	chunks = append(chunks, text[start:]) // Add the last chunk

	return chunks, nil
}
