package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	_ "strings"
	"sync"
)

const (
	fileTxt = "file.txt"
	ts      = "ts"
)

func init() {
	//删除ts
	pwd, _ := os.Getwd()
	pwd = filepath.Join(pwd, ts)
	fmt.Println(pwd)
	os.RemoveAll(pwd)
	os.Mkdir(pwd, 0777)
}

func main() {
	txt, err := os.Create("file.txt")

	defer txt.Close()
	defer func() {
		error1 := recover()
		if error1 != nil {
			fmt.Println(error1)
		}
	}()
	if err != nil {
		fmt.Println("create file.txt error", err)
		panic(err)
	}

	var wait sync.WaitGroup

	var url string

	index := 0
	for {
		if index < 10 {
			url = "https://www.pipiaass.com/filets/4022/00"
		} else if index < 100 {
			url = "https://www.pipiaass.com/filets/4022/0"
		} else {
			url = "https://www.pipiaass.com/filets/4022/"
		}
		fileName := strconv.Itoa(index) + ".ts"
		url += fileName
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("get url error")
		}
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			break
		}

		wait.Add(1)
		go func(fileName string, resp *http.Response) {

			fileName = "ts/" + fileName
			out, err := os.Create(fileName)
			if err != nil {
				fmt.Println("create  many ts error")
			}
			defer out.Close()
			_, err1 := io.Copy(out, resp.Body)
			if err1 != nil {
				fmt.Println("copy error")
			}
			defer resp.Body.Close()
			wait.Done()
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println("copy gorotinue error")
					wait.Done()
				}
			}()
		}(fileName, resp)
		basePwd, _ := os.Getwd()
		basePwd = filepath.Join(basePwd, ts)
		fileContent := "file " + "'" + basePwd + "/" + strconv.Itoa(index) + ".ts'\n"
		txt.WriteString(fileContent)
		index++
	}
	wait.Wait()
}
