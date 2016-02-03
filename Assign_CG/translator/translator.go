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
	var Ref_Map model.Ref_Maps

	((*Code).Libraries)=append(((*Code).Libraries),"#include <asm/unistd.h>")
	((*Code).Libraries)=append(((*Code).Libraries),"#include <syscall.h>")

	((*Code).Data_Section)=append(((*Code).Data_Section),".data")

	Non_Array_Variables,Array_Variables := model.VariableFind(instructions,leader[0],leader[leader_count-1])
// initialize all the map of r to v and v to r.	
	for ii=0;ii<len(Non_Array_Variables);ii++{
		Ref_Map.VtoR[Non_Array_Variables[ii]]="";
	}

	len_reg:=6
	for ii=0;ii<len_reg;ii++{
		Ref_Map.RtoV[model.Register[ii]]="";
	}

	data := (*Code).Data_Section
	for key := range Non_Array_Variables {
		data=append(data,Non_Array_Variables[i] + ":")
		data=append(data,".long " + strconv.Itoa(69))
		data=append(data,Non_Array_Variables[i] + "end:")
	}
	for key := range Array_Variables {
		data=append(data,Array_Variables[i] + ":")
		data=append(data,".rept 100")
		data=append(data,".long " + strconv.Itoa(69))
		data=append(data,".endr")
		data=append(data,Array_Variables[i] + "end:")
	}
	(*Code).Data_Section = data

	data = (*Code).Text_Section
	data = append(data,".text")
	((*Code).Text_Section) = data

	data = (*Code).Global_Section
	data=append((data,".globl _start")
	(*Code).Global_Section = data

	data = (*Code).Main_Code
	data=append((data,"_start")

	for i := 0; i < leader_count; i++ {
		table := cg_getreg.Preprocess(instructions,leader[i],leader[i+1]-1)
		for j := leader[i]; j < leader[i+1]; j++ {

			dest := instructions[j].Dest
			src1 := instructions[j].Src1
			src2 := instructions[j].Src2
			op   := instructions[j].Op
			jmp  := instructions[j].Jmp

			switch op{

			case "+", "-" :
	
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,dest,Ref_Map)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,src1,Ref_Map)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,src2,Ref_Map)

				(data=append((data,"movl" + " " + r2 + "," + r1 )
				(data=append((data,model.Arithmetic[op] + " " + r1 + "," + r3)
	

			case "*", "/", "%":
				// a=b/c or a=b*c
				// edx for a and then set it to 0 
				r4,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],dest,table,4)
				Free_Store(fresh,Old_Variable,Code,r4,dest,Ref_Map)

				(data=append((data,"movl" + "$0" + "," + r4 )
				// eax for a
				r1,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],dest,table,1)
				Load_and_Store(fresh,Old_Variable,Code,r1,dest,Ref_Map)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,src1,Ref_Map)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,src2,Ref_Map)

				// move b to eax
				(data=append((data,"movl" + " " + r2 + "," + r1 )
				// divl %(c waala register)
				(data=append((data,model.Arithmetic[op] + " " + r3)
				if op=="%" {
					(data=append((data,"movl" + " " + r4 + "," + r1 )
				}
				model.Set_Reg_Map(Ref_Map,r4,"")

			case "=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,dest,Ref_Map)
				// remove $

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r2,src1,Ref_Map)

				(data=append((data,"mov " + r1 + "," + r2)

			case "=[]" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,dest,Ref_Map)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,src2,Ref_Map)

				/// specifically for int 
				(data=append((data,"movl " + src1 + "(," + r3 ",4)" + "," + r1)

			case "[]=" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,src2,Ref_Map)

				r3,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],dest,table)
				Load_and_Store(fresh,Old_Variable,Code,r3,dest,Ref_Map)

				/// specifically for int 
				(data=append((data,"movl " + r1 + "," + src1 + "(," + r3 ",4)" )

			case "ifgoto" :
				r1,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src1,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,src1,Ref_Map)

				r2,fresh,Old_Variable = cg_getreg.Getreg(j-leader[i],src2,table)
				Load_and_Store(fresh,Old_Variable,Code,r1,src2,Ref_Map)
				
				(data=append((data,"cmp " + r1 + "," + r2 )
				(data=append((data, dest + " " + jmp)

			case "label" : 
				(data=append((data,"label " src1)

			case "ret" : 
				if src1!=""{
					r1,fresh,Old_Variable = cg_getreg.Getreg_Force(j-leader[i],src1,table,1)
					Free_Store(fresh,Old_Variable,Code,r1,src1,Ref_Map)

				}
				
				(data=append((data,"ret")

			case "call" :
				(data=append((data,"call " + " " + jmp)

			default :
			}			
		}
	}
	(*Code).Main_Code = data
	
}


func Load_and_Store(fresh bool, Old_Variable string,Code *model.Final_Code,reg *string,New_Variable string,Ref_Map *model.Ref_Maps) {
	if fresh==1 {
		if Old_Variable!="" {
			((*Code).Main_Code=append(((*Code).Main_Code,"Store " + *reg + " " + Old_Variable)
			model.Set_Var_Map(Ref_Map,Old_Variable,"");
		}
		((*Code).Main_Code=append(((*Code).Main_Code,"load " + *reg + " " + New_Variable)
		model.Set_Var_Map(Ref_Map,New_Variable,*reg);
	}else if fresh == 2 {
		*reg = "$" + New_Variable
	}
}

// used while dumping values in case of freeing a register
func Free_Store(fresh bool, Old_Variable string,Code *model.Final_Code,reg string,New_Variable string,Ref_Map *model.Ref_Maps) {
	if fresh==1 {
		if Old_Variable!="" {
			((*Code).Main_Code=append(((*Code).Main_Code,"Store " + *reg + " " + Old_Variable)
			model.Set_Var_Map(Ref_Map,Old_Variable,"");
		}
	}else if fresh == 2 {
		*reg = "$" + New_Variable
	}
}