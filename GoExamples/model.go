package model

import "reflect"

var Operator_Map =  map[string]int{
    "+": 1,
    "-": 2,
    "=": 3,
    "leq": 4,
    "geq": 5,
    "ifgoto": 6,
}


type Instruction struct{
	var Input1 string
	var Input2 string
	var Out string
	var Operator int
	var Jump string
}

/*
+,b,b,7
index 8
switch INS[8].Operator{

	INS[8].Operator_Map["+"] : fmt.printf(fd,"ADD ","$",getreg(INS[8].Out)," ","$",getreg(INS[8].Input1)," ")

}
*/


func (string *Input) Syntax() string{
		Type := reflect.TypeOf(Input)
		switch Type{
			case 
		}
	}