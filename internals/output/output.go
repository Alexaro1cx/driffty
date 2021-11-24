package output

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func Write(results map[string]interface{}, outputPath string) {
	name := outputPath + "/" + fmt.Sprint(time.Now().Unix()) + ".json"

	err := os.MkdirAll(outputPath, os.ModePerm)

	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(results, "", "  ")

	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(name, bytes, 0600); err != nil {
		panic(err)
	}
}
