package controller

import (
	"log"
	"sort"
)
const (
	Year = 365 //一年多少天
)
var (
	Countdown []Festival
	//公历节日
	NewYear                  Festival // 元旦
	ValentinesDay            Festival //情人节
	WomensDay                Festival //妇女节
	ArborDay                 Festival //植树节
	AprilFoolsDay            Festival //愚人节
	LaborDay                 Festival //劳动节
	YouthDay                 Festival //青年节
	ChildrensDay             Festival //儿童节
	NationalDay              Festival //国庆节
	NationalPovertyReliefDay Festival //国家扶贫日
	Halloween                Festival //万圣夜
	Double11                 Festival //
	NationalConstitutionDay  Festival ////国家宪法日
	ThanksgivingDay          Festival //感恩节
	Double12                 Festival //
	NationalMemorialDay      Festival //国家公祭日
	ChristmasEve             Festival //平安夜
	Christmas                Festival //圣诞节
	//农历节日
	ChineseNewYear      Festival //春节
	LanternFestival     Festival //元宵节
	ChingMingFestival   Festival //清明节
	MothersDay          Festival //母亲节
	DragonBoatFestival  Festival //端午节
	SingleDogDay        Festival //七夕节
	MidyearFestival     Festival //中元节
	MidAutumnFestival   Festival //中秋节
	DoubleNinthFestival Festival //重阳节
	LabaFestival        Festival //腊八节
	//其他
	WinterOlympics Festival //冬奥会
)
var HappyWeekDayMap = map[string]string{
	"Monday":    "星期一,还有四天周末",
	"Tuesday":   "星期二,还有三天周末",
	"Wednesday": "星期三,工作日过去一半了",
	"Thursday":  "肯德基疯狂星期四",
	"Friday":    "星期五,明天就放假了",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
}
var SadWeekDayMap = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五不能松懈,明天一定要去公司加班哦!",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
}
var (
	Lyingflat                  []Involution
	CollegeEntranceExamination Involution //高考
)
type Festival struct {
	ChineseName string `json:"chinese_name"` //节日中文名
	SubDay      int    `json:"sub_day"`      //节日剩余日期
	Date        string `json:"date"`         //节日日期
}
type Stand struct {
	Title []string `json:"title"`
	Count []string `json:"count"`
	Tail  []string `json:"tail"`
}
type Involution struct { //内卷
	ChineseName string `json:"chinese_name"` //日期中文名
	SubDay      int    `json:"sub_day"`      //日期倒计时
	Date        string `json:"date"`         //当天
}
type Happy struct {
	Title    string   `json:"title"`
	Contents []string `json:"contents"`
	Tail []string `json:"tail"`
}
type Sad struct {
	Title    string   `json:"title"`
	Contents []string `json:"contents"`
}
type HappySlice []Festival
type SadSlice []Involution
func init() {
	NewYear.SetChineseName("元旦")
	NewYear.SetDate("01-01")
	NewYear.SetSubDay(allInSolar(NewYear.GetDate()))
	Countdown = append(Countdown, NewYear)

	ValentinesDay.SetChineseName("情人节")
	ValentinesDay.SetDate("02-14")
	ValentinesDay.SetSubDay(allInSolar(ValentinesDay.GetDate()))
	Countdown = append(Countdown, ValentinesDay)

	WomensDay.SetChineseName("妇女节")
	WomensDay.SetDate("03-08")
	WomensDay.SetSubDay(allInSolar(WomensDay.GetDate()))
	Countdown = append(Countdown, WomensDay)

	ArborDay.SetChineseName("植树节")
	ArborDay.SetDate("03-12")
	ArborDay.SetSubDay(allInSolar(ArborDay.GetDate()))
	Countdown = append(Countdown, ArborDay)

	AprilFoolsDay.SetChineseName("愚人节")
	AprilFoolsDay.SetDate("04-01")
	AprilFoolsDay.SetSubDay(allInSolar(AprilFoolsDay.GetDate()))
	Countdown = append(Countdown, AprilFoolsDay)

	LaborDay.SetChineseName("劳动节")
	LaborDay.SetDate("05-01")
	LaborDay.SetSubDay(allInSolar(LaborDay.GetDate()))
	Countdown = append(Countdown, LaborDay)

	YouthDay.SetChineseName("青年节")
	YouthDay.SetDate("05-04")
	YouthDay.SetSubDay(allInSolar(YouthDay.GetDate()))
	Countdown = append(Countdown, YouthDay)

	ChildrensDay.SetChineseName("儿童节")
	ChildrensDay.SetDate("06-01")
	ChildrensDay.SetSubDay(allInSolar(ChildrensDay.GetDate()))
	Countdown = append(Countdown, ChildrensDay)

	NationalDay.SetChineseName("国庆节")
	NationalDay.SetDate("10-01")
	NationalDay.SetSubDay(allInSolar(NationalDay.GetDate()))
	Countdown = append(Countdown, NationalDay)

	NationalPovertyReliefDay.SetChineseName("国家扶贫日")
	NationalPovertyReliefDay.SetDate("10-17")
	NationalPovertyReliefDay.SetSubDay(allInSolar(NationalPovertyReliefDay.GetDate()))
	Countdown = append(Countdown, NationalPovertyReliefDay)

	Halloween.SetChineseName("万圣夜")
	Halloween.SetDate("10-31")
	Halloween.SetSubDay(allInSolar(Halloween.GetDate()))
	Countdown = append(Countdown, Halloween)

	Double11.SetChineseName("双十一")
	Double11.SetDate("11-11")
	Double11.SetSubDay(allInSolar(Double11.GetDate()))
	Countdown = append(Countdown, Double11)

	NationalConstitutionDay.SetChineseName("国家宪法日")
	NationalConstitutionDay.SetDate("12-04")
	NationalConstitutionDay.SetSubDay(allInSolar(NationalConstitutionDay.GetDate()))
	Countdown = append(Countdown, NationalConstitutionDay)

	ThanksgivingDay.SetChineseName("感恩节")
	ThanksgivingDay.SetDate("11-25")
	ThanksgivingDay.SetSubDay(allInSolar(ThanksgivingDay.GetDate()))
	Countdown = append(Countdown, ThanksgivingDay)

	Double12.SetChineseName("双十二")
	Double12.SetDate("12-12")
	Double12.SetSubDay(allInSolar(Double12.GetDate()))
	Countdown = append(Countdown, Double12)

	NationalMemorialDay.SetChineseName("国家公祭日")
	NationalMemorialDay.SetDate("12-13")
	NationalMemorialDay.SetSubDay(allInSolar(NationalMemorialDay.GetDate()))
	Countdown = append(Countdown, NationalMemorialDay)

	ChristmasEve.SetChineseName("平安夜")
	ChristmasEve.SetDate("12-24")
	ChristmasEve.SetSubDay(allInSolar(ChristmasEve.GetDate()))
	Countdown = append(Countdown, ChristmasEve)

	Christmas.SetChineseName("圣诞节")
	Christmas.SetDate("12-25")
	Christmas.SetSubDay(allInSolar(Christmas.GetDate()))
	Countdown = append(Countdown, Christmas)

	ChineseNewYear.SetChineseName("除夕")
	ChineseNewYear.SetDate("01-01")
	//ToDo 农历的春节时间暂时使用当年1号实现,年底如果发现问题再更改
	ChineseNewYear.SetSubDay(allInSpring(ChineseNewYear.GetDate()))
	Countdown = append(Countdown, ChineseNewYear)

	LanternFestival.SetChineseName("元宵节")
	LanternFestival.SetDate("01-15")
	LanternFestival.SetSubDay(allInLuna(LanternFestival.GetDate()))
	Countdown = append(Countdown, LanternFestival)

	ChingMingFestival.SetChineseName("清明节")
	ChingMingFestival.SetDate("03-05")
	ChingMingFestival.SetSubDay(allInLuna(ChingMingFestival.GetDate()))
	Countdown = append(Countdown, ChingMingFestival)

	MothersDay.SetChineseName("母亲节")
	MothersDay.SetDate("04-08")
	MothersDay.SetSubDay(allInLuna(MothersDay.GetDate()))
	Countdown = append(Countdown, MothersDay)

	DragonBoatFestival.SetChineseName("端午节")
	DragonBoatFestival.SetDate("05-05")
	DragonBoatFestival.SetSubDay(allInLuna(DragonBoatFestival.GetDate()))
	Countdown = append(Countdown, DragonBoatFestival)

	SingleDogDay.SetChineseName("七夕节")
	SingleDogDay.SetDate("07-07")
	SingleDogDay.SetSubDay(allInLuna(SingleDogDay.GetDate()))
	Countdown = append(Countdown, SingleDogDay)

	MidyearFestival.SetChineseName("中元节")
	MidyearFestival.SetDate("07-15")
	MidyearFestival.SetSubDay(allInLuna(MidyearFestival.GetDate()))
	Countdown = append(Countdown, MidyearFestival)

	MidAutumnFestival.SetChineseName("中秋节")
	MidAutumnFestival.SetDate("08-15")
	MidAutumnFestival.SetSubDay(allInLuna(MidAutumnFestival.GetDate()))
	Countdown = append(Countdown, MidAutumnFestival)

	DoubleNinthFestival.SetChineseName("重阳节")
	DoubleNinthFestival.SetDate("09-09")
	DoubleNinthFestival.SetSubDay(allInLuna(DoubleNinthFestival.GetDate()))
	Countdown = append(Countdown, DoubleNinthFestival)

	LabaFestival.SetChineseName("腊八节")
	LabaFestival.SetDate("12-08")
	LabaFestival.SetSubDay(allInLuna(LabaFestival.GetDate()))
	Countdown = append(Countdown, LabaFestival)

	WinterOlympics.SetChineseName("2022冬奥会")
	WinterOlympics.SetDate("2022-02-04")
	WinterOlympics.SetSubDay(allInThisYear(WinterOlympics.GetDate()))
	Countdown = append(Countdown, WinterOlympics)

	sort.Sort(HappySlice(Countdown))
}
func init() {
	log.SetFlags(8 | 5)
}
func init() {
	CollegeEntranceExamination.SetChineseName("高考")
	CollegeEntranceExamination.SetDate("06-07")
	CollegeEntranceExamination.SetSubDay(allInSolar(CollegeEntranceExamination.GetDate()))
	Lyingflat = append(Lyingflat, CollegeEntranceExamination)

	sort.Sort(SadSlice(Lyingflat))
}
func (c HappySlice) Len() int { // 重写 Len() 方法
	return len(c)
}
func (c HappySlice) Swap(i, j int) { // 重写 Swap() 方法
	c[i], c[j] = c[j], c[i]
}
func (c HappySlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return c[i].SubDay < c[j].SubDay
}
func (s SadSlice) Len() int { // 重写 Len() 方法
	return len(s)
}
func (s SadSlice) Swap(i, j int) { // 重写 Swap() 方法
	s[i], s[j] = s[j], s[i]
}
func (s SadSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return s[i].SubDay < s[j].SubDay
}
func (f *Festival) SetChineseName(s string) {
	f.ChineseName = s
	return
}
func (f Festival) GetChineseName() string {
	return f.ChineseName
}
func (f *Festival) SetSubDay(i int) {
	f.SubDay = i
	return
}
func (f Festival) GetSubDay() int {
	return f.SubDay
}
func (f *Festival) SetDate(s string) {
	f.Date = s
	return
}
func (f Festival) GetDate() string {
	return f.Date
}
func (i *Involution) SetChineseName(s string) {
	i.ChineseName = s
	return
}
func (i Involution) GetChineseName() string {
	return i.ChineseName
}
func (i *Involution) SetSubDay(n int) {
	i.SubDay = n
	return
}
func (i Involution) GetSubDay() int {
	return i.SubDay
}
func (i *Involution) SetDate(s string) {
	i.Date = s
	return
}
func (i Involution) GetDate() string {
	return i.Date
}