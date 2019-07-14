package os_file

import (
	"fmt"
	"io"
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
	fc := []byte("write by ioutil\n")
	err := ioutil.WriteFile(FileName, fc, 0644)
	if err != nil {
		panic(err)
	}
	// apend file
	fd, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		panic(err)
	}
	fd.WriteString("write by os.Open\n")
	fd.Close()
}

func ReadFile() {
	fd, err := os.OpenFile(FileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	var text = make([]byte, 1024)
	for {
		_, err := fd.Read(text)
		if err == io.EOF {
			break
		}
		fmt.Println(string(text))
	}
	fmt.Println("read file end")
}
