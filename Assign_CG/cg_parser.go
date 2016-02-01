package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Instr_struct struct {
	Op   string
	Dest string
	Src1 string
	Src2 string
	Jmp  string
}

var file string
var instructions = make([]*Instr_struct, 0, 5)
var leader = make([]int, 0, 5)

func removeDuplicates(a []int) []int { 
        result := []int{} 
        seen := map[int]int{} 
        for _, val := range a { 
                if _, ok := seen[val]; !ok { 
                        result = append(result, val) 
                        seen[val] = val 
                } 
        } 
        return result 
} 

func initialize_instr(instr *Instr_struct, Op, Dest, Src1, Src2, Jmp string) {
	instr.Op = Op
	instr.Dest = Dest
	instr.Src1 = Src1
	instr.Src2 = Src2
	instr.Jmp = Jmp
	//fmt.Println(instr, "parsed")
}

func parse_line(str string, line int) {
	instr := new(Instr_struct)

	s := strings.Split(str, ", ")

	if s[0] != strconv.Itoa(line) {
		log.Fatal("file line and instruction line no do not match")
	}

	if strconv.Itoa(line) != s[0] {
		log.Fatal("Line mismatch, in sorce file and cg_parser")
	}

	switch s[1] {
	case "+", "-", "*", "/":
		initialize_instr(instr, s[1], s[2], s[3], s[4], "0")
	case "=":
		initialize_instr(instr, s[1], s[2], s[3], "", "0")
	case "ifgoto":
		initialize_instr(instr, s[2], "", s[3], s[4], s[5])
		s,err := strconv.Atoi(s[5])
		if err!=nil {
			log.Fatal("Invalid Jump Target")
		}
		leader=append(leader,s);
	case "call":
		initialize_instr(instr, s[1], "", "", "", s[2])
	case "ret":
		initialize_instr(instr, s[1], "", "", "", "-1")
	case "print":
		initialize_instr(instr, s[1], "", s[2], "", "-2")
	case "label":
		initialize_instr(instr, s[1], "", s[2], "", "-3")
		leader=append(leader,line);
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

	leader=append(leader,1);

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		parse_line(str, line)
		line += 1
	}

	leader=append(leader,line);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	leader = removeDuplicates(leader) 
	sort.Ints(leader)

	defer file.Close()
}

func main() {

	parser(os.Args[1])
	fmt.Println(leader,instructions)
	return
}
