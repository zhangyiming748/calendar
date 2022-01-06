package calendar

import (
	"github.com/nosixtools/solarlunar"
	"testing"
)

func TestUnit(t *testing.T) {
	thisYear()
	Calendar()
	HappyDay()
	SadDay()

}
func TestUnits(t *testing.T) {
	day := solarlunar.LunarToSolar("2022-01-01", false)
	t.Logf("day = %v", day)

}
func TestMaster(t *testing.T) {

}
func BenchmarkHeart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calendar()
	}
}
func TestWeb(t *testing.T) {
	ShowWeb()
}
