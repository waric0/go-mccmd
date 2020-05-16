package output

import uuid "github.com/google/uuid"

// Header : Structure definition for "header" in "manifest.json"
type Header struct {
	Description      string    `json:"description"`
	Name             string    `json:"name"`
	UUID             uuid.UUID `json:"uuid"`
	Version          [3]int    `json:"version"`
	MinEngineVersion [3]int    `json:"min_engine_version"`
}

// Modules : Structure definition for "modules" in "manifest.json"
type Modules struct {
	Description string    `json:"description"`
	Type        string    `json:"data"`
	UUID        uuid.UUID `json:"uuid"`
	Version     [3]int    `json:"version"`
}

// Dependencies : Structure definition for "dependencies" in "manifest.json"
type Dependencies struct {
}

// Metadata : Structure definition for "metadata" in "manifest.json"
type Metadata struct {
	Authors []string `json:"authors"`
	URL     string   `json:"url"`
	License string   `json:"license"`
}

// Manifest : "manifest.json" Overall Structure Definition
////// Memo : Minimum Configuration + Metadata //////
type Manifest struct {
	FormatVersion int          `json:"format_version"`
	Header        Header       `json:"header"`
	Modules       Modules      `json:"modules"`
	Dependencies  Dependencies `json:"dependencies"`
	Metadata      Metadata     `json:"metadata"`
}
