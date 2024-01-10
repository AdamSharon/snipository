package common

type SnippetsResponse struct {
	// using a response struct to allow for future additions - like saving the data online,
	// Gist, server, etc.
	Snippets []Snippet
}

type Snippet struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type SnippetMap struct {
	Snippets map[string]string
}
