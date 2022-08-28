package downloader

import (
	"QQHQMusic/song"
	"QQHQMusic/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func getBaseDownloadFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "./"
	}

	dir := path.Join(homeDir, "Music/QQHDMusic/")
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		log.Println("can not make dir : " + dir)
		return "./"
	}

	return dir
}

func GetDownloadFolderForSingleSong() string {
	base := getBaseDownloadFolder()

	dir := path.Join(base, "单曲")
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		log.Println("can not make dir : " + dir)
		return base
	}

	return dir
}

func GetDownloadFolderForAllSongs() string {
	base := getBaseDownloadFolder()
	sub := utils.CurrentKeyWord + "-" + time.Now().Format("2006.01.02")
	dir := path.Join(base, sub)

	err := os.MkdirAll(dir, 0750)
	if err != nil {
		log.Println("can not make dir : " + dir)
		return base
	}

	return dir
}

var queue = make(chan int, 10)

func DownloadAll(list []*song.Song, toDir string) {

	for _, s := range list {
		queue <- 1
		go workItemOfDownloadAll(s, toDir)
	}

	for true {
		time.Sleep(time.Second)
		//if queue is empty, exit
		if len(queue) == 0 {
			log.Println(utils.GreenText("DOWNLOAD ALL DONE"))
			break
		}
	}
}

func workItemOfDownloadAll(s *song.Song, toDir string) {
	Download(s, toDir)
	<-queue
}

func Download(s *song.Song, toDir string) {

	fileName := s.Title + "-" + s.Singer + "." + s.BestFormat

	outputFile := path.Join(toDir, fileName)
	if _, err := os.Stat(outputFile); err == nil {
		log.Println("File exists, skip: ", outputFile)
		return
	}

	url := s.GetDownloadLink()
	log.Println("Downloading: ", s.Title+"-"+s.Singer)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("error:" + err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("close body error:" + err.Error())
		}
	}(resp.Body)

	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error:" + err.Error())
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("close file error:" + err.Error())
		}
	}(file)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println("copy file error:" + err.Error())
		return
	}

	if _, err := os.Stat(outputFile); err == nil {
		log.Println("已下载: ", outputFile)
		return
	}

}
