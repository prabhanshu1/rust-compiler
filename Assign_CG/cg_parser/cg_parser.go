package cg_parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	"../InstructionStruct" 
)



func Parse_line(str string, line int, instructions *[]*model.Instr_struct,leader *[]int) {
	instr := new(model.Instr_struct)

	s := strings.Split(str, ", ")

	if s[0] != strconv.Itoa(line) {
		log.Fatal("file line and instruction line no do not match")
	}

	if strconv.Itoa(line) != s[0] {
		log.Fatal("Line mismatch, in sorce file and cg_parser")
	}

	switch s[1] {
	case "+", "-", "*", "/":
		model.Initialize_instr(instr, s[1], s[2], s[3], s[4], "0")
	case "=":
		model.Initialize_instr(instr, s[1], s[2], s[3], "", "0")
	case "ifgoto":
		model.Initialize_instr(instr, s[2], "", s[3], s[4], s[5])
		s,err := strconv.Atoi(s[5])
		if err!=nil {
			log.Fatal("Invalid Jump Target")
		}
		*leader=append(*leader,s);
	case "call":
		model.Initialize_instr(instr, s[1], "", "", "", s[2])
	case "ret":
		model.Initialize_instr(instr, s[1], "", "", "", "-1")
	case "print":
		model.Initialize_instr(instr, s[1], "", s[2], "", "-2")
	case "label":
		model.Initialize_instr(instr, s[1], "", s[2], "", "-3")
		*leader=append(*leader,line);
	default:
		fmt.Println(s[1], "hello")
	}

	*instructions = append(*instructions, instr)

	return
}

func Parser(file_name string, instructions *[]*model.Instr_struct,leader *[]int) {
	line := 1
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	*leader=append(*leader,1);

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		Parse_line(str, line, instructions, leader)
		line += 1
	}

	*leader=append(*leader,line);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	*leader = model.RemoveDuplicates(*leader) 
	sort.Ints(*leader)

	defer file.Close()
}


