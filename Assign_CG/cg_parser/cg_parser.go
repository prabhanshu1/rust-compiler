package cg_parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	"../model" 
)



func Parse_line(str string, line int, instructions *[]*model.Instr_struct,leader *[]int) {
	instr := new(model.Instr_struct)

	s := strings.Split(str, ", ")

	if s[0] != strconv.Itoa(line) {
		log.Fatal("file line and instruction line no do not match")
	}

	if strconv.Itoa(line) != s[0] {
		log.Fatal("Line mismatch, in sorce file and cg_parser")
	}

	switch s[1] {
	case "+", "-", "*", "/":
		model.Initialize_instr(instr, s[1], s[2], s[3], s[4], "0")
	case "=":
		model.Initialize_instr(instr, s[1], s[2], s[3], "", "0")
	case "ifgoto":
		model.Initialize_instr(instr, s[1], s[2], s[3], s[4], s[5])
		s,err := strconv.Atoi(s[5])
		if err!=nil {
			log.Fatal("Invalid Jump Target")
		}
		*leader=append(*leader,s-1);
	case "call":
		model.Initialize_instr(instr, s[1], "", "", "", s[2])
		*leader=append(*leader,line-1);
	case "ret":
		model.Initialize_instr(instr, s[1], "", "", "", "-1")
		*leader=append(*leader,line);
	case "print":
		model.Initialize_instr(instr, s[1], "", s[2], "", "-2")
	case "label":
		model.Initialize_instr(instr, s[1], "", s[2], "", "-3")
		*leader=append(*leader,line-1);
	default:
		fmt.Println(s[1], "hello")
	}

	*instructions = append(*instructions, instr)

	return
}

func Parser(file_name string, instructions *[]*model.Instr_struct,leader *[]int) {
	line := 1
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	*leader=append(*leader,0);

	tmp := make([]*model.Instr_struct, 0,5)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		Parse_line(str, line, &tmp, leader)
		line += 1
	}

	*leader=append(*leader,line-1);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	*leader = model.RemoveDuplicates(*leader) 
	sort.Ints(*leader)

	label_count := 1
	delta := 0
	label_total := len(*leader) -1 ;
	
	Old_Line_Number_To_New_Labels := make(map[int]string)

	for i := 0; i < label_total; i++ {
		var tmp_delta = 0 
		if tmp[(*leader)[i]].Op!="label" {

			instr := new(model.Instr_struct)
			model.Initialize_instr(instr, "label", " ", "label" + strconv.Itoa(label_count), "", "-3")
			*instructions=append(*instructions,instr)
			Old_Line_Number_To_New_Labels[(*leader)[i]] = "label" + strconv.Itoa(label_count)

			label_count++
			tmp_delta++
		}else{
			Old_Line_Number_To_New_Labels[(*leader)[i]] = tmp[(*leader)[i]].Src1
		}
		for j := (*leader)[i]; j < (*leader)[i+1]; j++ {
			*instructions=append(*instructions,tmp[j])
			//fmt.Println(j, (*leader)[i+1])
		}

		(*leader)[i]+=delta
		delta+=tmp_delta
	}
	
	(*leader)[label_total]+=delta
	end := new(model.Instr_struct)
	model.Initialize_instr(end, "label", " ", "end", "", "-3")
	*instructions=append(*instructions,end)
	
	for key := range *instructions {
		if (*instructions)[key].Op == "ifgoto" {
			tmp_jump,_ := strconv.Atoi((*instructions)[key].Jmp)
			(*instructions)[key].Jmp = Old_Line_Number_To_New_Labels[tmp_jump-1]
		}
	}

	defer file.Close()
}


