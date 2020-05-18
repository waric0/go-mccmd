package output

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/google/uuid"
)

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

func makeDir(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		err = os.Mkdir(name, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func makeManifest(man *Manifest, path string) error {
	b, err := json.Marshal(*man)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "	")
	if err != nil {
		return err
	}
	b = buf.Bytes()
	name := filepath.Join(path, "manifest.json")
	err = ioutil.WriteFile(name, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetMCPackName : Check and get mcpack directory name
func GetMCPackName(man *Manifest) (string, error) {
	name := man.Header.Name
	symbols := []string{" ", "."}
	if name == "" {
		return "", errors.New("No mcpack directory name specified")
	} else if strings.Contains(name, "/") || strings.Contains(name, "\\") {
		return "", errors.New("Contains symbols that cannot be specified as mcpack directory names")
	} else {
		for _, v := range symbols {
			if name[1:] == v || name[:1] == v {
				return "", errors.New("A symbol is in a position that cannot be specified as an mcpack directory name")
			}
		}
	}
	return name, nil
}

// MakeMCPack : Function to write a mcpack directory
func MakeMCPack(man *Manifest, path string) error {
	name, err := GetMCPackName(man)
	if err != nil {
		return err
	}
	packDir := filepath.Join(path, name)
	err = makeDir(packDir)
	if err != nil {
		return err
	}
	funcDir := filepath.Join(packDir, "functions")
	err = makeDir(funcDir)
	if err != nil {
		return err
	}
	err = makeManifest(man, packDir)
	if err != nil {
		return err
	}
	return nil
}
