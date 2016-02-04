package cg_getreg

import (
	"../model"
	"fmt"
	"strconv"
)

//Called at every Basic Block
func Preprocess(instructions []*model.Instr_struct, start int, end int, tables *[]model.Ref_Table) {
	size := end - start + 1
	vars, array_vars := model.VariableFind(instructions, start, end)
	//fmt.Println(vars)

	vars=append(vars,array_vars...)

	var base_table model.Ref_Table
	base_table.Ref_t = make(map[string]int)
	for _, v := range vars {
		base_table.Dead(v)
	}

	(*tables)[size] = base_table
	//fmt.Println(*base_table)

	for i := size - 1; i >= 0; i-- {
		(*tables)[i].Ref_t = make(map[string]int)
		(*tables)[i].Ref_t = (*tables)[i+1].Ref_t
		//fmt.Println(*tables[i], *tables[i+1])
		ModifyTable(*instructions[i], &((*tables)[i]), i)
		//fmt.Println(*tables[i], *tables[i+1])
	}
	//fmt.Println("Here")
	/*for _, v := range *tables {
		fmt.Println(v)
	}*/
	return
}

//Return Values:
//The allocated register
//The return code: 0-> alreary in register | 1->NextUse Applied | 2->Not a variable
func Getreg(pos int, str string, table *[]model.Ref_Table, Ref_Map *model.Ref_Maps)(string, int, string) {
	max := 0
	var max_val string
	_, ok := (*Ref_Map).VtoR[str]
	
	if !ok {
		return str, 2, ""
	} else if (*Ref_Map).VtoR[str] != "" {
		return (*Ref_Map).VtoR[str], 0, ""
	} else {
		for key, value := range (*table)[pos].Ref_t {
			if value > max {
				max = value
				max_val = key
			}
		}
		return (*Ref_Map).VtoR[max_val], 1, max_val
	}

}

func Getreg_Force(pos int, str string, table *[]model.Ref_Table, Ref_Map *model.Ref_Maps, reg int) (string, int, string){
	_, ok := (*Ref_Map).VtoR[str]

	if !ok {
		return str, 2, ""
	} else if (*Ref_Map).VtoR[str] == model.Registers[reg] {
		return (*Ref_Map).VtoR[str], 0, ""
	} else {
		return model.Registers[reg], 1, (*Ref_Map).RtoV[model.Registers[reg]]
	}

}

func UseCheck(s string, table *model.Ref_Table, instr int) {
	if s != "" {
		_, err := strconv.Atoi(s)
		if err != nil {
			table.Use(s, instr)
		}
	}
	//fmt.Println(*vars, len(*vars))
}

func ModifyTable(instruction model.Instr_struct, table *model.Ref_Table, i int) {
	oper := instruction.Op
	dest := instruction.Dest
	src1 := instruction.Src1
	src2 := instruction.Src2
	switch oper {
	case "call", "ret", "label":
		return
	case "=":
		table.Dead(dest)
		UseCheck(src1, table, i)
		//UseCheck(src2, table, i)
		//fmt.Println(src1, src2)
	case "+", "-", "*", "/", "%":
		table.Dead(dest)
		UseCheck(src1, table, i)
		UseCheck(src2, table, i)
		//fmt.Println(src1, src2, dest)
	case "print":
		UseCheck(src1, table, i)
	case "ifgoto":
		UseCheck(src1, table, i)
		UseCheck(src2, table, i)
		//fmt.Println(src1, src2)
	}
	//fmt.Println(*table)
}
