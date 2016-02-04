package cg_getreg

import (
	"../model"
	"fmt"
	//	"reflect"
	"strconv"
)

//Called at every Basic Block
func Preprocess(instructions []*model.Instr_struct, start int, end int, tables *[]model.Ref_Table) {
	size := end - start + 1
	vars, array_vars := model.VariableFind(instructions, start, end)
	//fmt.Println(vars)

	vars = append(vars, array_vars...)
	fmt.Println(end, "  e  s  ", start, "\n")
	var base_table model.Ref_Table
	base_table.Ref_t = make(map[string]int)
	for _, v := range vars {
		base_table.Dead(v)
	}

	(*tables)[size] = base_table

	for i := size - 1; i >= 0; i-- {
		(*tables)[i].Ref_t = make(map[string]int)
		for key, value := range (*tables)[i+1].Ref_t {
			(*tables)[i].Ref_t[key] = value
		}
		ModifyTable(*instructions[i+start], &((*tables)[i]), i)
		// fmt.Println(i, "  ", (*tables)[i].Ref_t, (*tables)[i+1].Ref_t, "\n")
		// fmt.Printf("0x%x\n", reflect.ValueOf(((*tables)[i].Ref_t)).Pointer())
		// fmt.Printf("0x%x\n", reflect.ValueOf(((*tables)[i+1].Ref_t)).Pointer())

	}
	return
}

func ModifyTable(instruction model.Instr_struct, table *model.Ref_Table, i int) {
	oper := instruction.Op
	dest := instruction.Dest
	src1 := instruction.Src1
	src2 := instruction.Src2
	//	fmt.Println("oper dest s1 s2", oper, " ", dest, " ", src1, " ", src2)
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

//Return Values:
//The allocated register
//The return code: 0-> alreary in register | 1->NextUse Applied | 2->Not a variable
func Getreg(pos int, str string, table *[]model.Ref_Table, Ref_Map *model.Ref_Maps) (string, int, string) {
	fmt.Println(pos, " position  ")
	fmt.Println((*Ref_Map).VtoR[str])
	fmt.Println((*table)[pos].Ref_t)

	max := 0
	var max_val string
	_, ok := (*Ref_Map).VtoR[str]

	if !ok {
		return str, 2, ""
	} else if (*Ref_Map).VtoR[str] != "" {
		return (*Ref_Map).VtoR[str], 0, ""
	} else {
		for key, value := range (*Ref_Map).RtoV {
			if value == "" {
				fmt.Println(key, "in getreg key", value)
				return key, 1, ""
			}
		}

		for _, value := range (*Ref_Map).RtoV {
			if (*table)[pos].Ref_t[value] > max {
				max = (*table)[pos].Ref_t[value]
				max_val = value
			}
		}

		return (*Ref_Map).VtoR[max_val], 1, max_val
	}
}

func Getreg_Force(data *[]string,pos int, str string, table *[]model.Ref_Table, Ref_Map *model.Ref_Maps, reg int) (string, int, string) {
	_, ok := (*Ref_Map).VtoR[str]

	if !ok {
		return str, 2, ""
	} else if (*Ref_Map).VtoR[str] == model.Registers[reg] {
		return (*Ref_Map).VtoR[str], 0, ""
	} else {
		if (*Ref_Map).VtoR[str] != ""{
			*data = append(*data, "Store "+(*Ref_Map).VtoR[str]+" "+str)
			model.Set_Reg_Map(Ref_Map, (*Ref_Map).VtoR[str], "")
		}
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
