package model

import (
	"fmt"
	//"math"
	"strconv"
)

type Instr_struct struct {
	Op   string
	Dest string
	Src1 string
	Src2 string
	Jmp  string
}

type Final_Code struct {
	Libraries      []string
	Global_Section []string
	Data_Section   []string
	Text_Section   []string
	Main_Code      []string
}

type Ref_Table struct {
	Ref_t map[string]int
}

func (table *Ref_Table) Initialize(s string) {
	(*table).Ref_t[s] = 99999999
	return
}

//Variable to register
//Register variable

//getreg freereg wrong wrong wrong wrong
type Ref_Maps struct {
	VtoR map[string]string
	RtoV map[string]string
}

func Set_Reg_Map(Ref_Map *Ref_Maps, Reg string, Val string) {
	(*Ref_Map).RtoV[Reg] = Val
}

func Set_Var_Map(Ref_Map *Ref_Maps, Var string, Val string) {
	(*Ref_Map).VtoR[Var] = Val
}

// can be wrong

func (table *Ref_Table) Use(s string, instr int) {
	(*table).Ref_t[s] = instr
	return
}

func (table *Ref_Table) Dead(s string) {
	(*table).Ref_t[s] = -1
	return
}

var Registers = map[int]string{
	1: "%eax",
	2: "%ebx",
	3: "%ecx",
	4: "%edx",
	5: "%edi",
	6: "%esi",
}
var Arithmetic = map[string]string{
	"+": "addl",
	"-": "subl",
	"*": "mull",
	"/": "divl",
	"%": "divl",
}

/*var Registers = map[int]string{
  1 : "eax"
  2 : "ebx"


*/
func RemoveDuplicates(a []int) []int {
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

func Initialize_instr(instr *Instr_struct, Op, Dest, Src1, Src2, Jmp string) {
	instr.Op = Op
	instr.Dest = Dest
	instr.Src1 = Src1
	instr.Src2 = Src2
	instr.Jmp = Jmp
	//fmt.Println(instr, "parsed")
}

/*func Initialize_table_row(entry *Ref_Table_row, Variable string){
        entry.Variable = Variable
        entry.Last = -1                   //-1 corresponds to dead state
        entry.Next = -1                   //-1 corresponds to dead state
}
*/
/*func Copy(input []Ref_Table_row) []Ref_Table_row{
        output:= make([]Ref_Table_row, len(input))
        for i,v:=range input{
                output[i].Variable = v.Variable
                output[i].Next = v.Next
                output[i].Last = v.Last
        }
        return output
}*/

func VariableFind(instructions []*Instr_struct, start int, end int) ([]string, []string,[]string) {
	m := make(map[string]bool) //To keep track of what has already been inserted
	vars := make([]string, 0)

	array_m := make(map[string]bool) //To keep track of what has already been inserted
	array_vars := make([]string, 0)


	string_m := make(map[string]bool) //To keep track of what has already been inserted
	string_vars := make([]string, 0)

	for i := start; i <= end; i++ {
		if instructions[i].Op == "=[]" {
			AppendCheck(instructions[i].Dest, m, &vars)
			AppendCheck(instructions[i].Src1, array_m, &array_vars)
			AppendCheck(instructions[i].Src2, m, &vars)
		} else if instructions[i].Op == "[]=" {
			AppendCheck(instructions[i].Dest, m, &vars)
			AppendCheck(instructions[i].Src1, array_m, &array_vars)
			AppendCheck(instructions[i].Src2, m, &vars)
		} else if instructions[i].Op == "ifgoto" {
			AppendCheck(instructions[i].Src1, m, &vars)
			AppendCheck(instructions[i].Src2, m, &vars)
		} else if instructions[i].Op == "string" {
			AppendCheck(instructions[i].Dest, string_m, &string_vars)
			AppendCheck(instructions[i].Src1, string_m, &string_vars)
		} else if instructions[i].Op != "label" && instructions[i].Op != "arg"  {
			AppendCheck(instructions[i].Dest, m, &vars)
			AppendCheck(instructions[i].Src1, m, &vars)
			AppendCheck(instructions[i].Src2, m, &vars)
		}else if instructions[i].Op == "call" {
			AppendCheck(instructions[i].Dest, m, &vars)
		}
	}
	return vars, array_vars, string_vars
}

func AppendCheck(s string, m map[string]bool, vars *[]string) {
	if s != "" {
		_, err := strconv.Atoi(s)
		if err != nil { // error indicates NAN, hence it is a variable
			_, ok := m[s] //OK = true indicates already in map
			if !ok {
				m[s] = true
				*vars = append(*vars, s)
			}
		}
		//fmt.Println(*vars, len(*vars))
	}
}

func FormattedStringPrint(s []string) {
	for _, b := range s {
		fmt.Println(b)
	}
	fmt.Println("\n")
}
