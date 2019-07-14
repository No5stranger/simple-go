package os_file

import (
	"io/ioutil"
	"os"
)

var FileName string = "test_os_file"

func CreateFile() {
	fd, err := os.Create(FileName)
	if err != nil {
		panic(err)
	}
	fd.WriteString("new line in file")
	// remeber to close file
	fd.Close()
}

func WriteFile() {
	fc := []byte("write by ioutil")
	err := ioutil.WriteFile(FileName, fc, 0644)
	if err != nil {
		panic(err)
	}
	// apend file
	fd, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		panic(err)
	}
	fd.WriteString("write by os.Open")
	fd.Close()
}
