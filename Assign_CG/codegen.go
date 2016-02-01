package main

import (
	"fmt"
	"os"
	"./model"
	"./cg_parser"
	//"./translator"
)

var instructions = make([]*model.Instr_struct, 0, 5)
var leader = make([]int, 0, 5)
var assembly = make([]string,0,5)

func main() {

	cg_parser.Parser(os.Args[1],&instructions,&leader)

	//translator.Translate(&assembly,&instructions,&leader)

	for key := range instructions {
		fmt.Println(leader,key, instructions[key])
	}

	return
}