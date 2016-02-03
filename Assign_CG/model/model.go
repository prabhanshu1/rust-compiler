package model

import (
    "strconv"
)

type Instr_struct struct {
	Op   string
	Dest string
	Src1 string
	Src2 string
	Jmp  string
}

type Ref_Table_row struct{
        Variable string
        Last int
        Next int
}

type Ref_Table struct{
        Ref_t []Ref_Table_row
}

func (table *Ref_Table) Use(s string, instr int){
        for i, flag:= 0, 0; flag == 0 && i<len(table.Ref_t); i++{
                if (*table).Ref_t[i].Variable == s {
                        if((*table).Ref_t[i].Last == -1){(*table).Ref_t[i].Last = instr}
                        (*table).Ref_t[i].Next = instr
                        //fmt.Println("Here use")
                        flag = 1
                }
                //fmt.Println("Here no use use")
        }   
        return    
}


func (table *Ref_Table) Dead(s string){
        for i, flag := 0, 0; flag == 0; i++{
                if (*table).Ref_t[i].Variable == s {
                        (*table).Ref_t[i].Last = -1
                        (*table).Ref_t[i].Next = -1
                        flag = 1
                        //fmt.Println("Here Dead")
                }
        }
        return
}

var Registers = map[int]string{
        1 : "eax",
        2 : "ebx",

}
var Arithmetic = map[string]string{
        "+" : "add",
        "-" : "sub",
        "*" : "mul",
        "/" : "div",
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

func Initialize_table_row(entry *Ref_Table_row, Variable string){
        entry.Variable = Variable
        entry.Last = -1                   //-1 corresponds to dead state
        entry.Next = -1                   //-1 corresponds to dead state                                                                      
}

func Copy(input []Ref_Table_row) []Ref_Table_row{
        output:= make([]Ref_Table_row, len(input))
        for i,v:=range input{
                output[i].Variable = v.Variable
                output[i].Next = v.Next
                output[i].Last = v.Last
        }
        return output
}


func VariableFind(instructions []*Instr_struct, start int, end int)([]string){
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