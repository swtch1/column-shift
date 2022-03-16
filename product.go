package main

// product is the desired format we must extract from our input data.
type product struct {
	name              string // required
	sku               string // required
	productCode       string
	price             string // required
	cost              string
	category          string // required
	taxName           string // required
	ebt               bool
	vendorName        string
	partialQuantity   bool
	autoItemDiscounts string
	quantityOnHand    string
	minReorderPoint   string
	maxReorderPoint   string
	otherItemFeatures string
}
