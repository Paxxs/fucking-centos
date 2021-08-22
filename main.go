package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

// cspell:words Centos
// go:embed rpm

var f embed.FS

func main() {
	fmt.Println("😀当前系统为", getCentosVersion("major"), "版本")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

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
