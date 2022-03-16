package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func validateSheet(f *excelize.File, sheets ...string) error {
	for _, want := range sheets {
		var match bool
		for _, have := range f.GetSheetList() {
			if have == want {
				match = true
				break
			}
		}

		if !match {
			return fmt.Errorf("could not find sheet %s in spreadsheet", want)
		}
	}

	return nil
}
