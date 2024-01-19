package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 定义一个新的字符串参数，name为"f", 默认值为"domain.txt", 使用方法说明为"指定文件"
	filePath := flag.String("f", "domain.txt", "指定文件")

	// 使用flag参数
	err := RemoveDuplicates(*filePath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Run successfully.")
	}
}

func RemoveDuplicates(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	seen := map[string]bool{}
	scanner := bufio.NewScanner(file)
	outfile, err := os.Create("outfile.txt")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outfile)

	for scanner.Scan() {
		line := scanner.Text()
		// 检查是否为空行
		if strings.TrimSpace(line) != "" && !seen[line] {
			seen[line] = true
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	err = outfile.Close()
	if err != nil {
		return err
	}

	// 显式地关闭旧文件
	err = file.Close()
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	err = os.Rename("outfile.txt", filePath)

	return err
}
