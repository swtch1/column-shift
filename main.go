package main

import (
	"fmt"

	"github.com/homedepot/flop"
	"github.com/xuri/excelize/v2"
)

func main() {
	toProcess := []string{
		"./sample/InventorySnapshot.xlsx",
	}
	template := "./template/template.xlsx"
	out := "./inventory.xlsx"

	err := flop.Copy(template, out, flop.Options{MkdirAll: true})

	outFile, err := excelize.OpenFile(out)
	if err != nil {
		fatal(err, "failed to open template file: %v", err)
	}
	defer outFile.Close()

	for _, fileName := range toProcess {
		err := processFile(fileName, outFile)
		if err != nil {
			fatal(err, "cannot process %s", fileName)
		}
		info("processed %s", fileName)
	}

	err = outFile.Save()
	if err != nil {
		fatal(err, "failed to save out file %s: %v", out, err)
	}
}

func processFile(inFile string, out *excelize.File) error {
	f, err := excelize.OpenFile(inFile, excelize.Options{})
	if err != nil {
		fatal(err, "failed to open file %s", inFile)
	}
	defer f.Close()

	sheet := "InventorySnapshot"

	err = validateSheet(f, sheet)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	products, err := extractRows(rows)
	if err != nil {
		return err
	}

	err = writeTemplate(out, products)
	if err != nil {
		return fmt.Errorf("failed to write to template: %w", err)
	}

	return nil
}
