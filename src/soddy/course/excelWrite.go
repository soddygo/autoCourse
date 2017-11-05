package course

import (
	"github.com/xuri/excelize"
	"fmt"
	"strconv"
)

func WriteExcel(filePath string, termLabs [][]*Lab, labs []string) {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	//index := xlsx.NewSheet("Sheet2")
	//// Set value of a cell.
	//xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	//xlsx.SetCellValue("Sheet1", "B2", 100)

	//保存到excel里
	//xlsx := excelize.NewFile()
	var sheetName = "Sheet1"
	// Create a new sheet.
	//index :=xlsx.NewSheet(sheetName)

	xlsx.SetCellValue(sheetName, "A1", "天")

	//实验课标题
	var letter rune = 'B'
	var letterInt = int(letter)
	for i, labStr := range labs {
		xlsx.SetCellValue(sheetName, string(rune(letterInt+i))+strconv.Itoa(1), labStr)
	}
	//排课内容
	var letter2 rune = 'A'
	var letterInt2 = int(letter2)
	for i, labAry := range termLabs {
		//fmt.Println("show:",string(rune(letterInt2))+strconv.Itoa(i+2))
		xlsx.SetCellValue(sheetName, string(rune(letterInt2))+strconv.Itoa(i+2), i+1)
		for _, lab := range labAry {
			var index = getIndex(labs, lab.GetLab())

			//fmt.Println("column:", string(rune(letterInt2+index+1))+strconv.Itoa(i+2))

			xlsx.SetCellValue(sheetName, string(rune(letterInt2+index+1))+strconv.Itoa(i+2), lab.sClass)
		}
	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(0)
	// Save xlsx file by the given path.
	//err := xlsx.SaveAs("./Workbook.xlsx")
	fmt.Println("filePath：",filePath)
	err := xlsx.SaveAs(filePath)
	if err != nil {
		fmt.Println(err)
	}

}

func getIndex(labs []string, labStr string) int {
	for index, lab := range labs {
		if lab == labStr {
			return index
		}
	}
	panic("匹配不到此实验课")
	return 0
}
