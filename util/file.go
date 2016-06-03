package util

import (
	"io/ioutil"
	"os"
)

//文件夹下数量
func SizeofDir(dirPth string) int {
	fielinfo, _ := os.Stat(dirPth)
	if fielinfo.IsDir() {
		files, _ := ioutil.ReadDir(dirPth)
		return len(files)
	}
	return 0
}
