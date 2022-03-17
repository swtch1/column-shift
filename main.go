package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

var (
	//go:embed template/template.xlsx
	template     []byte
	saveOutputTo = "./inventory.xlsx"
)

func main() {
	// open the template
	outFile, err := excelize.OpenReader(bytes.NewReader(template))
	if err != nil {
		fatal(err, "failed to open template file: %v", err)
	}
	defer outFile.Close()

	// process all input files into template
	for _, fileName := range currentDirExcelFiles() {
		err := processFile(fileName, outFile)
		if err != nil {
			fatal(err, "cannot process %s", fileName)
		}
		info("processed %s", fileName)
	}

	// write our work to a separate location
	err = outFile.SaveAs(saveOutputTo)
	if err != nil {
		fatal(err, "failed to save out file %s: %v", saveOutputTo, err)
	}

	enterToContinue()
	os.Exit(0)
}

func enterToContinue() {
	fmt.Println()
	fmt.Println("Press enter to continue...")
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
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

func currentDirExcelFiles() []string {
	var xlsxFiles []string
	for _, f := range filesInCurrentDir() {
		if !strings.HasSuffix(f, ".xlsx") {
			continue
		}
		if f == saveOutputTo {
			continue
		}
		xlsxFiles = append(xlsxFiles, f)
	}
	return xlsxFiles
}

func filesInCurrentDir() []string {
	entries, err := os.ReadDir(".")
	if err != nil {
		fatal(err, "failed to read current directory")
	}

	var files []string
	for _, e := range entries {
		if e.Type().IsRegular() {
			files = append(files, e.Name())
		}
	}

	return files
}
