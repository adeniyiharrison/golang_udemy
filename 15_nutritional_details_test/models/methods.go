package models

import "strings"

// UpdateProductName used to clean up product names
func (up *UniversalProducts) UpdateProductName() {
	for i, v := range up.Results {
		up.Results[i].ProductName = strings.ToLower(v.ProductName)
		up.Results[i].DisplayName = strings.ToLower(v.DisplayName)
	}
}
