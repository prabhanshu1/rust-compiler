%{
package main
import "../Assign_IR/src/symtable"
import "fmt"
import "log"
  /* import "os" */
import "strconv"
import "strings"
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
  b.value=a.value;
  for(a.next!=nil){
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

%right '=' '!'  OP_SHLEQ OP_SHREQ  OP_ADDEQ OP_OREQ OP_ANDEQ  OP_MULEQ OP_DIVEQ OP_MODEQ OP_XOREQ 
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

Code : rust {print_ircode($1.code);}   
;

rust
: STRUCT IDENTIFIER marker_1 struct_expr rust  {$$.code=$4.code;p:=list_end(&$$.code);p.next=$5.code;}
| item_or_view_item rust  {$$.code=$1.code;p:=list_end(&$$.code);p.next=$2.code;}
| USE func_identifier ';' rust 
|  
;

marker_1:   {  fmt.Println("in rust _marker "+$0.s);$$.mp=symtab.Make_entry($0.s);$$.mp["type"]="struct"; }
;
item_or_view_item
: item_fn {$$.code=$1.code;}
;

item_fn
: FN IDENTIFIER fn_decl inner_attrs_and_block  
{$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="func"+$2.s;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;  
  $$.code=new (node);$$.code.value="jmp, "+$$.mp["after"];$$.code.next=new(node);$$.code.next.value="label, "+$$.mp["begin"] + ", " + $$.mp["funargs"];$$.code.next.next=new(node);$$.code.next.next=$4.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="label, "+$$.mp["after"];

}
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
: SYM_OPEN_CURLY maybe_inner_attrs maybe_stmts SYM_CLOSE_CURLY  {$$.code=$3.code; }   
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
  }  //maybe incomplete
;

maybe_stmts
: stmts       {$$.code=$1.code;}           
| /* empty */                  
;

stmts
: stmts stmt {$$.code=$1.code;p:=list_end(&$$.code);p.next=$2.code;}                 
| stmt    {$$.code=$1.code;}              
;

stmt
: let ';' {$$.code=$1.code;}                 
| item_or_view_item ';'         //incomplete                  
| expr_stmt  {$$.code=$1.code;}                   
| expr ';'   {$$.code=$1.code;$$.mp=$1.mp; }
;

// Things that can be an expr or a stmt, no semi required.
expr_stmt
: expr_match  {$$.code=$1.code;}
| expr_if   {$$.code=$1.code;}
| expr_while  {$$.code=$1.code;}
| expr_loop  {$$.code=$1.code;}
| expr_for  {$$.code=$1.code;}
| expr_return  {$$.code=$1.code;}
;

expr_return
: RETURN SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND ';' // incomplete
| RETURN ';'                   {$$.code=new(node);$$.code.value="return";}
| RETURN lit ';'               {$$.code=$2.code; p:=list_end(&$$.code);p.next=new(node);p.next.value="return, "+$2.mp["value"];}
| RETURN IDENTIFIER ';'        {$$.code=new(node); $2.mp =symtab.Find_id($2.s);
   if($2.mp ==nil){log.Fatal("Returning undefined identifier; ");};
   $$.code.value="return, "+$2.mp["value"];
 }
;

expr_match
: MATCH exp marker_2 SYM_OPEN_CURLY match_clauses SYM_CLOSE_CURLY ','  {
  $$.code=$2.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&$$.code);q.next=$5.code;
  
 }
;

marker_2
:  {  $$.mp=symtab.Make_entry("case_exp");$$.code=new(node);
   $$.code.value="=, "+$$.mp["cae_exp"]+", "+$0.mp["value"];
  $$.mp["after_match"]="label"+strconv.Itoa(label_num);label_num+=1;
}
;

match_clauses
: match_clause   {$$.code=$1.code;}
| match_clauses ',' match_clause  {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;}
;

match_clause
: pats_or maybe_guard OP_FAT_ARROW match_body  {
  
 temp:=symtab.Find_id("case_exp");
  if ($$.s=="_") {
    $$.code=$4.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="lable"+$$.mp["after_match"];
  }
  
  if($$.code==nil) {
    $$.code=new(node);$$.code.value="ifgoto, jne, "+temp["value"]+", "+$1.mp["value"]+", label"+strconv.Itoa(label_num);
    $$.code.next=$4.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="jmp, "+temp["after_match"];
    p.next.next=new(node);p.next.next.value="label, "+strconv.Itoa(label_num);
  }
  
  
 } 
;

match_body
: expr    {$$.code=$1.code;$$.mp=$1.mp;}
| expr_stmt  {$$.code=$1.code;$$.mp=$1.mp;}
;

maybe_guard
: IF expr  
|    
;


expr_if 
: IF exp block  {  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;
   fmt.Println("printing exp $2.mp", $2.mp);
  if($2.mp==nil) {log.Fatal("Bad If   block;;;")};

  $$.code=new(node);p:=$$.code;
  if($2.code!=nil){
  o:=copy_nodes($2.code,$$.code);o.next=new(node);p=o.next; }else{p=$$.code;}
  p.value="ifgoto, je, "+$2.mp["value"]+", 0, "+$$.mp["after"];p.next=new(node); p.next=$3.code;q:=list_end(&$$.code);q.next=new(node);q.next.value="label, "+$$.mp["after"];


 }

| IF exp block ELSE block_or_if  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["true"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;
   if($2.mp==nil) {log.Fatal("Expression or block  not declared in IF statement")};
   $$.code=$2.code;p:=list_end(&$$.code);
     p.next=new(node); p.next.value="ifgoto, je, "+$2.mp["value"]+", 1, "+$$.mp["true"];
  p.next.next=new(node);
  p.next.next=$5.code;q:=list_end(&$$.code);q.next=new(node);
  q.next.value="jmp, "+$$.mp["after"];q.next.next=new(node);
  q.next.next.value="label, "+$$.mp["true"];q.next.next.next=new(node);
  q.next.next.next=$3.code; r:=list_end(&$$.code);r.next=new(node);
  r.next.value="label, "+$$.mp["after"];
 }
;

block_or_if
: block   {  $$.mp=$1.mp;$$.code=$1.code; } 
| expr_if {   $$.mp=$1.mp;$$.code=$1.code; }
;

block
: SYM_OPEN_CURLY maybe_stmts SYM_CLOSE_CURLY { $$.code=$2.code;$$.mp=$2.mp;}  
;

expr_while
: WHILE exp block   {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1; 
  $$.code=new(node);$$.code.value="label, "+$$.mp["begin"];$$.code.next=$2.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="ifgoto, je, "+$2.mp["value"]+", 0, "+$$.mp["after"]; p.next.next=$3.code;r:=list_end(&$$.code);r.next=new(node);r.next.value="jmp, "+$$.mp["begin"];r.next.next=new(node);r.next.next.value="label, "+$$.mp["after"];}
;

expr_loop
: LOOP block {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;  if($2.code==nil) {log.Fatal("variable not declared")};
  $$.code=new (node);$$.code.value="label, "+$$.mp["begin"];$$.code.next=new(node);p:=copy_nodes($2.code,$$.code.next);p.next=new(node);p.next.value="jmp, "+$$.mp["begin"];}  
;

expr_for
: FOR exp IN exp block  {
$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1; 
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
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1; 
$$.code=$2.code;q:=list_end(&$$.code);q.next=new(node);
 q.next.value="=, "+$2.mp["value"]+", "+$4.mp["start"];
  q.next.next=new(node);q.next.next.value="label, "+$$.mp["begin"] ;
 r:=q.next.next;r.next=new(node);r.next.value="ifgoto, jg, "+$2.mp["value"]+", "+$4.mp["end"]+", "+$$.mp["after"];

  //r.next.next=$5.code;
  s:=list_end(&r);s.next=new(node);
  s.next=$5.code;
  s=list_end(&s);  s.next=new(node);s.next.value="+, "+$2.mp["value"]+", "+$2.mp["value"]+", "+"1";
  s.next.next=new(node);s.next.next.value="jmp, "+$$.mp["begin"];
 t:=s.next.next;t.next=new(node);t.next.value="label, "+$$.mp["after"];

 
  
}
| FOR SYM_OPEN_ROUND maybe_assignment ';' exp ';' maybe_assignment SYM_CLOSE_ROUND block  {

$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["begin"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1; 
$$.code=$3.code;p:=list_end(&$$.code);p.next=new(node);p.next.value="label, "+$$.mp["begin"];;p.next.next=$5.code;q:=list_end(&p);q.next=new(node);
 q.next.value="ifgoto, je, "+$5.mp["value"]+", 0, "+$$.mp["after"];

 q.next.next=$9.code;s:=list_end(&q);
 s.next=$7.code;t:=list_end(&s);t.next=new(node);
 t.next.value="jmp, "+$$.mp["begin"]; u:=t.next;u.next=new(node);
 u.next.value="label, "+$$.mp["after"];

}
;

let   // incomplete for array and struct => both have $4.map != nil;;
: LET maybe_mut pat  maybe_ty_ascription maybe_init_expr   {
  fmt.Println("in let",$4.mp,$5.s);
            fmt.Println("OOOOOOOOOO",$5)
  if($3.mp==nil) {log.Fatal("Variable name not present in let");}
    if($4.mp==nil){ 
    if($4.s!=""){
      if($5.mp!=nil){
          /*let mut y:i32 = expr */
        fmt.Println($5.mp["type"],$4.s)
          if($5.mp["type"]!=$4.s) {log.Fatal("Type mismatch in let ;;");}
          $3.mp["type"]=$2.s+$5.mp["type"];
      fmt.Println("MMMMMMMMMMMMMMMMMMMM",$4.s)
          $$.code=new(node);
          if($5.code!=nil) {
          p:=copy_nodes($5.code,$$.code);p.next=new(node);
          if $5.mp["Array"] == "true" {
                p2:=&p.next;
              if $5.mp["args"]!="" {
                s2 := strings.Split($5.mp["args"], ", ")
                for i := 0; i < $5.n; i++ {
                  (*p2).value="[]=, "+strconv.Itoa(i)+", "+$3.mp["value"] +", "+s2[i];
                  (*p2).next=new(node)
                  p2=&((*p2).next)
                }
              }else{
                for i := 0; i < $5.n; i++ {
                  (*p2).value="[]=, "+strconv.Itoa(i)+", "+$3.mp["value"] +", "+$5.mp["value"];
                  (*p2).next=new(node)
                  p2=&((*p2).next)
                }
              }
            }else{
            p.next.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
          }

          }else {
            if $5.mp["Array"] == "true" {
                p2:=&$$.code;
              if $5.mp["args"]!="" {
                s2 := strings.Split($5.mp["args"], ", ")
                for i := 0; i < $5.n; i++ {
                  (*p2).value="[]=, "+strconv.Itoa(i)+", "+$3.mp["value"] +", "+s2[i];
                  (*p2).next=new(node)
                  p2=&((*p2).next)
                }
              }else{
                for i := 0; i < $5.n; i++ {
                  (*p2).value="[]=, "+strconv.Itoa(i)+", "+$3.mp["value"] +", "+$5.mp["value"];
                  (*p2).next=new(node)
                  p2=&((*p2).next)
                }
              }
            }else{
            $$.code.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
          }
                      
          }
      } else{/*let  y:i32 */
        $3.mp["type"]=$2.s+$4.s;
       
      }
    } else{ /* let y = 5 */
      fmt.Println("FFFFFFFFFFFFFFFFFFF")
        print_ircode($5.code)
        fmt.Println("FFFFFFFFFFFFFFFFFFF")
      if($5.mp==nil) {log.Fatal("incomplete let expression  ;");}
      $3.mp["type"]=$2.s+$5.mp["type"];
      $$.code=new(node);$$.code=$5.code; p:=list_end(&$$.code);p.next=new(node);p.next.value="=, "+$3.mp["value"]+", "+$5.mp["value"];
              print_ircode($5.code)
        fmt.Println("FFFFFFFFFFFFFFFFFFF")
    }
    }else{

      if($4.mp["type"]!="struct") {log.Fatal("struct "+$4.mp["value"]+"not defined prior to use;");}
    str_slice := strings.Split($5.s, ",");
      $$.code=$5.code;p:=list_end(&$$.code);
    temp:=symtab.Make_entry($3.mp["value"]+"_"+str_slice[0]);
      for i := 0; i < len(str_slice); i+=2 {
        
        temp=symtab.Make_entry($3.mp["value"]+"_"+str_slice[i]);
        p.next=new(node);p.next.value="=, "+temp["value"]+", "+str_slice[i+1];
        p=p.next;
	}
      fmt.Println("in let, elssssss",$5.s);
      print_ircode($$.code);
            fmt.Println("in let, elssssss",$5.s);
    }

}
;

maybe_ty_ascription
: ':' ty   {if ($2.mp==nil) {$$.s=$2.s;$$.mp=nil;$$.code=nil;}else{$$.code=$2.code;$$.mp=$2.mp;  };}
| /* empty */ {$$.s="";$$.mp=nil;$$.code=nil; }
;

maybe_init_expr

: '=' round_exp   { fmt.Println("jjdddlsddd");$$.code=$2.code; $$.mp=$2.mp;}
| '='  struct_init ';' { fmt.Println("jjdddlsddddqqqqqq");$$.code=$2.code;$$.s=$2.s;}  //struct

| '=' SYM_OPEN_SQ exprs SYM_CLOSE_SQ  { fmt.Println("jjdddlsdddww");$$.code=$3.code;$$.mp=$3.mp;$$.n=$3.n;}//array

| '=' SYM_OPEN_SQ round_exp ';' LIT_INT SYM_CLOSE_SQ { fmt.Println("jjdddlsdddeeeeeee");$$.code=$3.code;$$.mp=$3.mp;$$.n=$5.n;$$.mp["Array"]="true";$$.mp["type"]="Array_"+$$.mp["type"]+"_"+strconv.Itoa($$.n)}//array  

| OPEQ_INT  opeq_ops  { fmt.Println("jjdddlsdddyyyyyyyy");
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.code=new(node);$$.code=$2.code;p:=list_end(&$$.code);p.next=new(node);
    if $2.mp["op"] == "" {
    p.next.value="="+", "+$$.mp["value"]+", "+strconv.Itoa($1.n);
  }else{
    p.next.value=$2.mp["op"]+", "+$$.mp["value"]+", "+strconv.Itoa($1.n)+", "+$2.mp["value"];
  }
  $$.mp["type"]="int";
}

| OPEQ_FLOAT opeq_ops  {fmt.Println("jjdddlsdddiiiii");}
| /* empty */ {$$.s="" ;fmt.Println("jjdddlsdddmmmmm");}
;

pats_or
: pat  {$$.mp=$1.mp;$$.code=nil;$$.s="";}
| lit {$$.code=$1.code;$$.mp=$1.mp;$$.s="";}
| '_' {$$.s="_";}
| range_tri  {$$.code=$1.code;$$.mp=$1.mp;}
| pats_or '|' pat  {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;}
| pats_or '|' lit   
| pats_or '|' range_tri   
;

range_tri
: lit OP_DOTDOTDOT lit  {$$.code=$1.code; p:=list_end(&$$.code);p.next=$3.code;}   


range_di
: LIT_INT OP_DOTDOT LIT_INT  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;;$$.mp["start"]=strconv.Itoa($1.n);$$.mp["end"]=strconv.Itoa($3.n);}
| LITERAL_CHAR OP_DOTDOT LITERAL_CHAR  {$$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["start"]=strconv.Itoa((int)(([]rune($1.s))[0]));$$.mp["end"]=strconv.Itoa((int)(([]rune($3.s))[0]));}


pat
: IDENTIFIER {

  $1.mp =symtab.Find_id($1.s);
  if($1.mp==nil){
    $1.mp=symtab.Make_entry($1.s);}
$$.mp=$1.mp;}
;


tys
: ty  
| tys ',' ty  
;




ty
: path  { $$.s=$1.s; }
| '~' ty { $$.s="~" + $2.s; }
| '*' maybe_mut ty  { $$.s="*" + $2.s + $3.s; }
| '&' maybe_mut ty { $$.s="&" + $2.s + $3.s; }
| OP_POWER maybe_mut ty { $$.s="**" + $2.s + $3.s; }
| SYM_OPEN_ROUND tys SYM_CLOSE_ROUND  
;

maybe_mut
: MUT   {$$.s=$1.s+"_";}
| /* empty */  {$$.s="";}
;


var_types
: VAR_TYPE  {$$.mp=nil; if($1.s=="i8")||($1.s=="i16")||($1.s=="i32")||($1.s=="i64")||($1.s=="isize")||($1.s=="u8")||($1.s=="u16")||($1.s=="u32")||($1.s=="u64")||($1.s=="usize"){$$.s="int";}  else{$$.s="str";}  }

| IDENTIFIER {

  $$.mp =symtab.Find_id($1.s);
  fmt.Println("in var_type",$$.mp);
  if($$.mp==nil){
     log.Fatal("var_type not defined,")}  
  }

path
: var_types {$$.s=$1.s; $$.mp=$1.mp;}
| SYM_OPEN_SQ var_types maybe_size SYM_CLOSE_SQ {$$.s="Array_"+$2.s+"_"+$3.s;}
;

maybe_size
: ';' LIT_INT {$$.s=strconv.Itoa($2.n)}
|                      //maybe incomplete
;

maybe_exprs
: exprs {$$.code=$1.code;$$.mp=$1.mp;}
| /* empty */ 
;

exprs
: expr {$$.code=$1.code;$$.mp["args"]=$1.mp["value"] +", ";$$.n=1;$$.mp["type"]=$1.mp["type"];}
| exprs ',' expr   { $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code;   
  $$.mp["args"]=$1.mp["args"]+$3.mp["value"] + ", ";
  if len($1.mp["type"])>5 && ($1.mp["type"])[0:5]=="Array" {
    sss:=strings.Split($1.mp["type"],"_");
    $1.mp["type"]=sss[1];

  }
  $$.n=$1.n+1;$$.mp["type"]="Array_"+$1.mp["type"]+"_"+strconv.Itoa($$.n); fmt.Println("LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL",$$.mp["args"]);$$.mp["Array"]="true";

}
;


maybe_assignment
: assignment {$$.code=$1.code;$$.mp=$1.mp;}
| 
;

hole
: IDENTIFIER   {
 p:=symtab.Find_id($1.s);
  if(p==nil){
    $1.mp=symtab.Make_entry($1.s);
    $$.mp=$1.mp;  
  }else{$$.mp=p;}
    }                        
| IDENTIFIER SYM_OPEN_SQ round_exp SYM_CLOSE_SQ       {
 p:=symtab.Find_id($1.s);
  if(p==nil){
    $1.mp=symtab.Make_entry($1.s);
    $1.mp["value2"]=$3.mp["value"];
    $$.mp=$1.mp;  
  }else{$$.mp=p;}
}


| IDENTIFIER '.' hole {
 p:=symtab.Find_id($1.s+"_"+$3.mp["value"]);
  if(p==nil){
    $1.mp=symtab.Make_entry($1.s+"_"+$3.mp["value"]);
    $$.mp=$1.mp;  
  }else{$$.mp=p;}

}  //incomplete

;

assignment
: hole '=' round_exp  {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p);q.next=new(node);if($1.mp["value2"]=="") {q.next.value="=, "+$1.mp["value"]+", "+$3.mp["value"];}else{q.next.value="[]=, "+$1.mp["value2"]+", "+$1.mp["value"] +", "+$3.mp["value"];}}

| hole OP_ADDEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="+, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_SUBEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="-, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }

| hole OP_MULEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="*, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_DIVEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="/, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_MODEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="%, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_ANDEQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="&, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }

| hole OP_SHLEQ round_exp 
| hole OP_SHREQ round_exp 

| hole OP_OREQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="|, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }
| hole OP_XOREQ round_exp {$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code;q:=list_end(&p.next);;q.next=new(node);q.next.value="^, "+$1.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"]; }

| hole OP_EQEQ round_exp 
| hole OP_NOTEQ round_exp 

| hole OPEQ_INT opeq_ops {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
  $$.code=new(node);$$.code=$3.code;p:=list_end(&$$.code);
  p.next=new(node);
  if $3.mp["op"] == "" {
    p.next.value="="+", "+$$.mp["value"]+", "+strconv.Itoa($2.n);
  }else{
    p.next.value=$3.mp["op"]+", "+$$.mp["value"]+", "+strconv.Itoa($2.n)+", "+$3.mp["value"];
  }
  p.next.next=new(node);
  p.next.next.value="=, "+$1.mp["value"]+", "+$$.mp["value"];
}
 
| hole OPEQ_FLOAT opeq_ops 


;



opeq_ops
:  '+' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="+";} 
| '-' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="-";} 
| '&' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="&";} 
| '|' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="|";} 
| '^' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="^";} 
| '/' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="/";} 
| '*' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="*";} 
| '>' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]=">";} 
| '<' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="<";} 
| '%' expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="%";} 
| '.' expr {fmt.Println("LLLLLLLLLLLLLLLL");$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]=".";} //incorrect 
| OP_RSHIFT expr 
| OP_LSHIFT expr 
| OP_ANDAND expr {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="&&";} //incorrect
| OP_OROR expr  {$$.code=$2.code;$$.mp=$2.mp; $$.mp["op"]="||";} //incorrect 
| OP_POWER expr         
|   { $$.s=""; }
;

expr 
: round_exp {fmt.Println("hello in expr");$$.code=$1.code;$$.mp=$1.mp;}
| assignment {fmt.Println("sadsad");$$.code=$1.code;$$.mp=$1.mp;}
;

//$$struct remaining

exp
: lit {$$.mp=$1.mp;$$.code=$1.code;}

| IDENTIFIER     {
   fmt.Println("jjdddlsvvvvvvv");
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
  $$.code.value="=[], "+$$.mp["value"]+", "+ $1.s + ", " + $3.mp["value"];
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
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));
 temp_num+=1;
 $$.mp["type"]=$1.mp["type"];
 $$.mp["true"]="label"+strconv.Itoa(label_num);
 label_num+=1;
 $$.mp["after"]="label"+strconv.Itoa(label_num);
 label_num+=1;
      $$.code=$1.code;
      p:=list_end(&$$.code); 
      p.next=new(node);
      p.next=$3.code;
      q:=list_end(&$$.code);

      //q:=copy_nodes(p.next,$3.code);
      q.next=new(node);
      if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
      if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};
    q.next.value="ifgoto, jg, "+$1.mp["value"]+", "+$3.mp["value"]+", "+$$.mp["true"];
    r:=new(node);
    q.next.next=r;r.value="=, "+$$.mp["value"]+", "+"0";r.next=new(node);
    r.next.value="jmp, "+$$.mp["after"];
    r.next.next=new(node);s:=r.next.next;s.value="label, "+$$.mp["true"];s.next=new(node);
    s.next.value="=, "+$$.mp["value"]+", "+"1";s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"]; }
  

| round_exp '<' round_exp   {
 $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));
 temp_num+=1;
 $$.mp["type"]=$1.mp["type"];
 $$.mp["true"]="label"+strconv.Itoa(label_num);
 label_num+=1;
 $$.mp["after"]="label"+strconv.Itoa(label_num);
 label_num+=1;
      $$.code=$1.code;
      p:=list_end(&$$.code); 
      p.next=new(node);
      p.next=$3.code;
      q:=list_end(&$$.code);

      //q:=copy_nodes(p.next,$3.code);
      q.next=new(node);
      if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
      if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};
    q.next.value="ifgoto, jl, "+$1.mp["value"]+", "+$3.mp["value"]+", "+$$.mp["true"];
    r:=new(node);
    q.next.next=r;r.value="=, "+$$.mp["value"]+", "+"0";r.next=new(node);
    r.next.value="jmp, "+$$.mp["after"];
    r.next.next=new(node);s:=r.next.next;s.value="label, "+$$.mp["true"];s.next=new(node);
    s.next.value="=, "+$$.mp["value"]+", "+"1";s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"]; }

| round_exp OP_LEQ round_exp  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));
 temp_num+=1;
 $$.mp["type"]=$1.mp["type"];
 $$.mp["true"]="label"+strconv.Itoa(label_num);
 label_num+=1;
 $$.mp["after"]="label"+strconv.Itoa(label_num);
 label_num+=1;
      $$.code=$1.code;
      p:=list_end(&$$.code); 
      p.next=new(node);
      p.next=$3.code;
      q:=list_end(&$$.code);

      //q:=copy_nodes(p.next,$3.code);
      q.next=new(node);
      if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
      if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};
    q.next.value="ifgoto, jle, "+$1.mp["value"]+", "+$3.mp["value"]+", "+$$.mp["true"];
    r:=new(node);
    q.next.next=r;r.value="=, "+$$.mp["value"]+", "+"0";r.next=new(node);
    r.next.value="jmp, "+$$.mp["after"];
    r.next.next=new(node);s:=r.next.next;s.value="label, "+$$.mp["true"];s.next=new(node);
    s.next.value="=, "+$$.mp["value"]+", "+"1";s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"]; }
| round_exp OP_GEQ round_exp  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));
 temp_num+=1;
 $$.mp["type"]=$1.mp["type"];
 $$.mp["true"]="label"+strconv.Itoa(label_num);
 label_num+=1;
 $$.mp["after"]="label"+strconv.Itoa(label_num);
 label_num+=1;
      $$.code=$1.code;
      p:=list_end(&$$.code); 
      p.next=new(node);
      p.next=$3.code;
      q:=list_end(&$$.code);

      //q:=copy_nodes(p.next,$3.code);
      q.next=new(node);
      if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
      if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};
    q.next.value="ifgoto, jge, "+$1.mp["value"]+", "+$3.mp["value"]+", "+$$.mp["true"];
    r:=new(node);
    q.next.next=r;r.value="=, "+$$.mp["value"]+", "+"0";r.next=new(node);
    r.next.value="jmp, "+$$.mp["after"];
    r.next.next=new(node);s:=r.next.next;s.value="label, "+$$.mp["true"];s.next=new(node);
    s.next.value="=, "+$$.mp["value"]+", "+"1";s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"]; }
      
| round_exp '%' round_exp    {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="%, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }

| round_exp '.' round_exp   //incorrect 
{fmt.Println("in a.b");
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&$$.code);  q.next=new(node);
  q.next.value="=, "+$$.mp["value"]+", "+$1.mp["value"]+"_"+$3.mp["value"];
  
   }
| round_exp OP_RSHIFT round_exp   //incorrect
    {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value=">>, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_LSHIFT round_exp { //incorrect
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="<<, "+$$.mp["value"]+", "+$1.mp["value"]+", "+$3.mp["value"];
   }
| round_exp OP_ANDAND round_exp  {

  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"]; $$.mp["false"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;
    $$.code=$1.code;p:=list_end(&$1.code);
    p.next=new(node);p.next=$3.code;  q:=list_end(&$$.code);
    q.next=new(node);
    if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
  if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};
  q.next.value="ifgoto, je, "+$1.mp["value"]+", 0, "+$$.mp["false"];
 r:=new(node);q.next.next=r;
  r.value="ifgoto, je, "+$3.mp["value"]+", 0, "+$$.mp["false"];
  r.next=new(node);rr:=r.next; rr.value="=, "+$$.mp["value"]+", "+"1";
  rr.next=new(node);rr.next.value="jmp, "+$$.mp["after"];
  rr.next.next=new(node);s:=rr.next.next;s.value="label, "+$$.mp["false"];
  s.next=new(node);s.next.value="=, "+$$.mp["value"]+", "+"0";
  s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"];
}


| round_exp OP_OROR round_exp  {
  $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"]; $$.mp["true"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;$$.mp["after"]="label"+strconv.Itoa(label_num);label_num+=1;label_num+=1;

$$.code=$1.code;p:=list_end(&$1.code);
    p.next=new(node);p.next=$3.code;  q:=list_end(&$$.code);
    q.next=new(node);
    if($1.mp==nil)||($3.mp==nil) {log.Fatal("variable not declared")};
    if($3.mp["type"]!=$1.mp["type"]) {log.Fatal("Type Mismatch")};

    q.next.value="ifgoto, je, "+$1.mp["value"]+", 1, "+$$.mp["true"];r:=new(node);
    q.next.next=r;r.value="ifgoto, je, "+$3.mp["value"]+", 1, "+$$.mp["true"];
    r.next=new(node);rr:=r.next; rr.value="=, "+$$.mp["value"]+", "+"0";
    rr.next=new(node);rr.next.value="jmp, "+$$.mp["after"];rr.next.next=new(node);
    s:=rr.next.next;s.value="label, "+$$.mp["true"];s.next=new(node);
    s.next.value="=, "+$$.mp["value"]+", "+"1";s.next.next=new(node);s.next.next.value="label, "+$$.mp["after"];
}


| round_exp OP_POWER round_exp { //incorrect
  
   }
| func_identifier SYM_OPEN_ROUND maybe_exprs SYM_CLOSE_ROUND  {
    $$.mp=symtab.Make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;$$.mp["type"]=$1.mp["type"];$$.code=$1.code;  p:=list_end(&$$.code);  p.next=$3.code; q:=list_end(&p.next);  q.next=new(node);
  q.next.value="push, "+$3.mp["args"];q.next.next=new(node);q.next.next.value="call, " + $1.mp["value"]+ ", ";
}

| CONTINUE     
| CONTINUE IDENTIFIER  
| UNSAFE block    
| block   {$$.mp=$1.mp;$$.code=$1.code;}
| BREAK {}
| BREAK IDENTIFIER {}
;

round_exp 
: exp { fmt.Println("jjdddlsdddddcccccc");$$.mp=$1.mp;$$.code=$1.code;}
|  SYM_OPEN_ROUND round_exp SYM_CLOSE_ROUND  {$$.mp=$1.mp;$$.code=$1.code;}
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
: SYM_OPEN_CURLY field_exprs default_field_expr SYM_CLOSE_CURLY  {

  
} 
;

field_exprs
: field_expr 
| field_exprs ',' field_expr 
;

field_expr
: maybe_mut IDENTIFIER ':' path  {}
;

default_field_expr
: ','	
| ',' OP_DOTDOT expr 
| /* empty */ 
;

struct_init
: IDENTIFIER SYM_OPEN_CURLY field_inits SYM_CLOSE_CURLY   {
  $$.s=$3.s;$$.code=$3.code;
  
}
;

field_inits
: field_init          {$$.s=$1.s;$$.code=$1.code;}
| field_inits ',' field_init { $$.s=$1.s+","+$3.s;$$.code=$1.code;p:=list_end(&$$.code);p.next=$3.code; }
;

field_init
: IDENTIFIER ':'  exp {$$.s=$1.s+","+$3.mp["value"];$$.code=$3.code;}
;



