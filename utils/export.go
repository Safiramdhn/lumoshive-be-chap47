package utils

import (
	"fmt"
	"golang-chap47/models"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

func ExportOrdersToExcel(orders []models.Order, filePath string) error {
	currentTime := time.Now()
	log.Printf("Running export orders to %s : %v", filePath, currentTime)
	f := excelize.NewFile()
	sheet := "Orders"

	// Set up main sheet headers
	f.SetSheetName("Sheet1", sheet)
	f.SetCellValue(sheet, "A1", "Order ID")
	f.SetCellValue(sheet, "B1", "Order Date")
	f.SetCellValue(sheet, "C1", "Total Quantity")
	f.SetCellValue(sheet, "D1", "Total Amount")
	f.SetCellValue(sheet, "E1", "Order Status")
	f.SetCellValue(sheet, "F1", "Product Name")
	f.SetCellValue(sheet, "G1", "Quantity")
	f.SetCellValue(sheet, "H1", "Price")

	row := 2 // Start from row 2 for data
	for _, order := range orders {
		startRow := row // Track starting row for merging
		for _, item := range order.OrderItems {
			// Write order item details
			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), item.ProductName) // Product Name
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), item.Quantity)    // Quantity
			f.SetCellValue(sheet, fmt.Sprintf("H%d", row), item.Price)       // Price
			row++
		}

		// Write order-level data and merge cells for order columns
		f.SetCellValue(sheet, fmt.Sprintf("A%d", startRow), order.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", startRow), order.CreatedAt.Format("2006-01-02"))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", startRow), order.TotalQuantity)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", startRow), order.Total)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", startRow), order.Status)
		f.MergeCell(sheet, fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", row-1))
		f.MergeCell(sheet, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", row-1))
		f.MergeCell(sheet, fmt.Sprintf("C%d", startRow), fmt.Sprintf("C%d", row-1))
		f.MergeCell(sheet, fmt.Sprintf("D%d", startRow), fmt.Sprintf("D%d", row-1))
		f.MergeCell(sheet, fmt.Sprintf("E%d", startRow), fmt.Sprintf("E%d", row-1))
	}

	// Save the Excel file
	return f.SaveAs(filePath)
}
