package course

import (
	"fmt"
	"github.com/xuri/excelize"
	"strings"
)

type TermInfo struct {
	teachers []string
	labs     []string
	sClass   []string
}

func (self *TermInfo)GetTeacher() []string {
	return self.teachers
}
func (self *TermInfo) GetLab()[]string  {
	return self.labs
}
func (self *TermInfo) GetSclass() []string {
	return self.sClass
}

func ReadCourse(filePath string) *TermInfo {
	fmt.Println("开始读取excel内容")

	//xlsx, err := excelize.OpenFile("./Workbook.xlsx")
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Get value from cell by given worksheet name and axis.
	//cell := xlsx.GetCellValue("Sheet1", "B2")
	//fmt.Println(cell)
	// Get all the rows in the Sheet1.
	//rows := xlsx.GetRows("Sheet1")
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	//老师列表
	var tearcherArray = make([]string, 12)
	//实验课列表
	var labArray = make([]string, 12)
	//班级列表
	var sClassArray = make([]string, 20)

	// Get all the rows in the Sheet1.
	var sheetName = xlsx.GetSheetName(1)
	fmt.Println(sheetName, "-excel的Sheet名称")
	rows := xlsx.GetRows(sheetName)
	fmt.Println(len(rows), "行数")

	for i, row := range rows {
		if i == 0 {
			//第一行是标题，忽略
			continue
		}
		for j, colCell := range row {
			if j == 0 && strings.Count(strings.TrimSpace(colCell), "") > 1 {
				if i >= len(tearcherArray)+1 {
					tearcherArray = append(tearcherArray, colCell)
				} else {
					tearcherArray[i-1] = colCell
				}
			}
			if j == 2 && strings.Count(strings.TrimSpace(colCell), "") > 1 {
				if i >= len(labArray)+1 {
					labArray = append(labArray, colCell)
				} else {
					labArray[i-1] = colCell
				}
			}
			if j == 4 && strings.Count(strings.TrimSpace(colCell), "") > 1 {
				if i >= len(sClassArray)+1 {
					sClassArray = append(sClassArray, colCell)
				} else {
					sClassArray[i-1] = colCell
				}
			}
		}
	}

	termInfo := &TermInfo{teachers: tearcherArray, sClass: sClassArray, labs: labArray}

	//test
	//for _, v := range tearcherArray {
	//	fmt.Println("老师：" + v)
	//}
	//for _, v := range labArray {
	//	fmt.Println("实验课：" + v)
	//}
	//for _, v := range sClassArray {
	//	fmt.Println("班级：" + v)
	//}

	return termInfo

}
