package flatten

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Flatten flattens an excel file using the specified delimiter and returns the flat file
func Flatten(filePath string, delimiter rune) ([]os.File, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return make([]os.File, 0), err
	}

	files := make([]os.File, 0)
	for _, name := range file.GetSheetMap() {
		rows, err := file.GetRows(name)
		if err != nil {
			return make([]os.File, 0), err
		}

		directory := filepath.Dir(filePath)
		os.Chdir(directory)
		newFile, err := os.Create(fmt.Sprintf("%s.csv", name)) //TODO: doesn't seem to work
		if err != nil {
			return make([]os.File, 0), err
		}

		writer := csv.NewWriter(newFile)
		writer.Comma = delimiter

		for _, row := range rows {
			if err := writer.Write(row); err != nil {
				return make([]os.File, 0), err
			}
		}

		writer.Flush()
		newFile.Close()

		files = append(files, *newFile)
	}

	return files, nil
}
