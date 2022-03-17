package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const headerLen = 4

var row = headerLen

// writeTemplate writes products to the template file.  We know exactly what
// this looks like so we can make a lot of assumptions.
func writeTemplate(file *excelize.File, products []product) error {
	for _, p := range products {
		err := setRowValues(file, row, p)
		if err != nil {
			return fmt.Errorf("failed to set values on row %d: %w", row, err)
		}
		row += 1
	}

	return nil
}

func setRowValues(f *excelize.File, row int, p product) error {
	et := &errTracker{}

	et.capture(f.SetCellValue("Products", fmt.Sprintf("A%d", row), p.name))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("B%d", row), p.sku))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("C%d", row), p.productCode))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("D%d", row), p.price))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("E%d", row), p.cost))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("F%d", row), p.category))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("G%d", row), p.taxName))
	et.capture(f.SetCellBool("Products", fmt.Sprintf("H%d", row), p.ebt))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("I%d", row), p.vendorName))
	et.capture(f.SetCellBool("Products", fmt.Sprintf("J%d", row), p.partialQuantity))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("K%d", row), p.autoItemDiscounts))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("L%d", row), p.autoDiscountExprDate))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("M%d", row), p.quantityOnHand))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("N%d", row), p.minReorderPoint))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("O%d", row), p.maxReorderPoint))
	et.capture(f.SetCellValue("Products", fmt.Sprintf("P%d", row), p.otherItemFeatures))

	return et.Error()
}

type errTracker struct {
	err error
}

func (t *errTracker) capture(err error) {
	if t.err != nil {
		return
	}
	if err != nil {
		t.err = err
	}
}
func (t *errTracker) Error() error {
	return t.err
}
