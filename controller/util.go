package controller

import (
	"bufio"
	"fmt"
	"github.com/nosixtools/solarlunar"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func HappyDayJson() []string {
	days := []string{}
	for _, v := range Countdown {
		day := fmt.Sprintf("距离%v还有%v天", v.GetChineseName(), v.GetSubDay())
		days = append(days, day)
	}
	return days
}
func SadDayJson() []string {
	days := []string{}
	for _, v := range Lyingflat {
		day := fmt.Sprintf("距离%v还有%v天", v.GetChineseName(), v.GetSubDay())
		days = append(days, day)
	}
	return days
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

func Gift() []string {
	src := "./users.txt"
	users := readLine(src)
	lines := make([]string, 0)
	for _, user := range users {
		lines = append(lines, fmt.Sprintf("帝豪全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("帝豪全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("豪庭全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("君怡全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("情思全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("蓝天全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("年代全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("5号馆全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("享乐汇全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("新海员全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("新太阳全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("水悦轩全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("金海湾全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("水玲珑全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("水疗会全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("水善坊全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("水云间全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("嘉联华全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("新境界全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("新名仕全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("龙湖阁全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("紫汕坊全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("新美龙全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("云顶天境全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("鹏发理疗全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("衡山水疗全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("龙湖夜总会全体妈咪预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("四季酒店全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
		lines = append(lines, fmt.Sprintf("八号公馆全体技师预祝大客户%v节日快乐,期待您的下次光临", user))
	}
	return lines
}
func readLine(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Printf("打开用户目录文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Printf("关闭用户目录文件失败: %s\n", err)
		}
	}()
	names := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		names = append(names, string(a))
		//log.Printf("读取到的用户(%s)\n", string(a))
	}
	return names
}
