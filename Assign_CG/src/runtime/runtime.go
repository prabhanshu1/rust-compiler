package runtime

import (
	"../model"
)

func ContextSave(data  *model.FinalCode.Data_Section, Ref_Map *model.Ref_Maps){
	temp := Ref_Map.RtoV["%eax"]
	data = append (data, "movl %eax, "+temp);   //Return Value
	data = append (data, "pushl %ebx");			
	data = append (data, "pushl %ecx");			
	data = append (data, "pushl %edx");			
	data = append (data, "pushl %esi");			
	data = append (data, "pushl %edi");						
}

func Clear_Maps(Ref_Map *model.Ref_Maps) {
	for key, value := range (*Ref_Map).RtoV {
		model.Set_Reg_Map(Ref_Map, key, "")
	}
	for key, _ := range (*Ref_Map).VtoR {
		model.Set_Var_Map(Ref_Map, key, "")
	}
}
