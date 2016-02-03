package translator

import (
	"../model"
	"../cg_getreg"
	"strconv"
)

func Translate((*Code) *model.Final_Code,instructions *[]*model.Instr_struct,leader *[]int) {
	leader_count = len(leader) -1 ;

	var fresh int
	var Old_Variable string
	var r1,r2,r3,r4 string

	((*Code).Libraries)=append(((*Code).Libraries),"#include <asm/unistd.h>")
	((*Code).Libraries)=append(((*Code).Libraries),"#include <syscall.h>")

	((*Code).Data_Section)=append(((*Code).Data_Section),".data")

	Non_Array_Variables,Array_Variables := model.VariableFind(instructions,leader[0],leader[leader_count]-1)
	for key := range Non_Array_Variables {
		((*Code).Data_Section)=append(((*Code).Data_Section),Non_Array_Variables[i] + ":")
		((*Code).Data_Section)=append(((*Code).Data_Section),".long " + strconv.Itoa(69))
		((*Code).Data_Section)=append(((*Code).Data_Section),Non_Array_Variables[i] + "end:")
	}
	for key := range Array_Variables {
		((*Code).Data_Section)=append(((*Code).Data_Section),Array_Variables[i] + ":")
		((*Code).Data_Section)=append(((*Code).Data_Section),".rept 100")
		((*Code).Data_Section)=append(((*Code).Data_Section),".long " + strconv.Itoa(69))
		((*Code).Data_Section)=append(((*Code).Data_Section),".endr")
		((*Code).Data_Section)=append(((*Code).Data_Section),Array_Variables[i] + "end:")
	}

	((*Code).Text_Section)=append(((*Code).Text_Section),".text")
	((*Code).Global_Section)=append(((*Code).Global_Section),".globl _start")
	((*Code).Main_Code=append(((*Code).Main_Code,"_start")

	for i := 0; i < leader_count; i++ {
		table := cg_getreg.Preprocess(instructions,leader[i],leader[i+1]-1)
		for j := leader[i]; j < leader[i+1]; j++ {
			switch instructions[j].Op{

			case "+", "-" :
	
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Dest)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,instructions[j].Src1)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,instructions[j].Src2)

				((*Code).Main_Code=append(((*Code).Main_Code,"movl" + " " + r2 + "," + r1 )
				((*Code).Main_Code=append(((*Code).Main_Code,model.Arithmetic[instructions[j].Op] + " " + r1 + "," + r3)
	

			case "*", "/", "%":
				// a=b/c or a=b*c
				// edx for a and then set it to 0 
				r4,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],instructions[j].Dest,table,4)
				Load_and_Store(fresh,Old_Variable,Code,r4,instructions[j].Dest)

				((*Code).Main_Code=append(((*Code).Main_Code,"movl" + "$0" + "," + r4 )
				// eax for a
				r1,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],instructions[j].Dest,table,1)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Dest)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,instructions[j].Src1)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,instructions[j].Src2)

				// move b to eax
				((*Code).Main_Code=append(((*Code).Main_Code,"movl" + " " + r2 + "," + r1 )
				// divl %(c waala register)
				((*Code).Main_Code=append(((*Code).Main_Code,model.Arithmetic[instructions[j].Op] + " " + r3)
				if instructions[j].Op=="%" {
					((*Code).Main_Code=append(((*Code).Main_Code,"movl" + " " + r4 + "," + r1 )
				}
				cg_getreg.Free_Reg(r4)

			case "=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Dest)
				// remove $

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,instructions[j].Src1)

				((*Code).Main_Code=append(((*Code).Main_Code,"mov " + r1 + "," + r2)

			case "=[]" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Dest)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,instructions[j].Src2)

				/// specifically for int 
				((*Code).Main_Code=append(((*Code).Main_Code,"movl " + instructions[j].Src1 + "(," + r3 ",4)" + "," + r1)

			case "[]=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Src2)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,instructions[j].Dest)

				/// specifically for int 
				((*Code).Main_Code=append(((*Code).Main_Code,"movl " + r1 + "," + instructions[j].Src1 + "(," + r3 ",4)" )

			case "ifgoto" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Src1)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],instructions[j].Src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,instructions[j].Src2)
				
				((*Code).Main_Code=append(((*Code).Main_Code,"cmp " + r1 + "," + r2 )
				((*Code).Main_Code=append(((*Code).Main_Code, instructions[j].Dest + " " + instructions[j].jmp)

			case "label" : 
				((*Code).Main_Code=append(((*Code).Main_Code,"label " instructions[j].Src1)

			case "ret" : 
				((*Code).Main_Code=append(((*Code).Main_Code,"ret")

			case "call" :
				((*Code).Main_Code=append(((*Code).Main_Code,"call " + " " + instructions[j].Jmp)

			default :
			}			
		}
	}
	
}


func Load_and_Store(fresh bool, Old_Variable string,Code *model.Final_Code,reg string,New_Variable string) {
	if fresh==1 {
		if Old_Variable!="" {
			((*Code).Main_Code=append(((*Code).Main_Code,"Store " + reg + " " + Old_Variable)
		}
		((*Code).Main_Code=append(((*Code).Main_Code,"load " + reg + " " + New_Variable)
	}else if fresh == 2 {
		reg = "$" + New_Variable
	}
}