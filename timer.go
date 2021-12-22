package calendar

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HappyTimer() {
	now := time.Now().Format("今天是2006年1月2日\nPM3点04分05秒")
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	now = strings.Replace(now, "AM", "上午", -1)
	now = strings.Replace(now, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		fmt.Println("早上好,摸鱼人!")
	}
	if inthour == 12 {
		fmt.Println("中午好,摸鱼人!")
	}
	if inthour >= 13 && inthour <= 17 {
		fmt.Println("下午好,摸鱼人!")
	}
	if inthour >= 18 && inthour <= 22 {
		fmt.Println("晚上好,摸鱼人!")
	}
	if inthour >= 23 && inthour <= 5 {
		fmt.Println("晚上抓紧睡觉,休息好,第二天才有精神摸鱼!")
	}
	fmt.Println(now)
	if inthour == 15 {
		fmt.Println("喂!三点几咧!做......做撚啊做!")
		fmt.Println("饮茶先啊!")
		fmt.Println("三点几 饮......饮茶先啊!")
		fmt.Println("做咁多都冇用嘅,老细唔锡你嘅咧!")
		fmt.Println("喂!饮下茶先啊!三点几咧!")
		fmt.Println("做碌鸠啊做!")
	}
	if inthour == 18 {
		fmt.Println("喂!朋友!")
		fmt.Println("做乜嘢咁多啦?差唔多七点咧!")
		fmt.Println("放工啦!唔使做咁多啦!")
		fmt.Println("做咁多,钱带去边度?")
		fmt.Println("差唔多七点咧!放工!")
		fmt.Println("落去茶室,饮下靓靓嘅杯......麦啤酒、黑啤酒,OK?")
		fmt.Println("Happy下,唔使做咁多.")
		fmt.Println("死咗都冇用咧,银纸冇得带去咧!")
		fmt.Println("Happy下,饮酒,OK?")
	}
	fmt.Println("【摸鱼办】提醒您")
	fmt.Println("工作再累,一定不要忘记摸鱼哦!")
	fmt.Println("有事没事起身去茶水间,去厕所,去廊道走走")
	fmt.Println("吃饭时间就吃饭")
	fmt.Println("午休时间就午休")
	fmt.Println("别老在工位上坐着,钱是老板的,但命是自己的")
}
func SadTimer() {
	now := time.Now().Format("今天是2006年1月2日\nPM3点04分05秒")
	hour := time.Now().Format("15")
	inthour, _ := strconv.Atoi(hour)
	now = strings.Replace(now, "AM", "上午", -1)
	now = strings.Replace(now, "PM", "下午", -1)
	if inthour >= 6 && inthour <= 11 {
		fmt.Println("早上好,开始内卷的一天!")
	}
	if inthour == 12 {
		fmt.Println("中午好,别人休息的时候正是你内卷的好时候!")
	}
	if inthour >= 13 && inthour <= 17 {
		fmt.Println("下午好,继续卷!")
	}
	if inthour >= 18 && inthour <= 22 {
		fmt.Println("晚上好,坚持住!")
	}
	if inthour >= 23 && inthour <= 5 {
		fmt.Println("晚上抓紧卷,不睡觉你才能获得比别人更多的时间!")
	}
	fmt.Println(now)
	fmt.Println("【内卷办】提醒您")
}
