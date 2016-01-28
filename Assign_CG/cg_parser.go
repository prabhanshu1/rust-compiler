package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instr_struct struct {
	Op   string
	Dest string
	Src1 string
	Src2 string
	Jmp  bool
}

var file string
var instructions = make([]*Instr_struct, 0, 5)

func parse_line(str string, line int) {
	instr := new(Instr_struct)

	s := strings.Split(str, ", ")

	if s[0] != strconv.Itoa(line) {
		log.Fatal("file line and instruction line no do not match")
	}

	switch s[1] {
	case "+", "-", "*", "/":
		instr.Op = s[1]
		instr.Dest = s[2]
		instr.Src1 = s[3]
		instr.Src2 = s[4]
		instr.Jmp = false
		fmt.Println(instr, "case")
	default:
		fmt.Println(s[1], "hello")
	}

	instructions = append(instructions, instr)
	return
}

func parser(file_name string) {
	line := 1
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		parse_line(str, line)
		line += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func main() {

	parser(os.Args[1])
	return
}
