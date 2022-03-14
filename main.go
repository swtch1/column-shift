package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	fileName := "./sample/InventorySnapshot.xls"

	f, err := excelize.OpenFile(fileName, excelize.Options{})
	if err != nil {
		fatal(err, "failed to open file %s", fileName)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		fmt.Println("sheet:", sheet)
	}
}
