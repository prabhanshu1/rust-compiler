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

func space(a string,i int)int{
for ;a[i]==' ';i++{}
return i
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


//%debug

%token MUT 
%token IDENTIFIER
%token OP_INSIDE
%token FOR
%token RETURN
%token AS
%token SYM_OPEN_SQ
%token SYM_CLOSE_SQ
%token SYM_OPEN_ROUND
%token SYM_CLOSE_ROUND
%token SYM_OPEN_CURLY
%token SYM_CLOSE_CURLY
%token ANDAND
%token BINOPEQ
%token DOTDOT
%token DOTDOTDOT
%token EQEQ
%token FAT_ARROW
%token GE
%token LE
%token LIFETIME
%token LIT_CHAR
%token FLOAT //
%token LIT_FLOAT_UNSUFFIXED
%token LIT_INT_UNSUFFIXED
%token LITERAL_STR
%token LITERAL_CHAR
%token LIT_UINT
%token MOD_SEP
%token NE
%token OROR
%token SHL
%token SHR
%token UNDERSCORE

%token KEYWORD
%token VAR_TYPE
%token LIT_INT
%token OPEQ_INT
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
%token OP_DOTDOTDOT
%token OP_SUB
%token OP_ADD
%token OP_AND
%token OP_LEQ
%token OP_GEQ
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
%token OP_FAT_ARROW
%token SYM_COLCOL
%token SYM_HASH
%token SYM_COMMA
%token SYM_SEMCOL
%token IDENTIFIER
%token FINISH
%token NEWLINE

%token ABSTRACT 
%token ALIGNOF 
%token AS 
%token BECOME 
%token BOX 
%token BREAK 
%token CONST 
%token CONTINUE 
%token CRATE 
%token DO 
%token ELSE 
%token ENUM 
%token EXTERN 
%token FALSE 
%token FINAL 
%token FN 
%token FOR 
%token IF 
%token IMPL 
%token IN 
%token LET 
%token LOOP 
%token MACRO 
%token MATCH 
%token MOD 
%token MOVE 
%token OFFSETOF 
%token OVERRIDE 
%token PRIV 
%token PROC 
%token PUB 
%token PURE 
%token REF 
%token RETURN 
%token SELF 
%token SELF 
%token SIZEOF 
%token STATIC 
%token STRUCT 
%token SUPER 
%token TRAIT 
%token TRUE 
%token TYPE 
%token TYPEOF 
%token UNSAFE 
%token UNSIZED 
%token USE 
%token VIRTUAL 
%token WHERE 
%token WHILE 
%token YIELD 
%token PRINTLN 
%token MACRO_RULES
// keywords

%token SHEBANG
%token STATIC_LIFETIME



 //%expect 0

%nonassoc CONTINUE
%nonassoc IDENTIFIER
%nonassoc SYM_OPEN_ROUND
%nonassoc SYM_OPEN_CURLY











// In where clauses, "for" should have greater precedence when used as
// a higher ranked constraint than when used as the beginning of a
// for_in_type (which is a ty)


// Binops & unops, and their precedences


// RETURN needs to be lower-precedence than tokens that start
// prefix_exprs

%right '=' OP_SHLEQ OP_SHREQ  OP_ADDEQ OP_OREQ OP_ANDEQ  OP_MULEQ OP_DIVEQ OP_MODEQ OP_XOREQ
%left OP_OROR
%left OP_ANDAND
%left OP_EQEQ OP_NOTEQ
%left '<' '>' OP_GEQ OP_LEQ
%left '|'
%left '^'
%left '&'
%left OP_LSHIFT OP_RSHIFT
%left '+' '-'

%left '*' '/' '%'




%start rust

%%

/// println, print macro support => standard macros
rust
:STRUCT IDENTIFIER struct_expr
|item_or_view_item
;


item_or_view_item
: item_fn
;

item_fn
: FN IDENTIFIER fn_decl inner_attrs_and_block  {fmt.Println("REACHING fn")}
;

fn_decl
: fn_args ret_ty {fmt.Println("REACHING fn_decl")}
;

fn_args
: SYM_OPEN_ROUND maybe_args_general SYM_CLOSE_ROUND	{fmt.Println("REACHING fn_args")}
;

maybe_args_general
: args_general
| /* empty */		{fmt.Println("REACHING maybe_args_general")}
;

args_general
: arg_general
| args_general ',' arg_general
;

arg_general
: pat ':' ty
;

ret_ty
: OP_INSIDE '!'
| OP_INSIDE ty
| OP_INSIDE SYM_OPEN_ROUND SYM_CLOSE_ROUND
| /* empty */
;

inner_attrs_and_block
: SYM_OPEN_CURLY maybe_inner_attrs maybe_stmts SYM_CLOSE_CURLY   {fmt.Println("REACHING inner_attrs_and_block")}
;

maybe_inner_attrs
: inner_attrs {fmt.Println("REACHING maybe_inner_attrs")}
| /* empty */ {fmt.Println("REACHING maybe_inner_attrs2")}
;

inner_attrs
: inner_attr {fmt.Println("REACHING inner_attrs")}
| inner_attrs inner_attr {fmt.Println("REACHING inner_attrs2")}
;

inner_attr
: SHEBANG '[' meta_item ']' {fmt.Println("REACHING inner_attr")}
;


meta_item	
: IDENTIFIER
| IDENTIFIER '=' lit
| IDENTIFIER SYM_OPEN_ROUND meta_seq SYM_CLOSE_ROUND
;

meta_seq
: meta_item
| meta_seq ',' meta_item
;

maybe_mod_items
: mod_items
| /* empty */
;

mod_items
: mod_item
| mod_items mod_item
;

mod_item
: maybe_outer_attrs item_or_view_item    { $$ = $2; }
;


maybe_outer_attrs
: outer_attrs
| /* empty */
;

outer_attrs
: outer_attr
| outer_attrs outer_attr
;

outer_attr
: '#' '[' meta_item ']'
;

lit
: LIT_CHAR
| LIT_INT {fmt.Println("REACHING LIT_INT")}
| LIT_UINT
| LIT_INT_UNSUFFIXED
| FLOAT
| LIT_FLOAT_UNSUFFIXED
| LITERAL_STR
| LITERAL_CHAR
| TRUE
| FALSE
;

maybe_stmts
: stmts {fmt.Println("REACHING maybe_stmts1")}
| /* empty */ {fmt.Println("REACHING maybe_stmts2")}
;

stmts
: stmts stmt 
| stmt 
;

stmt
: let ';' {fmt.Println("REACHING LET")} 
| item_or_view_item ';'
| expr_stmt     
| expr ';'
;

// Things that can be an expr or a stmt, no semi required.
expr_stmt
: expr_match
| expr_if {fmt.Println("REACHING IF")}
| expr_while
| expr_loop
| expr_for
;

expr_match
: MATCH IDENTIFIER SYM_OPEN_CURLY match_clauses SYM_CLOSE_CURLY
| MATCH IDENTIFIER SYM_OPEN_CURLY match_clauses ',' SYM_CLOSE_CURLY
;

match_clauses
: match_clause
| match_clauses ',' match_clause
;

match_clause
: pats_or maybe_guard OP_FAT_ARROW match_body {fmt.Println("REACHING match_clause")}
;

match_body
: expr
| expr_stmt
;

maybe_guard
: IF expr
| // empty
;

expr_if
: IF exp block 
| IF exp block ELSE block_or_if
;

block_or_if
: block
| expr_if
;

block
: SYM_OPEN_CURLY maybe_stmts SYM_CLOSE_CURLY {fmt.Println("REDUCING TO BLOCK")}
;

expr_while
: WHILE exp block
;

expr_loop
: LOOP block
;

expr_for
: FOR exp IN exp block
| FOR exp IN range_di block
| FOR SYM_OPEN_ROUND maybe_assignment ';' exp ';' maybe_assignment SYM_CLOSE_ROUND block
;

let
: LET maybe_mut pat maybe_ty_ascription maybe_init_expr 
;

maybe_ty_ascription
: ':' ty
| /* empty */
;

maybe_init_expr
: '=' expr 
| OPEQ_INT  opeq_ops {fmt.Println("REACHING maybe_init_expr	")}
| OPEQ_FLOAT opeq_ops
| /* empty */
;

pats_or
: pat
| lit
| '_'
| range_tri
| pats_or '|' pat
| pats_or '|' lit
| pats_or '|' range_tri
;

range_tri
: LIT_INT OP_DOTDOTDOT LIT_INT
| LITERAL_CHAR OP_DOTDOTDOT LITERAL_CHAR

range_di
: LIT_INT OP_DOTDOT LIT_INT
| LITERAL_CHAR OP_DOTDOT LITERAL_CHAR


pat
: IDENTIFIER
;


tys
: ty
| tys ',' ty
;

ty
: path
| '~' ty
| '*' maybe_mut ty
| SYM_OPEN_ROUND tys SYM_CLOSE_ROUND
;

maybe_mut
: MUT
| /* empty */
;

path
: VAR_TYPE
;
maybe_exprs
: exprs
| /* empty */
;

exprs
: expr
| exprs ',' expr
;

//// $$$opeq+int doesn't work

maybe_assignment
: assignment
|
;

assignment
: IDENTIFIER '=' expr 
| IDENTIFIER OPEQ_INT opeq_ops
| IDENTIFIER OPEQ_FLOAT opeq_ops

;

opeq_ops
:  '+' expr     {fmt.Println("REACHING opeq_ops	")}  
|  '-' expr 
| '<' expr   
| '>' expr       
| OP_LSHIFT expr       
| OP_RSHIFT expr       
|
;

expr
: exp
| assignment
;

//$$struct remaining

exp
: lit 
| IDENTIFIER     {fmt.Println("REACHED IDENTIFIER in exp")}                       
//| IDENTIFIER struct_expr                
| exp '+' exp       
| exp '-' exp       
| exp '<' exp   
| exp '>' exp       
| exp OP_LSHIFT exp       
| exp OP_RSHIFT exp       
| exp SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND         
| CONTINUE                         
| CONTINUE IDENTIFIER                   
| UNSAFE block                     
| block                            
;


struct_expr
: SYM_OPEN_CURLY field_inits default_field_init SYM_CLOSE_CURLY
;

field_inits
: field_init
| field_inits ',' field_init
;

field_init
: maybe_mut IDENTIFIER ':' expr
;

default_field_init
: ','
| ',' OP_DOTDOT expr
| /* empty */
;
