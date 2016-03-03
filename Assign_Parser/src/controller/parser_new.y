%{
package main
import "fmt"
//import "math"
import "os"
%}

%debug

%token SHL
%token SHR
%token LE
%token EQEQ
%token NE
%token GE
%token ANDAND
%token OROR
%token BINOPEQ
%token OP_DOTDOT  //
%token DOTDOTDOT
%token MOD_SEP
%token OP_INSIDE  //
%token OP_FAT_ARROW //
%token LIT_CHAR
%token INTEGER   //
%token LIT_UINT
%token LIT_INT_UNSUFFIXED
%token FLOAT   //
%token LIT_FLOAT_UNSUFFIXED
%token LITERAL  //
%token LIT_STR_RAW
%token IDENTIFIER   //
%token UNDERSCORE
%token LIFETIME

// keywords
%token SELF
%token STATIC
%token SUPER
%token AS
%token BREAK
%token CRATE
%token "else"
%token ENUM
%token EXTERN
%token "false"
%token "fn"
%token "for"
%token "if"
%token IMPL
%token "in"
%token "let"
%token "loop"
%token "match"
%token MOD
%token "mut"
%token ONCE
%token PRIV
%token PUB
%token REF
%token RETURN
%token "struct"
%token "true"
%token TRAIT
%token TYPE
%token UNSAFE
%token "use"
%token "while"
%token CONTINUE  // whether to change to string or not b/c of nonassoc CONTINUE.
%token PROC
%token BOX
%token TYPEOF

%token SHEBANG
%token STATIC_LIFETIME



%expect 0

%nonassoc CONTINUE
%nonassoc IDENTIFIER
%nonassoc '('
%nonassoc '{'
%left '+' '-'

%start rust

%%

/// println, print macro support => standard macros
rust
:"struct" IDENTIFIER struct_expr
|item_or_view_item
;


item_or_view_item
: item_fn
;

item_fn
: "fn" IDENTIFIER fn_decl inner_attrs_and_block  { $$ = mk_node("fn", 1, $3); }
;

fn_decl
: fn_args ret_ty
;

fn_args
: '(' maybe_args_general ')'
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
| /* empty */
;

inner_attrs_and_block
: '{' maybe_inner_attrs maybe_stmts '}'   { $$ = $2; }
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
| IDENTIFIER '(' meta_seq ')'
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
| INTEGER
| LIT_UINT
| LIT_INT_UNSUFFIXED
| FLOAT
| LIT_FLOAT_UNSUFFIXED
| LITERAL
| LIT_STR_RAW
| "true"
| "false"
;

maybe_stmts
: stmts
| /* empty */
;

stmts
: stmts stmt
| stmt
;

stmt
: let
| item_or_view_item
| expr_stmt
| expr
;

// Things that can be an expr or a stmt, no semi required.
expr_stmt
: expr_match
| expr_if
| expr_while
| expr_loop
| expr_for
;

expr_match
: "match" expr '{' match_clauses '}'
| "match" expr '{' match_clauses ',' '}'
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
: "if" expr
| // empty
;

expr_if
: "if" expr block
| "if" expr block "else" block_or_if
;

block_or_if
: block
| expr_if
;

block
: '{' maybe_stmts '}'
;

expr_while
: "while" expr block
;

expr_loop
: "loop" block
;

expr_for
: "for" expr "in" expr block
;

let
: "let" maybe_mut pat maybe_ty_ascription maybe_init_expr
;

maybe_ty_ascription
: ':' ty
| /* empty */
;

maybe_init_expr
: '=' expr
| /* empty */
;


pats_or
: pat
| pats_or '|' pat
;

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
| '(' tys ')'
;

maybe_mut
: "mut"
| /* empty */
;

path:
maybe_exprs
: exprs
| /* empty */
;

exprs
: expr
| exprs ',' expr
;

expr     // add  other operations.
: lit
| IDENTIFIER                            { $$ = mk_node("IDENTIFIER", 0); }
| IDENTIFIER struct_expr                { $$ = mk_node("struct", 1, $1); }
| expr '+' expr                    { $$ = mk_node("+", 2, $1, $2); }
| expr '(' maybe_exprs ')'         { $$ = mk_node("call", 1, $1); }
| CONTINUE                         { $$ = mk_node("continue", 0); }
| CONTINUE IDENTIFIER                   { $$ = mk_node("continue-label", 0); }
| UNSAFE block                     { $$ = mk_node("unsafe-block", 0); }
| block                            { $$ = mk_node("block", 0); }
;


struct_expr
: '{' field_inits default_field_init '}'
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











