package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

var sec int
var cmd string

func init() {
	flag.IntVar(&sec, "time", 60, "Set timeout(sec) value.")
	flag.StringVar(&cmd, "cmd", "", "Set command or scripts.")
}

func execCmd(ch chan string) {

	args := []string{"/c", cmd}
	exe := exec.Command("cmd", args...)
	out, err := exe.CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		close(ch)
	} else {
		ch <- string(out)
	}
}

func setTimeOut() bool {
	flag.Parse()

	if cmd == "" {
		flag.Usage()
		return false
	}

	ch := make(chan string, 1)
	go execCmd(ch)

	select {
	case res, ok := <-ch:
		if ok != true {
			return false
		}
		fmt.Println(res)
		return true

	case <-time.After(time.Second * time.Duration(sec)):
		fmt.Println("Expired timeout.")
		return false
	}
}

func main() {
	setTimeOut()
}
