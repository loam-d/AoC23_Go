package util

import (
	"bufio"
	"fmt"
	"os"
)

type linereader struct {
	reader *bufio.Scanner
	_file  *os.File
	line   string
}

func getLineReader(inputFile string) (linereader, error) {
	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return linereader{}, err
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)
	reader := linereader{filescanner, readFile, ""}

	return reader, nil
}

func (lr *linereader) Close() {
	lr._file.Close()
}

func (lr *linereader) Line() string {
	return lr.line
}

func (lr *linereader) Next() bool {
	if lr.reader.Scan() {
		lr.line = lr.reader.Text()
		return true
	}
	return false
}

func GetInputFileReader(day int, part int) (linereader, error) {
	return getLineReader("./input/day1_1.txt")
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
