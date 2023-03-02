package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var PrefixChan chan string
var IndexMap map[string]int

func init() {
	IndexMap = make(map[string]int)
	PrefixChan = make(chan string, 10000000)
	SetupIndexes()
}

func SetupIndexes() {
	v := 0
	for i := 97; i < 123; i++ {

		IndexMap[string(rune(i))] = v
		v++
	}
	for i := 65; i < 91; i++ {
		IndexMap[string(rune(i))] = v
		v++
	}
	for i := 48; i < 58; i++ {
		IndexMap[string(rune(i))] = v
		v++
	}
	fmt.Println(IndexMap)
}

func LoadConfig() (string, error) {

	fileName := os.Getenv("File_Name")
	if fileName == "" {
		return "", errors.New("empty filename configuration")
	}

	return fileName, nil
}

func SetupConfig(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	err = scanFile(f)
	return err
}

func scanFile(f *os.File) error {
	s := bufio.NewScanner(f)
	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
		return err
	}

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		PrefixChan <- line
	}
	close(PrefixChan)
	return nil
}
