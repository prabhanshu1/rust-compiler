package translator

import (
	"../model"
	"../cg_getreg"
	"strconv"
)

func Translate(assembly *[]string,instructions *[]*model.Instr_struct,leader *[]int) {
	leader_count = len(leader) -1 ;

	var fresh bool
	var Old_Variable string
	var r1,r2,r3,r4 string

	*assembly=append(*assembly,"#include <asm/unistd.h>")
	*assembly=append(*assembly,"#include <syscall.h>")
	*assembly=append(*assembly,".data")

	All_Variables := model.VariableFind(instructions,leader[0],leader[leader_count]-1)
	for key := range All_Variables {
		*assembly=append(*assembly,All_Variables[i] + ":")
		*assembly=append(*assembly,".long " + strconv.Itoa(69))
		*assembly=append(*assembly,All_Variables[i] + "end:")
	}

	*assembly=append(*assembly,".text")
	*assembly=append(*assembly,".globl _start")
	*assembly=append(*assembly,"_start")

	for i := 0; i < leader_count; i++ {
		table := cg_getreg.Preprocess(instructions,leader[i],leader[i+1]-1)
		for j := leader[i]; j < leader[i+1]; j++ {
			switch instructions[j].Op{

			case "+", "-" :
	
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,assembly,r1,instructions[j].Dest)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,assembly,r2,instructions[j].Src1)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,assembly,r3,instructions[j].Src2)

				*assembly=append(*assembly,"movl" + " " + r2 + "," + r1 )
				*assembly=append(*assembly,model.Arithmetic[instructions[j].Op] + " " + r1 + "," + r3)
	

			case "*", "/", "%":
				// a=b/c or a=b*c
				// edx for a and then set it to 0 
				r4,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],instructions[j].Dest,table,4)
				Load_and_Store(fresh,Old_Variable,assembly,r4,instructions[j].Dest)

				*assembly=append(*assembly,"movl" + "$0" + "," + r4 )
				// eax for a
				r1,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],instructions[j].Dest,table,1)
				Load_and_Store(fresh,Old_Variable,assembly,r1,instructions[j].Dest)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,assembly,r2,instructions[j].Src1)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,assembly,r3,instructions[j].Src2)

				// move b to eax
				*assembly=append(*assembly,"movl" + " " + r2 + "," + r1 )
				// divl %(c waala register)
				*assembly=append(*assembly,model.Arithmetic[instructions[j].Op] + " " + r3)
				if instructions[j].Op=="%" {
					*assembly=append(*assembly,"movl" + " " + r4 + "," + r1 )
				}
				cg_getreg.Free_Reg(r4)

			case "=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,assembly,r1,instructions[j].Dest)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,assembly,r2,instructions[j].Src1)

				*assembly=append(*assembly,"mov " + r1 + "," + r2)

			case "ifgoto" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,assembly,r1,instructions[j].Src2)
				
				*assembly=append(*assembly,"cmp $" + " " + instructions[j].Src1 + " " + r1)
				*assembly=append("j" + *assembly,instructions[j].Dest + " " + instructions[j].jmp)

			case "label" : 
				*assembly=append(*assembly,"label " instructions[j].Src1)

			case "ret" : 
				*assembly=append(*assembly,"ret")

			case "call" :
				*assembly=append(*assembly,"call " + " " + instructions[j].Jmp)

			default :
			}			
		}
	}
	
}


func Load_and_Store(fresh bool, Old_Variable string,assembly *[]string,reg string,New_Variable string) {
	if fresh {
		if Old_Variable!="" {
			*assembly=append(*assembly,"Store " + reg + " " + Old_Variable)
		}
		*assembly=append(*assembly,"load " + reg + " " + New_Variable)
	}
}