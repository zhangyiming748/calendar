package calendar

import "testing"

func TestUnit(t *testing.T) {
	thisYear()
	chineseNewYear()
}
func TestMaster(t *testing.T) {
	SubDay()
}
func TestNextYear(t *testing.T) {
	ret:=nextYear()
	t.Log(ret)
}

