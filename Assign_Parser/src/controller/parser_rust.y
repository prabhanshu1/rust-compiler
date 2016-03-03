
%{
package main
import "fmt"
//import "math"
import "os"
var line = 0

func err_neg(a int64) {
	if a<0 {
		fmt.Println("IT has to be non-negative")
	}
}

func btoi(a bool)int64 {
	if a==false {
		return 0
	}
	return 1
}

func itob(a int64)bool {
	if a==0 {
		return false
	}
	return true
}

type node struct {
    Token string
    Value string
    children []int
}

var tree []node

func make_node(n node) int {
	tree = append(tree,n)
	own_number := len(tree) - 1
/*	for i,_ := range tree[own_number].children{
		tree[i].Parent = own_number
	}*/
	return own_number
}

func writer(n int,fo *os.File) {
	if _, err := fo.Write([]byte("{")); err != nil {
	            panic(err)
	        }

	if _, err := fo.Write([]byte("\"name\": \"" + tree[n].Token + " " + tree[n].Value + "\"" )); err != nil {
	            panic(err)
	        }	

	if len(tree[n].children)!=0 {
		if _, err := fo.Write([]byte(",\"children\": [")); err != nil {
	            panic(err)
	        }
	    }

	for ii,i := range tree[n].children{
		writer(i,fo)
		if ii!=(len(tree[n].children)-1) {
			if _, err := fo.Write([]byte(",")); err != nil {
	            panic(err)
	        }
		}
		if _, err := fo.Write([]byte("\n")); err != nil {
	            panic(err)
	        }
	}

	if len(tree[n].children)!=0{
		if _, err := fo.Write([]byte("]")); err != nil {
	            panic(err)
	        }
	    }

	if _, err := fo.Write([]byte("}")); err != nil {
	            panic(err)
	        }
	
}

func make_json(n int) {
	fo, err := os.Create("code.json")
    if err != nil {
        panic(err)
    }
    
    writer(n,fo)
    // close fo on exit and check for its returned error
  
    if err := fo.Close(); err != nil {
        panic(err)
    }
   

	
}



%}



%union {
  nn int // node_number
  n int
  n64 int64
  f64 float64
  s string
  b bool
}

%token KEYWORD
%token VAR_TYPE
%token INTEGER
%token OPEQ_INTEGER
%token HEX
%token OCTAL
%token BINARY
%token FLOAT
%token OPEQ_FLOAT
%token LITERAL
%token OP_EQ
%token OP_RSHIFT
%token OP_LSHIFT
%token OP_ADDEQ
%token OP_SUBEQ
%token OP_MULEQ
%token OP_DIVEQ
%token OP_MODEQ
%token OP_INSIDE
%token OP_EQEQ
%token OP_NOTEQ
%token OP_ANDAND
%token OP_OROR
%token OP_POWER
%token OP_DOTDOT
%token OP_SUB
%token OP_ADD
%token OP_AND
%token OP_OR
%token OP_XOR
%token OP_FSLASH
%token OP_NOT
%token OP_COLUMN
%token OP_MUL
%token OP_GTHAN
%token OP_LTHAN
%token OP_MOD
%token OP_EQ
%token OP_DOT
%token OP_APOSTROPHE
%token SYM_COLCOL
%token SYM_HASH
%token SYM_OPEN_SQ
%token SYM_CLOSE_SQ
%token SYM_OPEN_ROUND
%token SYM_CLOSE_ROUND
%token SYM_OPEN_CURLY
%token SYM_CLOSE_CURLY
%token SYM_COMMA
%token SYM_SEMCOL
%token IDENTIFIER
%token FINISH

/*%expect 0

// fake-precedence symbol to cause '|' bars in lambda context to parse
// at low precedence, permit things like |x| foo = bar, where '=' is
// otherwise lower-precedence than '|'. Also used for proc() to cause
// things like proc() a + b to parse as proc() { a + b }.
%precedence LAMBDA

%precedence SELF

// MUT should be lower precedence than IDENT so that in the pat rule,
// "& MUT pat" has higher precedence than "binding_mode ident [@ pat]"
%precedence MUT

// IDENT needs to be lower than '{' so that 'foo {' is shifted when
// trying to decide if we've got a struct-construction expr (esp. in
// contexts like 'if foo { .')
//
// IDENT also needs to be lower precedence than '<' so that '<' in
// 'foo:bar . <' is shifted (in a trait reference occurring in a
// bounds list), parsing as foo:(bar<baz>) rather than (foo:bar)<baz>.
%precedence IDENTIFIER

// A couple fake-precedence symbols to use in rules associated with +
// and < in trailing type contexts. These come up when you have a type
// in the RHS of operator-AS, such as "foo as bar<baz>". The "<" there
// has to be shifted so the parser keeps trying to parse a type, even
// though it might well consider reducing the type "bar" and then
// going on to "<" as a subsequent binop. The "+" case is with
// trailing type-bounds ("foo as bar:A+B"), for the same reason.
%precedence SHIFTPLUS

%precedence MOD_SEP
%precedence RARROW ':'

// In where clauses, "for" should have greater precedence when used as
// a higher ranked constraint than when used as the beginning of a
// for_in_type (which is a ty)
%precedence FORTYPE
%precedence FOR

// Binops & unops, and their precedences
%precedence BOX
%precedence BOXPLACE
%nonassoc DOTDOT

// RETURN needs to be lower-precedence than tokens that start
// prefix_exprs
%precedence RETURN

%right '=' SHLEQ SHREQ MINUSEQ ANDEQ OREQ PLUSEQ STAREQ SLASHEQ CARETEQ PERCENTEQ
%right LARROW
%left OROR
%left ANDAND
%left EQEQ NE
%left '<' '>' LE GE
%left '|'
%left '^'
%left '&'
%left SHL SHR
%left '+' '-'
%precedence AS
%left '*' '/' '%'
%precedence '!'

%precedence '{' '[' '(' '.'

%precedence RANGE*/
%%

Code : Statements  {$$.nn=make_node(node{"Code","",[]int{$1.nn}});make_json($$.nn);}

;

Statements : expr Statements  {$$.nn=make_node(node{"Statements","",[]int{$1.nn,$2.nn}});}
| Decl_stmt Statements  {$$.nn=make_node(node{"Statements","",[]int{$1.nn,$2.nn}});}
| FINISH {$$.nn=make_node(node{"Statements","",[]int{make_node(node{"FINISH","",[]int{}})}})}
;

Decl_stmt : item {???????????}
| let_decl {$$.nn=make_node(node{"Decl_stmt","",[]int{$1.nn}})}
;

let_decl : item {?????????????????}
;

expr : arith {$$.nn=make_node(node{"expr","",[]int{$1.nn}}	)}
;

arith : IDENTIFIER OP_ADD arith {$$.nn=make_node(node{"arith","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{"OP_ADD","+",[]int{}}),$3.nn}});}
| IDENTIFIER SYM_SEMCOL {$$.nn=make_node(node{"arith","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{"SYM_SEMCOL",";",[]int{}})}});fmt.Println("BBBBBBBBBB",$1.s)}
;


%%