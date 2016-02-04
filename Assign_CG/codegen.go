package main

import (
	"fmt"
	"os"
	"./model"
	"./cg_parser"
	//"./cg_getreg"
	"./translator"
)


func main() {

	var instructions = make([]*model.Instr_struct, 0, 5)
	var leader = make([]int, 0, 5)
	var assembly model.Final_Code

	cg_parser.Parser(os.Args[1], &instructions, &leader)

	translator.Translate(&assembly,instructions,leader)

	
	fmt.Println(assembly)

	/*for key := range instructions {
		fmt.Println(leader,key, instructions[key])
	}*/

	return
}