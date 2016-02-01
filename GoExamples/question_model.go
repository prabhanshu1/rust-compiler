package question_model

import "reflect"
//import "labix.org/v2/mgo/bson"
//this is models file for qustion types

// types as saved in Database
type Question struct{
	Text string
	Images []string
}
func (raw *Question) Validate() string{
	if reflect.TypeOf(raw.Text).String()!="string"{
		return "Question Text should be a string"
	}
	if reflect.TypeOf(raw.Images).String()!="[]string"{
		return "Question Images should be an array of strings"
	}
	return ""
}
type Option struct{
	Text string
	Images []string
}
func (raw *Option) Validate() string{
	if reflect.TypeOf(raw.Text).String()!="string"{
		return "Options Text should be a string"
	}
	if reflect.TypeOf(raw.Images).String()!="[]string"{
		return "Options Images should be an array of strings"
	}
	return ""
}
type Solution struct{
	Soltext string
	Solimages []string
	Solvedios []string	
}
func (raw *Solution) Validate() string{
	if reflect.TypeOf(raw.Soltext).String()!="string"{
		return "Solution Text should be a string"
	}
	if reflect.TypeOf(raw.Solimages).String()!="[]string"{
		return "Solution Images should be an array of strings"
	}
	if reflect.TypeOf(raw.Solvedios).String()!="[]string"{
		return "Solution Vedios should be an array of strings"
	}
	return ""
}
type Isused struct{
	Attempt bool
}
type Info struct{
	Qid int
	Lo_id []int
	Languages []string
	Type int
	Rating float64
	Attempt bool
}
func (raw *Info) Validate() string{
	if reflect.TypeOf(raw.Languages).String()!="[]string"{
		return "Languages List should be an array of strings"
	}
	if reflect.TypeOf(raw.Attempt).String()!="bool"{
		return "Attempt value should be a boolean"
	}
	if reflect.TypeOf(raw.Qid).String()!="int"&&reflect.TypeOf(raw.Qid).String()!="int32"&&reflect.TypeOf(raw.Qid).String()!="int64" {
		return "Question ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_id).String()!="[]int"&&reflect.TypeOf(raw.Lo_id).String()!="[]int32"&&reflect.TypeOf(raw.Lo_id).String()!="[]int64" {
		return "Learning Objective ID should be an array of Integers"
	}
	if reflect.TypeOf(raw.Type).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64" {
		return "Question Type should be an Integer"
	}
	if raw.Type<1||raw.Type>7{
		return "Question Type should range from 1 to 7"
	}
	if reflect.TypeOf(raw.Rating).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64"&&reflect.TypeOf(raw.Rating).String()!="float"&&reflect.TypeOf(raw.Rating).String()!="float32"&&reflect.TypeOf(raw.Rating).String()!="float64" {
		return "Question Rating should be a float"
	}
	if raw.Rating<0||raw.Type>10{
		return "Question Type should range from 0 to 10"
	}
	return ""
}
type Scoring struct{
	Type int
	Onwrong float64
	Oncorrect float64
}
func (raw *Scoring) Validate() string{
	if reflect.TypeOf(raw.Onwrong).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64"&&reflect.TypeOf(raw.Onwrong).String()!="float"&&reflect.TypeOf(raw.Onwrong).String()!="float32"&&reflect.TypeOf(raw.Onwrong).String()!="float64" {
		return "Score On Wrong should be a float"
	}
	if reflect.TypeOf(raw.Oncorrect).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64"&&reflect.TypeOf(raw.Oncorrect).String()!="float"&&reflect.TypeOf(raw.Oncorrect).String()!="float32"&&reflect.TypeOf(raw.Oncorrect).String()!="float64" {
		return "Score On Correct should be a float"
	}
	if reflect.TypeOf(raw.Type).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64" {
		return "Scoring Type should be an Integer"
	}
	return ""
}	
//single correct
type Data1 struct{
	Que Question
	Opt []Option
	Sol Solution
}
func (raw *Data1) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Opt {
    	str=raw.Opt[key].Validate()
    	if str!="" {
		return str
		}
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	return ""
}
type Qtype1 struct{
	Inf Info 
	Data map[string]Data1 
	Ans int
	Sco Scoring
}
func (raw *Qtype1) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp:=raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	if reflect.TypeOf(raw.Ans).String()!="int"&&reflect.TypeOf(raw.Ans).String()!="int32"&&reflect.TypeOf(raw.Ans).String()!="int64" {
		return "Answer should be an Integer"
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}
//multiple correct
type Data2 struct{
	Que Question
	Opt []Option
	Sol Solution
}
func (raw *Data2) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Opt {
    	str=raw.Opt[key].Validate()
    	if str!="" {
		return str
		}
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	return ""
}
type Qtype2 struct{
	Inf Info 
	Data map[string]Data2
	Ans []int
	Sco Scoring
}
func (raw *Qtype2) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp:=raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	if reflect.TypeOf(raw.Ans).String()!="[]int"&&reflect.TypeOf(raw.Ans).String()!="[]int32"&&reflect.TypeOf(raw.Ans).String()!="[]int64" {
		return "Answer should be an array of Integers"
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}
//matrix match
type Data3 struct{
	Que Question
	Opl []Option
	Opr []Option
	Sol Solution
}
func (raw *Data3) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Opl {
    	str=raw.Opl[key].Validate()
    	if str!="" {
		return str
		}
	}
	for key,_ := range raw.Opr {
    	str=raw.Opr[key].Validate()
    	if str!="" {
		return str
		}
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	return ""
}
type Qtype3 struct{
	Inf Info
	Data map[string]Data3
	Ans [][]int
	Sco Scoring
}
func (raw *Qtype3) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp := raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	if reflect.TypeOf(raw.Ans).String()!="[][]int"&&reflect.TypeOf(raw.Ans).String()!="[][]int32"&&reflect.TypeOf(raw.Ans).String()!="[][]int64" {
		return "Answer should be a 2-d array of Integers"
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}
//integer type
type Data4 struct{
	Que Question
	Sol Solution
}
func (raw *Data4) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	return ""
}
type Qtype4 struct{
	Inf Info
	Data map[string]Data4
	Ans int
	Sco Scoring
}
func (raw *Qtype4) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp :=raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	if reflect.TypeOf(raw.Ans).String()!="int"&&reflect.TypeOf(raw.Ans).String()!="int32"&&reflect.TypeOf(raw.Ans).String()!="int64" {
		return "Answer should be an Integer"
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}


//puzzels
type Data6 struct{
	Tit string
	Que Question
	Sol Solution
}
func (raw *Data6) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	if reflect.TypeOf(raw.Tit).String()!="string"{
		return "Title should be a string"
	}
	return ""
}
type Qtype6 struct{
	Inf Info
	Data map[string]Data6
	Ans int
	Sco Scoring
}
func (raw *Qtype6) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp := raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	if reflect.TypeOf(raw.Ans).String()!="int"&&reflect.TypeOf(raw.Ans).String()!="int32"&&reflect.TypeOf(raw.Ans).String()!="int64" {
		return "Answer should be an Integer"
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}
//programming
type Data7 struct{
	Tit string
	Que Question
	Inp string
	Out string
	Con string
	Sin string
	Sou string
	Sex string
	Sol Solution
}
func (raw *Data7) Validate() string{
	str:=raw.Que.Validate()
	if str!="" {
		return str
	}
	str =raw.Sol.Validate()
	if str!="" {
		return str
	}
	if reflect.TypeOf(raw.Tit).String()!="string"{
		return "Title should be a string"
	}
	if reflect.TypeOf(raw.Inp).String()!="string"{
		return "Input Format should be a string"
	}
	if reflect.TypeOf(raw.Out).String()!="string"{
		return "Output Format should be a string"
	}
	if reflect.TypeOf(raw.Con).String()!="string"{
		return "Constraints should be a string"
	}
	if reflect.TypeOf(raw.Sin).String()!="string"{
		return "Sample Input should be a string"
	}
	if reflect.TypeOf(raw.Sou).String()!="string"{
		return "Sample Output be a string"
	}
	if reflect.TypeOf(raw.Sex).String()!="string"{
		return "Sample Explaination be a string"
	}
	return ""
}
type Qtype7 struct{
	Inf Info
	Data map[string]Data7
	Sco Scoring
}
func (raw *Qtype7) Validate() string{
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	for key,_ := range raw.Data {
		temp := raw.Data[key]
    	str=temp.Validate()
    	if str!="" {
			return str
		}
	}
	str =raw.Sco.Validate()
	if str!="" {
		return str
	}
	return ""
}

type Qtype5 struct{
	Inf Info
	Paragraph string
	Q1 []Qtype1
	Q2 []Qtype2
	Q3 []Qtype3
	Q4 []Qtype4
	Q6 []Qtype6
	Q7 []Qtype7
}
func (raw *Qtype5) Validate() string{
	if reflect.TypeOf(raw.Paragraph).String()!="string"{
		return "Paragraph should be a string"
	}
	str:=raw.Inf.Validate()
	if str!="" {
		return str
	}
	if raw.Inf.Type==1 {
		for key,_ := range raw.Q1{
			return raw.Q1[key].Validate()
		}
	}
	if raw.Inf.Type==2 {
		for key,_ := range raw.Q2{
			return raw.Q2[key].Validate()
		}
	}
	if raw.Inf.Type==3 {
		for key,_ := range raw.Q3{
			return raw.Q3[key].Validate()
		}
	}
	if raw.Inf.Type==4 {
		for key,_ := range raw.Q4{
			return raw.Q4[key].Validate()
		}
	}
	if raw.Inf.Type==6 {
		for key,_ := range raw.Q6{
			return raw.Q6[key].Validate()
		}
	}
	if raw.Inf.Type==7 {
		for key,_ := range raw.Q7{
			return raw.Q7[key].Validate()
		}
	}
	if raw.Inf.Type<1||raw.Inf.Type>7||raw.Inf.Type==5 {
		return "Question Type should range from 1 to 7 excluding 5"
	}
	return ""
}
type  Superquestion struct{
	Exam int
	Type int
	Q1 Qtype1
	Q2 Qtype2
	Q3 Qtype3
	Q4 Qtype4
	Q5 Qtype5
	Q6 Qtype6
	Q7 Qtype7
}
func (raw *Superquestion) Validate() string {
	if reflect.TypeOf(raw.Exam).String()!="int"&&reflect.TypeOf(raw.Exam).String()!="int32"&&reflect.TypeOf(raw.Exam).String()!="int64" {
		return "Exam ID should be an Integer"
	}
	if reflect.TypeOf(raw.Type).String()!="int"&&reflect.TypeOf(raw.Type).String()!="int32"&&reflect.TypeOf(raw.Type).String()!="int64" {
		return "Type should be an Integer"
	}
	if raw.Type==1 {
		return raw.Q1.Validate()
	}
	if raw.Type==2 {
		return raw.Q2.Validate()
	}
	if raw.Type==3 {
		return raw.Q3.Validate()
	}
	if raw.Type==4 {
		return raw.Q4.Validate()
	}
	if raw.Type==5 {
		return raw.Q5.Validate()
	}
	if raw.Type==6 {
		return raw.Q6.Validate()
	}
	if raw.Type==7 {
		return raw.Q7.Validate()
	}
	if raw.Type<1||raw.Type>7 {
		return "Question Type should range from 1 to 7"
	}
	return ""
}

/*type NewQuestion struct{
	Exam int
	Inf Info
	Question interface{}
}*/
//types as used as inputs in queries
type Qsolved struct{
	Lo_id int
	Qid int
	Userid int
	Language string
	Correctness float64
	Timespent float64
	Exam int
	Marked_Answer []int
	Marked_Answer_Matrix [][]int
	Lo_status int
	Qcount int
	Unattemptedcount int
	Parents []int
	Dep_Lo []int
	Dep_Count []int
	Custom int
}
type Qsolvedcustom struct{
	Lo_id int
	Qid int
	Userid int
	Language string
	Correctness float64
	Timespent float64
	Exam int
	Marked_Answer []int
	Lo_status int
	Qcount int
	Unattemptedcount int
	Lo_List []int
	Parents []int
	Dep_Lo []int
	Dep_Count []int
}
//types as used as outputs in queries
type NextQuestion struct{
	Lo_id int
	Qid int
	Userid int
	Status int
}

//APIquery structure
type ApiQueryData struct{
	Origin string
	Setid int
	Qid int
	Correctness float64
	Userid int
	Timespent float64
	Type int
	Token string
	Qcount int
}
type ApiQueryCustomData struct{
	Origin string
	Setid int
	Qid int
	Correctness float64
	Userid int
	Timespent float64
	Type int
	Token string
	Qcount int
	List []int
}
//APIoutput structure
type ApiOutputData struct{
	Userid int 
	Setid int 
	Qid int 
	Rating float64 //user Rating
}

//Answer_Marked_Attribute
type Marked_Info struct{
	Userid int
	Lo_id int
	Attempts int
	Type int
}
type Answer_Marked_Attribute_Question struct{
	Qid int
	Correctness float64
	Timespent float64
	Marked_Answer []string
}
type Attempt struct{
	Que []Answer_Marked_Attribute_Question
}
type Answer_Marked_Attribute struct{
	Inf Marked_Info
	Att []Attempt
}

//models for answer marked queries
type Attempt_Data_Query struct{
	Exam int
	Userid int
	Lo_id int
	Attemptno int
}

//requests for question add
type ApiQuestionAddData struct{
	Origin string
	Qid int
	Setid int
	Rating float64
	Token string
}
type Qclear struct{
	Lo_id int
	Userid int	
	Exam int
}
func (raw *Qclear) Validate() string{
	if reflect.TypeOf(raw.Userid).String()!="int"&&reflect.TypeOf(raw.Userid).String()!="int32"&&reflect.TypeOf(raw.Userid).String()!="int64" {
		return "User ID should be an Integer"
	}
	if reflect.TypeOf(raw.Exam).String()!="int"&&reflect.TypeOf(raw.Exam).String()!="int32"&&reflect.TypeOf(raw.Exam).String()!="int64" {
		return "Exam ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_id).String()!="int"&&reflect.TypeOf(raw.Lo_id).String()!="int32"&&reflect.TypeOf(raw.Lo_id).String()!="int64" {
		return "Learning Objective ID should be an Integer"
	}
	return ""
}
type ApiQclearData struct{
	Origin string
	Setid int
	Qid int
	Userid int
	Token string
}

type Apiuseradd struct{
	Origin string
	Userid int
	Setid int
	Rating float64
	Mode int
	Token string
	SortedRating float64
}
type Reply struct{
	Question interface{}
	Lo_status int
	Unlocked []int
}
func (raw *Qsolved) Validate() string{
	if reflect.TypeOf(raw.Userid).String()!="int"&&reflect.TypeOf(raw.Userid).String()!="int32"&&reflect.TypeOf(raw.Userid).String()!="int64" {
		return "User ID should be an Integer"
	}
	if reflect.TypeOf(raw.Exam).String()!="int"&&reflect.TypeOf(raw.Exam).String()!="int32"&&reflect.TypeOf(raw.Exam).String()!="int64" {
		return "Exam ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_id).String()!="int"&&reflect.TypeOf(raw.Lo_id).String()!="int32"&&reflect.TypeOf(raw.Lo_id).String()!="int64" {
		return "Learning Objective ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_status).String()!="int"&&reflect.TypeOf(raw.Lo_status).String()!="int32"&&reflect.TypeOf(raw.Lo_status).String()!="int64" {
		return "Learning Objective Status should be an Integer"
	}
	if reflect.TypeOf(raw.Qid).String()!="int"&&reflect.TypeOf(raw.Qid).String()!="int32"&&reflect.TypeOf(raw.Qid).String()!="int64" {
		return "Question ID should be an Integer"
	}
	if reflect.TypeOf(raw.Parents).String()!="[]int"&&reflect.TypeOf(raw.Parents).String()!="[]int32"&&reflect.TypeOf(raw.Parents).String()!="[]int64" {
		return "Parents List should be an Integer array"
	}
	if reflect.TypeOf(raw.Dep_Lo).String()!="[]int"&&reflect.TypeOf(raw.Dep_Lo).String()!="[]int32"&&reflect.TypeOf(raw.Dep_Lo).String()!="[]int64" {
		return "Dependent Lo List should be an Integer array"
	}
	if reflect.TypeOf(raw.Dep_Count).String()!="[]int"&&reflect.TypeOf(raw.Dep_Count).String()!="[]int32"&&reflect.TypeOf(raw.Dep_Count).String()!="[]int64" {
		return "Dependent On Count list should be an Integer array"
	}
	if reflect.TypeOf(raw.Qcount).String()!="int"&&reflect.TypeOf(raw.Qcount).String()!="int32"&&reflect.TypeOf(raw.Qcount).String()!="int64" {
		return "Count of Questions in given Learning Objective should be an Integer"
	}
	if reflect.TypeOf(raw.Unattemptedcount).String()!="int"&&reflect.TypeOf(raw.Unattemptedcount).String()!="int32"&&reflect.TypeOf(raw.Unattemptedcount).String()!="int64" {
		return "Count of Unattempted Questions in given Learning Objective should be an Integer"
	}
	if raw.Unattemptedcount<0 {
		return "Count of Unattempted Questions in given Learning Objective should be non-negative"
	}
	if reflect.TypeOf(raw.Language).String()!="string" {
		return "Language name should be a string"
	}
	if reflect.TypeOf(raw.Correctness).String()!="float64"&&reflect.TypeOf(raw.Correctness).String()!="float32"&&reflect.TypeOf(raw.Correctness).String()!="int"&&reflect.TypeOf(raw.Correctness).String()!="int32"&&reflect.TypeOf(raw.Correctness).String()!="int64" {
		return "Score should be a float"
	}
	if reflect.TypeOf(raw.Timespent).String()!="float64"&&reflect.TypeOf(raw.Timespent).String()!="float32"&&reflect.TypeOf(raw.Timespent).String()!="int"&&reflect.TypeOf(raw.Timespent).String()!="int32"&&reflect.TypeOf(raw.Timespent).String()!="int64" {
		return "Time spent should be a float"
	}
	if reflect.TypeOf(raw.Marked_Answer).String()!="[]int"&&reflect.TypeOf(raw.Marked_Answer).String()!="[]int32"&&reflect.TypeOf(raw.Marked_Answer).String()!="[]int64" {
		return "Marked Answers should be array of integers"
	}
	if reflect.TypeOf(raw.Marked_Answer_Matrix).String()!="[][]int"&&reflect.TypeOf(raw.Marked_Answer_Matrix).String()!="[][]int32"&&reflect.TypeOf(raw.Marked_Answer).String()!="[][]int64" {
		return "Marked Answers should be a 2-d array of integers"
	}
	if len(raw.Marked_Answer) <1&&len(raw.Marked_Answer_Matrix)<1{
		return "Please mark atleast one answer"
	}
	if raw.Qcount <0{
		return "Count of Questions in given Learning Objective should be non-negative"
	}
	if raw.Correctness>1.0003||raw.Correctness<(-0.0003) {
		return "Score should range from 0 to 1"
	}
	return ""
}
func (raw *Qsolvedcustom) Validate() string{
	if reflect.TypeOf(raw.Userid).String()!="int"&&reflect.TypeOf(raw.Userid).String()!="int32"&&reflect.TypeOf(raw.Userid).String()!="int64" {
		return "User ID should be an Integer"
	}
	if reflect.TypeOf(raw.Exam).String()!="int"&&reflect.TypeOf(raw.Exam).String()!="int32"&&reflect.TypeOf(raw.Exam).String()!="int64" {
		return "Exam ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_id).String()!="int"&&reflect.TypeOf(raw.Lo_id).String()!="int32"&&reflect.TypeOf(raw.Lo_id).String()!="int64" {
		return "Learning Objective ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_status).String()!="int"&&reflect.TypeOf(raw.Lo_status).String()!="int32"&&reflect.TypeOf(raw.Lo_status).String()!="int64" {
		return "Learning Objective Status should be an Integer"
	}
	if reflect.TypeOf(raw.Qid).String()!="int"&&reflect.TypeOf(raw.Qid).String()!="int32"&&reflect.TypeOf(raw.Qid).String()!="int64" {
		return "Question ID should be an Integer"
	}
	if reflect.TypeOf(raw.Parents).String()!="[]int"&&reflect.TypeOf(raw.Parents).String()!="[]int32"&&reflect.TypeOf(raw.Parents).String()!="[]int64" {
		return "Parents List should be an Integer array"
	}
	if reflect.TypeOf(raw.Lo_List).String()!="[]int"&&reflect.TypeOf(raw.Lo_List).String()!="[]int32"&&reflect.TypeOf(raw.Lo_List).String()!="[]int64" {
		return "Included Lo List should be an Integer array"
	}
	if reflect.TypeOf(raw.Dep_Lo).String()!="[]int"&&reflect.TypeOf(raw.Dep_Lo).String()!="[]int32"&&reflect.TypeOf(raw.Dep_Lo).String()!="[]int64" {
		return "Dependent Lo List should be an Integer array"
	}
	if reflect.TypeOf(raw.Dep_Count).String()!="[]int"&&reflect.TypeOf(raw.Dep_Count).String()!="[]int32"&&reflect.TypeOf(raw.Dep_Count).String()!="[]int64" {
		return "Dependent On Count list should be an Integer array"
	}
	if reflect.TypeOf(raw.Qcount).String()!="int"&&reflect.TypeOf(raw.Qcount).String()!="int32"&&reflect.TypeOf(raw.Qcount).String()!="int64" {
		return "Count of Questions in given Learning Objective should be an Integer"
	}
	if reflect.TypeOf(raw.Unattemptedcount).String()!="int"&&reflect.TypeOf(raw.Unattemptedcount).String()!="int32"&&reflect.TypeOf(raw.Unattemptedcount).String()!="int64" {
		return "Count of Unattempted Questions in given Learning Objective should be an Integer"
	}
	if raw.Unattemptedcount<0 {
		return "Count of Unattempted Questions in given Learning Objective should be non-negative"
	}
	if reflect.TypeOf(raw.Language).String()!="string" {
		return "Language name should be a string"
	}
	if reflect.TypeOf(raw.Correctness).String()!="float64"&&reflect.TypeOf(raw.Correctness).String()!="float32"&&reflect.TypeOf(raw.Correctness).String()!="int"&&reflect.TypeOf(raw.Correctness).String()!="int32"&&reflect.TypeOf(raw.Correctness).String()!="int64" {
		return "Score should be a float"
	}
	if reflect.TypeOf(raw.Timespent).String()!="float64"&&reflect.TypeOf(raw.Timespent).String()!="float32"&&reflect.TypeOf(raw.Timespent).String()!="int"&&reflect.TypeOf(raw.Timespent).String()!="int32"&&reflect.TypeOf(raw.Timespent).String()!="int64" {
		return "Time spent should be a float"
	}
	if reflect.TypeOf(raw.Marked_Answer).String()!="[]int"&&reflect.TypeOf(raw.Marked_Answer).String()!="[]int32"&&reflect.TypeOf(raw.Marked_Answer).String()!="[]int64" {
		return "Marked Answers should be array of Integers"
	}
	if len(raw.Marked_Answer) <1{
		return "Please mark atleast one answer"
	}
	if raw.Qcount <1{
		return "Count of Questions in given Learning Objective should be positive"
	}
	if raw.Correctness>1.0003||raw.Correctness<(-0.0003) {
		return "Score should range from 0 to 1"
	}
	return ""
}
func (raw *Attempt_Data_Query) Validate() string {
	if reflect.TypeOf(raw.Userid).String()!="int" {
		return "User ID should be an Integer"
	}
	if reflect.TypeOf(raw.Exam).String()!="int" {
		return "Exam ID should be an Integer"
	}
	if reflect.TypeOf(raw.Lo_id).String()!="int" {
		return "Learning Objective ID should be an Integer"
	}
	if reflect.TypeOf(raw.Attemptno).String()!="int" {
		return "Attempt No. should be an Integer"
	}
	if raw.Attemptno<1 {
		return "Attempts Count should be positive"
	}
	return ""
}
// type Attempt_Data_Query struct{
// 	Exam int
// 	Userid int
// 	Lo_id int
// 	Attemptno int
// }
