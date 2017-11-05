package main

import (
	"soddy/course"
	"fmt"
)

func main() {
	filePath := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo.xlsx"
	filePath2 := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo2.xlsx"
	termInfo := course.ReadCourse(filePath)


	//course
	courseObj := &course.Course{}
	courseObj.Init(termInfo.GetTeacher(),termInfo.GetSclass(),termInfo.GetLab(),nil,nil)
	courseObj.AllCourseAndClass()
	var schoolTerm = courseObj.RandomCourse()
	var labArray = courseObj.SortClass(schoolTerm)


	for i,labInfos := range labArray{
		fmt.Println("index:",i,"\t")
		for _,labInfo := range labInfos {
			fmt.Print(labInfo.GetSclass(),"\t")
			fmt.Print(labInfo.GetLab(),"\t")
			//fmt.Print(labInfo.GetFlag(),"\t")
			fmt.Println()
		}

	}


	course.WriteExcel(filePath2,labArray,termInfo.GetLab())
}
