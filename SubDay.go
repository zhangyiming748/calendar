package calendar

import (
	"fmt"
	"time"
)

const (
	NewYear             = "2022-01-01" //新年
	ValentinesDay       = "2021-02-14" //情人节
	WomensDay           = "2021-03-08" //妇女节
	ArborDay            = "2021-03-12" //植树节
	AprilFoolsDay       = "2021-04-01" //愚人节
	LaborDay            = "2021-05-01" //劳动节
	YouthDay            = "2021-05-04" //青年节
	ChildrensDay        = "2021-06-01" //儿童节
	NationalDay         = "2021-10-01" //国庆节
	Halloween           = "2021-10-31" //万圣夜
	ThanksgivingDay     = "2021-11-25" //感恩节
	NationalMemorialDay = "2021-12-13" //国家公祭日
	ChristmasEve        = "2021-12-24" //平安夜
	Christmas           = "2021-12-25" //圣诞节
)

//获取当前日期
func SubDay() {
	fmt.Println(time.Now().Format("早上好，摸鱼人!今天是2006年01月02日"))
	fmt.Println("【摸鱼办】提醒您")
	fmt.Println("工作再累,一定不要忘记摸鱼哦!")
	fmt.Println("有事没事起身去茶水间,去厕所,去廊道走走")
	fmt.Println("别老在工位上坐着,钱是老板的,但命是自己的")
	defer nextNewYear()
	defer func() {
		fmt.Println("上班是帮老板赚钱,摸鱼是赚老板的钱!")
		fmt.Println("最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天")
	}()
	subValentinesDay()
	subWomensDay()
	subArborDay()
	subAprilFoolsDay()
	subLaborDay()
	subYouthDay()
	subChildrensDay()
	subNationalDay()
	subHalloween()
	subThanksgivingDay()
	subNationalMemorialDay()
	subChristmasEve()
	subChristmas()
}

//情人节
func subValentinesDay() {
	ret, _ := time.Parse("2006-01-02", ValentinesDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离情人节还有%v天\n", int(unsub.Hours())/24)
}

//妇女节
func subWomensDay() {
	ret, _ := time.Parse("2006-01-02", WomensDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离妇女节还有%v天\n", int(unsub.Hours())/24)
}

//植树节
func subArborDay() {
	ret, _ := time.Parse("2006-01-02", ArborDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离植树节还有%v天\n", int(unsub.Hours())/24)
}

//愚人节
func subAprilFoolsDay() {
	ret, _ := time.Parse("2006-01-02", AprilFoolsDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离愚人节还有%v天\n", int(unsub.Hours())/24)
}

//劳动节
func subLaborDay() {
	ret, _ := time.Parse("2006-01-02", LaborDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离劳动节还有%v天\n", int(unsub.Hours())/24)
}

//青年节
func subYouthDay() {
	ret, _ := time.Parse("2006-01-02", YouthDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离青年节还有%v天\n", int(unsub.Hours())/24)
}

//儿童节
func subChildrensDay() {
	ret, _ := time.Parse("2006-01-02", ChildrensDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离儿童节还有%v天\n", int(unsub.Hours())/24)
}

//国庆节
func subNationalDay() {
	ret, _ := time.Parse("2006-01-02", NationalDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国庆节还有%v天\n", int(unsub.Hours())/24)
}

//感恩节
func subThanksgivingDay() {
	ret, _ := time.Parse("2006-01-02", ThanksgivingDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离感恩节还有%v天\n", int(unsub.Hours())/24)
}

//国家公祭日
func subNationalMemorialDay() {
	ret, _ := time.Parse("2006-01-02", NationalMemorialDay)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国家公祭日还有%v天\n", int(unsub.Hours())/24)
}

//万圣夜
func subHalloween() {
	ret, _ := time.Parse("2006-01-02", Halloween)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离万圣夜还有%v天\n", int(unsub.Hours())/24)
}

//平安夜
func subChristmasEve() {
	ret, _ := time.Parse("2006-01-02", ChristmasEve)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离平安夜还有%v天\n", int(unsub.Hours())/24)
}

//圣诞节
func subChristmas() {
	ret, _ := time.Parse("2006-01-02", Christmas)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离圣诞节还有%v天\n", int(unsub.Hours())/24)
}

//计算和元旦的差值
func nextNewYear() {
	ret, _ := time.Parse("2006-01-02", NewYear)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离下一年还有%v天\n", int(unsub.Hours())/24)
}