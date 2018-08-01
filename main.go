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
		writer.WriteString(texts[idx] + "\r\n")
	}

	writer.Flush()
}

func main() {
	var filename = os.Getenv("LOCALAPPDATA") + `/FortniteGame/Saved/Config/WindowsClient/GameUserSettings.ini`
	fmt.Println(os.Getenv("LOCALAPPDATA"))
	var ret = readfile(filename)

	for idx := range ret {
		if ret[idx] == "bDisableMouseAcceleration=False" {
			fmt.Println("exist bDisableMouseAcceleration=False :", idx)
			ret[idx] = "bDisableMouseAcceleration=True"
			fmt.Println("replace bDisableMouseAcceleration=True")
		}
	}
	writefile(filename, ret)
}
