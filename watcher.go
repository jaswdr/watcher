package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Usage return a usage message showing the user how to use this package
func Usage() string {
	return `Usage:
  watcher <command> <files>

Example:
	watcher ls .`
}

// Watch start watching a list of directories
func Watch(command string, directories []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	executionId := 1
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("--- Execution %d (%s) ---\n", executionId, time.Now())
					executionId += 1
					cmd := exec.Command("/bin/sh", "-c", command)

					stdout, _ := cmd.StdoutPipe()
					stderr, _ := cmd.StderrPipe()
					cmd.Start()

					errScanner := bufio.NewScanner(stderr)
					for errScanner.Scan() {
						fmt.Println(errScanner.Text())
					}

					outScanner := bufio.NewScanner(stdout)
					for outScanner.Scan() {
						fmt.Println(outScanner.Text())
					}

					cmd.Wait()
				}
			case err := <-watcher.Errors:
				fmt.Println(err)
			}
		}
	}()

	for _, dir := range directories {
		go func(dir string) {
			watcher.Add(dir)
		}(dir)
	}

	<-done
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage())
		os.Exit(1)
	}

	Watch(os.Args[1], os.Args[2:])
}
