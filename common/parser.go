package common

import (
	"encoding/json"
	"fmt"
)

// GetSnippetResponse is a snippet unmarshaller.
func GetSnippetResponse(data []byte) (*SnippetsResponse, error) {

	var snippets []Snippet
	err := json.Unmarshal(data, &snippets)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling snippets: %w", err)
	}

	return &SnippetsResponse{Snippets: snippets}, nil
}

// ConvertToMap converts a SnippetsResponse to a map.
func ConvertToMap(response *SnippetsResponse) map[string]string {
	snippetMap := make(map[string]string)
	for _, snippet := range response.Snippets {
		snippetMap[snippet.Name] = snippet.Command
	}

	return snippetMap
}
