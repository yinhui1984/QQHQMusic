package ui

import (
	"QQHQMusic/searcher"
	"QQHQMusic/song"
	"QQHQMusic/utils"
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strings"
)

func AskForKeyword() {
	for true {
		utils.ClearScreen()
		fmt.Println(utils.GreenText("输入关键词就行搜索: "))
		fmt.Print(utils.RedText("=> "))
		reader := bufio.NewReader(os.Stdin)
		keyWord, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		keyWord = strings.TrimSpace(keyWord)

		songList := searcher.Search(keyWord)
		ShowSongList(songList)
		AskForCommand(songList)
	}
}

func ShowSongList(songList []*song.Song) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Singer", "Format", "Size(M)", "Album"})

	for i, s := range songList {

		title := s.Title
		if len(s.Title) > 40 {
			title = s.Title[0:37] + "..."
		}

		album := s.Album
		if len(s.Album) > 40 {
			album = s.Album[0:37] + "..."
		}

		singer := s.Singer
		if len(singer) > 25 {
			singer = s.Singer[0:22] + "..."
		}

		t.AppendRow(table.Row{
			i, title, singer, s.BestFormat, int(s.BestSize / 1024 / 1024), "《" + album + "》",
		})

	}

	t.SetStyle(table.StyleBold)
	colorBOnW := text.Colors{text.BgWhite, text.FgBlack}
	// set colors using Colors/ColorsHeader/ColorsFooter
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Colors: text.Colors{text.FgYellow}, ColorsHeader: colorBOnW},
		{Number: 2, Colors: text.Colors{text.FgHiRed}, ColorsHeader: colorBOnW},
		{Number: 3, Colors: text.Colors{text.FgHiWhite}, ColorsHeader: colorBOnW, ColorsFooter: colorBOnW},
		{Number: 4, Colors: text.Colors{text.FgGreen}, ColorsHeader: colorBOnW, ColorsFooter: colorBOnW},
		{Number: 5, Colors: text.Colors{text.FgCyan}, ColorsHeader: colorBOnW},
	})

	t.Render()

	////大多数情况下, 显示列表后第一件事情就是试听第一首, 不如直接播放
	//if len(songList) > 0 {
	//	songList[0].Play()
	//}
}

func AskForCommand(songList []*song.Song) {
	fmt.Println("\n\n选择和执行命令")
	fmt.Println("    比如播放第1首:p 1")
	fmt.Println("    比如下载第1首:d 1")
	fmt.Println("    比如下载全部:da")
	fmt.Println("    退出循环:q")

	for true {
		fmt.Print(utils.RedText("\n=> "))

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			break
		}

		PlayCommand{}.Execute(input, songList)
		DownloadCommand{}.Execute(input, songList)
		DownloadAllCommand{}.Execute(input, songList)

	}
}
