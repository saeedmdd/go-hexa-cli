package modifier

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/saeedmdd/go-hexa-cli/utils"
)

func findAndReplace(filePath, oldText, newText string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	modifiedContent := strings.ReplaceAll(utils.B2S(content), oldText, newText)
	os.WriteFile(filePath, utils.S2B(modifiedContent), 0644)
	return nil
}
func ModifyFiles(rootDir, oldText, newText string) error {
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		err = findAndReplace(path, oldText, newText)
		if err != nil {
			fmt.Printf("error while modifying: %v", err)
			return err
		}
		return nil
	})
	return err
}
