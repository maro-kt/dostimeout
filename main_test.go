package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestTimeoutExpired(t *testing.T) {
	sec = 10
	cmd = "ping localhost -n 15 && hello "
	result := setTimeOut()

	if result == true {
		t.Errorf("Timeout did not occured. Result is %v", result)
	}
}

func TestSuccessful(t *testing.T) {
	sec = 10
	cmd = "hello "
	result := setTimeOut()

	if result == false {
		t.Errorf("Timeout occured. Result is %v", result)
	}
}

func TestCommandWithSpace(t *testing.T) {
	sec = 10
	cmd = filepath.Join("hoge bat", "hoge.bat")
	result := setTimeOut()
	if result == false {
		t.Errorf("Timeout occured. Result is %v", result)
	}
}

func TestUnknonwCommand(t *testing.T) {
	sec = 60
	cmd = "aaaa"
	result := setTimeOut()
	if result == false {
		fmt.Println("TestUnknonwCommand is OK")
	} else {
		t.Errorf("Test fails: %v", result)
	}
}
