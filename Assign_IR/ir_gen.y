%{
package main
import "fmt"
import "log"
import "os"
var line = 0
func list_end(l node)node {
  while(l.next!=NULL){l=l.next;}
  return l;
}
  
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


type node struct {
  var value [100]char
  next *node
}

%union {
  code *node=NULL
  var map map[string]string =NULL
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
%token OP_ANDMUT
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
%token FINISH
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

%right '=' '!' OP_SHLEQ OP_SHREQ  OP_ADDEQ OP_OREQ OP_ANDEQ  OP_MULEQ OP_DIVEQ OP_MODEQ OP_XOREQ 
%left OP_EQEQ OP_NOTEQ
%left '<' '>' OP_GEQ OP_LEQ
%left '|' 
%left '^' 
%left '&' OP_ANDMUT
%left OP_LSHIFT OP_RSHIFT OP_ANDAND OP_OROR OP_POWER
%left '+' '-' '.'

%left '*' '/' '%' 




%start Code

%%

/// println, print macro support => standard macros

Code : rust   {$$.nn=make_node(node{"Code","",[]int{$1.nn}});make_json($$.nn);}
;

rust
: STRUCT IDENTIFIER struct_expr rust  {$$.nn=make_node(node{"rust","",[]int{make_node(node{"STRUCT","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}}),$3.nn,$4.nn}})}
| item_or_view_item rust  {$$.nn=make_node(node{"rust","",[]int{$1.nn,$2.nn}})}
| USE func_identifier ';' rust {$$.nn=make_node(node{"rust","",[]int{make_node(node{"USE","",[]int{}}),$2.nn,make_node(node{";","",[]int{}}),$4.nn}})}
|  {$$.nn=make_node(node{"rust","",[]int{}});}
;


item_or_view_item
: item_fn {$$.nn=make_node(node{"item_or_view_item","",[]int{$1.nn}})}
;

item_fn
: FN IDENTIFIER fn_decl inner_attrs_and_block  {$$.nn=make_node(node{"item_fn","",[]int{make_node(node{"FN","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}}),$3.nn,$4.nn}})}
;

fn_decl
: fn_args ret_ty {$$.nn=make_node(node{"fn_decl","",[]int{$1.nn,$2.nn}})}
;

fn_args
: SYM_OPEN_ROUND maybe_args_general SYM_CLOSE_ROUND	{$$.nn=make_node(node{"fn_args","",[]int{make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$2.nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
;

maybe_args_general
: args_general      {$$.nn=make_node(node{"maybe_args_general","",[]int{$1.nn}})}
| /* empty */       {$$.nn=make_node(node{"maybe_args_general","",[]int{}})}
;

args_general
: arg_general        {$$.nn=make_node(node{"args_general","",[]int{$1.nn}})}
| args_general ',' arg_general {$$.nn=make_node(node{"args_general","",[]int{$1.nn,make_node(node{",","",[]int{}}),$3.nn}})}
;

arg_general
: pat ':' ty  {$$.nn=make_node(node{"arg_general","",[]int{$1.nn,make_node(node{":","",[]int{}}),$3.nn}})}
;

ret_ty
: OP_INSIDE '!'  {$$.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",$1.s,[]int{}}),make_node(node{"!","",[]int{}})}})}
| OP_INSIDE ty  {$$.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",$1.s,[]int{}}),$2.nn}})}
| OP_INSIDE SYM_OPEN_ROUND  SYM_CLOSE_ROUND {$$.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",$1.s,[]int{}}),make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
| /* empty */ {$$.nn=make_node(node{"ret_ty","",[]int{}})}
;

inner_attrs_and_block
: SYM_OPEN_CURLY maybe_inner_attrs maybe_stmts SYM_CLOSE_CURLY   {$$.nn=make_node(node{"inner_attrs_and_block","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),$2.nn,$3.nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
;

maybe_inner_attrs
: inner_attrs {$$.nn=make_node(node{"maybe_inner_attrs","",[]int{$1.nn}})}
| /* empty */ {$$.nn=make_node(node{"maybe_inner_attrs","",[]int{}})}
;

inner_attrs
: inner_attr {$$.nn=make_node(node{"inner_attrs","",[]int{$1.nn}})}
| inner_attrs inner_attr {$$.nn=make_node(node{"inner_attrs","",[]int{$1.nn,$2.nn}})}
;

inner_attr
: SHEBANG '[' meta_item ']'  {$$.nn=make_node(node{"inner_attr","",[]int{make_node(node{"SHEBANG","",[]int{}}),make_node(node{"[","",[]int{}}),$3.nn,make_node(node{"]","",[]int{}})}})}
;


meta_item	
: IDENTIFIER {$$.nn=make_node(node{"meta_item","",[]int{$1.nn}})}
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
: LIT_CHAR  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_CHAR",$1.s,[]int{}})}})}} 
| LIT_INT  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_INT",$1.s,[]int{}})}})}} 
| LIT_UINT  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_UINT",$1.s,[]int{}})}})}}
| LIT_INT_UNSUFFIXED  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_INT_UNSUFFIXED",$1.s,[]int{}})}})}}
| FLOAT  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"FLOAT",$1.s,[]int{}})}})}}
| LIT_FLOAT_UNSUFFIXED  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_FLOAT_UNSUFFIXED",$1.s,[]int{}})}})}}
| LITERAL_STR  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LITERAL_STR",$1.s,[]int{}})}})}}
| LITERAL_CHAR  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"LITERAL_CHAR",$1.s,[]int{}})}})}}
| TRUE  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"TRUE",$1.s,[]int{}})}})}}
| FALSE  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"FALSE",$1.s,[]int{}})}})}}
| VAR_TYPE  {{$$.nn=make_node(node{"lit","",[]int{make_node(node{"VAR_TYPE",$1.s,[]int{}})}})}}                 
;

maybe_stmts
: stmts  {$$.nn=make_node(node{"maybe_stmts","",[]int{$1.nn}})}                
| /* empty */ {$$.nn=make_node(node{"maybe_stmts","",[]int{}})}                 
;

stmts
: stmts stmt {$$.nn=make_node(node{"stmts","",[]int{$1.nn,$2.nn}})}                 
| stmt {$$.nn=make_node(node{"stmts","",[]int{$1.nn}})}                 
;

stmt
: let ';' {$$.nn=make_node(node{"stmt","",[]int{$1.nn,make_node(node{";","",[]int{}})}})}                 
| item_or_view_item ';' {{$$.nn=make_node(node{"stmt","",[]int{$1.nn,make_node(node{";","",[]int{}})}})}}                 
| expr_stmt     {$$.nn=make_node(node{"stmt","",[]int{$1.nn}})}                 
| expr ';' {$$.nn=make_node(node{"stmt","",[]int{$1.nn,make_node(node{";","",[]int{}})}})}                 
;

// Things that can be an expr or a stmt, no semi required.
expr_stmt
: expr_match  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})}
| expr_if  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})} 
| expr_while  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})}
| expr_loop  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})}
| expr_for  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})}
| expr_return  {$$.nn=make_node(node{"expr_stmt","",[]int{$1.nn}})}
;

expr_return
: RETURN SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND ';' {$$.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$3.nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
| RETURN ';' {$$.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{";","",[]int{}})}})}
| RETURN lit ';'  {$$.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),$2.nn,make_node(node{";","",[]int{}})}})}
| RETURN IDENTIFIER ';'{$$.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}}),make_node(node{";","",[]int{}})}})}
;

expr_match
: MATCH IDENTIFIER SYM_OPEN_CURLY match_clauses SYM_CLOSE_CURLY  {$$.nn=make_node(node{"expr_match","",[]int{make_node(node{"MATCH","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}}),make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),$4.nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
| MATCH IDENTIFIER SYM_OPEN_CURLY match_clauses ',' SYM_CLOSE_CURLY {$$.nn=make_node(node{"expr_match","",[]int{make_node(node{"MATCH","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}}),make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),$4.nn,make_node(node{",","",[]int{}}),make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
;

match_clauses
: match_clause  {$$.nn=make_node(node{"match_clauses","",[]int{$1.nn}})}
| match_clauses ',' match_clause {$$.nn=make_node(node{"match_clauses","",[]int{$1.nn,make_node(node{",","",[]int{}}),$3.nn}})}
;

match_clause
: pats_or maybe_guard OP_FAT_ARROW match_body {$$.nn=make_node(node{"match_clauses","",[]int{$1.nn,$2.nn,make_node(node{"OP_FAT_ARROW","=>",[]int{}}),$4.nn}})}
;

match_body
: expr   {$$.nn=make_node(node{"match_body","",[]int{$1.nn}})}
| expr_stmt {$$.nn=make_node(node{"match_body","",[]int{$1.nn}})}
;

maybe_guard
: IF expr  {$$.nn=make_node(node{"match_guard","",[]int{make_node(node{"IF","",[]int{}}),$2.nn}})}
|    {$$.nn=make_node(node{"match_guard","",[]int{}})}
;

expr_if
: IF exp block  {$$.nn=make_node(node{"match_guard","",[]int{make_node(node{"IF","",[]int{}}),$2.nn,$3.nn}})}
| IF exp block ELSE block_or_if  {$$.nn=make_node(node{"match_guard","",[]int{make_node(node{"IF","",[]int{}}),$2.nn,$3.nn,make_node(node{"ELSE","",[]int{}}),$5.nn}})}
;

block_or_if
: block   {$$.nn=make_node(node{"block_or_if","",[]int{$1.nn}})}
| expr_if {$$.nn=make_node(node{"block_or_if","",[]int{$1.nn}})}
;

block
: SYM_OPEN_CURLY maybe_stmts SYM_CLOSE_CURLY  {$$.nn=make_node(node{"block","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),$2.nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
;

expr_while
: WHILE exp block   {$$.nn=make_node(node{"expr_while","",[]int{make_node(node{"WHILE","",[]int{}}),$2.nn,$3.nn}})}
;

expr_loop
: LOOP block  {$$.nn=make_node(node{"expr_loop","",[]int{make_node(node{"WHILE","",[]int{}}),$2.nn}})}
;

expr_for
: FOR exp IN exp block   {$$.nn=make_node(node{"expr_for","",[]int{make_node(node{"FOR","",[]int{}}),$2.nn,make_node(node{"IN","",[]int{}}),$4.nn,$5.nn}})}
| FOR exp IN range_di block  {$$.nn=make_node(node{"expr_for","",[]int{make_node(node{"FOR","",[]int{}}),$2.nn,make_node(node{"IN","",[]int{}}),$4.nn,$5.nn}})}
| FOR SYM_OPEN_ROUND maybe_assignment ';' exp ';' maybe_assignment SYM_CLOSE_ROUND block   {$$.nn=make_node(node{"expr_for","",[]int{make_node(node{"FOR","",[]int{}}),make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$3.nn,make_node(node{";","",[]int{}}),$5.nn,make_node(node{";","",[]int{}}),$7.nn,make_node(node{"SYM_CLOSE_ROUND","(",[]int{}}),$9.nn}})}
;

let
: LET maybe_mut pat maybe_ty_ascription maybe_init_expr   {$$.nn=make_node(node{"let","",[]int{make_node(node{"LET","",[]int{}}),$2.nn,$3.nn,$4.nn,$5.nn}})}
;

maybe_ty_ascription
: ':' ty   {$$.nn=make_node(node{"maybe_ty_ascription","",[]int{make_node(node{";","",[]int{}}),$2.nn}})}
| /* empty */ {$$.nn=make_node(node{"maybe_ty_ascription","",[]int{}})}
;

maybe_init_expr
: '=' expr   {$$.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),$2.nn}})}
| '=' SYM_OPEN_SQ exprs SYM_CLOSE_SQ  {$$.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),$3.nn,make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})}
| '=' SYM_OPEN_SQ round_exp ';' LIT_INT SYM_CLOSE_SQ  {$$.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),$3.nn,make_node(node{";","",[]int{}}),make_node(node{"LIT_INT",$5.s,[]int{}}),make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})} 
| OPEQ_INT  opeq_ops {$$.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"OPEQ_INT","",[]int{}}),$2.nn}})}
| OPEQ_FLOAT opeq_ops{$$.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"OPEQ_FLOAT","",[]int{}}),$2.nn}})}
| /* empty */{$$.nn=make_node(node{"maybe_init_expr","",[]int{}})}
;

pats_or
: pat  {$$.nn=make_node(node{"pats_or","",[]int{$1.nn}})}
| lit {$$.nn=make_node(node{"pats_or","",[]int{$1.nn}})}
| '_' {$$.nn=make_node(node{"pats_or","",[]int{make_node(node{"_","",[]int{}})}})}
| range_tri {$$.nn=make_node(node{"pats_or","",[]int{$1.nn}})}
| pats_or '|' pat  {$$.nn=make_node(node{"pats_or","",[]int{$1.nn,make_node(node{"|","",[]int{}}),$3.nn}})}
| pats_or '|' lit   {$$.nn=make_node(node{"pats_or","",[]int{$1.nn,make_node(node{"|","",[]int{}}),$3.nn}})}
| pats_or '|' range_tri   {$$.nn=make_node(node{"pats_or","",[]int{$1.nn,make_node(node{"|","",[]int{}}),$3.nn}})}
;

range_tri
: LIT_INT OP_DOTDOTDOT LIT_INT   {$$.nn=make_node(node{"range_tri","",[]int{make_node(node{"LIT_INT",$1.s,[]int{}}),make_node(node{"OP_DOTDOTDOT",$2.s,[]int{}}),make_node(node{"LIT_INT",$3.s,[]int{}}),}})}
| LITERAL_CHAR OP_DOTDOTDOT LITERAL_CHAR  {$$.nn=make_node(node{"range_tri","",[]int{make_node(node{"LITERAL_CHAR",$1.s,[]int{}}),make_node(node{"OP_DOTDOTDOT",$2.s,[]int{}}),make_node(node{"LITERAL_CHAR",$3.s,[]int{}}),}})}

range_di
: LIT_INT OP_DOTDOT LIT_INT  {$$.nn=make_node(node{"range_di","",[]int{make_node(node{"LIT_INT",$1.s,[]int{}}),make_node(node{"OP_DOTDOT",$2.s,[]int{}}),make_node(node{"LIT_INT",$3.s,[]int{}}),}})}
| LITERAL_CHAR OP_DOTDOT LITERAL_CHAR  {$$.nn=make_node(node{"range_di","",[]int{make_node(node{"LITERAL_CHAR",$1.s,[]int{}}),make_node(node{"OP_DOTDOT",$2.s,[]int{}}),make_node(node{"LITERAL_CHAR",$3.s,[]int{}}),}})}


pat
: IDENTIFIER  {$$.nn=make_node(node{"pats_or","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}})}})}
;


tys
: ty  {$$.nn=make_node(node{"tys","",[]int{$1.nn}})}
| tys ',' ty  {$$.nn=make_node(node{"tys","",[]int{$1.nn,make_node(node{",","",[]int{}}),$3.nn}})}
;

ty
: path  {$$.nn=make_node(node{"ty","",[]int{$1.nn}})}
| '~' ty
| '*' maybe_mut ty
| '&' maybe_mut ty
| OP_POWER maybe_mut ty
| SYM_OPEN_ROUND tys SYM_CLOSE_ROUND  {$$.nn=make_node(node{"ty","",[]int{make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$2.nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
;

maybe_mut
: MUT   {$$.nn=make_node(node{"maybe_mut","",[]int{make_node(node{"MUT","",[]int{}})}})}
| /* empty */  {$$.nn=make_node(node{"maybe_mut","",[]int{}})}
;


var_types
: VAR_TYPE  {$$.nn=make_node(node{"var_types","",[]int{make_node(node{"VAR_TYPE",$1.s,[]int{}})}})}                       
| IDENTIFIER {$$.nn=make_node(node{"var_types","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}})}})}                       

path
: var_types {$$.nn=make_node(node{"path","",[]int{$1.nn}})}
| SYM_OPEN_SQ var_types maybe_size SYM_CLOSE_SQ {$$.nn=make_node(node{"path","",[]int{make_node(node{"SYM_OPEN_SQ","{",[]int{}}),$2.nn,$3.nn,make_node(node{"SYM_CLOSE_SQ","}",[]int{}})}})}
;

maybe_size
: ';' LIT_INT {$$.nn=make_node(node{"maybe_size","",[]int{make_node(node{";","",[]int{}}),make_node(node{"LIT_INT",$2.s,[]int{}})}})}
| {$$.nn=make_node(node{"maybe_size","",[]int{}})}
;

maybe_exprs
: exprs {$$.nn=make_node(node{"maybe_exprs","",[]int{$1.nn}})}
| /* empty */ {$$.nn=make_node(node{"maybe_exprs","",[]int{}})}
;

exprs
: expr {$$.nn=make_node(node{"exprs","",[]int{$1.nn}})}
| exprs ',' expr {$$.nn=make_node(node{"exprs","",[]int{$1.nn,make_node(node{",","",[]int{}}),$3.nn}})}
;

//// $$$opeq+int doesn't work

maybe_assignment
: assignment {$$.nn=make_node(node{"maybe_assignment","",[]int{$1.nn}})}
| {$$.nn=make_node(node{"maybe_assignment","",[]int{}})}
;

hole
: IDENTIFIER {$$.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}})}})}                       
| IDENTIFIER SYM_OPEN_SQ round_exp SYM_CLOSE_SQ {$$.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),$3.nn,make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})}                       
| IDENTIFIER '.' hole {$$.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{".","",[]int{}}),$3.nn}})}

assignment
: hole '=' expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"=","",[]int{}}),$3.nn}})}
| hole OP_ADDEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"+=","OP_ADDEQ",[]int{}}),$3.nn}})}
| hole OP_SUBEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"-=","OP_SUBEQ",[]int{}}),$3.nn}})}
| hole OP_LEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"<=","OP_LEQ",[]int{}}),$3.nn}})}
| hole OP_GEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{">=","OP_GEQ",[]int{}}),$3.nn}})}
| hole OP_MULEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"*=","OP_MULEQ",[]int{}}),$3.nn}})}
| hole OP_DIVEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"/=","OP_DIVEQ",[]int{}}),$3.nn}})}
| hole OP_MODEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"%=","OP_MODEQ",[]int{}}),$3.nn}})}
| hole OP_ANDEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"&=","OP_ANDEQ",[]int{}}),$3.nn}})}
| hole OP_SHLEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"<<=","OP_SHLEQ",[]int{}}),$3.nn}})}
| hole OP_SHREQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{">>=","OP_SHREQ",[]int{}}),$3.nn}})}
| hole OP_OREQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"|=","OP_OREQ",[]int{}}),$3.nn}})}
| hole OP_XOREQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"^=","OP_XOREQ",[]int{}}),$3.nn}})}
| hole OP_EQEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"==","OP_EQEQ",[]int{}}),$3.nn}})}
| hole OP_NOTEQ expr {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"!=","OP_NOTEQ",[]int{}}),$3.nn}})}
| hole OPEQ_INT opeq_ops {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"=int","OPEQ_INT",[]int{}}),$3.nn}})}
| hole OPEQ_FLOAT opeq_ops {$$.nn=make_node(node{"assignment","",[]int{$1.nn,make_node(node{"=float","OPEQ_FLOAT",[]int{}}),$3.nn}})}


;



opeq_ops
:  '+' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"+","",[]int{}}),$2.nn}})}
| '-' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"-","",[]int{}}),$2.nn}})}
| '&' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"&","",[]int{}}),$2.nn}})}
| '|' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"|","",[]int{}}),$2.nn}})}
| '^' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"^","",[]int{}}),$2.nn}})}
| '/' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"/","",[]int{}}),$2.nn}})}
| '*' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"*","",[]int{}}),$2.nn}})}
| '>' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{">","",[]int{}}),$2.nn}})}
| '<' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"<","",[]int{}}),$2.nn}})}
| '%' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"%","",[]int{}}),$2.nn}})}
| '.' expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{".","",[]int{}}),$2.nn}})}
| OP_RSHIFT expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_RSHIFT",">>",[]int{}}),$2.nn}})}
| OP_LSHIFT expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_LSHIFT","<<",[]int{}}),$2.nn}})}
| OP_ANDAND expr {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_ANDAND","&&",[]int{}}),$2.nn}})}
| OP_OROR expr  {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_OROR","||",[]int{}}),$2.nn}})}
| OP_POWER expr     {$$.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_POWER","**",[]int{}}),$2.nn}})}    
|
;

expr
: round_exp {$$.nn=make_node(node{"expr","",[]int{$1.nn}})}
| assignment {$$.nn=make_node(node{"expr","",[]int{$1.nn}})}
;

//$$struct remaining

exp
: lit {$$.map["value"]=$1.map["value"];$$.map["place"]=$1.map["place"];$$.map["type"]=$1.map["type"];}
| IDENTIFIER     {$$.map["value"]=$1.map["value"];$$.map["place"]=$1.map["place"];$$.map["type"]=$1.map["type"];$$.code=$1.code;}
| IDENTIFIER SYM_OPEN_SQ round_exp SYM_CLOSE_SQ     {}
| IDENTIFIER ':' struct_expr  {$$.nn=make_node(node{"exp","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{":","",[]int{}}),$3.nn}})}              
| '!' round_exp      {}
| '&' round_exp      {$$.nn=make_node(node{"exp","",[]int{make_node(node{"&","",[]int{}}),$2.nn}})}
| OP_ANDMUT round_exp      {$$.nn=make_node(node{"exp","",[]int{make_node(node{"OP_ANDMUT","&mut",[]int{}}),$2.nn}})}
| round_exp '-' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="-, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '+' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="+, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '&' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="&, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '|' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="|, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '^' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="^, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '/' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="/, $$.map["value"],$1.map["value"], $3.map["value"]"; }
| round_exp '*' round_exp       {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="*, $$.map["value"],$1.map["value"], $3.map["value"]"; }    
| round_exp '>' round_exp       
| round_exp '<' round_exp  
| round_exp '%' round_exp  {$$.code=$1.code; p=list_end($$.code); p.next=$3.code;q=list_end($1.code);q.next=new(node);if($$.map==NULL)||($1.map==NULL)||($3.map==NULL) {log.Fatal("variable not declared")}; if($$.map["type"]!=$1.map["type"])||($$.map["type"]!=$3.map["type"])||($3.map["type"]!=$1.map["type"]) {log.Fatal("Type Mismatch")}; q.next.code="%, $$.map["value"],$1.map["value"], $3.map["value"]"; } 
| round_exp '.' round_exp {$$.nn=make_node(node{"exp","",[]int{$1.nn,make_node(node{".","",[]int{}}),$3.nn}})}
| round_exp OP_RSHIFT round_exp       
| round_exp OP_LSHIFT round_exp 
| round_exp OP_ANDAND round_exp {$$.nn=make_node(node{"exp","",[]int{$1.nn,make_node(node{"OP_ANDAND","&&",[]int{}}),$3.nn}})}
| round_exp OP_OROR round_exp {$$.nn=make_node(node{"exp","",[]int{$1.nn,make_node(node{"OP_OROR","||",[]int{}}),$3.nn}})}
| round_exp OP_POWER round_exp {$$.nn=make_node(node{"exp","",[]int{$1.nn,make_node(node{"OP_POWER","**",[]int{}}),$3.nn}})}
| func_identifier SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND  {$$.nn=make_node(node{"exp","",[]int{$1.nn,make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$3.nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}       
| CONTINUE     {$$.nn=make_node(node{"exp","",[]int{make_node(node{"CONTINUE","",[]int{}})}}) }                                    
| CONTINUE IDENTIFIER  {$$.nn=make_node(node{"exp","",[]int{make_node(node{"CONTINUE","",[]int{}}),make_node(node{"IDENTIFIER",$2.s,[]int{}})}}) }                
| UNSAFE block    {$$.nn=make_node(node{"exp","",[]int{make_node(node{"UNSAFE","",[]int{}}),$2.nn}}) }                
| block   {$$.nn=make_node(node{"exp","",[]int{$1.nn}})  }                       
;

round_exp 
: SYM_OPEN_ROUND round_exp SYM_CLOSE_ROUND {$$.nn=make_node(node{"round_exp","",[]int{make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),$1.nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
| exp {$$.nn=make_node(node{"round_exp","",[]int{$1.nn}})}
;

func_identifier 
: IDENTIFIER {$$.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}})}})}
| IDENTIFIER SYM_COLCOL func_identifier {$$.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{"SYM_COLCOL","::",[]int{}}),$3.nn}})}
| IDENTIFIER '!' {$$.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",$1.s,[]int{}}),make_node(node{"!","",[]int{}})}})}
;

struct_expr
: SYM_OPEN_CURLY field_inits default_field_init SYM_CLOSE_CURLY {$$.nn=make_node(node{"struct_expr","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),$2.nn,$3.nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
;

field_inits
: field_init {$$.nn=make_node(node{"field_inits","",[]int{$1.nn}})}
| field_inits ',' field_init {$$.nn=make_node(node{"field_inits","",[]int{$1.nn,make_node(node{",","",[]int{}}),$3.nn}})}
;

field_init
: maybe_mut IDENTIFIER ':' expr {$$.nn=make_node(node{"field_init","",[]int{$1.nn,make_node(node{"IDENTIFIER",$2.s,[]int{}}),make_node(node{":","",[]int{}}),$4.nn}})}
;

default_field_init
: ','	{$$.nn=make_node(node{"default_field_init","",[]int{make_node(node{",","",[]int{}})}})}
| ',' OP_DOTDOT expr {$$.nn=make_node(node{"default_field_init","",[]int{make_node(node{",","",[]int{}}),make_node(node{"OP_DOTDOT","..",[]int{}}),$3.nn}})}
| /* empty */ {$$.nn=make_node(node{"default_field_init","",[]int{}})}
;
