package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nosixtools/solarlunar"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Happy struct {
	Title    string   `json:"title"`
	Contents []string `json:"contents"`
}
type Sad struct {
	Title    string   `json:"title"`
	Contents []string `json:"contents"`
}
type HappySlice []Festival
type SadSlice []Involution

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

const (
	Year = 365 //一年多少天
)

type Festival struct {
	ChineseName string `json:"chinese_name"` //节日中文名
	SubDay      int    `json:"sub_day"`      //节日剩余日期
	Date        string `json:"date"`         //节日日期
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

type Involution struct { //内卷
	ChineseName string `json:"chinese_name"` //日期中文名
	SubDay      int    `json:"sub_day"`      //日期倒计时
	Date        string `json:"date"`         //当天
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

var (
	Lyingflat                  []Involution
	CollegeEntranceExamination Involution //高考

)

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
func allInSolar(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}

//特殊日期 如奥运会
func allInThisYear(date string) int {
	day := date
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}
func allInLuna(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	convert := solarlunar.LunarToSolar(day, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}
func allInSpring(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	convert := solarlunar.LunarToSolar(day, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours())/24 - 1
}

//获取当前年份
func thisYear() string {
	ret := fmt.Sprint(time.Now().Format("2006"))
	return ret
}
func nextYear() string {
	this := thisYear()
	i, _ := strconv.Atoi(this)
	i += 1
	s := strconv.Itoa(i)
	return s
}
func Calendar() {
	defer AnniversaryDay()
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) > 3 { //0-9
		HappyDay()
	} else {
		SadDay()
	}
}
func HappyDay() {
	HappyTimer()
	defer func() {
		fmt.Println("上班是帮老板赚钱,摸鱼是赚老板的钱!")
		fmt.Println("该休息就休息,该放松就放松")
		Gift()
		fmt.Println("最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天")
	}()
	for _, v := range Countdown {
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			//fmt.Printf("明天是%v\n", v.GetChineseName())
			continue
		}
	}
	for _, v := range Countdown {
		if v.GetSubDay() < 0 {
			continue
		}

		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			fmt.Printf("明天是%v\n", v.GetChineseName())
			continue
		}
		fmt.Printf("距离%v还有%v天\n", v.GetChineseName(), v.GetSubDay())
	}
}
func HappyDayJson() []string {
	days := []string{}
	for _, v := range Countdown {
		day := fmt.Sprintf("距离%v还有%v天", v.GetChineseName(), v.GetSubDay())
		days = append(days, day)
	}
	return days
}
func SadDay() {
	SadTimer()
	for _, v := range Lyingflat {
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			fmt.Printf("\t%v\n", v.GetChineseName())
		}
	}
	for _, v := range Lyingflat {
		if v.GetSubDay() < 0 {
			continue
		}
		if v.GetSubDay() == 0 || v.GetSubDay() == 365 {
			//fmt.Printf("明天是%v\n", v.GetChineseName())
			continue
		}
		fmt.Printf("距离%v还有%v天\n", v.GetChineseName(), v.GetSubDay())
	}
}
func SadDayJson() []string {
	days := []string{}
	for _, v := range Lyingflat {
		day := fmt.Sprintf("距离%v还有%v天", v.GetChineseName(), v.GetSubDay())
		days = append(days, day)
	}
	return days
}

//计算和元旦的差值
func nextNewYear() {
	thisYearInt, _ := strconv.Atoi(thisYear())
	nextYearInt := thisYearInt + 1
	nextYearStr := strconv.Itoa(nextYearInt)
	day := strings.Join([]string{nextYearStr, "01-01"}, "-") //2021-01-01
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离明年还有%v天\n", int(unsub.Hours())/24)
	if int(unsub.Hours())/24 <= 31 {
		fmt.Println("过几天又会有人发\"新的一年,新的自己\"这种自欺欺人的话")
	}
}

// ToDo 实现通过网页请求

func ShowWeb() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context，封装了request和response
	r.GET("/HappyCount", HappyDayCountdown)
	r.GET("/SadCount", SadDayCountdown)
	r.GET("/", StandRequest)
	//3.监听端口，默认在8080
	//Run("里面不指定端口号默认为8080")
	if err := r.SetTrustedProxies([]string{"0.0.0.0"}); err != nil {
		return
	}
	//gin.SetMode(gin.ReleaseMode)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func HappyDayCountdown(ctx *gin.Context) {
	content := Happy{
		Title:    "Happy",
		Contents: HappyDayJson(),
	}
	ctx.JSON(http.StatusOK, content)
}
func SadDayCountdown(ctx *gin.Context) {
	content := Sad{
		Title:    "Sad",
		Contents: SadDayJson(),
	}
	ctx.JSON(http.StatusOK, content)
}

type stand struct {
	Title []string `json:"title"`
	Count []string `json:"count"`
	Tail  []string `json:"tail"`
}

func StandRequest(ctx *gin.Context) {
	content := stand{
		Title: StandHeadJson(),
		Count: HappyDayJson(),
		Tail:  StandTailJson(),
	}
	ctx.JSON(http.StatusOK, content)
}
func StandHeadJson() []string {
	Heads := []string{}
	nowDay := time.Now().Format("今天是2006年1月2日")
	//Heads=append(Heads,nowDay)
	nowTime := time.Now().Format("PM 3点04分05秒")
	//Heads=append(Heads,nowTime)
	week := time.Now().Weekday().String()
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	nowTime = strings.Replace(nowTime, "AM", "上午", -1)
	nowTime = strings.Replace(nowTime, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		Heads = append(Heads, fmt.Sprint("早上好,摸鱼人!"))
	}
	if inthour == 12 {
		Heads = append(Heads, fmt.Sprint("中午好,摸鱼人!"))
	}
	if inthour >= 13 && inthour <= 17 {
		Heads = append(Heads, fmt.Sprint("下午好,摸鱼人!"))
	}
	if inthour >= 18 && inthour <= 22 {
		Heads = append(Heads, fmt.Sprint("晚上好,摸鱼人!"))
	}
	if inthour >= 23 && inthour <= 5 {
		Heads = append(Heads, fmt.Sprint("晚上抓紧睡觉,休息好,第二天才有精神摸鱼!"))
	}
	Heads = append(Heads, fmt.Sprint(nowDay, HappyWeekDayMap[week], nowTime))
	Heads = append(Heads, fmt.Sprint("上班是帮老板赚钱,摸鱼是赚老板的钱!"))
	Heads = append(Heads, fmt.Sprint("该休息就休息,该放松就放松"))

	if inthour == 15 {
		Heads = append(Heads, fmt.Sprint("喂!三点几咧!做......做撚啊做!"))
		Heads = append(Heads, fmt.Sprint("饮茶先啊!"))
		Heads = append(Heads, fmt.Sprint("三点几 饮......饮茶先啊!"))
		Heads = append(Heads, fmt.Sprint("做咁多都冇用嘅,老细唔锡你嘅咧!"))
		Heads = append(Heads, fmt.Sprint("喂!饮下茶先啊!三点几咧!"))
		Heads = append(Heads, fmt.Sprint("做碌鸠啊做!"))
	}
	if inthour == 18 {
		Heads = append(Heads, fmt.Sprint("喂!朋友!"))
		Heads = append(Heads, fmt.Sprint("做乜嘢咁多啦?差唔多七点咧!"))
		Heads = append(Heads, fmt.Sprint("放工啦!唔使做咁多啦!"))
		Heads = append(Heads, fmt.Sprint("做咁多,钱带去边度?"))
		Heads = append(Heads, fmt.Sprint("差唔多七点咧!放工!"))
		Heads = append(Heads, fmt.Sprint("落去茶室,饮下靓靓嘅杯......麦啤酒、黑啤酒,OK?"))
		Heads = append(Heads, fmt.Sprint("Happy下,唔使做咁多."))
		Heads = append(Heads, fmt.Sprint("死咗都冇用咧,银纸冇得带去咧!"))
		Heads = append(Heads, fmt.Sprint("Happy下,饮酒,OK?"))
	}
	Heads = append(Heads, fmt.Sprint("【摸鱼办】提醒您"))
	Heads = append(Heads, fmt.Sprint("工作再累,一定不要忘记摸鱼哦!"))
	Heads = append(Heads, fmt.Sprint("有事没事起身去茶水间,去厕所,去廊道走走"))
	Heads = append(Heads, fmt.Sprint("吃饭时间就吃饭"))
	Heads = append(Heads, fmt.Sprint("午休时间就午休"))
	Heads = append(Heads, fmt.Sprint("别老在工位上坐着,钱是老板的,但命是自己的"))
	Heads = append(Heads)

	return Heads
}
func StandTailJson() []string {
	tails := []string{}
	tails = append(tails, fmt.Sprint("最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天"))
	return tails
}
