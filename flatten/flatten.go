package flatten

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Flatten flattens an excel file using the specified delimiter and returns the flat file
func Flatten(filePath string, delimiter string) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, name := range file.GetSheetMap() {
		fmt.Printf("Sheet name: %s\n", name)

		rows, err := file.GetRows(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		newFile, err := os.Create(fmt.Sprintf("C:\\Users\\John\\Desktop\\%s.csv", name)) //TODO: doesn't seem to work
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		w := csv.NewWriter(newFile)
		w.Comma = []rune(delimiter)[0]

		for _, row := range rows {
			if err := w.Write(row); err != nil {
				log.Fatalln("Error writing record to file: ", err)
			}
		}

		w.Flush()
	}

	return
}
