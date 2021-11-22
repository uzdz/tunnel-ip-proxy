package utils

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func Output(cmd *exec.Cmd, c Charset) {
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	stderr, _ := cmd.StderrPipe()
	//stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Println("exec the cmd failed")
	}

	// 正常日志
	//logScan := bufio.NewScanner(stdout)
	//for logScan.Scan() {
	//	log.Println("info - " + ConvertByte2String(logScan.Bytes(), c))
	//}

	// 错误日志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		s := scan.Text()
		log.Println("build error: ", s)
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}
	// 等待命令执行完
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		// 执行失败，返回错误信息
		log.Println(errBuf.String() + "")
	}
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
