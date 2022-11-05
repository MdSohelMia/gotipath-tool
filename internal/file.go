package internal

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dustin/go-humanize"
	"github.com/iafan/cwalk"
)

var wg sync.WaitGroup

var (
	concurrent    = runtime.GOMAXPROCS(runtime.NumCPU())
	semaphoreChan = make(chan struct{}, concurrent)
)

type File struct {
	size  string
	total int
}

func NewFile() *File {
	return &File{
		size:  "0Kb",
		total: 0,
	}
}

func LsFiles(dir string) {
	// block while full
	semaphoreChan <- struct{}{}

	go func() {
		defer func() {
			// read to release a slot
			<-semaphoreChan
			wg.Done()
		}()

		file, err := os.Open(dir)
		if err != nil {
			fmt.Println("error opening directory", err.Error())
		}
		defer file.Close()
		files, err := file.Readdir(-1) // Loads all children files into memory. A more efficient way?
		if err != nil {
			fmt.Println("error reading directory", err.Error())
		}
		for _, f := range files {
			fmt.Println(dir + "/" + f.Name())
			if f.IsDir() {
				wg.Add(1)
				go LsFiles(dir + "/" + f.Name())
			}
		}
	}()
}

func Walk(dir string) {
	defer TimeTrack(time.Now(), "File Count")
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)

	s.Start()
	s.Prefix = "File Calculating: "
	var size int64
	count := 0

	err := cwalk.Walk(dir, func(path string, info os.FileInfo, err error) error {
		count++
		size += info.Size()
		s.UpdateCharSet(spinner.CharSets[35])
		return nil
	})
	s.Stop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("================================================")
	fmt.Println("Total File Size :", humanize.Bytes(uint64(size)))
	fmt.Println("================================================")
	fmt.Println("Total File count:", count)
	fmt.Println("================================================")
}
