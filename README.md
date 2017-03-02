# setTimerCmd

- This is the code that set timeout for external command, execute files or script, like unix command "timeout".
- Its code is for Windows DOS Prompt.
- It is written in Go.
- Stdout/Stderr in Japanese, Chinese and Korean is not Supported.

## Usage

### Install

`$go get -u github.com/maro-kt/setTimerCmd`

### Command

`setTimerCmd.exe -time=<Time value(sec)> -cmd=<commands>`

### Option

| Option | Type | Description |
| ---- | ---- | ---- |
| -time | int | Timeout value.(unit:second, default:60)  |
| -cmd | string | Command or scripts path.|


### Example

Basic usages.  
`>setTimerCmd.exe -time=90 -cmd="hoge.exe"`

You can use arguments in -cmd.  
`>setTimerCmd.exe -time=90 -cmd="hoge.exe /b foo"`
