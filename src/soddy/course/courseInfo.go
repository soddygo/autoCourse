package course

import (
	"time"
	"fmt"
	"strings"
)

//课程信息对象
type Course struct {
	teacher []string
	studentClass []string
	lab []string
	startDate []string
	endDate []string

}

//计算实验课时间，周一到周五
func WorkDay(startDate string,endDate string) (dayArray []time.Time){
	startTime,_ := time.Parse("2016-01-02",startDate)
	endTime,_ := time.Parse("2016-01-02",endDate)

	//计算时间相差天数
	//sub := endTime.Sub(startTime)
	//subDay := sub.Hours()/24
	//dayArray = [subDay]time.Time{}
	//循环所有日期，计算周一到周五的日期
	tempDate := startTime
	for i :=0 ;tempDate.Before(endTime)==true;i++{
		weekday := tempDate.Weekday().String()
		if strings.Compare(weekday,"Sunday") ==0 ||strings.Compare(weekday,"Saturday") ==0 {
			continue
		}
		dayArray[i] = tempDate
		tempDate.AddDate(0,0,1)
		tempDate.Format("2006-01-02")
		fmt.Printf(tempDate.String())//打印时间
	}
	return dayArray
}

//初始化
func (self *Course)init(teacher []string,studentClass []string,lab []string,startDate []string,endDate []string)  {
	self.teacher = teacher
	self.studentClass = studentClass
	self.lab = lab
	self.startDate = startDate
	self.endDate = endDate
}
//组合班级和实验课的集合
func (self *Course) AllCourseAndClass()  map[string][len(self.studentClass)]string{
	var labMap  map[string][len(self.studentClass)]string
	for _,lab := range self.lab{
		students := [len(self.studentClass)]string{}

		for j,student := range self.studentClass{
			students[j]= student
		}
		labMap[lab] = students
	}

	return labMap

}

//按照上下学期，随机给出课程和班级。每个班级，上学期上6个实验课，下学期上6个实验课
func (self *Course)RandomCourse()  {

}