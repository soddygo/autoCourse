package main

import (
	"soddy/course"
	"fmt"
	"os"
	"path/filepath"
	"log"
	"strings"
)

func main() {

	cmd := course.PaseCmd()
	if cmd.GetVersion() {
		fmt.Println("version 0.0.1")
	} else if cmd.GetHelp() {
		fmt.Println("必须输入模板excel的绝对路径， -in /xxx/xx/输入模板.xlsx")
	} else {
		course.PrintUsage()
	}

	var inPath = cmd.GetIn()
	var outPath = cmd.GetOut()
	fmt.Println("cmd.GetIn():", inPath)
	fmt.Println("cmd.GetOut():", outPath)

	existPath := getTruePath(inPath, false)
	outPathTrue := getTruePath(outPath, true)

	if strings.Count(existPath, "") > 1 {
		//filePath := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo.xlsx"
		//filePath2 := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo2.xlsx"
		fmt.Println("existPath:", existPath)
		termInfo := course.ReadCourse(existPath)

		//course
		courseObj := &course.Course{}
		courseObj.Init(termInfo.GetTeacher(), termInfo.GetSclass(), termInfo.GetLab(), nil, nil)
		courseObj.AllCourseAndClass()
		var schoolTerm = courseObj.RandomCourse()
		var labArray = courseObj.SortClass(schoolTerm)

		for i, labInfos := range labArray {
			fmt.Println("index:", i, "\t")
			for _, labInfo := range labInfos {
				fmt.Print(labInfo.GetSclass(), "\t")
				fmt.Print(labInfo.GetLab(), "\t")
				//fmt.Print(labInfo.GetFlag(),"\t")
				fmt.Println()
			}

		}

		fmt.Println("outPathTrue:", outPathTrue)
		course.WriteExcel(outPathTrue, labArray, termInfo.GetLab())

	}

}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getTruePath(inPath string, createFlag bool) string {
	var existPath = ""
	flag, error := pathExists(inPath)
	if !flag {
		fmt.Println(error, "开始尝试相对路径")
		if createFlag {
			_, _ = os.Create(inPath) //创建文件
			fmt.Println("文件不存在,在路径：", inPath, "创建文件")
			existPath = inPath
		} else {
			if strings.Count(inPath, "") == 1 {
				return existPath
			}
			//尝试相对路径
			flag2, error2 := pathExists(getCurrentDirectory() + inPath)
			if flag2 {
				existPath = getCurrentDirectory() + inPath
				fmt.Println("flag2:", existPath)

			} else {
				fmt.Println(error2)
			}
		}

	} else {
		//文件存在
		existPath = inPath
	}

	return existPath
}
