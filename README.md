# dostimeout
This is the program that set timeout for external command or script, like unix command "timeout".  　
This program is written in Go.  
# Usage
go run main.go -time=60 -cmd="hello.exe"

-- cmd string  
        Set command or scripts.  
-- time int  
        Set timeout(sec) value. (default 60)  
