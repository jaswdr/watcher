package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
)

func Usage() {
	usage := `Usage:
  watcher <command> <files>

Example:
	watcher ls .
	`

	fmt.Println(usage)
}

func main() {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	files := os.Args[2:]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					cmd := exec.Command("/bin/sh", "-c", os.Args[1])
					out, _ := cmd.CombinedOutput()
					fmt.Println(string(out))
				}
			case err := <-watcher.Errors:
				fmt.Println(err)
			}
		}
	}()

	for _, file := range files {
		err := watcher.Add(file)
		if err != nil {
			fmt.Println(err)
		}
	}

	<-done
}
