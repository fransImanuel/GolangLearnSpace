package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

type FileInfo struct {
	FilePath  string // file location
	Content   []byte // file content
	Sum       string // md5 sum of content
	IsRenamed bool   // indicator whether the particular file is renamed already or not
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

			//if there is an error, return immidiately
			if err != nil {
				return err
			}

			//if it is a sub directory, return immidiately
			if info.IsDir() {
				return nil
			}

			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil

		})
		if err != nil {
			log.Println("ERROR:", err.Error())
		}

		close(chanOut)
	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}

		close(chanOut)
	}()

	return chanOut
}

func main() {
	log.Println("start")
	start := time.Now()

	//pipeline 1: loop all files and read it
	chanFileContent := readFiles()

	//pipeline 2: calculate md5sum
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	//pipeline 3:rename files
	chanRenam1 := rename(chanFileSum)
	chanRenam2 := rename(chanFileSum)
	chanRenam3 := rename(chanFileSum)
	chanRenam4 := rename(chanFileSum)
	chanRenam := mergeChanFileInfo(chanRenam1, chanRenam2, chanRenam3, chanRenam4)

	//print output
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRenam {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
