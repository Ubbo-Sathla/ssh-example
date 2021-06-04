package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func handlerErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)
	}
}
func main() {
	//ssh 服务地址home.mojotv.cn:22
	client, err := ssh.Dial("tcp", "10.127.253.187:22", &ssh.ClientConfig{
		User:            "root",                                         //ssh 用户名
		Auth:            []ssh.AuthMethod{ssh.Password("KPFpu5zhB3Qt")}, //ssh 密码
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	handlerErr(err, "dial")
	session, err := client.NewSession()
	handlerErr(err, "ssh session 创建")
	defer session.Close()

	//当前机器的terminal 连接到ssh-session的为终端
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	// 配置pty
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	//这里设置的固定的terminal size 25x100
	//当terminal 窗口尺寸改变的时候 会导致终端显示错位
	err = session.RequestPty("xterm", 25, 80, modes)
	handlerErr(err, "请求PTY为终端")

	err = session.Shell()
	handlerErr(err, "开始 shell")

	err = session.Wait()
	handlerErr(err, "执行完毕")
}
