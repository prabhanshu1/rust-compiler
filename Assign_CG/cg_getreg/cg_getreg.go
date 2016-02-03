package cg_getreg

import (
	"../model" 
	"strconv"
	"fmt"
)


//Called at every Basic Block
func Preprocess(instructions []*model.Instr_struct, start int, end int) (*[]*model.Ref_Table) {
	size := end - start + 1
	var tables = make([]*model.Ref_Table, size+1)
	vars := VariableFind(instructions, start, end)
	fmt.Println(vars)

	base_table := new(model.Ref_Table)
	for _, v := range vars{
		tr := new(model.Ref_Table_row)
		model.Initialize_table_row(tr, v)
		base_table.Ref_t = append(base_table.Ref_t, *tr)		
	}

	tables[size] = base_table
	//fmt.Println(*base_table)

	for i:=size-1; i >= 0; i-- {
		tables[i] = new(model.Ref_Table) 
		(*tables[i]).Ref_t = model.Copy((*tables[i+1]).Ref_t)
		//fmt.Println(*tables[i], *tables[i+1])
		ModifyTable(*instructions[i], tables[i], i)
		//fmt.Println(*tables[i], *tables[i+1])
	}
	//fmt.Println("Here")
	for _, v:= range tables{
		fmt.Println(*v)
	}
	return &tables
}

func VariableFind(instructions []*model.Instr_struct, start int, end int)([]string){
	m:= make(map[string]bool)   //To keep track of what has already been inserted
	vars := make([]string, 0);

	for i:=start; i <= end; i++{
		if(instructions[i].Op != "call" && instructions[i].Op != "label"){
			AppendCheck(instructions[i].Dest, m, &vars)
			AppendCheck(instructions[i].Src1, m, &vars)
			AppendCheck(instructions[i].Src2, m, &vars)
		}
	}
	return vars
}

func AppendCheck(s string, m map[string]bool, vars *[]string){
	if s!=""{
		_, err:= strconv.Atoi(s)  
		if err != nil {        // error indicates NAN, hence it is a variable
			_, ok:= m[s]       //OK = true indicates already in map
			if !ok {
				m[s] = true
				*vars = append(*vars, s)
			}
		}
	//fmt.Println(*vars, len(*vars))
	}
}

func UseCheck(s string, table *model.Ref_Table, instr int){
	if s!=""{
		_, err:= strconv.Atoi(s)  
		if err != nil {table.Use(s, instr)}
		}
	//fmt.Println(*vars, len(*vars))
}

func ModifyTable(instruction model.Instr_struct, table *model.Ref_Table, i int){
	oper:= instruction.Op
	dest:= instruction.Dest
	src1:= instruction.Src1
	src2:= instruction.Src2
	switch oper{
		case "call", "ret", "label":
			return
		case "=":
			table.Dead(dest)
			UseCheck(src1, table, i)
			UseCheck(src2, table, i)
			//fmt.Println(src1, src2)
		case "+", "-", "*", "/":
			UseCheck(src1, table, i)
			UseCheck(src2, table, i)
			UseCheck(dest, table, i)
			//fmt.Println(src1, src2, dest)
		case "print":
			UseCheck(src1, table, i)
		case "leq", "geq":
			UseCheck(src1, table, i)
			UseCheck(src2, table, i)
			//fmt.Println(src1, src2)
	}
	//fmt.Println(*table)
}

