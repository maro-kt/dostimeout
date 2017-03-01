# SetTimeout for Windows DOS Prompt
- This is the code that set timeout for external command, execute files or script, like unix command "timeout".
- Its code is for Windows DOS Prompt.
- It is written in Go.  

# Usage
-cmd string  
        Set command or scripts.  
-time int  
        Set timeout(sec) value. (default 60)  

## example
`
go run main.go -time=60 -cmd="hello.exe"
`
