package course

import (
	"time"
	"fmt"
)

//课程信息对象
type CourseInfo struct {
	teacher []string
	student []string
	lab []string
	startDate []string
	endDate []string

}

//计算实验课时间，周一到周五
func WorkDay(startDate string,endDate string) []string{
	startTime,_ := time.Parse("2016-01-02",startDate)
	endTime,_ := time.Parse("2016-01-02",endDate)

	//计算时间相差天数
	sub := endTime.Sub(startTime)
	subDay := sub.Hours()/24
	dayArray := [subDay]time.Time{}
	//循环所有日期，计算周一到周五的日期
	tempDate := startTime
	for i :=0 ;tempDate.Before(endTime)==true;i++{
		dayArray[i] = tempDate
		tempDate.AddDate(0,0,1)
		tempDate.Format("2006-01-02")
		fmt.Printf(tempDate.String())//打印时间
		//TODO
	}

	time.Now()
	//return time.Weekday().String()
	return nil
}