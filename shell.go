package util

import (
	"os/exec"
	"runtime"
)

// ExecShell 阻塞式的执行外部shell命令的函数,等待执行完毕并返回所有输出
func ExecShell(s string) (out []byte, err error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("/bin/sh", "-c", s)
		out, err = cmd.CombinedOutput()
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", s)
		out, err = cmd.CombinedOutput()
	}
	return out, err
}
