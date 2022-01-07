package calendar

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestUnit(t *testing.T) {
	thisYear()
	//chineseNewYear()
}
func TestMaster(t *testing.T) {
	Calendar()
}
func BenchmarkHeart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calendar()
	}
}
func TestHappy(t *testing.T) {
	HappyDay()
}
func TestSad(t *testing.T) {
	SadDay()
}
func TestNextYear(t *testing.T) {
	ret := nextYear()
	t.Log(ret)
}
func TestStructs(t *testing.T) {
	var NewYear Festival
	NewYear.SetChineseName("元旦")
	t.Log(NewYear.GetChineseName())
	NewYear.SetDate("01-01")
	t.Log(NewYear.GetChineseName())
	NewYear.SetSubDay(allInSolar(NewYear.GetDate()))
	fmt.Println(NewYear.GetDate())
}

//英语音标
func TestUniCode(t *testing.T) {
	fmt.Println("\u0250")
	fmt.Println("\u0251")
	fmt.Println("\u0252")
	fmt.Println("\u0253")
	fmt.Println("\u0254")
	fmt.Println("\u0255")
	fmt.Println("\u0256")
	fmt.Println("\u0257")
	fmt.Println("\u0258")
	fmt.Println("\u0259")
	fmt.Println("\u025A")
	fmt.Println("\u025B")
	fmt.Println("\u025C")
	fmt.Println("\u025D")
	fmt.Println("\u025E")
	fmt.Println("\u025F")
	fmt.Println("\u0260")
	fmt.Println("\u0261")
	fmt.Println("\u0262")
	fmt.Println("\u0263")
	fmt.Println("\u0264")
	fmt.Println("\u0265")
	fmt.Println("\u0266")
	fmt.Println("\u0267")
	fmt.Println("\u0268")
	fmt.Println("\u0269")
	fmt.Println("\u026A")
	fmt.Println("\u026B")
	fmt.Println("\u026C")
	fmt.Println("\u026D")
	fmt.Println("\u026E")
	fmt.Println("\u026F")
	fmt.Println("\u0270")
	fmt.Println("\u0271")
	fmt.Println("\u0272")
	fmt.Println("\u0273")
	fmt.Println("\u0274")
	fmt.Println("\u0275")
	fmt.Println("\u0276")
	fmt.Println("\u0277")
	fmt.Println("\u0278")
	fmt.Println("\u0279")
	fmt.Println("\u027A")
	fmt.Println("\u027B")
	fmt.Println("\u027C")
	fmt.Println("\u027D")
	fmt.Println("\u027E")
	fmt.Println("\u027F")
	fmt.Println("\u0280")
	fmt.Println("\u0281")
	fmt.Println("\u0282")
	fmt.Println("\u0283")
	fmt.Println("\u0284")
	fmt.Println("\u0285")
	fmt.Println("\u0286")
	fmt.Println("\u0287")
	fmt.Println("\u0288")
	fmt.Println("\u0289")
	fmt.Println("\u028A")
	fmt.Println("\u028B")
	fmt.Println("\u028C")
	fmt.Println("\u028D")
	fmt.Println("\u028E")
	fmt.Println("\u028F")
	fmt.Println("\u0290")
	fmt.Println("\u0291")
	fmt.Println("\u0292")
	fmt.Println("\u0293")
	fmt.Println("\u0294")
	fmt.Println("\u0295")
	fmt.Println("\u0296")
	fmt.Println("\u0297")
	fmt.Println("\u0298")
	fmt.Println("\u0299")
	fmt.Println("\u029A")
	fmt.Println("\u029B")
	fmt.Println("\u029C")
	fmt.Println("\u029D")
	fmt.Println("\u029E")
	fmt.Println("\u029F")
	fmt.Println("\u02A0")
	fmt.Println("\u02A1")
	fmt.Println("\u02A2")
	fmt.Println("\u02A3")
	fmt.Println("\u02A4")
	fmt.Println("\u02A5")
	fmt.Println("\u02A6")
	fmt.Println("\u02A7")
	fmt.Println("\u02A8")
	fmt.Println("\u02A9")
	fmt.Println("\u02AA")
	fmt.Println("\u02AB")
	fmt.Println("\u02AC")
	fmt.Println("\u02AD")
	fmt.Println("\u02AE")
	fmt.Println("\u02AF")

}

func seed() {
	rand.Seed(32)
	fmt.Println("第一轮")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %v", rand.Intn(10))
	}
	fmt.Println("第一轮")
	rand.Seed(64)
	fmt.Println("第二轮")
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10))
	}
	fmt.Println("第二轮")

	rand.Seed(128)
	fmt.Println("第三轮")
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10))
	}
	fmt.Println("第三轮")

}
func BenchmarkSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		{
			seed()
		}
	}
}

func TestGift(t *testing.T) {
	Gift()
}
