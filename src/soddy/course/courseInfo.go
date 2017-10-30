package course

import "time"

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

	time.Now()
	//return time.Weekday().String()
	return nil
}