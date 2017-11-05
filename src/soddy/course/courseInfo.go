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

func (self *Lab) GetSclass() string {
	return self.sClass
}
func (self *Lab) GetCompanyMap() map[int]bool {
	return self.companyMap
}
func (self *Lab) GetLab() string {
	return self.lab
}
func (self *Lab) GetFlag() bool {
	return self.flag
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
func (self *Course) Init(teacher []string, studentClass []string, lab []string, startDate []string, endDate []string) {
	self.teacher = teacher
	self.studentClass = studentClass
	self.lab = lab
	self.startDate = startDate
	self.endDate = endDate
}

//组合班级和实验课的集合,班级作为key，实验课lab作为values
func (self *Course) AllCourseAndClass() map[string][]string {
	//var labMap map[string][]string
	var labMap = make(map[string][]string)

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
	randLabs := make([]string, 6)
	lab1Count := 0
	if len(self.lab) >= 12 {
		for {
			r := rand.Intn(len(self.lab) - 1)
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
			if lab1Count >= 6 {
				break
			}
		}
	}

	//2 剩下的实验课，放入另外的数组里,用于下学期使用
	randLabs2 := make([]string, 6)
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
		//studentMap[sClass] = labTerm

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
		allTerm := make([]*Lab, len(labTerm)+len(labTerm2))
		copy(allTerm, labTerm)
		copy(allTerm[len(labTerm):], labTerm2)
		studentMap[sClass] = allTerm
	}

	//上学期对象
	var term1 = &SchoolTerm{}
	term1.index = 1
	term1.sClassMapLab = studentMap

	return term1
}

//排课，按照天为单位，顺序列出
//一天10-12个实验课，一学期实验课时间是6周，一周是3天实验课时间
func (self *Course) SortClass(term *SchoolTerm) [][]*Lab {

	//先按照1天12个实验课排，一个班级，一周只上一天
	//先按照一个班级一个实验来排，2个学期共36天实验课
	var termLabs = make([][]*Lab, 0)
	termCount := 0
	//直到所有的课程和班级安排完，跳出
	for {

		var labCount = 0
		var dayLabs = make([]*Lab, 0)
		for _, lab := range self.lab {
			//遍历
			for _, v := range term.sClassMapLab {
				//遍历这个班级的实验课，是否已被安排
				flag, labObj := self.hasContainLab(v, lab, false)
				if !flag {
					//dayLabs[i] = labObj
					dayLabs = append(dayLabs, labObj)
					//修改flag 标记
					labObj.flag = true
					labCount++
					break
				}
			}

		}

		//判断2个学期的课程，是否安排完成了
		//所有班级课程已经安排完成，无命中课程，跳出
		if labCount == 0 {
			break
		}
		//一天的实验课安排满，跳出12个lab循环，重新下一天的课程
		termLabs = append(termLabs, dayLabs)
		termCount++
	}
	return termLabs
}

//是否包含实验课判断，company是连队信息
func (self *Course) hasContainLab(labArray []*Lab, labStr string, complanyFlag bool, company ...int) (bool, *Lab) {
	flag := false
	labObj := &Lab{}
	if !complanyFlag {
		for _, lab := range labArray {
			if lab.lab == labStr {
				if lab.flag == true {
					flag = true
				}
				labObj = lab
				break
			}
		}
	} else {
		//需要对连队进行判断
		for _, lab := range labArray {
			if lab.lab == labStr {
				if lab.companyMap[company[0]] == true {
					flag = true
				}
				labObj = lab
				break
			}
		}
	}
	return flag, labObj
}
