package course

import (
	"time"
	"fmt"
	"strings"
	"math/rand"
)

//课程信息对象
type Course struct {
	teacher      []string
	studentClass []string
	lab          []string
	startDate    []string
	endDate      []string
	labMap       map[string][]string
}

//学期,课程
type SchoolTerm struct {
	index        int
	sClassMapLab map[string][]*Lab
}

//实验课对象
type Lab struct {
	sClass     string
	companyMap map[int]bool //连队，一个班级2个连队,bool false，表示此连队没上过这个实验课
	lab        string
	flag       bool
}

//计算实验课时间，周一到周五
func WorkDay(startDate string, endDate string) (dayArray []time.Time) {
	startTime, _ := time.Parse("2016-01-02", startDate)
	endTime, _ := time.Parse("2016-01-02", endDate)

	//计算时间相差天数
	//sub := endTime.Sub(startTime)
	//subDay := sub.Hours()/24
	//dayArray = [subDay]time.Time{}
	//循环所有日期，计算周一到周五的日期
	tempDate := startTime
	for i := 0; tempDate.Before(endTime) == true; i++ {
		weekday := tempDate.Weekday().String()
		if strings.Compare(weekday, "Sunday") == 0 || strings.Compare(weekday, "Saturday") == 0 {
			continue
		}
		dayArray[i] = tempDate
		tempDate.AddDate(0, 0, 1)
		tempDate.Format("2006-01-02")
		fmt.Printf(tempDate.String()) //打印时间
	}
	return dayArray
}

//初始化
func (self *Course) init(teacher []string, studentClass []string, lab []string, startDate []string, endDate []string) {
	self.teacher = teacher
	self.studentClass = studentClass
	self.lab = lab
	self.startDate = startDate
	self.endDate = endDate
}

//组合班级和实验课的集合,班级作为key，实验课lab作为values
func (self *Course) AllCourseAndClass() map[string][]string {
	var labMap map[string][]string

	for _, student := range self.studentClass {
		labslice := make([]string, 12)
		for j, lab := range self.lab {
			labslice[j] = lab
		}
		labMap[student] = labslice
	}
	self.labMap = labMap
	return labMap

}

//按照上下学期，随机给出课程和班级。每个班级，上学期上6个实验课，下学期上6个实验课
func (self *Course) RandomCourse() *SchoolTerm {
	//每个班级，随机抽6个实验课
	//1 先随机抽6个实验课
	randLabs := [6]string{}
	lab1Count := 0
	if len(self.lab) >= 12 {
		for {
			r := rand.Intn(12)
			curLab := self.lab[r]
			//判断数组里是否包含了此实验课，不包含则加入
			containFlag := false
			for _, lab := range randLabs {
				if lab == curLab {
					containFlag = true
					break
				}
			}
			if !containFlag {
				//没有此实验课，放入随机实验课的数组里ß
				randLabs[lab1Count] = self.lab[r]
				lab1Count++
			}
			//当满足随机实验课够6个的时候，就跳出循环
			if len(randLabs) >= 6 {
				break
			}
		}
	}

	//2 剩下的实验课，放入另外的数组里,用于下学期使用
	randLabs2 := [6]string{}
	lab2Count := 0
	for _, lab2 := range self.lab {
		containFlag := false
		for _, lab1 := range randLabs {
			if lab2 == lab1 {
				containFlag = true
				break
			}
		}
		if !containFlag {
			randLabs2[lab2Count] = lab2
			lab2Count++
		}
	}

	var studentMap = map[string][]*Lab{}
	//抽取一半的班级，上半学期，去学randLabs里面的实验课

	for _, sClass := range self.studentClass {
		//上半学期的课程
		var labTerm = make([]*Lab, 6)
		for j, labStr := range randLabs {
			var lab = &Lab{}
			lab.sClass = sClass
			lab.lab = labStr
			//连队初始化
			lab.companyMap = map[int]bool{0: false, 1: false}
			lab.flag = false
			labTerm[j] = lab
		}
		studentMap[sClass] = labTerm

		//下半学期的课程
		var labTerm2 = make([]*Lab, 6)
		for j, labStr := range randLabs2 {
			var lab = &Lab{}
			lab.sClass = sClass
			lab.lab = labStr
			//连队初始化
			lab.companyMap = map[int]bool{0: false, 1: false}
			lab.flag = false
			labTerm2[j] = lab
		}
		studentMap[sClass] = labTerm2
	}

	//上学期对象
	var term1 = &SchoolTerm{}
	term1.index = 1
	term1.sClassMapLab = studentMap

	return term1
}

//排课，按照天为单位，顺序列出
//一天10-12个实验课，一学期实验课时间是6周，一周是3天实验课时间
func (self *Course) sortClass(term *SchoolTerm) {
	//先按照1天12个实验课排，一个班级，一周只上一天
	//先按照一个班级一个实验来排
	var termLabs = make([][]*Lab, 18)
	var dayLabs = make([]*Lab, 12)
	//直到所有的课程和班级安排完，跳出
	for {

		for _, lab := range self.lab {

			//遍历
			for k, v := range term.sClassMapLab {
				//遍历l这个班级的实验课，是否已被安排

			}

		}

	}

}

//是否包含实验课判断，company是连队信息
func hasContainLab(labArray []*Lab, labStr string, complanyFlag bool, company int) bool {
	var flag = false

	if !complanyFlag {
		for _, lab := range labArray {
			if lab.lab == labStr && lab.flag == true {
				flag = true
				break
			}

		}
	} else {
		//需要对连队进行判断
		for _, lab := range labArray {
			if lab.lab == labStr && lab.companyMap[company] == true {
				flag = true
				break
			}
		}
	}

	return flag
}
