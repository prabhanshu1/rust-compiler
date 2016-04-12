package runtime

func ContextSave{

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
