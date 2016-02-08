package translator

import (
	"../cg_getreg"
	"../model"
	//"fmt"
	"strconv"
)

func Translate(Code *model.Final_Code, instructions []*model.Instr_struct, leader []int) {
	leader_count := len(leader) - 1

	var fresh int
	var Old_Variable string
	var r1, r2, r3, r4 string
	var Ref_Map model.Ref_Maps
	Ref_Map.VtoR = make(map[string]string)
	Ref_Map.RtoV = make(map[string]string)

	/*((*Code).Libraries) = append(((*Code).Libraries), "#include <asm/unistd.h>")
	((*Code).Libraries) = append(((*Code).Libraries), "#include <syscall.h>")*/

	((*Code).Data_Section) = append(((*Code).Data_Section), ".section .data")

	Non_Array_Variables, Array_Variables, String_Variables := model.VariableFind(instructions, leader[0], leader[leader_count])
	// initialize all the map of r to v and v to r.
	Non_Array_Variables = append(Non_Array_Variables, "temporary_compiler_variable")
	for i := 0; i < len(Non_Array_Variables); i++ {
		Ref_Map.VtoR[Non_Array_Variables[i]] = ""
	}

	len_reg := 6
	for i := 1; i <= len_reg; i++ {
		Ref_Map.RtoV[model.Registers[i]] = ""
	}

	data := (*Code).Data_Section
	for i := range Non_Array_Variables {
		data = append(data, Non_Array_Variables[i]+":")
		data = append(data, ".long "+strconv.Itoa(69))
		data = append(data, Non_Array_Variables[i]+"end:")
	}

	for i := 0; i < len(String_Variables); i++ {
		data = append(data, String_Variables[i]+":")
		data = append(data, ".ascii "+"\""+String_Variables[i+1]+"\"")
		data = append(data, String_Variables[i]+"end:")
		i++
	}

	for i := range Array_Variables {
		data = append(data, Array_Variables[i]+":")
		data = append(data, ".rept 100")
		data = append(data, ".long "+strconv.Itoa(69))
		data = append(data, ".endr")
		data = append(data, Array_Variables[i]+"end:")
	}
	(*Code).Data_Section = data

	data = (*Code).Text_Section
	data = append(data, ".section .text")
	((*Code).Text_Section) = data

	data = (*Code).Global_Section
	data = append(data, ".globl main")
	(*Code).Global_Section = data

	data = (*Code).Main_Code
	data = append(data, "main:")

	for i := 0; i < leader_count; i++ {
		table := make([]model.Ref_Table, leader[i+1]-leader[i]+2)
		cg_getreg.Preprocess(instructions, leader[i], leader[i+1]-1, &table)
		for j := leader[i]; j < leader[i+1]; j++ {

			dest := instructions[j].Dest
			src1 := instructions[j].Src1
			src2 := instructions[j].Src2
			op := instructions[j].Op
			jmp := instructions[j].Jmp
			//fmt.Println("\n ", instructions[j])
			switch op {

			case "+","&","^","|":
				// to prevent the case a=b+a
				if src2 == dest {
					src2, src1 = src1, src2
				}
				r2, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
				//fmt.Println("fuck", r2, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r2, src1, &Ref_Map)
				//fmt.Println("fuck2", r2, fresh, Old_Variable)

				r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				//fmt.Println("fuck3", r3, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)
				//fmt.Println("fuck4", r3, fresh, Old_Variable)

				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
				//fmt.Println("fuck5", r1, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
				//fmt.Println("fuck6", r1, fresh, Old_Variable)

				data = append(data, "movl"+" "+r2+","+r1)

				data = append(data, model.Arithmetic[op]+" "+r3+","+r1)

			case "-":
				r2, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r2, src1, &Ref_Map)

				r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)

				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)

				if src2 == dest {
					//a=b-a  case
					src2 = "temporary_compiler_variable"
					r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
					//fmt.Println("fuck7", r3, fresh, Old_Variable)
					Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)
					data = append(data, "movl"+" "+r1+","+r3)

				}

				data = append(data, "movl"+" "+r2+","+r1)
				data = append(data, model.Arithmetic[op]+" "+r3+","+r1)

			case "*", "/", "%":
				// a=b/c or a=b*c
				// edx for a and then set it to 0
				r4, fresh, Old_Variable = cg_getreg.Getreg_Force(&data, j-leader[i], dest, &table, &Ref_Map, 4)
				//fmt.Println("fuck", r4, fresh, Old_Variable)

				Special_Store(fresh, Old_Variable, &data, &r4, dest, &Ref_Map)
				//fmt.Println("fuck2", Ref_Map.RtoV, Ref_Map.VtoR)
				Hold_Reg(r4, &Ref_Map)
				//fmt.Println("fuck3", Ref_Map.RtoV)

				data = append(data, "movl "+"$0"+","+r4)
				// eax for a
				r1, fresh, Old_Variable = cg_getreg.Getreg_Force(&data, j-leader[i], dest, &table, &Ref_Map, 1)
				//fmt.Println("fuck4", r1, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
				Hold_Reg(r1, &Ref_Map)
				//fmt.Println("fuck5", Ref_Map.RtoV)

				r2, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
				//fmt.Println("fuck6", r2, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r2, src1, &Ref_Map)

				r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				//fmt.Println("fuck7", r3, fresh, Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)

				if src2 == dest {
					src2 = "temporary_compiler_variable"
					r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
					//fmt.Println("fuck7", r3, fresh, Old_Variable)
					Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)
					data = append(data, "movl"+" "+r1+","+r3)

				}

				// move b to eax
				data = append(data, "movl"+" "+r2+","+r1)
				//move c to edx
				// divl %(c waala register)
				data = append(data, model.Arithmetic[op]+" "+r3)
				if op == "%" {
					data = append(data, "movl"+" "+r4+","+r1)
				}

				Free_Reg(r4, &Ref_Map, "")
				Free_Reg(r1, &Ref_Map, dest)

			case "=":

				r2, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r2, src1, &Ref_Map)

				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
				////fmt.Println(r1, " r1  ", fresh, "  old var", Old_Variable)
				Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
				// remove $
				data = append(data, "movl "+r2+","+r1)

			case "addof":


			r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
			////fmt.Println(r1, " r1  ", fresh, "  old var", Old_Variable)
			Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
			// remove $
			data = append(data, "movl $"+src1+","+r1)

			case "=[]":

				r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r3, src2, &Ref_Map)

				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
				/// specifically for int
				data = append(data, "movl "+src1+"(,"+r3+",4)"+","+r1)

			case "[]=":
				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r1, src2, &Ref_Map)

				r3, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r3, dest, &Ref_Map)
				//dest * 4
				/// specifically for int
				data = append(data, "movl "+r1+","+src1+"(,"+r3+",4)")

			case "ifgoto":
				r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
				////fmt.Println("ZZZZZZZZZZ",r1)
				Load_and_Store(fresh, Old_Variable, &data, &r1, src1, &Ref_Map)
				////fmt.Println("ZZZZZZZZZZ2",r1)

				r2, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
				Load_and_Store(fresh, Old_Variable, &data, &r2, src2, &Ref_Map)

				////fmt.Println("XXXXXXXXXXXXXX",src1,"YYY",src2,"YYY",r1,"YYY",r2)
				/*				if fresh==2 {
								r1,r2=r2,r1
							}*/
				Free_reg_at_end(&data, &Ref_Map)
				data = append(data, "cmpl "+r1+","+r2)
				data = append(data, dest+" "+jmp)

			case "label":
				data = append(data, src1+":")

			case "ret":
				if src1 != "" {
					r1, fresh, Old_Variable = cg_getreg.Getreg_Force(&data, j-leader[i], src1, &table, &Ref_Map, 1)
					Load_and_Store(fresh, Old_Variable, &data, &r1, src1, &Ref_Map)

				}
				//Free_reg_at_end(&data, &Ref_Map)

				data = append(data, "ret")

				/*			case "print":
							r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
							////fmt.Println("ZZZZZZZZZZ",r1)
							Load_and_Store(fresh, Old_Variable, &data, &r1, src1, &Ref_Map)

							data = append(data, "push "+r1)
							data = append(data, "print")*/

			case "arg":

				if jmp != "" {
					data = append(data, "pushl "+jmp)
				}
				if src2 != "" {
					r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src2, &table, &Ref_Map)
					Load_and_Store(fresh, Old_Variable, &data, &r1, src2, &Ref_Map)
					data = append(data, "pushl "+r1)
				}
				if src1 != "" {
					r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], src1, &table, &Ref_Map)
					Load_and_Store(fresh, Old_Variable, &data, &r1, src1, &Ref_Map)
					data = append(data, "pushl "+r1)
				}
				if dest != "" {
					r1, fresh, Old_Variable = cg_getreg.Getreg(j-leader[i], dest, &table, &Ref_Map)
					Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
					data = append(data, "pushl "+r1)
				}

			case "call":
				data = append(data, "call "+" "+jmp)
				if dest!="" {
					data=append(data,"movl %eax,"+dest)
					r1, fresh, Old_Variable = cg_getreg.Getreg_Force(&data, j-leader[i], dest, &table, &Ref_Map, 1)
					Load_and_Store(fresh, Old_Variable, &data, &r1, dest, &Ref_Map)
				}

			case "exit":
				Free_reg_at_end(&data, &Ref_Map)
				data = append(data, "movl $1,%eax")
				data = append(data, "movl $0,%ebx")
				data = append(data, "int $0x80")
			case "jmp":
				Free_reg_at_end(&data, &Ref_Map)
				data = append(data, op+" "+jmp)
			default:
			}
		}
		Free_reg_at_end(&data, &Ref_Map)
	}

	(*Code).Main_Code = data
}

func Load_and_Store(fresh int, Old_Variable string, data *[]string, reg *string, New_Variable string, Ref_Map *model.Ref_Maps) {
	if fresh == 1 {
		if Old_Variable != "" {
			*data = append(*data, "movl "+*reg+","+Old_Variable)
			model.Set_Var_Map(Ref_Map, Old_Variable, "")

		}
		*data = append(*data, "movl "+New_Variable+", "+*reg)
		model.Set_Var_Map(Ref_Map, New_Variable, *reg)
		model.Set_Reg_Map(Ref_Map, *reg, New_Variable)
	} else if fresh == 2 {
		*reg = "$" + New_Variable
	}
}

// used while dumping values in case of freeing a register
func Special_Store(fresh int, Old_Variable string, data *[]string, reg *string, New_Variable string, Ref_Map *model.Ref_Maps) {
	if fresh == 1 {
		if Old_Variable != "" {
			*data = append(*data, "movl "+*reg+","+Old_Variable)
			model.Set_Var_Map(Ref_Map, Old_Variable, "")
		}
	} else if fresh == 2 {
		*reg = "$" + New_Variable
	}
}
func Free_reg_at_end(data *[]string, Ref_Map *model.Ref_Maps) {
	for key, value := range (*Ref_Map).RtoV {
		if value != "" {
			*data = append(*data, "movl "+key+", "+value)
		}
		model.Set_Reg_Map(Ref_Map, key, "")
	}
	for key, _ := range (*Ref_Map).VtoR {
		model.Set_Var_Map(Ref_Map, key, "")
	}

}

func Hold_Reg(reg string, Ref_Map *model.Ref_Maps) {
	model.Set_Reg_Map(Ref_Map, reg, "@@@@")
}

func Free_Reg(reg string, Ref_Map *model.Ref_Maps, Old_Variable string) {
	model.Set_Reg_Map(Ref_Map, reg, Old_Variable)
}
