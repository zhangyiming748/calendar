package calendar

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	AS  AnniversarySlice
	DSU Anniversary //string = "1991年12月25日" //苏联解体
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
	DSU.SetName("苏联解体")
	DSU.SetDate("1991年12月25日")
	DSU.SetSubDay(CountDay(DSU.GetDate()))
	AS = append(AS, DSU)

	sort.Sort(AS)
}

func CountDay(date string) int {
	ret, _ := time.Parse("2006年01月02日", date)
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
	for _, v := range AS {
		fmt.Printf("距离%v已经过去了%v周年\n", v.GetName(), v.GetSubDay())
	}
}
