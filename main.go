package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	toProcess := []string{
		"./sample/InventorySnapshot.xlsx",
	}

	for _, fileName := range toProcess {
		err := processFile(fileName)
		if err != nil {
			fatal(err, "cannot process %s", fileName)
		}
		info("processed %s", fileName)
	}

}

func processFile(fileName string) error {
	f, err := excelize.OpenFile(fileName, excelize.Options{})
	if err != nil {
		fatal(err, "failed to open file %s", fileName)
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

	err = writeTemplate(products)
	if err != nil {
		return fmt.Errorf("failed to write to template: %w", err)
	}

	return nil
}

func writeTemplate(products []product) error {
	for _, p := range products {
		fmt.Printf(" -> JMTDEBUG: %s: %+v\n", "product", p) // FIXME: (JMT) testing
	}
	return nil
}
