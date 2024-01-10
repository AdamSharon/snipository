package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"snipository/common"
)

func PushSnippetToFile(snippetToAdd common.Snippet, pathToDataFile string) error {
	fmt.Printf("Pushing to local file: %s -- %s\n", snippetToAdd.Name, snippetToAdd.Command)
	data, err := GetDataFromFile(pathToDataFile)
	if len(data) == 0 { // If the file is empty, set it to an empty array
		data = []byte("[]")
	}

	var snippets []common.Snippet
	err = json.Unmarshal(data, &snippets)
	if err != nil {
		return fmt.Errorf("error unmarshalling snippets: %w", err)
	}

	snippets = append(snippets, snippetToAdd)

	snippetBytes, err := json.MarshalIndent(snippets, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling snippet file: %w", err)
	}

	err = os.WriteFile(pathToDataFile, snippetBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Successfully pushed to file")

	return nil
}

func GetDataFromFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist. please make to create a store file at '%s'", path)
		}
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return data, nil
}
