package model

type Instr_struct struct {
	Op   string
	Dest string
	Src1 string
	Src2 string
	Jmp  string
}

var Registers = map[int]string{
        1 : "eax"
        2 : "ebx"
}

var Arithmetic = map[string]string{
        "+" : "add"
        "-" : "sub"
        "*" : "mul"
        "/" : "div"
}

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
