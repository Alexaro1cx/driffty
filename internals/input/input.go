package input

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func RawRead(path string) map[string]interface{} {
	buff, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		panic(err)
	}

	var out map[string]interface{}

	if err := json.Unmarshal(buff, &out); err != nil {
		panic(err)
	}

	return out
}
