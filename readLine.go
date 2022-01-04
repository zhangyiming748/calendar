package calendar

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadLine(src string) []string {
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
