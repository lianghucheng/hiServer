package ccsu

import (
	"regexp"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"hiServer/models"
	//"github.com/astaxie/beego/orm"
)

type TemporaryCourse struct {
	order      string
	weekday    string
	courseBody string
	name       string
	class      string
	existsweek string
	position   string
	teacher    string
}

var Course = regexp.MustCompile(`<div id="([\d])-([\d])-2" >(.+)</div>`)
var Detail = regexp.MustCompile(`([^<]+)<br>([^<]+)<br>[\[]([^\]]+)[\]]<br>[\(]([^\)]+)周[\)]<br>([^<]+)<br>`)

func CcsuListParser(body string) {

	getCourseBodyMatches := Course.FindAllStringSubmatch(body, -1)
	var temporaryCourses, resultCourse []TemporaryCourse

	for _, courseBodyMatch := range getCourseBodyMatches {

		order := courseBodyMatch[1]
		weekday := courseBodyMatch[2]
		courseBody := courseBodyMatch[3]

		temporarycourse := TemporaryCourse{
			order:      order,
			weekday:    weekday,
			courseBody: courseBody,
		}

		temporaryCourses = append(temporaryCourses, temporarycourse)
	}

	//beego.Debug(temporaryCourse)
	for _, temporaryCourse := range temporaryCourses {
		detailMatches := Detail.FindAllStringSubmatch(temporaryCourse.courseBody, -1)

		for _, detailMatch := range detailMatches {
			name := detailMatch[1]
			class := detailMatch[2]
			teacher := detailMatch[3]
			position := detailMatch[5]



			existweek:=ExtractExistweek(detailMatch[4])
			//beego.Debug(existweek)

			temporaryCourse.name = name
			temporaryCourse.class = class
			temporaryCourse.teacher = teacher
			temporaryCourse.existsweek = existweek
			temporaryCourse.position = position

			/*
			beego.Debug("name---",name)
			beego.Debug("class---",class)
			beego.Debug("existsweek---",existweek)
			beego.Debug("position---",position)
			beego.Debug("teacher---",teacher)
			beego.Debug("========================")
			beego.Debug("order",temporaryCourse.order)
			beego.Debug("weekday",temporaryCourse.weekday)
			beego.Debug("name",temporaryCourse.name)
			beego.Debug("class",temporaryCourse.class)
			beego.Debug("existsweek",temporaryCourse.existsweek)
			beego.Debug("position",temporaryCourse.position)
			beego.Debug("teacher",temporaryCourse.teacher)
			*/

			resultCourse = append(resultCourse, temporaryCourse)
		}
	}
	//beego.Debug("resultCourse",temporaryCourse)

	//课表解析的结果展示
	//printTemporaryResultCourse(resultCourse)

	user:=models.User{
		StuNum :"B20160304219",
		Password :"134531",
		Power :models.Normal,
		Identity :models.Student,
	}
	var courses []models.Course
	for _,resultCourseValue:=range resultCourse{
		course:=models.Course{
			Order :resultCourseValue.order,
			Weekday:resultCourseValue.weekday,
			Name:resultCourseValue.name,
			Class:resultCourseValue.class,
			Teacher:resultCourseValue.teacher,
			ExistWeek:resultCourseValue.existsweek,
			Position:resultCourseValue.position,
			User:&user,
		}
		courses=append(courses,course)
	}
	//printModelResultCourse(courses)
	/*for _,courseValue:=range courses{
		beego.Debug()
		num,err:=courseValue.Insert()
		if err!=nil{
			beego.Debug()
			beego.Error(err)
		}
		beego.Debug(num)
	}*/

	courseRead:=new(models.Course)
	//o:=orm.NewOrm()
	//o.QueryTable("course").Filter("name","数据库系统原理").One(courseRead)
	beego.Debug(courseRead)

	return
}

func splitComma(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}

func splitBar(s rune) bool {
	if s == '-' {
		return true
	}
	return false
}

func ExtractExistweek(existweekMatch string)(string){
	var resultExistweek string
	var resultExistweekSlice []string

	/*for _, detail := range detailMatch {
		beego.Debug("order,weekday:", temporaryCourse.order, temporaryCourse.weekday, detail)
		beego.Debug()
	}*/
	//beego.Debug(len(detailMatch))

	existweekSlice := strings.FieldsFunc(existweekMatch, splitComma)
	//显示切片原始数据
	//beego.Debug("existweek---", existweekSlice)
	for _, existweekSliceElement := range existweekSlice {
		if len(existweekSliceElement) >= 3 {
			weekRange := strings.FieldsFunc(existweekSliceElement, splitBar)
			weekRangeStart:=weekRange[0]
			weekRangeEnd:=weekRange[1]
			start, err := strconv.Atoi(weekRangeStart)
			if err != nil {
				beego.Error(err)
				return ""
			}
			end, err := strconv.Atoi(weekRangeEnd)
			if err != nil {
				beego.Error(err)
				return ""
			}
			var reduction []string
			for i := start; i <= end; i++ {
				tem_reduction := strconv.Itoa(i)
				reduction = append(reduction, tem_reduction)
			}
			//beego.Debug(reduction)
			resultExistweekSlice = append(resultExistweekSlice, reduction...)
		} else {
			resultExistweekSlice = append(resultExistweekSlice, existweekSliceElement)
		}
		//beego.Debug(resultExistweekSlice)
	}
	//beego.Debug(resultExistweekSlice)
	resultExistweek = strings.Join(resultExistweekSlice, ",")
	return resultExistweek
}

func printTemporaryResultCourse(resultCourse []TemporaryCourse){
	for _,courseValue:=range resultCourse{
		beego.Debug("==========================================")
		beego.Debug("==========================================")
		beego.Debug("order",courseValue.order)
		beego.Debug("weekday",courseValue.weekday)
		beego.Debug("name",courseValue.name)
		beego.Debug("class",courseValue.class)
		beego.Debug("existsweek",courseValue.existsweek)
		beego.Debug("position",courseValue.position)
		beego.Debug("teacher",courseValue.teacher)
	}
}

func printModelResultCourse(courses []models.Course){
	for _,courseValue:=range courses{
		beego.Debug("==========================================")
		beego.Debug("==========================================")
		beego.Debug("order",courseValue.Order)
		beego.Debug("weekday",courseValue.Weekday)
		beego.Debug("name",courseValue.Name)
		beego.Debug("class",courseValue.Class)
		beego.Debug("existsweek",courseValue.ExistWeek)
		beego.Debug("position",courseValue.Position)
		beego.Debug("teacher",courseValue.Teacher)
	}
}