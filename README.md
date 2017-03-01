# SetTimeout
- This is the code that set timeout for external command, execute files or script, like unix command "timeout".
- Its code is for Windows DOS Prompt.
- It is written in Go.  

# Usage
## Command
`
SetTimeout.exe -time=<Time value(sec)> -cmd=<commands>
`

## Option
-cmd : Set command or scripts.  (string)
-time : Set timeout(sec) value. (int, default:60) 

## Example
`>SetTimeout.exe -time=90 -cmd="hoge.exe"`

You can use arguments in -cmd.  
`>SetTimeout.exe -time=90 -cmd="hoge.exe /b foo"`


