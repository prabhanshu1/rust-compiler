package symtab

var symtab map[string](map[string]string)

var qqq = 0

func Init() {
	symtab = make(map[string](map[string]string))
}
func Make_entry(s string) map[string]string {

	if qqq == 0 {
		Init()
		qqq = qqq + 1
	}

	var newmap map[string]string = make(map[string]string)
	newmap["value"] = s
	symtab[s] = newmap
	return newmap
}
