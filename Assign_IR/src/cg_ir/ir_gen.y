%{
package main
import "../Assign_IR/src/symtable"
import "fmt"
import "log"
  /* import "os" */
import "strconv"
  var line = 0
  var temp_num=0;
 var label_num=0;

 func print_ircode(a *node){
   if(a==nil) {fmt.Println("No code to print");return ;}
   fmt.Println(a.value);
   for (a.next!=nil){
     a=a.next;
     fmt.Println(a.value);
   }
   } 
 

 func list_end(l **node) *node {
   if((*l)==nil) {(*l)=new(node); return *l;}
 p:=*l;
   for (p.next!=nil){p=p.next}
   return p;
}

type node struct {
  value string
  next *node
}

func copy_nodes(a *node, b *node) *node {
  if(a==nil){return b;}
  fmt.Println(a,b)
  b.value=a.value;
  for(a.next!=nil){
    fmt.Println(a.value)
    b.next=new(node);
    b=b.next;
    a=a.next;
    b.value=a.value;
  }
  return b; 
}

func space(a string,i int)int{
for ;a[i]==' ';i++{}
return i
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

%}




%struct {
  code *node
  mp map[string]string
  n int
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

Code : rust {}   
;

rust
: STRUCT IDENTIFIER struct_expr rust  
| item_or_view_item rust  
| USE func_identifier ';' rust 
|  
;


item_or_view_item
: item_fn 
;

item_fn
: FN IDENTIFIER fn_decl inner_attrs_and_block  
{$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+$2.s;$$.mp["after"]=strconv.Itoa(label_num);  
  $$.code=new (node);$$.code.value="jmp, "+$$.mp["after"];$$.code.next=new(node);$$.code.next.value="label, "+$$.mp["begin"];$$.code.next.next=new(node);p:=copy_nodes($4.code,$$.code.next.next);p.next=new(node);p.next.value="label, "+$$.mp["after"];print_ircode($$.code)}
;

fn_decl
: fn_args ret_ty 
;

fn_args
: SYM_OPEN_ROUND maybe_args_general SYM_CLOSE_ROUND	
;

maybe_args_general
: args_general      
| /* empty */       
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
| OP_INSIDE SYM_OPEN_ROUND  SYM_CLOSE_ROUND 
| /* empty */ 
;

inner_attrs_and_block
: SYM_OPEN_CURLY maybe_inner_attrs maybe_stmts SYM_CLOSE_CURLY   
;

maybe_inner_attrs
: inner_attrs 
| /* empty */ 
;

inner_attrs
: inner_attr 
| inner_attrs inner_attr 
;

inner_attr
: SHEBANG '[' meta_item ']'  
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
lit
: LIT_CHAR  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="str";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+$1.s; }   
| LIT_INT    {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="int";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+strconv.Itoa($1.n); }   
| LIT_UINT   {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="int";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+strconv.Itoa($1.n); }   
| LIT_INT_UNSUFFIXED  
| FLOAT  
| LIT_FLOAT_UNSUFFIXED  
| LITERAL_STR  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="str";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+$1.s; }   
| LITERAL_CHAR  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="str";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+$1.s; }   
| TRUE    {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="int";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", 1"; }   
| FALSE   {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="int";$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", 0"; }   
| VAR_TYPE   {if($1.s=="i8")||($1.s=="i16")||($1.s=="i32")||($1.s=="i64")||($1.s=="isize")||($1.s=="u8")||($1.s=="u16")||($1.s=="u32")||($1.s=="u64")||($1.s=="usize"){$$.s="int";}  else{$$.s="str";}
{$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]="str";$$.mp["value"]=$1.s;$$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+$1.s; }
  }
;

maybe_stmts
: stmts       {$$.code=$1.code;print_ircode($1.code);}           
| /* empty */                  
;

stmts
: stmts stmt {$$.code=$1.code;p:=list_end(&$1.code);p.next=$2.code;}                 
| stmt    {$$.code=$1.code;}              
;

stmt
: let ';' {if($1.code==nil) {fmt.Println("in stmt let code is nil");};$$.code=$1.code;}                 
| item_or_view_item ';'                  
| expr_stmt                      
| expr ';'   {$$.code=$1.code;$$.mp=$1.mp; }
;

// Things that can be an expr or a stmt, no semi required.
expr_stmt
: expr_match  
| expr_if   
| expr_while  
| expr_loop  
| expr_for  
| expr_return  
;

expr_return
: RETURN SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND ';' 
| RETURN ';' 
| RETURN lit ';'  
| RETURN IDENTIFIER ';'
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
: pats_or maybe_guard OP_FAT_ARROW match_body 
;

match_body
: expr   
| expr_stmt 
;

maybe_guard
: IF expr  
|    
;


expr_if 
: IF exp block  { fmt.Println("iff0"); $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);
   fmt.Println("iff1",$2,$3);
  if($2.mp==nil) {log.Fatal("Bad If   block;;;")};
  fmt.Println("iff2");
  $$.code=new(node);fmt.Println("ssdd");p:=$$.code;
  if($2.code!=nil){
  o:=copy_nodes($2.code,$$.code);o.next=new(node);p=o.next;fmt.Println("sss00"); }else{p=$$.code;}
  p.value="ifgoto, je, "+$2.mp["value"]+", 0, "+$$.mp["after"];fmt.Println("iff2-1");p.next=new(node);if($3.code==nil){fmt.Println("$3.code nil in expr_if");}; q:=copy_nodes($3.code,p.next);q.next=new(node); fmt.Println("iff2-3");q.next.value="label, "+$$.mp["after"];
 fmt.Println("iff3");
  print_ircode($$.code);
  fmt.Println("iff4");
 }

| IF exp block ELSE block_or_if  { $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);$$.mp["true"]="label"+strconv.Itoa(label_num);
   if($2.mp==nil)||($5.mp==nil) {log.Fatal("Expression or block  not declared in IF statement")};
   $$.code=$2.code;p:=list_end(&$2.code);
     p.next=new(node); p.next.value="ifgoto, je, "+$2.mp["value"]+", 1, "+$$.mp["true"];
  p.next.next=new(node); q:=copy_nodes(p.next.next,$5.code);q.next=new(node);
  q.next.value="jmp, "+$$.mp["after"];q.next.next=new(node);
  q.next.next.value="label, "+$$.mp["true"];q.next.next.next=new(node);
  r:=copy_nodes(q.next.next.next,$3.code);r.next=new(node);
  r.next.value="label, "+$$.mp["after"];}
;

block_or_if
: block   { $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.mp=$1.mp;$$.code=$1.code;} 
| expr_if { $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.mp=$1.mp;$$.code=$1.code;}
;

block
: SYM_OPEN_CURLY maybe_stmts SYM_CLOSE_CURLY {$$.code=$2.code}  
;

expr_while
: WHILE exp block   {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);$$.mp["after"]="label"+strconv.Itoa(label_num); 
  $$.code=new(node);$$.code.value="label, "+$$.mp["begin"];$$.code.next=$2.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="ifgoto, je, "+$2.mp["value"]+", 0, "+$$.mp["after"]; p.next.next=$3.code;r:=list_end(&$$.code);r.next=new(node);r.next.value="jmp, "+$$.mp["begin"];r.next.next=new(node);r.next.next.value="label, "+$$.mp["after"];fmt.Println("Printing WHILE");print_ircode($$.code)}
;

expr_loop
: LOOP block {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);$$.mp["after"]="label"+strconv.Itoa(label_num);  if($2.code==nil) {log.Fatal("variable not declared")};
  $$.code=new (node);$$.code.value="label, "+$$.mp["begin"];$$.code.next=new(node);p:=copy_nodes($2.code,$$.code.next);p.next=new(node);p.next.value="jmp, "+$$.mp["begin"];print_ircode($$.code)}  
;

expr_for
: FOR exp IN exp block  {
$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);$$.mp["after"]="label"+strconv.Itoa(label_num); 
  $$.code=new (node);p:=copy_nodes($2.code,$$.code);p.next=new(node);q:=copy_nodes($4.code,p.next);
 tmp:=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  
  q.next=new(node);q.next.value="=, "+tmp["value"]+", "+"0";
  q.next.next=new(node);q.next.next.value="label, "+$$.mp["begin"] ;
 r:=q.next.next;r.next=new(node);r.next.value="ifgoto, jg, "+tmp["value"]+", "+$4.mp["size"]+", "+$$.mp["after"];
  r.next.next=new(node);r.next.next.value="=, "+$2.mp["value"]+", "+$4.mp[tmp["value"]];r.next.next.next=new(node);
  s:=copy_nodes($5.code,r.next.next.next);s.next=new(node);
  s.next.value="+, "+tmp["value"]+", "+tmp["value"]+", "+"1";
  s.next.next=new(node);s.next.next.value="jmp, "+$$.mp["begin"];
 t:=s.next.next;t.next=new(node);t.next.value="label, "+$$.mp["after"];
   }
| FOR exp IN range_di block  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);$$.mp["after"]="label"+strconv.Itoa(label_num); 
$$.code=$2.code;q:=list_end(&$$.code);q.next=new(node);
 q.next.value="=, "+$2.mp["value"]+", "+$4.mp["start"];
  q.next.next=new(node);q.next.next.value="label, "+$$.mp["begin"] ;
 r:=q.next.next;r.next=new(node);r.next.value="ifgoto, jg, "+$2.mp["value"]+", "+$4.mp["end"]+", "+$$.mp["after"];

  //r.next.next=$5.code;
  s:=list_end(&r);fmt.Println("AAAAAAAAAAAA",$5,$2);s.next=new(node);
  s.next=$5.code;
  s=list_end(&s);  s.next=new(node);s.next.value="+, "+$2.mp["value"]+", "+$2.mp["value"]+", "+"1";
  s.next.next=new(node);s.next.next.value="jmp, "+$$.mp["begin"];
 t:=s.next.next;t.next=new(node);t.next.value="label, "+$$.mp["after"];
fmt.Println("FFFFFFFFFFF")
 print_ircode($$.code);
  
}
| FOR SYM_OPEN_ROUND maybe_assignment ';' exp ';' maybe_assignment SYM_CLOSE_ROUND block  {
   fmt.Println("AT for inside start")
$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);$$.mp["after"]="label"+strconv.Itoa(label_num); 
$$.code=$3.code;p:=list_end(&$$.code);p.next=$5.code;q:=list_end(&p);q.next=new(node);
 q.next.value="ifgoto, je, "+$5.mp["value"]+", 0, "+$$.mp["after"];
   fmt.Println("AT for inside mid")
 q.next.next=$9.code;s:=list_end(&q);
 s.next=$7.code;t:=list_end(&s);t.next=new(node);
 t.next.value="jmp, "+$$.mp["begin"]; u:=t.next;u.next=new(node);
 u.next.value="label, "+$$.mp["after"];
 fmt.Println("AT for inside end")
 print_ircode($$.code);
}
;

let
: LET maybe_mut pat  maybe_ty_ascription maybe_init_expr   {
  if($3.mp==nil) {log.Fatal("Variable name not present in let");}
  if($4.mp==nil){
    if($4.s!=""){
      if($5.mp!=nil){
          /*let mut y:i32 = expr */
          if($5.mp["type"]!=$4.s) {log.Fatal("Type mismatch in let ;;");}
          $3.mp["type"]=$5.mp["type"];
          $$.code=new(node);
          if($5.code!=nil) {
          p:=copy_nodes($5.code,$$.code);p.next=new(node);p.next.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
          }else {
            $$.code.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
          }
      } else{/*let  y:i32 */
        $3.mp["type"]=$4.s;
      }
    } else{ /* let y = 5 */
      if($5.mp==nil) {log.Fatal("incomplete let expression  ;");}
      $3.mp["type"]=$5.mp["type"];
      $$.code=$5.code;p:=list_end(&$5.code);p.next=new(node);p.next.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
    }
        }
  fmt.Println("let after");
}
;

maybe_ty_ascription
: ':' ty   {if ($2.mp==nil) {$$.s=$2.s;}else{$$.code=$2.code;$$.mp=$2.mp;  }}
| /* empty */ {$$.s="";}
;

maybe_init_expr
: '=' expr   {$$.code=$2.code; $$.mp=$1.mp;}
| '=' SYM_OPEN_SQ exprs SYM_CLOSE_SQ  
| '=' SYM_OPEN_SQ round_exp ';' LIT_INT SYM_CLOSE_SQ  
| OPEQ_INT  opeq_ops  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.code=new(node);$$.code.value="=, "+$$.mp["value"]+", "+strconv.Itoa($1.n);fmt.Println($1.n)
  $$.mp["type"]="int"
  if($2.code!=nil){
    $$.code=$2.code;p:=list_end(&$2.code);p.next=new(node);p.next.value="+, "+$$.mp["value"]+", "+strconv.Itoa($1.n)+", "+$2.mp["value"];}else{
    $$.s="";}
}
| OPEQ_FLOAT opeq_ops 
| /* empty */ {$$.s=""}
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
: LIT_INT OP_DOTDOT LIT_INT  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;;$$.mp["start"]=strconv.Itoa($1.n);$$.mp["end"]=strconv.Itoa($3.n);}
| LITERAL_CHAR OP_DOTDOT LITERAL_CHAR  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["start"]=strconv.Itoa((int)(([]rune($1.s))[0]));$$.mp["end"]=strconv.Itoa((int)(([]rune($3.s))[0]));}


pat
: IDENTIFIER {
 $1.mp =symtab.Find_id($1.s);
  if($1.mp==nil){
    $1.mp=symtab.Make_entry($1.s);}
$$.mp=$1.mp;fmt.Println($1.mp,"sdaaa",$$.mp);}
;


tys
: ty  
| tys ',' ty  
;

ty
: path  {if($1.code==nil) {$$.s=$1.s;} else{ $$.code=$1.code;$$.mp=$1.mp; } }
| '~' ty
| '*' maybe_mut ty
| '&' maybe_mut ty
| OP_POWER maybe_mut ty
| SYM_OPEN_ROUND tys SYM_CLOSE_ROUND  
;

maybe_mut
: MUT   {$$.s=$1.s;}
| /* empty */  {$$.s="";}
;


var_types
: VAR_TYPE  {if($1.s=="i8")||($1.s=="i16")||($1.s=="i32")||($1.s=="i64")||($1.s=="isize")||($1.s=="u8")||($1.s=="u16")||($1.s=="u32")||($1.s=="u64")||($1.s=="usize"){$$.s="int";}  else{$$.s="str";}}
| IDENTIFIER {$$.s=$1.s;}

path
: var_types {$$.s=$1.s;}
| SYM_OPEN_SQ var_types maybe_size SYM_CLOSE_SQ 
;

maybe_size
: ';' LIT_INT 
| 
;

maybe_exprs
: exprs {$$.code=$1.code;$$.mp=$1.mp;}
| /* empty */ 
;

exprs
: expr {$$.code=$1.code;$$.mp["args"]=$1.mp["value"] +", ";}
| exprs ',' expr   {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code;   
  $$.mp["args"]=$1.mp["args"]+$3.mp["value"] + ", ";
}
;

//// $$$opeq+int doesn't work

maybe_assignment
: assignment {fmt.Println("In maybe assignment");$$.code=$1.code;print_ircode($1.code);fmt.Println("In maybe assignment")}
| 
;

hole
: IDENTIFIER   {
 p:=symtab.Find_id($1.s);fmt.Println("(in exp ) Identifier =",p["value"]);
  if(p==nil){fmt.Println("(in exp )new Identifier ",$1.s);
    $1.mp=symtab.Make_entry($1.s);
    $$.mp=$1.mp;  
  }else{$$.mp=p;}
    }                        
| IDENTIFIER SYM_OPEN_SQ round_exp SYM_CLOSE_SQ                        
| IDENTIFIER '.' hole 

assignment
: hole '=' expr  {fmt.Println("After func call");$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;fmt.Println("hole");print_ircode($$.code);fmt.Println("hole2");q:=list_end(&p);q.next=new(node);q.next.value="=, "+$1.mp["value"]+", "+$3.mp["value"];fmt.Println("hole3");print_ircode($$.code);fmt.Println("hole4");fmt.Println("DONE func call");}

| hole OP_ADDEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="+, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_SUBEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="-, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_LEQ expr 
| hole OP_GEQ expr 
| hole OP_MULEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="*, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_DIVEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="/, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_MODEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="%, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_ANDEQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="&, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_SHLEQ expr 
| hole OP_SHREQ expr 
| hole OP_OREQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="|, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_XOREQ expr {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="^, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; print_ircode($$.code);}
| hole OP_EQEQ expr 
| hole OP_NOTEQ expr 
| hole OPEQ_INT opeq_ops {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.code=new(node);$$.code.value="=, "+$1.mp["value"]+", "+strconv.Itoa($2.n);fmt.Println($1.n)
    if($3.code!=nil){
      $$.code.next=new(node);$$.code.next=$2.code;p:=list_end(&$$.code.next);p.next=new(node);p.next.value="+, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];}else{
    $$.s="";}
}
 
| hole OPEQ_FLOAT opeq_ops 


;



opeq_ops
:  '+' expr 
| '-' expr 
| '&' expr 
| '|' expr 
| '^' expr 
| '/' expr 
| '*' expr 
| '>' expr 
| '<' expr 
| '%' expr 
| '.' expr 
| OP_RSHIFT expr 
| OP_LSHIFT expr 
| OP_ANDAND expr 
| OP_OROR expr  
| OP_POWER expr         
|   { $$.s=""; }
;

expr
: round_exp {$$.code=$1.code;$$.mp=$1.mp;}
| assignment {$$.code=$1.code;$$.mp=$1.mp;}
;

//$$struct remaining

exp
: lit {$$.mp=$1.mp;}

| IDENTIFIER     {
 p:=symtab.Find_id($1.s);
  if(p==nil){
    $1.mp=symtab.Make_entry($1.s);
    $$.mp=$1.mp;  
  }else{$$.mp=p;}

    }
| IDENTIFIER SYM_OPEN_SQ round_exp SYM_CLOSE_SQ     {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.mp["type"]=$1.mp["type"];
  $$.code=new(node);
  $$.code.value="=, "+$$.mp["value"]+", "+$1.mp[$3.mp["value"]];
}
| IDENTIFIER ':' struct_expr  
| '!' round_exp      {}         /* not of conditional */
| '&' round_exp      
| OP_ANDMUT round_exp      
| round_exp '-' round_exp      {  
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="-, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
}

| round_exp '+' round_exp        {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="+, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
  }

| round_exp '&' round_exp        {
   $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="&, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp '|' round_exp        {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="|, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }

| round_exp '^' round_exp        {
   $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="^, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp '/' round_exp        {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="/, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; 
}

| round_exp '*' round_exp    {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="*, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
  }
    
| round_exp '>' round_exp       {
   $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value=">, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
 }

| round_exp '<' round_exp   {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="<, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
      
| round_exp '%' round_exp    {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="%, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }

| round_exp '.' round_exp 
{
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="., "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_RSHIFT round_exp   
    {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value=">>, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_LSHIFT round_exp {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="<<, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_ANDAND round_exp  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="&&, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }      
| round_exp OP_OROR round_exp {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="||, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_POWER round_exp {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="**, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| func_identifier SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND  {
    $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="push, "+$3.mp["args"];q.next.next=new(node);q.next.next.value="call, " + $1.mp["value"]+ ", "; print_ircode($$.code)
}
| CONTINUE     
| CONTINUE IDENTIFIER  
| UNSAFE block    
| block   {$$.mp=$1.mp;$$.code=$1.code;}
| BREAK {}
| BREAK IDENTIFIER {}
;

round_exp 
: SYM_OPEN_ROUND round_exp SYM_CLOSE_ROUND  {$$.mp=$1.mp;$$.code=$1.code;}
| exp {$$.mp=$1.mp;$$.code=$1.code;fmt.Println("DODODIDODODO ");print_ircode($$.code)}
;

func_identifier 
: IDENTIFIER  {
 p:=symtab.Find_id($1.s);
  if(p==nil){
    $1.mp=symtab.Make_entry($1.s);
    $$.mp=$1.mp;  
  }else{$$.mp=p;}

    }
| IDENTIFIER SYM_COLCOL func_identifier  {
 p:=symtab.Find_id($1.s+"_"+$3.s);
  if(p==nil){
    $$.mp=symtab.Make_entry($1.s+"_"+$3.s);
  }else{$$.mp=p;}

    }
| IDENTIFIER '!'  {
 p:=symtab.Find_id($1.s+"_");
  if(p==nil){
    $$.mp=symtab.Make_entry($1.s+"_");
  }else{$$.mp=p;}

    }
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
