package course

import (
	"github.com/xuri/excelize"
	"fmt"
)

func WriteExcel(filePath string) {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	//err := xlsx.SaveAs("./Workbook.xlsx")
	err := xlsx.SaveAs(filePath)
	if err != nil {
		fmt.Println(err)
	}
}