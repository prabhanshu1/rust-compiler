package main

import (
	"fmt"
	"os"
	"./model"
	"./cg_parser"
)

var instructions = make([]*model.Instr_struct, 0, 5)
var leader = make([]int, 0, 5)


func main() {

	cg_parser.Parser(os.Args[1],&instructions,&leader)
	//fmt.Println(leader,instructions[0])
	return
}