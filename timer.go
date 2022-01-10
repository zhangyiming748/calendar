package calendar

import (
	"github.com/zhangyiming748/calendar/v2/util/log"
	"strconv"
	"strings"
	"time"
)

var HappyWeekDayMap = map[string]string{
	"Monday":    "星期一,还有四天周末",
	"Tuesday":   "星期二,还有三天周末",
	"Wednesday": "星期三,工作日过去一半了",
	"Thursday":  "肯德基疯狂星期四",
	"Friday":    "周五就应该有周五的态度,有什么事下周一再说",
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

func HappyTimer() {
	nowDay := time.Now().Format("今天是2006年1月2日")
	nowTime := time.Now().Format("PM 3点04分05秒")
	week := time.Now().Weekday().String()
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	nowTime = strings.Replace(nowTime, "AM", "上午", -1)
	nowTime = strings.Replace(nowTime, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		log.Debug.Println("早上好,摸鱼人!")
	}
	if inthour == 12 {
		log.Debug.Println("中午好,摸鱼人!")
	}
	if inthour >= 13 && inthour <= 17 {
		log.Debug.Println("下午好,摸鱼人!")
	}
	if inthour >= 18 && inthour <= 22 {
		log.Debug.Println("晚上好,摸鱼人!")
	}
	if inthour >= 23 && inthour <= 5 {
		log.Debug.Println("晚上抓紧睡觉,休息好,第二天才有精神摸鱼!")
	}
	log.Debug.Println(nowDay, HappyWeekDayMap[week], nowTime)
	log.Debug.Println("周一不摸鱼,脑子有问题")
	log.Debug.Println("周二不摸鱼,上班没心情")
	log.Debug.Println("周三不摸鱼,你是头傻驴")
	log.Debug.Println("周四不摸鱼,谁请肯德基")
	log.Debug.Println("周五不摸鱼,浪费好心情")
	log.Debug.Println("周六不摸鱼,活着真多余")

	if inthour == 15 {
		log.Debug.Println("喂!三点几咧!做......做撚啊做!")
		log.Debug.Println("饮茶先啊!")
		log.Debug.Println("三点几 饮......饮茶先啊!")
		log.Debug.Println("做咁多都冇用嘅,老细唔锡你嘅咧!")
		log.Debug.Println("喂!饮下茶先啊!三点几咧!")
		log.Debug.Println("做碌鸠啊做!")
	}
	if inthour == 18 {
		log.Debug.Println("喂!朋友!")
		log.Debug.Println("做乜嘢咁多啦?差唔多七点咧!")
		log.Debug.Println("放工啦!唔使做咁多啦!")
		log.Debug.Println("做咁多,钱带去边度?")
		log.Debug.Println("差唔多七点咧!放工!")
		log.Debug.Println("落去茶室,饮下靓靓嘅杯......麦啤酒、黑啤酒,OK?")
		log.Debug.Println("Happy下,唔使做咁多.")
		log.Debug.Println("死咗都冇用咧,银纸冇得带去咧!")
		log.Debug.Println("Happy下,饮酒,OK?")
	}
	log.Debug.Println("【摸鱼办】提醒您")
	log.Debug.Println("工作再累,一定不要忘记摸鱼哦!")
	log.Debug.Println("有事没事起身去茶水间,去厕所,去廊道走走")
	log.Debug.Println("吃饭时间就吃饭")
	log.Debug.Println("午休时间就午休")
	log.Debug.Println("别老在工位上坐着,钱是老板的,但命是自己的")
}
func SadTimer() {
	nowDay := time.Now().Format("今天是2006年1月2日")
	nowTime := time.Now().Format("PM3点04分05秒")
	week := time.Now().Weekday().String()
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	nowTime = strings.Replace(nowTime, "AM", "上午", -1)
	nowTime = strings.Replace(nowTime, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		log.Debug.Println("早上好,开始内卷的一天!")
	}
	if inthour == 12 {
		log.Debug.Println("中午好,别人休息的时候正是你内卷的好时候!")
	}
	if inthour >= 13 && inthour <= 17 {
		log.Debug.Println("下午好,继续卷!")
	}
	if inthour >= 18 && inthour <= 22 {
		log.Debug.Println("晚上好,坚持住!")
	}
	if inthour >= 23 && inthour <= 5 {
		log.Debug.Println("晚上抓紧卷,不睡觉你才能获得比别人更多的时间!")
	}
	log.Debug.Println(nowDay, nowTime, SadWeekDayMap[week])
	log.Debug.Println("【内卷办】提醒您")
}
