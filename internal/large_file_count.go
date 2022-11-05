package internal

import (
	"fmt"
	"os"
	"time"
)

func LargeFileCount(path string) {
	defer TimeTrack(time.Now(), "file count.")
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Errors:", err.Error())
	}
	println("Total File:", len(dir))
}
