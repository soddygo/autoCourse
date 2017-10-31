package main

import "soddy/course"

func main() {
	filePath := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo.xlsx"
	filePath2 := "/Users/soddygosongguochao/Documents/Go_Project/autoCourse/src/files/initInfo2.xlsx"
	course.ReadCourse(filePath)
	course.WriteExcel(filePath2)
}
