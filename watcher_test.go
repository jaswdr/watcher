package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestUsageIsNotEmpty(t *testing.T) {
	usage := Usage()
	if len(usage) == 0 {
		t.Error("Fail expecting that usage was not empty, but it was")
	}
}

func TestWriteSingleWatchedFile(t *testing.T) {
	tmp1, _ := ioutil.TempFile("/tmp", "watcher_")
	tmp2, _ := ioutil.TempFile("/tmp", "watcher_")
	cmd := fmt.Sprintf("echo 'bar' > %s", tmp2.Name())

	t.Logf("Watching %s", tmp1.Name())
	go Watch(cmd, []string{tmp1.Name()})
	time.Sleep(1 * time.Second)

	t.Logf("Writing to %s", tmp1.Name())
	ioutil.WriteFile(tmp1.Name(), []byte("foo"), 0666)
	time.Sleep(1 * time.Second)

	t.Logf("Reading to %s", tmp2.Name())
	content, _ := ioutil.ReadFile(tmp2.Name())

	cmp := strings.TrimSpace(string(content))

	if cmp != "bar" {
		t.Errorf("Fail when watching a single file, expected 'bar' found '%s'", content)
	}
}

func TestWatchedDir(t *testing.T) {
	tmp1, _ := ioutil.TempFile("/tmp", "watcher_")
	tmp2, _ := ioutil.TempFile("/tmp", "watcher_")
	cmd := fmt.Sprintf("echo 'bar' > %s", tmp2.Name())

	t.Logf("Watching %s", tmp1.Name())
	go Watch(cmd, []string{"/tmp"})
	time.Sleep(1 * time.Second)

	t.Logf("Writing to %s", tmp1.Name())
	ioutil.WriteFile(tmp1.Name(), []byte("foo"), 0666)
	time.Sleep(1 * time.Second)

	t.Logf("Reading to %s", tmp2.Name())
	content, _ := ioutil.ReadFile(tmp2.Name())

	cmp := strings.TrimSpace(string(content))

	if cmp != "bar" {
		t.Errorf("Fail when watching a single created file, expected 'bar' found '%s'", content)
	}
}

func TestCreationInsideWatchedDir(t *testing.T) {
	tmp1, _ := ioutil.TempFile("/tmp", "watcher_")
	tmp2, _ := ioutil.TempFile("/tmp", "watcher_")
	cmd := fmt.Sprintf("echo 'bar' > %s", tmp2.Name())

	os.Remove(tmp1.Name())

	t.Logf("Watching %s", tmp1.Name())
	go Watch(cmd, []string{"/tmp"})
	time.Sleep(1 * time.Second)

	t.Logf("Writing to %s", tmp1.Name())
	ioutil.WriteFile(tmp1.Name(), []byte("foo"), 0666)
	time.Sleep(1 * time.Second)

	t.Logf("Reading to %s", tmp2.Name())
	content, _ := ioutil.ReadFile(tmp2.Name())

	cmp := strings.TrimSpace(string(content))

	if cmp != "bar" {
		t.Errorf("Fail when watching a single created file, expected 'bar' found '%s'", content)
	}
}
