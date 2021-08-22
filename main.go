package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
)

// cspell:words Centos

//go:embed rpm

var f embed.FS

func main() {
	version := getCentosVersion("major")
	fmt.Println("😀当前系统为", version, "版本")

	// 查看嵌入的资源
	// dirEntries, _ := f.ReadDir("rpm")
	// for _, de := range dirEntries {
	// 	fmt.Println(de.Name(), de.IsDir())
	// }

	switch version {
	case 8:
		fmt.Println("📦 install shell")

		fish, err := f.ReadFile(`rpm/8/fish-3.3.1-1.4.x86_64.rpm`)
		checkError(err)
		// todo: 目录问题
		err = ioutil.WriteFile(`/tmp/1.rpm`, fish, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/1.rpm`)
		checkError(err)

		fmt.Println("📦 install tmux")
		tmux, err := f.ReadFile(`rpm/8/tmux-2.6-1.11.x86_64.rpm`)
		checkError(err)
		err = ioutil.WriteFile(`/tmp/2.rpm`, tmux, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/2.rpm`)
		checkError(err)
	case 7:
		fmt.Println("📦 install shell")

		fish, err := f.ReadFile(`rpm/7/fish-3.3.1-1.1.x86_64.rpm`)
		checkError(err)
		// todo: 目录问题
		err = ioutil.WriteFile(`/tmp/1.rpm`, fish, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/1.rpm`)
		checkError(err)

		fmt.Println("📦 install tmux")
		tmux, err := f.ReadFile(`rpm/7/tmux-2.9a-4.4.x86_64.rpm`)
		checkError(err)
		err = ioutil.WriteFile(`/tmp/2.rpm`, tmux, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/2.rpm`)
		checkError(err)
	case 6:
		fmt.Println("📦 install shell")

		fish, err := f.ReadFile(`rpm/6/fish-3.1.2+1603.gff144a38d-2.1.x86_64.rpm`)
		checkError(err)
		// todo: 目录问题
		err = ioutil.WriteFile(`/tmp/1.rpm`, fish, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/1.rpm`)
		checkError(err)

		fmt.Println("📦 install tmux")
		tmux, err := f.ReadFile(`rpm/6/tmux-2.9a-4.1.x86_64.rpm`)
		checkError(err)
		err = ioutil.WriteFile(`/tmp/2.rpm`, tmux, 0777)
		checkError(err)
		err = runCommand(`/usr/bin/yum`, `install`, "-y", `/tmp/2.rpm`)
		checkError(err)
	}
}

// 报错
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 获取centos版本
func getCentosVersion(level string) int {
	filepath := "/etc/redhat-release"
	filecontent, err := ioutil.ReadFile((filepath))
	checkError(err)
	release_content := string(filecontent)
	// ([\d]+)
	// match, _ := regexp.MatchString("([\\d]+)", release_content)
	re := regexp.MustCompile(`(?m)([\d]+)`)
	versions := re.FindAllString(release_content, -1)
	if len(versions) == 0 {
		return 0
	}
	// 统一出口，用于返回的版本数字
	var ver_num = 0

	ver_major, _ := strconv.Atoi(versions[0])
	ver_minor, _ := strconv.Atoi(versions[1])

	switch level {
	case "major":
		ver_num = ver_major
	case "minor":
		ver_num = ver_minor
	default:
		ver_num = 0
	}
	return ver_num
}

// 执行cmd命令
func runCommand(name string, arg ...string) error {
	// var output []byte
	// var err error
	// cmd := exec.Command(name, arg...)
	// if output, err = cmd.Output(); err != nil {
	// 	return err
	// }
	// fmt.Println(string(output))
	// return nil

	// cmd := exec.Command(name, arg...)
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	return err
	// }
	// if err := cmd.Start(); err != nil {
	// 	return err
	// }
	// output_bytes, err := ioutil.ReadAll(stdout)
	// if err != nil {
	// 	return err
	// }
	// if err := cmd.Wait(); err != nil {
	// 	return err
	// }
	// fmt.Printf("=> \n\n %s", output_bytes)
	// return nil

	cmd := exec.Command(name, arg...)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return err
	}

	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
				return err
			}
			break
		}
		fmt.Printf("=> %s\n", string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
