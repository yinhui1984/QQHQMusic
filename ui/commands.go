package ui

import (
	"QQHQMusic/downloader"
	"QQHQMusic/song"
	"fmt"
	"strings"
)

// Commander 识别传入到Execute方法中的参数,
// 如果参数符合命令格式, 则干活
// 否则啥也不干
type Commander interface {
	Execute(str string, songList []*song.Song) bool
}

type PlayCommand struct {
}

func (play PlayCommand) Execute(str string, songList []*song.Song) bool {

	var (
		cmd   string
		index int
	)
	reader := strings.NewReader(str)
	_, err := fmt.Fscanf(reader, "%s %d", &cmd, &index)
	if err != nil {
		return false
	}

	if "p" == cmd || "play" == cmd {
		if index < 0 || index >= len(songList) {
			fmt.Println("id 超范围")
			return false
		}

		s := songList[index]
		s.Play()
	}

	return false
}

type DownloadCommand struct {
}

func (down DownloadCommand) Execute(str string, songList []*song.Song) bool {

	var (
		cmd   string
		index int
	)
	reader := strings.NewReader(str)
	_, err := fmt.Fscanf(reader, "%s %d", &cmd, &index)
	if err != nil {
		return false
	}

	if "d" == cmd || "download" == cmd {
		if index < 0 || index >= len(songList) {
			fmt.Println("id 超范围")
			return false
		}

		s := songList[index]
		downloader.Download(s, downloader.GetDownloadFolderForSingleSong())
	}

	return false
}

type DownloadAllCommand struct {
}

func (downAll DownloadAllCommand) Execute(str string, songList []*song.Song) bool {

	if "da" == str {
		downloader.DownloadAll(songList, downloader.GetDownloadFolderForAllSongs())
	}

	return false
}
