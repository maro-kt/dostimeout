# SetTimeout for Windows DOS Prompt
- This is the code that set timeout for external command, execute files or script, like unix command "timeout".
- Its code is for Windows DOS Prompt.
- It is written in Go.  

# Usage
## Command
`
SetTimeout -time=< Time value(sec) > -cmd=<commands>
`

## Option
-cmd : Set command or scripts.  (string)
-time : Set timeout(sec) value. (int, default:60)  

## example
`
go run main.go -time=60 -cmd="hello.exe"
`
You can execute the external code with argments.
`
go run main.go -time=60 -cmd="dir * /b"
`
