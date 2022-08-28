package downloader

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strings"
)

/*
这个文件用于显示下载进度
*/

// WriteCounter 计算写到它的字节数。它实现了io.Writer接口
// 我们可以将其传递给io.TeeReader()，它将报告每个写入周期的进度。
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	// 通过使用回车来清空该行，回到起点并删除  剩余的字符，用空格填充
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// 再次返回并打印当前的下载状态
	// 我们使用humanize包，以有意义的方式打印字节数（例如10MB）。
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
