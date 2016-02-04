package main

import (
	"./cg_parser"
	"./model"
	//"fmt"
	"os"
	//"./cg_getreg"
	"./translator"
)

func main() {

	var instructions = make([]*model.Instr_struct, 0, 5)
	var leader = make([]int, 0, 5)
	var assembly model.Final_Code

	cg_parser.Parser(os.Args[1], &instructions, &leader)
	// for i := 0; i < len(instructions); i++ {
	// 	fmt.Println(i, "   ", instructions[i])
	// }
	translator.Translate(&assembly, instructions, leader)

	model.FormattedStringPrint(assembly.Libraries)
	model.FormattedStringPrint(assembly.Global_Section)
	model.FormattedStringPrint(assembly.Data_Section)
	model.FormattedStringPrint(assembly.Text_Section)
	model.FormattedStringPrint(assembly.Main_Code)

	/*for key := range instructions {
		fmt.Println(leader,key, instructions[key])
	}*/

	return
}
