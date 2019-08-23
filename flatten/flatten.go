package flatten

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Flatten flattens an excel file using the specified delimiter and returns the flat file
func Flatten(filePath string, delimiter string) []excelize.File {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return make([]excelize.File, 0)
	}

	for _, name := range file.GetSheetMap() {
		fmt.Printf("Sheet name: %s\n", name)

		rows, err := file.GetRows(name)
		if err != nil {
			fmt.Println(err)
			return make([]excelize.File, 0)
		}

		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, delimiter)
			}
			fmt.Println()
		}
	}

	files := make([]excelize.File, 0)
	files = append(files, *file)
	return files
}
