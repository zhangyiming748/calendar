package calendar

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	AS           AnniversarySlice
	DSU          Anniversary //苏联解体
	QVOD         Anniversary //快播关闭
	XC           Anniversary //辛丑条约
	PY           Anniversary //溥仪即位
	XH           Anniversary //辛亥革命
	SY           Anniversary //九一八事变
	Чорнобиль    Anniversary //切尔诺贝利
	福島第一原子力発電所事故 Anniversary //福岛核泄漏
	GodBirth     Anniversary //birth
	Apple Anniversary//iphone发布
)

type Anniversary struct {
	name   string //中文名
	subday int    //剩余日期
	date   string //日期
}

func (a *Anniversary) SetName(s string) {
	a.name = s
	return
}
func (a Anniversary) GetName() string {
	return a.name
}
func (a *Anniversary) SetSubDay(i int) {
	a.subday = i
	return
}
func (a Anniversary) GetSubDay() int {
	return a.subday
}
func (a *Anniversary) SetDate(s string) {
	a.date = s
	return
}
func (a Anniversary) GetDate() string {
	return a.date
}

type AnniversarySlice []Anniversary

func (s AnniversarySlice) Len() int { // 重写 Len() 方法
	return len(s)
}
func (s AnniversarySlice) Swap(i, j int) { // 重写 Swap() 方法
	s[i], s[j] = s[j], s[i]
}
func (s AnniversarySlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return s[i].subday < s[j].subday
}
func init() {


	QVOD.SetName("快播关闭")
	QVOD.SetDate("2014年4月16日")
	QVOD.SetSubDay(CountDay(QVOD.GetDate()))
	AS = append(AS, QVOD)


	Apple.SetName("Iphone发布")
	Apple.SetDate("2007年1月9日")
	Apple.SetSubDay(CountDay(Apple.GetDate()))
	AS=append(AS,Apple)

	sort.Sort(AS)
}

func CountDay(date string) int {
	ret, _ := time.Parse("2006年1月2日", date)
	old := fmt.Sprintf("%v", ret)
	oldYear := strings.Split(old, "-")[0]
	thisYear := fmt.Sprint(time.Now().Format("2006"))
	oi, _ := strconv.Atoi(oldYear)
	ti, _ := strconv.Atoi(thisYear)
	sub := ti - oi
	//fmt.Printf("old is %v\nthis is %v\noldYear is %v\nsub is %v\n", old, thisYear, oldYear, sub)
	return sub
}
func AnniversaryDay() {
	//fmt.Println("P.S.")
	for _, v := range AS {
		//fmt.Printf("距离%v已经过去了%v周年\n", v.GetName(), v.GetSubDay())
		switch v.GetName() {
		case "快播关闭":
			fmt.Printf("距离%v已经过去了%v周年\t但你还欠王欣一个年费会员\n", v.GetName(), v.GetSubDay())
		default:
			fmt.Printf("距离%v已经过去了%v周年\n", v.GetName(), v.GetSubDay())
		}
	}
}
