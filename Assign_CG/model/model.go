package model


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

/*var Registers = map[int]string{
        1 : "eax"
        2 : "ebx"
=======
var Registers = map[int]string{
        1 : "eax",
        2 : "ebx",
>>>>>>> ffbe63154fa48782dc38b81dc1c414515658dee0
}

var Arithmetic = map[string]string{
        "+" : "add",
        "-" : "sub",
        "*" : "mul",
        "/" : "div",
}
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
