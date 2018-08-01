package main

import (
	"bufio"
	"fmt"
	"os"
)

func readfile(filename string) []string {
	fp, err := os.Open(filename)
	var ret []string
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func writefile(filename string, texts []string) {
	var writer *bufio.Writer

	file, _ := os.OpenFile(filename, os.O_WRONLY, 0644)
	writer = bufio.NewWriter(file)

	for idx := range texts {
		writer.WriteString(texts[idx] + "\n")
	}

	writer.Flush()
}

func main() {
	var filename = `GameUserSettings.ini`
	fmt.Println(filename)
	var ret = readfile(filename)

	for idx := range ret {
		if ret[idx] == "bDisableMouseAcceleration=False" {
			ret[idx] = "bDisableMouseAcceleration=True"
			fmt.Println(ret[idx])
		}
	}
}
