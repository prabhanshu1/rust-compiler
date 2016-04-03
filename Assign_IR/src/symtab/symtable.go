package symtab

symtab map[string](map[string]string) := make map[string](map[string]string) 

func make_entry(s string) {
	var newmap map[string]string = make map[string]string
	newmap["value"]=s
	symtab[s] = newmap
	return newmap
}
package symtab