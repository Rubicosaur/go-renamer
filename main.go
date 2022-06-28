package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// provide path like C:\ ...
	mainPath := `.\`
	renameMultipleFiles(mainPath)

}

// Function to loop through nested folders recursively and rename specific file

func renameMultipleFiles(mainFolderPath string) {

	x, err := os.Open(mainFolderPath)

	if err != nil {
		log.Fatal(err)
	}

	files, err := x.ReadDir(-1)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range files {

		var currName = item.Name()
		var currentPath = fmt.Sprintf(`%s\%s`, mainFolderPath, currName)

		if item.Type().IsDir() {

			renameMultipleFiles(currentPath)

		} else if strings.Contains(currName, "arkusz") && strings.Contains(currName, ".xlsx") {

			// pattern below in variable newName

			var newName = fmt.Sprintf(`10 %s`, currName)
			var newPath = fmt.Sprintf(`%s\%s`, mainFolderPath, newName)

			err := os.Rename(currentPath, newPath)

			if err != nil {
				log.Fatal(err)
			}

		}

	}

}
