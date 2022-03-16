package main

import (
	"fmt"
	"strings"
)

func extractRows(rows [][]string) ([]product, error) {
	startsAt := -1
	for i, row := range rows {
		// skip the first couple until we get to what we want
		if len(row) == 0 {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(row[0]), "item name") {
			startsAt = i + 1
		}
	}
	if startsAt == -1 {
		return nil, fmt.Errorf(`we looked for the header row starting with "Item Name" but never found it`)
	}

	var products []product
	for i := startsAt; i < len(rows); i++ {
		row := rows[i]

		if len(row) < 1 {
			return nil, fmt.Errorf("invalid number of columns in row %d", i+1)
		}
		if strings.EqualFold(strings.TrimSpace(row[0]), "total") {
			break
		}

		product, err := extractRow(row)
		if err != nil {
			return nil, fmt.Errorf("failed to extract row %d: %w", i+1, err)
		}
		products = append(products, product)

		// break // FIXME: (JMT) testing
	}

	return products, nil
}

func extractRow(row []string) (product, error) {
	if len(row) < 20 {
		return product{}, fmt.Errorf("invalid number of columns in row")
	}

	p := product{
		name:            row[0],
		sku:             row[2],
		price:           row[14],
		cost:            row[15],
		category:        row[3],
		taxName:         row[0],
		vendorName:      row[6],
		quantityOnHand:  row[7],
		minReorderPoint: row[8],
		maxReorderPoint: row[8],
	}

	if p.minReorderPoint == "" {
		p.minReorderPoint = "0"
	}
	if p.maxReorderPoint == "" {
		p.maxReorderPoint = "0"
	}

	return p, nil
}
