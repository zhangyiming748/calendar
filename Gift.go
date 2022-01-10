package calendar

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	words = []string{"年轻的时候千万不要因为没钱而绝望,因为你要知道,你以后没钱的日子还很多","铁打的身体,磁铁打的床"}
)

func Gift() {
	src := "./users.txt"
	if isExist(src){
		users := readLine(src)
		for _, user := range users {
			fmt.Printf("帝豪全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("豪庭全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("君怡全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("情思全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("蓝天全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("年代全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("5号馆全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("享乐汇全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("新海员全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("新太阳全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("水悦轩全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("金海湾全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("水玲珑全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("水疗会全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("水善坊全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("水云间全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("嘉联华全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("新境界全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("新名仕全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("龙湖阁全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("紫汕坊全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("新美龙全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("云顶天境全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("鹏发理疗全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("衡山水疗全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("龙湖夜总会全体妈咪预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("四季酒店全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
			fmt.Printf("八号公馆全体技师预祝大客户%v节日快乐,期待您的下次光临\n", user)
		}
	}
	rand.Seed(time.Now().Unix())
	fmt.Println(words[rand.Intn(len(words))])
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
func isExist(fname string)bool {
	_, err := os.Stat(fname)
	if err == nil{
		//fmt.Println("File exist")
		return true
	}
	if os.IsNotExist(err){
		//fmt.Println("File not exist")
		return false
	}
	return false
}