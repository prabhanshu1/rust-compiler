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
	var r1,r2,r3 string

	for i := 0; i < leader_count; i++ {
		table := cg_getreg.Preprocess(instructions,leader[i],leader[i+1]-1)
		for j := leader[i]; j < leader[i+1]; j++ {
			switch instructions[j].Op{

			case "+", "-", "*", "/":
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r1 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r1 + " " + instructions[j].Dest)
				}
				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r2 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r2 + " " + instructions[j].Src1)
				}
				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r3 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r3 + " " + instructions[j].Src2)
				}
				*assembly=append(*assembly,model.Arithmetic[instructions[j].Op] + r1 + "," + r2 + "," + r3)

			case "=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r1 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r1 + " " + instructions[j].Dest)
				}
				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r2 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r2 + " " + instructions[j].Src1)
				}
				*assembly=append(*assembly,"mov " + r1 + "," + r2)

			case "ifgoto" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				if fresh {
					if Old_Variable!="" {
						*assembly=append(*assembly,"Store " + r1 + " " + Old_Variable)
					}
					*assembly=append(*assembly,"load " + r1 + " " + instructions[j].Src2)
				}
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


