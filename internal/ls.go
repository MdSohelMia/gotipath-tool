package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iafan/cwalk"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Ls(dir string) {
	defer TimeTrack(time.Now(), "File Count")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "File Name", "Type", "Size", "Permision"})
	var size int64
	count := 0

	err := cwalk.Walk(dir, func(path string, info os.FileInfo, err error) error {
		count++
		size += info.Size()
		sz := humanize.Bytes(uint64(info.Size()))
		var fileType string

		if info.IsDir() {
			fileType = "Directory"
		} else {
			fileType = "File"
		}

		t.AppendRow([]interface{}{count, info.Name(), fileType, sz, info.Mode()})
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	t.AppendFooter(table.Row{"", "", "", "", "", "Total File", count})
	t.Render()
}
