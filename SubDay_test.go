package calendar

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	thisYear()
	//chineseNewYear()
}
func TestMaster(t *testing.T) {
	SubDay()
}
func TestNextYear(t *testing.T) {
	ret:=nextYear()
	t.Log(ret)
}
func TestStructs(t *testing.T) {
	//NewYear := Festival{}
	//NewYear:=new(Festival)
	var NewYear Festival
	NewYear.SetChineseName("元旦")
	t.Log(NewYear.GetChineseName())
	NewYear.SetDate("01-01")
	t.Log(NewYear.GetChineseName())
	NewYear.SetSubDay(allInOne(NewYear.GetDate()))
	fmt.Println(NewYear.GetDate())

}
