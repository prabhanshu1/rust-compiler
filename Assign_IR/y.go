
//line ./src/cg_ir/ir_gen.y:2
package main
import __yyfmt__ "fmt"
//line ./src/cg_ir/ir_gen.y:2
		import "fmt"
import "log"
import "os"
import "strconv"
  var line = 0
  var temp_num=0;
 var label_num=0;
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

func copy_nodes(a node, b node)node{
  b.value=a.value;
  while(a.next!=NULL){
    b.next=new (node);
    b=b.next;
    a=a.next;
    b.value=a.value;
  }
  return b; 
}

type node struct {
  var value [100]char
  next *node
}


//line ./src/cg_ir/ir_gen.y:134
type yySymType struct {
	yys int
  code *node=NULL
  var map map[string]string =NULL
}

const MUT = 57346
const IDENTIFIER = 57347
const OP_INSIDE = 57348
const FOR = 57349
const RETURN = 57350
const AS = 57351
const SYM_OPEN_SQ = 57352
const SYM_CLOSE_SQ = 57353
const SYM_OPEN_ROUND = 57354
const SYM_CLOSE_ROUND = 57355
const SYM_OPEN_CURLY = 57356
const SYM_CLOSE_CURLY = 57357
const ANDAND = 57358
const BINOPEQ = 57359
const DOTDOT = 57360
const DOTDOTDOT = 57361
const EQEQ = 57362
const FAT_ARROW = 57363
const GE = 57364
const LE = 57365
const LIFETIME = 57366
const LIT_CHAR = 57367
const FLOAT = 57368
const LIT_FLOAT_UNSUFFIXED = 57369
const LIT_INT_UNSUFFIXED = 57370
const LITERAL_STR = 57371
const LITERAL_CHAR = 57372
const LIT_UINT = 57373
const MOD_SEP = 57374
const NE = 57375
const OROR = 57376
const SHL = 57377
const SHR = 57378
const UNDERSCORE = 57379
const KEYWORD = 57380
const VAR_TYPE = 57381
const LIT_INT = 57382
const OPEQ_INT = 57383
const HEX = 57384
const OCTAL = 57385
const BINARY = 57386
const OPEQ_FLOAT = 57387
const LITERAL = 57388
const OP_EQ = 57389
const OP_RSHIFT = 57390
const OP_LSHIFT = 57391
const OP_ADDEQ = 57392
const OP_SUBEQ = 57393
const OP_MULEQ = 57394
const OP_DIVEQ = 57395
const OP_MODEQ = 57396
const OP_ANDMUT = 57397
const OP_EQEQ = 57398
const OP_NOTEQ = 57399
const OP_ANDAND = 57400
const OP_OROR = 57401
const OP_POWER = 57402
const OP_DOTDOT = 57403
const OP_DOTDOTDOT = 57404
const OP_SUB = 57405
const OP_ADD = 57406
const OP_AND = 57407
const OP_LEQ = 57408
const OP_GEQ = 57409
const OP_OR = 57410
const OP_XOR = 57411
const OP_FSLASH = 57412
const OP_NOT = 57413
const OP_COLUMN = 57414
const OP_MUL = 57415
const OP_GTHAN = 57416
const OP_LTHAN = 57417
const OP_MOD = 57418
const OP_DOT = 57419
const OP_APOSTROPHE = 57420
const OP_FAT_ARROW = 57421
const SYM_COLCOL = 57422
const SYM_HASH = 57423
const SYM_COMMA = 57424
const SYM_SEMCOL = 57425
const FINISH = 57426
const NEWLINE = 57427
const ABSTRACT = 57428
const ALIGNOF = 57429
const BECOME = 57430
const BOX = 57431
const BREAK = 57432
const CONST = 57433
const CONTINUE = 57434
const CRATE = 57435
const DO = 57436
const ELSE = 57437
const ENUM = 57438
const EXTERN = 57439
const FALSE = 57440
const FINAL = 57441
const FN = 57442
const IF = 57443
const IMPL = 57444
const IN = 57445
const LET = 57446
const LOOP = 57447
const MACRO = 57448
const MATCH = 57449
const MOD = 57450
const MOVE = 57451
const OFFSETOF = 57452
const OVERRIDE = 57453
const PRIV = 57454
const PROC = 57455
const PUB = 57456
const PURE = 57457
const REF = 57458
const SELF = 57459
const SIZEOF = 57460
const STATIC = 57461
const STRUCT = 57462
const SUPER = 57463
const TRAIT = 57464
const TRUE = 57465
const TYPE = 57466
const TYPEOF = 57467
const UNSAFE = 57468
const UNSIZED = 57469
const USE = 57470
const VIRTUAL = 57471
const WHERE = 57472
const WHILE = 57473
const YIELD = 57474
const PRINTLN = 57475
const MACRO_RULES = 57476
const SHEBANG = 57477
const STATIC_LIFETIME = 57478
const OP_SHLEQ = 57479
const OP_SHREQ = 57480
const OP_OREQ = 57481
const OP_ANDEQ = 57482
const OP_XOREQ = 57483

var yyToknames = []string{
	"MUT",
	"IDENTIFIER",
	"OP_INSIDE",
	"FOR",
	"RETURN",
	"AS",
	"SYM_OPEN_SQ",
	"SYM_CLOSE_SQ",
	"SYM_OPEN_ROUND",
	"SYM_CLOSE_ROUND",
	"SYM_OPEN_CURLY",
	"SYM_CLOSE_CURLY",
	"ANDAND",
	"BINOPEQ",
	"DOTDOT",
	"DOTDOTDOT",
	"EQEQ",
	"FAT_ARROW",
	"GE",
	"LE",
	"LIFETIME",
	"LIT_CHAR",
	"FLOAT",
	"LIT_FLOAT_UNSUFFIXED",
	"LIT_INT_UNSUFFIXED",
	"LITERAL_STR",
	"LITERAL_CHAR",
	"LIT_UINT",
	"MOD_SEP",
	"NE",
	"OROR",
	"SHL",
	"SHR",
	"UNDERSCORE",
	"KEYWORD",
	"VAR_TYPE",
	"LIT_INT",
	"OPEQ_INT",
	"HEX",
	"OCTAL",
	"BINARY",
	"OPEQ_FLOAT",
	"LITERAL",
	"OP_EQ",
	"OP_RSHIFT",
	"OP_LSHIFT",
	"OP_ADDEQ",
	"OP_SUBEQ",
	"OP_MULEQ",
	"OP_DIVEQ",
	"OP_MODEQ",
	"OP_ANDMUT",
	"OP_EQEQ",
	"OP_NOTEQ",
	"OP_ANDAND",
	"OP_OROR",
	"OP_POWER",
	"OP_DOTDOT",
	"OP_DOTDOTDOT",
	"OP_SUB",
	"OP_ADD",
	"OP_AND",
	"OP_LEQ",
	"OP_GEQ",
	"OP_OR",
	"OP_XOR",
	"OP_FSLASH",
	"OP_NOT",
	"OP_COLUMN",
	"OP_MUL",
	"OP_GTHAN",
	"OP_LTHAN",
	"OP_MOD",
	"OP_DOT",
	"OP_APOSTROPHE",
	"OP_FAT_ARROW",
	"SYM_COLCOL",
	"SYM_HASH",
	"SYM_COMMA",
	"SYM_SEMCOL",
	"FINISH",
	"NEWLINE",
	"ABSTRACT",
	"ALIGNOF",
	"BECOME",
	"BOX",
	"BREAK",
	"CONST",
	"CONTINUE",
	"CRATE",
	"DO",
	"ELSE",
	"ENUM",
	"EXTERN",
	"FALSE",
	"FINAL",
	"FN",
	"IF",
	"IMPL",
	"IN",
	"LET",
	"LOOP",
	"MACRO",
	"MATCH",
	"MOD",
	"MOVE",
	"OFFSETOF",
	"OVERRIDE",
	"PRIV",
	"PROC",
	"PUB",
	"PURE",
	"REF",
	"SELF",
	"SIZEOF",
	"STATIC",
	"STRUCT",
	"SUPER",
	"TRAIT",
	"TRUE",
	"TYPE",
	"TYPEOF",
	"UNSAFE",
	"UNSIZED",
	"USE",
	"VIRTUAL",
	"WHERE",
	"WHILE",
	"YIELD",
	"PRINTLN",
	"MACRO_RULES",
	"SHEBANG",
	"STATIC_LIFETIME",
	" =",
	" !",
	"OP_SHLEQ",
	"OP_SHREQ",
	"OP_OREQ",
	"OP_ANDEQ",
	"OP_XOREQ",
	" <",
	" >",
	" |",
	" ^",
	" &",
	" +",
	" -",
	" .",
	" *",
	" /",
	" %",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 38,
	15, 209,
	-2, 118,
	-1, 89,
	12, 202,
	41, 131,
	45, 131,
	50, 131,
	51, 131,
	52, 131,
	53, 131,
	54, 131,
	56, 131,
	57, 131,
	66, 131,
	67, 131,
	137, 131,
	139, 131,
	140, 131,
	141, 131,
	142, 131,
	143, 131,
	-2, 171,
	-1, 149,
	12, 202,
	-2, 171,
	-1, 308,
	41, 132,
	45, 132,
	50, 132,
	51, 132,
	52, 132,
	53, 132,
	54, 132,
	56, 132,
	57, 132,
	66, 132,
	67, 132,
	137, 132,
	139, 132,
	140, 132,
	141, 132,
	142, 132,
	143, 132,
	-2, 172,
}

const yyNprod = 212
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1054

var yyAct = []int{

	69, 96, 78, 219, 68, 188, 222, 88, 275, 280,
	35, 72, 114, 266, 86, 111, 273, 221, 142, 143,
	62, 58, 13, 291, 57, 87, 38, 355, 144, 145,
	146, 216, 242, 132, 131, 141, 137, 136, 140, 338,
	319, 137, 136, 140, 45, 335, 289, 225, 224, 196,
	129, 128, 127, 77, 15, 316, 43, 323, 218, 317,
	364, 116, 123, 124, 283, 36, 63, 191, 16, 346,
	333, 122, 55, 309, 325, 344, 25, 53, 65, 115,
	112, 324, 34, 332, 152, 98, 102, 103, 101, 104,
	282, 100, 157, 109, 52, 148, 151, 185, 153, 108,
	281, 16, 322, 268, 367, 42, 54, 354, 7, 353,
	351, 271, 352, 23, 139, 138, 134, 135, 133, 132,
	131, 141, 137, 136, 140, 365, 17, 51, 3, 265,
	192, 193, 194, 61, 150, 150, 5, 150, 125, 159,
	121, 197, 59, 126, 180, 181, 182, 110, 120, 109,
	215, 315, 60, 217, 14, 214, 223, 220, 107, 17,
	29, 227, 228, 229, 230, 231, 232, 233, 234, 235,
	236, 237, 238, 239, 240, 241, 187, 55, 179, 366,
	80, 320, 339, 106, 223, 198, 199, 200, 201, 202,
	203, 204, 205, 206, 207, 208, 209, 210, 211, 212,
	213, 264, 263, 269, 261, 291, 142, 143, 159, 259,
	190, 54, 310, 290, 178, 50, 144, 145, 146, 49,
	93, 279, 278, 56, 177, 277, 10, 183, 267, 48,
	20, 260, 24, 285, 369, 270, 31, 27, 262, 189,
	36, 186, 184, 36, 292, 293, 294, 295, 296, 297,
	298, 299, 300, 301, 302, 303, 304, 305, 306, 307,
	4, 147, 39, 98, 102, 103, 101, 104, 282, 100,
	284, 11, 150, 12, 313, 311, 8, 108, 281, 142,
	143, 25, 117, 118, 119, 327, 318, 330, 331, 144,
	145, 146, 336, 37, 16, 328, 22, 195, 47, 113,
	314, 67, 272, 130, 334, 135, 133, 132, 131, 141,
	137, 136, 140, 286, 326, 360, 340, 321, 276, 274,
	76, 2, 75, 74, 350, 67, 9, 73, 278, 345,
	348, 277, 349, 347, 71, 21, 107, 26, 66, 64,
	312, 41, 223, 150, 40, 357, 33, 361, 358, 342,
	343, 362, 17, 89, 32, 83, 84, 30, 220, 363,
	85, 106, 109, 337, 19, 178, 28, 18, 368, 6,
	67, 179, 1, 98, 102, 103, 101, 104, 105, 100,
	132, 131, 141, 137, 136, 140, 0, 108, 99, 0,
	0, 89, 0, 83, 84, 359, 0, 0, 85, 279,
	109, 0, 0, 92, 0, 0, 0, 0, 0, 36,
	0, 98, 102, 103, 101, 104, 105, 100, 0, 356,
	0, 0, 0, 0, 0, 108, 99, 0, 0, 98,
	102, 103, 101, 104, 282, 100, 0, 0, 97, 0,
	94, 92, 0, 108, 281, 0, 107, 0, 7, 80,
	0, 329, 70, 82, 0, 79, 142, 143, 0, 0,
	0, 0, 0, 0, 0, 0, 144, 145, 146, 0,
	0, 106, 0, 308, 95, 0, 97, 0, 94, 81,
	0, 0, 0, 0, 107, 0, 90, 80, 142, 143,
	226, 82, 0, 79, 0, 0, 91, 0, 144, 145,
	146, 0, 107, 0, 0, 0, 0, 0, 55, 106,
	142, 143, 95, 53, 0, 46, 0, 81, 0, 0,
	144, 145, 146, 0, 90, 142, 143, 106, 0, 0,
	0, 0, 0, 0, 91, 144, 145, 146, 0, 0,
	55, 0, 54, 0, 0, 53, 0, 115, 0, 0,
	0, 0, 139, 138, 134, 135, 133, 132, 131, 141,
	137, 136, 140, 51, 0, 0, 0, 0, 0, 0,
	254, 255, 0, 0, 54, 0, 0, 0, 0, 0,
	256, 257, 258, 0, 139, 138, 134, 135, 133, 132,
	131, 141, 137, 136, 140, 51, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 139, 138, 134, 135,
	133, 132, 131, 141, 137, 136, 140, 0, 0, 0,
	0, 139, 138, 134, 135, 133, 132, 131, 141, 137,
	136, 140, 142, 143, 0, 0, 0, 0, 0, 0,
	0, 44, 144, 145, 146, 0, 0, 142, 143, 0,
	0, 50, 0, 0, 0, 49, 0, 144, 145, 146,
	0, 0, 0, 0, 0, 48, 251, 250, 246, 247,
	245, 243, 244, 253, 249, 248, 252, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 0, 49, 0, 89,
	0, 0, 0, 0, 341, 0, 85, 48, 109, 0,
	0, 0, 98, 102, 103, 101, 104, 105, 100, 98,
	102, 103, 101, 104, 105, 100, 108, 99, 0, 0,
	158, 0, 0, 108, 99, 0, 0, 155, 139, 138,
	134, 135, 133, 132, 131, 141, 137, 136, 140, 92,
	98, 102, 103, 101, 104, 105, 100, 133, 132, 131,
	141, 137, 136, 140, 108, 99, 0, 0, 0, 0,
	0, 0, 0, 0, 89, 0, 0, 0, 0, 0,
	0, 85, 0, 109, 97, 107, 94, 0, 0, 0,
	0, 0, 107, 0, 98, 102, 103, 101, 104, 105,
	100, 0, 0, 0, 0, 0, 0, 0, 108, 99,
	106, 0, 0, 0, 0, 0, 0, 106, 0, 0,
	95, 0, 0, 107, 92, 0, 0, 0, 0, 149,
	0, 0, 90, 142, 143, 0, 85, 0, 109, 0,
	0, 0, 91, 144, 145, 146, 0, 0, 106, 98,
	102, 103, 101, 104, 105, 100, 0, 0, 0, 97,
	0, 94, 0, 108, 99, 0, 0, 107, 0, 0,
	0, 0, 0, 0, 0, 0, 149, 0, 175, 92,
	156, 0, 176, 85, 0, 109, 0, 161, 162, 165,
	166, 167, 106, 173, 174, 95, 98, 102, 103, 101,
	104, 288, 100, 163, 164, 0, 0, 90, 0, 0,
	108, 287, 0, 0, 97, 0, 94, 91, 0, 0,
	149, 0, 107, 0, 0, 0, 92, 154, 0, 109,
	0, 134, 135, 133, 132, 131, 141, 137, 136, 140,
	98, 102, 103, 101, 104, 105, 100, 106, 0, 0,
	95, 0, 0, 0, 108, 99, 0, 0, 0, 0,
	0, 97, 90, 94, 0, 0, 0, 0, 0, 107,
	92, 0, 91, 0, 160, 0, 169, 170, 171, 168,
	172, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 106, 0, 0, 95, 0, 0,
	0, 0, 0, 0, 0, 97, 0, 94, 0, 90,
	0, 0, 0, 107, 0, 0, 0, 0, 0, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 106, 0,
	0, 95, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 90, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 91,
}
var yyPact = []int{

	8, -1000, -1000, 271, 8, 266, -1000, 268, 140, -1000,
	-101, -12, 218, 8, 277, 8, 266, -1000, 146, 230,
	235, -1000, -130, -1000, 257, -1000, -1000, -1000, -1000, -79,
	-1000, 503, 210, -132, -1000, -136, -1000, 127, 72, -137,
	348, -79, -1000, -143, -1000, -1000, 67, -1000, 535, 277,
	277, 277, -1000, 172, -1000, -1000, -1000, 235, 535, -1000,
	-1000, 759, 759, 123, 348, -1000, -103, -104, -1000, -105,
	277, -1000, -1000, -1000, -1000, -1000, -1000, 584, -1000, 256,
	814, 814, 135, 905, 715, 814, -1000, 827, -1000, 214,
	814, 814, 814, 215, 237, 135, -1000, 236, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 348,
	-1000, 234, -1000, 54, -1000, 535, -1000, 535, 535, 535,
	-106, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	235, 814, 814, 814, 814, 814, 814, 814, 814, 814,
	814, 814, 814, 814, 814, 814, 814, 141, 135, 21,
	584, 135, -1000, -45, 759, 759, -1000, -107, -108, 477,
	759, 759, 759, 759, 759, 759, 759, 759, 759, 759,
	759, 759, 759, 759, 759, 522, 522, 814, 233, 140,
	584, 231, 231, 759, -1000, -1000, -1000, 114, -146, 91,
	535, -1000, -1000, -1000, -1000, 224, 71, -141, -111, -111,
	231, 158, 599, -1000, -1000, 775, 775, -1000, -111, -116,
	-116, -116, -116, -116, 238, -31, 814, -1000, 861, -109,
	-1000, 200, -133, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 759, 759, 759, 759, 759, 759, 759,
	759, 759, 759, 759, 759, 759, 759, 759, 759, -1000,
	462, -1000, 63, -1000, 199, -1000, -1000, 677, 234, -1000,
	-1000, -1000, 14, 535, 25, -1000, -44, -1000, -1000, -1000,
	-1000, 19, 12, 79, 440, 135, 135, 22, 9, 814,
	-110, 759, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 814,
	-1000, -1000, 26, -1000, -1000, 684, 522, 522, -1000, -1000,
	60, -10, 404, 759, 70, 82, -1000, -1000, -1000, -1000,
	-1000, -1000, 69, 77, -128, -1000, -1000, 408, -1000, 234,
	-1000, 759, -1000, -1000, -1000, -1000, 386, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 233, -1000, -1000, 49, -30,
	-1000, -1000, -1000, 166, -1000, 64, 135, 223, -1000, -1000,
}
var yyPgo = []int{

	0, 372, 321, 22, 260, 220, 369, 367, 366, 364,
	357, 354, 346, 82, 10, 12, 344, 66, 341, 105,
	5, 7, 340, 340, 340, 340, 340, 340, 340, 339,
	78, 338, 4, 0, 334, 11, 327, 323, 322, 320,
	17, 319, 8, 318, 317, 315, 14, 1, 314, 313,
	3, 232, 302, 300, 6, 53, 32, 9, 299, 298,
	94, 297, 2, 25, 296, 293, 113,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 4, 6, 7, 9,
	11, 11, 12, 12, 13, 10, 10, 10, 10, 8,
	16, 16, 18, 18, 19, 20, 20, 20, 22, 22,
	23, 23, 24, 24, 25, 26, 26, 27, 27, 28,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 17, 17, 29, 29, 30, 30, 30, 30, 32,
	32, 32, 32, 32, 32, 39, 39, 39, 39, 34,
	34, 41, 41, 42, 45, 45, 44, 44, 35, 35,
	48, 48, 47, 36, 37, 38, 38, 38, 31, 52,
	52, 53, 53, 53, 53, 53, 53, 43, 43, 43,
	43, 43, 43, 43, 57, 57, 49, 49, 14, 58,
	58, 15, 15, 15, 15, 15, 15, 51, 51, 60,
	60, 59, 59, 61, 61, 40, 40, 54, 54, 50,
	50, 63, 63, 63, 62, 62, 62, 62, 62, 62,
	62, 62, 62, 62, 62, 62, 62, 62, 62, 62,
	62, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 33, 33,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	55, 55, 5, 5, 5, 3, 64, 64, 66, 65,
	65, 65,
}
var yyR2 = []int{

	0, 1, 4, 2, 4, 0, 1, 4, 2, 3,
	1, 0, 1, 3, 3, 2, 2, 3, 0, 4,
	1, 0, 1, 2, 4, 1, 3, 4, 1, 3,
	1, 0, 1, 2, 2, 1, 0, 1, 2, 4,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 2, 1, 2, 2, 1, 2, 1,
	1, 1, 1, 1, 1, 5, 2, 3, 3, 5,
	6, 1, 3, 4, 1, 1, 2, 0, 3, 5,
	1, 1, 3, 3, 2, 5, 5, 9, 5, 2,
	0, 2, 4, 6, 2, 2, 0, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 1, 2, 3, 3, 3, 3, 1, 0, 1,
	1, 1, 4, 2, 0, 1, 0, 1, 3, 1,
	0, 1, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 0, 1, 1,
	1, 1, 4, 3, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 1, 2, 2, 1, 1, 2,
	3, 1, 1, 3, 2, 4, 1, 3, 4, 1,
	3, 0,
}
var yyChk = []int{

	-1000, -1, -2, 120, -4, 128, -6, 100, 5, -2,
	-5, 5, 5, -3, 14, 155, 80, 138, -7, -9,
	12, -2, -64, -66, -51, 4, -2, -5, -8, 14,
	-10, 6, -11, -12, -13, -14, 5, -65, 156, 5,
	-16, -18, -19, 135, 138, -15, 12, -59, 162, 152,
	148, 60, -60, 10, 39, 5, 13, 156, 157, 15,
	-66, 61, 157, -17, -29, -30, -31, -4, -32, -33,
	104, -34, -35, -36, -37, -38, -39, -55, -62, 107,
	101, 131, 105, 7, 8, 12, -46, -63, -21, 5,
	138, 148, 55, -5, 92, 126, -47, 90, 25, 40,
	31, 28, 26, 27, 29, 30, 123, 98, 39, 14,
	-19, 158, 13, -58, -15, 12, -15, -51, -51, -51,
	-60, -13, -15, -33, -33, 15, -30, 155, 155, 155,
	-51, 150, 149, 148, 146, 147, 153, 152, 145, 144,
	154, 151, 48, 49, 58, 59, 60, 5, -46, 5,
	-55, -46, -47, -46, 12, 12, 155, -21, 5, -55,
	137, 50, 51, 66, 67, 52, 53, 54, 142, 139,
	140, 141, 143, 56, 57, 41, 45, 10, 151, 157,
	-55, -55, -55, 12, 5, -47, 5, -17, -20, 5,
	156, 13, -15, -15, -15, -61, 155, -14, -55, -55,
	-55, -55, -55, -55, -55, -55, -55, -55, -55, -55,
	-55, -55, -55, -55, 14, -47, 10, -47, 103, -50,
	-62, -40, -54, -33, 155, 155, 13, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -56, 149, 150, 148, 146, 147, 153, 152,
	145, 144, 154, 151, 48, 49, 58, 59, 60, -56,
	-55, -63, 5, -3, -40, 15, 159, 137, 12, -15,
	11, 40, -52, 157, -41, -42, -43, -14, -21, 161,
	-57, 40, 30, 95, -55, -46, -49, 40, 30, 155,
	13, 156, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, 11, 10,
	13, -21, -22, -20, -53, 137, 41, 45, -15, 15,
	156, -44, 146, 101, 62, 62, -48, -47, -35, 11,
	-47, -47, 61, 61, -46, 155, -33, -55, 13, 156,
	-33, 10, -56, -56, 15, -42, 79, -14, -21, -57,
	-33, 40, 30, 40, 30, 155, 11, -20, -54, -55,
	-45, -33, -32, -50, 11, 155, 13, 40, -47, 11,
}
var yyDef = []int{

	5, -2, 1, 0, 5, 0, 6, 0, 0, 3,
	0, 202, 0, 5, 118, 5, 0, 204, 0, 18,
	11, 2, 211, 206, 0, 117, 4, 203, 7, 21,
	8, 0, 0, 10, 12, 0, 108, 0, -2, 0,
	52, 20, 22, 0, 15, 16, 0, 111, 0, 118,
	118, 118, 121, 0, 119, 120, 9, 0, 0, 205,
	207, 0, 0, 0, 51, 54, 0, 0, 57, 0,
	118, 59, 60, 61, 62, 63, 64, 168, 169, 0,
	0, 0, 0, 0, 0, 0, 201, 0, 170, -2,
	0, 0, 0, 0, 194, 0, 197, 198, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 52,
	23, 0, 17, 0, 109, 0, 112, 0, 0, 0,
	124, 13, 14, 210, 208, 19, 53, 55, 56, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 201, -2,
	0, 201, 84, 201, 130, 126, 66, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 167, 167, 0, 0, 0,
	174, 175, 176, 126, 195, 196, 199, 0, 0, 25,
	0, 116, 113, 114, 115, 0, 0, 90, 177, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 190, 191, 192, 0, 78, 0, 83, 0, 0,
	129, 0, 125, 127, 67, 68, 200, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 150,
	0, 133, 131, 173, 0, 82, 24, 0, 0, 110,
	122, 123, 96, 0, 0, 71, 77, 97, 98, 99,
	100, 41, 47, 0, 0, 201, 0, 41, 47, 0,
	0, 0, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 165, 166, -2, 0,
	193, 26, 0, 28, 88, 0, 167, 167, 89, 69,
	0, 0, 0, 0, 0, 0, 79, 80, 81, 172,
	85, 86, 0, 0, 201, 65, 128, 0, 27, 0,
	91, 0, 94, 95, 70, 72, 0, 101, 102, 103,
	76, 104, 105, 106, 107, 130, 132, 29, 0, 168,
	73, 74, 75, 0, 92, 0, 0, 0, 87, 93,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 138, 3, 160, 3, 154, 148, 3,
	3, 3, 152, 149, 156, 150, 151, 153, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 157, 155,
	144, 137, 145, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 158, 3, 159, 147, 161, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 146, 3, 162,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 139, 140, 141, 142, 143,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line ./src/cg_ir/ir_gen.y:339
		{yyVAL.nn=make_node(node{"Code","",[]int{yyS[yypt-0].nn}});make_json(yyVAL.nn);}
	case 2:
		//line ./src/cg_ir/ir_gen.y:343
		{yyVAL.nn=make_node(node{"rust","",[]int{make_node(node{"STRUCT","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 3:
		//line ./src/cg_ir/ir_gen.y:344
		{yyVAL.nn=make_node(node{"rust","",[]int{yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 4:
		//line ./src/cg_ir/ir_gen.y:345
		{yyVAL.nn=make_node(node{"rust","",[]int{make_node(node{"USE","",[]int{}}),yyS[yypt-2].nn,make_node(node{";","",[]int{}}),yyS[yypt-0].nn}})}
	case 5:
		//line ./src/cg_ir/ir_gen.y:346
		{yyVAL.nn=make_node(node{"rust","",[]int{}});}
	case 6:
		//line ./src/cg_ir/ir_gen.y:351
		{yyVAL.nn=make_node(node{"item_or_view_item","",[]int{yyS[yypt-0].nn}})}
	case 7:
		//line ./src/cg_ir/ir_gen.y:355
		{yyVAL.nn=make_node(node{"item_fn","",[]int{make_node(node{"FN","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 8:
		//line ./src/cg_ir/ir_gen.y:359
		{yyVAL.nn=make_node(node{"fn_decl","",[]int{yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 9:
		//line ./src/cg_ir/ir_gen.y:363
		{yyVAL.nn=make_node(node{"fn_args","",[]int{make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
	case 10:
		//line ./src/cg_ir/ir_gen.y:367
		{yyVAL.nn=make_node(node{"maybe_args_general","",[]int{yyS[yypt-0].nn}})}
	case 11:
		//line ./src/cg_ir/ir_gen.y:368
		{yyVAL.nn=make_node(node{"maybe_args_general","",[]int{}})}
	case 12:
		//line ./src/cg_ir/ir_gen.y:372
		{yyVAL.nn=make_node(node{"args_general","",[]int{yyS[yypt-0].nn}})}
	case 13:
		//line ./src/cg_ir/ir_gen.y:373
		{yyVAL.nn=make_node(node{"args_general","",[]int{yyS[yypt-2].nn,make_node(node{",","",[]int{}}),yyS[yypt-0].nn}})}
	case 14:
		//line ./src/cg_ir/ir_gen.y:377
		{yyVAL.nn=make_node(node{"arg_general","",[]int{yyS[yypt-2].nn,make_node(node{":","",[]int{}}),yyS[yypt-0].nn}})}
	case 15:
		//line ./src/cg_ir/ir_gen.y:381
		{yyVAL.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",yyS[yypt-1].s,[]int{}}),make_node(node{"!","",[]int{}})}})}
	case 16:
		//line ./src/cg_ir/ir_gen.y:382
		{yyVAL.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",yyS[yypt-1].s,[]int{}}),yyS[yypt-0].nn}})}
	case 17:
		//line ./src/cg_ir/ir_gen.y:383
		{yyVAL.nn=make_node(node{"ret_ty","",[]int{make_node(node{"OP_INSIDE",yyS[yypt-2].s,[]int{}}),make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
	case 18:
		//line ./src/cg_ir/ir_gen.y:384
		{yyVAL.nn=make_node(node{"ret_ty","",[]int{}})}
	case 19:
		//line ./src/cg_ir/ir_gen.y:388
		{yyVAL.nn=make_node(node{"inner_attrs_and_block","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),yyS[yypt-2].nn,yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
	case 20:
		//line ./src/cg_ir/ir_gen.y:392
		{yyVAL.nn=make_node(node{"maybe_inner_attrs","",[]int{yyS[yypt-0].nn}})}
	case 21:
		//line ./src/cg_ir/ir_gen.y:393
		{yyVAL.nn=make_node(node{"maybe_inner_attrs","",[]int{}})}
	case 22:
		//line ./src/cg_ir/ir_gen.y:397
		{yyVAL.nn=make_node(node{"inner_attrs","",[]int{yyS[yypt-0].nn}})}
	case 23:
		//line ./src/cg_ir/ir_gen.y:398
		{yyVAL.nn=make_node(node{"inner_attrs","",[]int{yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 24:
		//line ./src/cg_ir/ir_gen.y:402
		{yyVAL.nn=make_node(node{"inner_attr","",[]int{make_node(node{"SHEBANG","",[]int{}}),make_node(node{"[","",[]int{}}),yyS[yypt-1].nn,make_node(node{"]","",[]int{}})}})}
	case 25:
		//line ./src/cg_ir/ir_gen.y:407
		{yyVAL.nn=make_node(node{"meta_item","",[]int{yyS[yypt-0].nn}})}
	case 34:
		//line ./src/cg_ir/ir_gen.y:428
		{ yyVAL = yyS[yypt-0]; }
	case 40:
		//line ./src/cg_ir/ir_gen.y:447
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_CHAR",yyS[yypt-0].s,[]int{}})}})}}
	case 41:
		//line ./src/cg_ir/ir_gen.y:448
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_INT",yyS[yypt-0].s,[]int{}})}})}}
	case 42:
		//line ./src/cg_ir/ir_gen.y:449
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_UINT",yyS[yypt-0].s,[]int{}})}})}}
	case 43:
		//line ./src/cg_ir/ir_gen.y:450
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_INT_UNSUFFIXED",yyS[yypt-0].s,[]int{}})}})}}
	case 44:
		//line ./src/cg_ir/ir_gen.y:451
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"FLOAT",yyS[yypt-0].s,[]int{}})}})}}
	case 45:
		//line ./src/cg_ir/ir_gen.y:452
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LIT_FLOAT_UNSUFFIXED",yyS[yypt-0].s,[]int{}})}})}}
	case 46:
		//line ./src/cg_ir/ir_gen.y:453
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LITERAL_STR",yyS[yypt-0].s,[]int{}})}})}}
	case 47:
		//line ./src/cg_ir/ir_gen.y:454
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"LITERAL_CHAR",yyS[yypt-0].s,[]int{}})}})}}
	case 48:
		//line ./src/cg_ir/ir_gen.y:455
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"TRUE",yyS[yypt-0].s,[]int{}})}})}}
	case 49:
		//line ./src/cg_ir/ir_gen.y:456
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"FALSE",yyS[yypt-0].s,[]int{}})}})}}
	case 50:
		//line ./src/cg_ir/ir_gen.y:457
		{{yyVAL.nn=make_node(node{"lit","",[]int{make_node(node{"VAR_TYPE",yyS[yypt-0].s,[]int{}})}})}}
	case 51:
		//line ./src/cg_ir/ir_gen.y:461
		{yyVAL.nn=make_node(node{"maybe_stmts","",[]int{yyS[yypt-0].nn}})}
	case 52:
		//line ./src/cg_ir/ir_gen.y:462
		{yyVAL.nn=make_node(node{"maybe_stmts","",[]int{}})}
	case 53:
		//line ./src/cg_ir/ir_gen.y:466
		{yyVAL.nn=make_node(node{"stmts","",[]int{yyS[yypt-1].nn,yyS[yypt-0].nn}})}
	case 54:
		//line ./src/cg_ir/ir_gen.y:467
		{yyVAL.nn=make_node(node{"stmts","",[]int{yyS[yypt-0].nn}})}
	case 55:
		//line ./src/cg_ir/ir_gen.y:471
		{yyVAL.nn=make_node(node{"stmt","",[]int{yyS[yypt-1].nn,make_node(node{";","",[]int{}})}})}
	case 56:
		//line ./src/cg_ir/ir_gen.y:472
		{{yyVAL.nn=make_node(node{"stmt","",[]int{yyS[yypt-1].nn,make_node(node{";","",[]int{}})}})}}
	case 57:
		//line ./src/cg_ir/ir_gen.y:473
		{yyVAL.nn=make_node(node{"stmt","",[]int{yyS[yypt-0].nn}})}
	case 58:
		//line ./src/cg_ir/ir_gen.y:474
		{yyVAL.nn=make_node(node{"stmt","",[]int{yyS[yypt-1].nn,make_node(node{";","",[]int{}})}})}
	case 59:
		//line ./src/cg_ir/ir_gen.y:479
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 60:
		//line ./src/cg_ir/ir_gen.y:480
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 61:
		//line ./src/cg_ir/ir_gen.y:481
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 62:
		//line ./src/cg_ir/ir_gen.y:482
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 63:
		//line ./src/cg_ir/ir_gen.y:483
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 64:
		//line ./src/cg_ir/ir_gen.y:484
		{yyVAL.nn=make_node(node{"expr_stmt","",[]int{yyS[yypt-0].nn}})}
	case 65:
		//line ./src/cg_ir/ir_gen.y:488
		{yyVAL.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),yyS[yypt-2].nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
	case 66:
		//line ./src/cg_ir/ir_gen.y:489
		{yyVAL.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{";","",[]int{}})}})}
	case 67:
		//line ./src/cg_ir/ir_gen.y:490
		{yyVAL.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),yyS[yypt-1].nn,make_node(node{";","",[]int{}})}})}
	case 68:
		//line ./src/cg_ir/ir_gen.y:491
		{yyVAL.nn=make_node(node{"expr_return","",[]int{make_node(node{"RETURN","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-1].s,[]int{}}),make_node(node{";","",[]int{}})}})}
	case 69:
		//line ./src/cg_ir/ir_gen.y:495
		{yyVAL.nn=make_node(node{"expr_match","",[]int{make_node(node{"MATCH","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-3].s,[]int{}}),make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
	case 70:
		//line ./src/cg_ir/ir_gen.y:496
		{yyVAL.nn=make_node(node{"expr_match","",[]int{make_node(node{"MATCH","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-4].s,[]int{}}),make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),yyS[yypt-2].nn,make_node(node{",","",[]int{}}),make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
	case 71:
		//line ./src/cg_ir/ir_gen.y:500
		{yyVAL.nn=make_node(node{"match_clauses","",[]int{yyS[yypt-0].nn}})}
	case 72:
		//line ./src/cg_ir/ir_gen.y:501
		{yyVAL.nn=make_node(node{"match_clauses","",[]int{yyS[yypt-2].nn,make_node(node{",","",[]int{}}),yyS[yypt-0].nn}})}
	case 73:
		//line ./src/cg_ir/ir_gen.y:505
		{yyVAL.nn=make_node(node{"match_clauses","",[]int{yyS[yypt-3].nn,yyS[yypt-2].nn,make_node(node{"OP_FAT_ARROW","=>",[]int{}}),yyS[yypt-0].nn}})}
	case 74:
		//line ./src/cg_ir/ir_gen.y:509
		{yyVAL.nn=make_node(node{"match_body","",[]int{yyS[yypt-0].nn}})}
	case 75:
		//line ./src/cg_ir/ir_gen.y:510
		{yyVAL.nn=make_node(node{"match_body","",[]int{yyS[yypt-0].nn}})}
	case 76:
		//line ./src/cg_ir/ir_gen.y:514
		{yyVAL.nn=make_node(node{"match_guard","",[]int{make_node(node{"IF","",[]int{}}),yyS[yypt-0].nn}})}
	case 77:
		//line ./src/cg_ir/ir_gen.y:515
		{yyVAL.nn=make_node(node{"match_guard","",[]int{}})}
	case 78:
		//line ./src/cg_ir/ir_gen.y:519
		{  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["after"]="label"+strconv.Itoa(label_num);
	  if(yyS[yypt-1].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-1].code,yyVAL.code); p.next=new(node); p.next.value="ifgoto, je, $2.map["value"], 0, $$.map["after"]";p.next.next=new(node); q=copy_nodes(p.next.next,yyS[yypt-0].code);q.next=new(node); q.next.value="label, $$.map["after"]"; }
	case 79:
		//line ./src/cg_ir/ir_gen.y:523
		{ yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["after"]="label"+strconv.Itoa(label_num);yyVAL.map["true"]="label"+strconv.Itoa(label_num);
	  if(yyS[yypt-3].map==NULL)||(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("Expression or block  not declared in IF statement")};
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-3].code,yyVAL.code); p.next=new(node); p.next.value="ifgoto, je, $2.map["value"], 1, $$.map["true"]";p.next.next=new(node); q=copy_nodes(p.next.next,yyS[yypt-0].code);q.next=new(node); q.next.value="jmp, $$.map["after"]";q.next.next=new(code); q.next.next.value="label, $$.map["true"]";q.next.next.next=new(node);r=copy_nodes(q.next.next.next,yyS[yypt-2].code);r.next=new(node);r.next.value="label, $$.map["after"]";}
	case 80:
		//line ./src/cg_ir/ir_gen.y:529
		{ yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
	  yyVAL.map=yyS[yypt-0].map;yyVAL.code=yyS[yypt-0].code;}
	case 81:
		//line ./src/cg_ir/ir_gen.y:531
		{ yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
	  yyVAL.map=yyS[yypt-0].map;yyVAL.code=yyS[yypt-0].code;}
	case 82:
		//line ./src/cg_ir/ir_gen.y:536
		{yyVAL.nn=make_node(node{"block","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
	case 83:
		//line ./src/cg_ir/ir_gen.y:540
		{yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["begin"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);  if(yyS[yypt-1].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  yyVAL.code=new (node);yyVAL.code.value="label, $$.map["begin"]";p=copy_codes(yyS[yypt-1].code,yyVAL.code);p.next=new(node);p.next.value="ifgoto, je, $2.map["value"], 0, $$.map["after"]";p.next.next=new(node); r=copy_codes(yyS[yypt-0].code,p.next.next);r.next=new(node);r.next.value="jmp, $$.map["begin"]";r.next.next=new(node);r.next.next.value="label, $$.map["after"]";}
	case 84:
		//line ./src/cg_ir/ir_gen.y:545
		{yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["begin"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);  if(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  yyVAL.code=new (node);yyVAL.code.value="label, $$.map["begin"]";yyVAL.next=new(node);p=copy_codes(yyS[yypt-0].code,yyVAL.next);p.next=new(node);p.next.value="jmp, $$.map["begin"]";}
	case 85:
		//line ./src/cg_ir/ir_gen.y:550
		{
	yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["begin"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);  if(yyS[yypt-3].map==NULL)||(yyS[yypt-1].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable/range_di /block not declared");};
	  yyVAL.code=new (node);p=copy_codes(yyS[yypt-3].code,yyVAL.code);p.next=new(node);q=copy_codes(yyS[yypt-1].code,p.next);
	  tmp=make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
	  
	  q.next=new(node);q.next.value="=, tmp.map["value"], 0";
	  q.next.next=new(node);q.next.next.value="label, $$.map["begin"]" ;
	  r=q.next.next;r.next=new(node);r.next.value="ifgoto, jg, tmp.map["value"], $4.map["size"], $$.map["after"]";
	  r.next.next=new(node);r.next.next.value="=, $2.map["value"], $4.map[tmp["value"]]";r.next.next.next=new(node);
	  s=copy_codes(yyS[yypt-0].code,r.next.next.next);s.next=new(node);
	  s.next.value="+, tmp.map["value"], tmp.map["value"], 1";
	  s.next.next=new(node);s.next.next.value="jmp, $$.map["begin"]";
	  t=s.next.next;t.next=new(node);t.next.value="label, $$.map["after"]";
	   }
	case 86:
		//line ./src/cg_ir/ir_gen.y:564
		{yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["begin"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);  if(yyS[yypt-3].map==NULL)||(yyS[yypt-1].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable/range_di /block not declared");};
	  yyVAL.code=new (node);p=copy_codes(yyS[yypt-3].code,yyVAL.code);p.next=new(node);q=copy_codes(yyS[yypt-1].code,p.next);
	  q.next=new(node);q.next.value="=, $2.map["value"], $4.map["start"]";
	  q.next.next=new(node);q.next.next.value="label, $$.map["begin"]" ;
	  r=q.next.next;r.next=new(node);r.next.value="ifgoto, jg, $2.map["value"], $4.map["end"], $$.map["after"]";
	  r.next.next=new(node);
	  s=copy_codes(yyS[yypt-0].code,r.next.next);s.next=new(node);s=copy_codes(yyS[yypt-0].code,s.next);
	  s.next=new(node);s.next.value="+, $2.map["value"], $2.map["value"], 1";
	  s.next.next=new(node);s.next.next.value="jmp, $$.map["begin"]";
	  t=s.next.next;t.next=new(node);t.next.value="label, $$.map["after"]";
	  
	}
	case 87:
		//line ./src/cg_ir/ir_gen.y:576
		{
	yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["begin"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);  (yyS[yypt-4].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable/range_di /block not declared");};
	 yyVAL.code=new (node);p=copy_codes(yyS[yypt-6].code,yyVAL.code);p.next=new(node);q=copy_codes(yyS[yypt-4].code,p.next);q.next=new(node);
	 q.next.value="ifgoto, je, $5.map["value"], 0, $$.map["after"]";r=q.next;r.next=new(node);
	 s=copy_codes(yyS[yypt-0].code,r.next);s.next=new(node);
	 t=copy_codes(yyS[yypt-2].code,s.next);t.next=new(node);
	 t.next.value="jmp, $$.map["begin"]"; u=t.next;u.next=new(temp);
	 u.next.value="label, $$.map["after"]";
	}
	case 88:
		//line ./src/cg_ir/ir_gen.y:588
		{
	  if(yyS[yypt-0]!=NULL) {yyS[yypt-2].map["type"]=yyS[yypt-0].map["type"];yyVAL.code=new(node);yyVAL.code.value="=, $3.map["value"], $5.map["value"]";
	    if(yyS[yypt-1]!=NULL) {if (yyS[yypt-1].map["type"]!=yyS[yypt-0].map["type"]) {log.Fatal("Type mismatch in let expression");} }
	    }
	  if (yyS[yypt-0]==NULL) &&(yyS[yypt-1]!=NULL) {yyS[yypt-2].map["type"]=yyS[yypt-1].map["type"];}
	  if(yyS[yypt-0]==NULL) ||(yyS[yypt-1]==NULL) {log.Fatal("unable to infer enough type information about `_`");}
	}
	case 89:
		//line ./src/cg_ir/ir_gen.y:598
		{}
	case 90:
		//line ./src/cg_ir/ir_gen.y:599
		{yyVAL.nn=make_node(node{"maybe_ty_ascription","",[]int{}})}
	case 91:
		//line ./src/cg_ir/ir_gen.y:603
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),yyS[yypt-0].nn}})}
	case 92:
		//line ./src/cg_ir/ir_gen.y:604
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})}
	case 93:
		//line ./src/cg_ir/ir_gen.y:605
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"=","",[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),yyS[yypt-3].nn,make_node(node{";","",[]int{}}),make_node(node{"LIT_INT",yyS[yypt-1].s,[]int{}}),make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})}
	case 94:
		//line ./src/cg_ir/ir_gen.y:606
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"OPEQ_INT","",[]int{}}),yyS[yypt-0].nn}})}
	case 95:
		//line ./src/cg_ir/ir_gen.y:607
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{make_node(node{"OPEQ_FLOAT","",[]int{}}),yyS[yypt-0].nn}})}
	case 96:
		//line ./src/cg_ir/ir_gen.y:608
		{yyVAL.nn=make_node(node{"maybe_init_expr","",[]int{}})}
	case 97:
		//line ./src/cg_ir/ir_gen.y:612
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-0].nn}})}
	case 98:
		//line ./src/cg_ir/ir_gen.y:613
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-0].nn}})}
	case 99:
		//line ./src/cg_ir/ir_gen.y:614
		{yyVAL.nn=make_node(node{"pats_or","",[]int{make_node(node{"_","",[]int{}})}})}
	case 100:
		//line ./src/cg_ir/ir_gen.y:615
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-0].nn}})}
	case 101:
		//line ./src/cg_ir/ir_gen.y:616
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-2].nn,make_node(node{"|","",[]int{}}),yyS[yypt-0].nn}})}
	case 102:
		//line ./src/cg_ir/ir_gen.y:617
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-2].nn,make_node(node{"|","",[]int{}}),yyS[yypt-0].nn}})}
	case 103:
		//line ./src/cg_ir/ir_gen.y:618
		{yyVAL.nn=make_node(node{"pats_or","",[]int{yyS[yypt-2].nn,make_node(node{"|","",[]int{}}),yyS[yypt-0].nn}})}
	case 104:
		//line ./src/cg_ir/ir_gen.y:622
		{yyVAL.nn=make_node(node{"range_tri","",[]int{make_node(node{"LIT_INT",yyS[yypt-2].s,[]int{}}),make_node(node{"OP_DOTDOTDOT",yyS[yypt-1].s,[]int{}}),make_node(node{"LIT_INT",yyS[yypt-0].s,[]int{}}),}})}
	case 105:
		//line ./src/cg_ir/ir_gen.y:623
		{yyVAL.nn=make_node(node{"range_tri","",[]int{make_node(node{"LITERAL_CHAR",yyS[yypt-2].s,[]int{}}),make_node(node{"OP_DOTDOTDOT",yyS[yypt-1].s,[]int{}}),make_node(node{"LITERAL_CHAR",yyS[yypt-0].s,[]int{}}),}})}
	case 106:
		//line ./src/cg_ir/ir_gen.y:626
		{yyVAL.nn=make_node(node{"range_di","",[]int{make_node(node{"LIT_INT",yyS[yypt-2].s,[]int{}}),make_node(node{"OP_DOTDOT",yyS[yypt-1].s,[]int{}}),make_node(node{"LIT_INT",yyS[yypt-0].s,[]int{}}),}})}
	case 107:
		//line ./src/cg_ir/ir_gen.y:627
		{yyVAL.nn=make_node(node{"range_di","",[]int{make_node(node{"LITERAL_CHAR",yyS[yypt-2].s,[]int{}}),make_node(node{"OP_DOTDOT",yyS[yypt-1].s,[]int{}}),make_node(node{"LITERAL_CHAR",yyS[yypt-0].s,[]int{}}),}})}
	case 108:
		//line ./src/cg_ir/ir_gen.y:631
		{yyS[yypt-0].map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map=yyS[yypt-0].map; }
	case 109:
		//line ./src/cg_ir/ir_gen.y:636
		{yyVAL.nn=make_node(node{"tys","",[]int{yyS[yypt-0].nn}})}
	case 110:
		//line ./src/cg_ir/ir_gen.y:637
		{yyVAL.nn=make_node(node{"tys","",[]int{yyS[yypt-2].nn,make_node(node{",","",[]int{}}),yyS[yypt-0].nn}})}
	case 111:
		//line ./src/cg_ir/ir_gen.y:641
		{yyVAL.nn=make_node(node{"ty","",[]int{yyS[yypt-0].nn}})}
	case 116:
		//line ./src/cg_ir/ir_gen.y:646
		{yyVAL.nn=make_node(node{"ty","",[]int{make_node(node{"SYM_OPEN_ROUND","(",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_ROUND",")",[]int{}})}})}
	case 117:
		//line ./src/cg_ir/ir_gen.y:650
		{yyVAL.nn=make_node(node{"maybe_mut","",[]int{make_node(node{"MUT","",[]int{}})}})}
	case 118:
		//line ./src/cg_ir/ir_gen.y:651
		{yyVAL.nn=make_node(node{"maybe_mut","",[]int{}})}
	case 119:
		//line ./src/cg_ir/ir_gen.y:656
		{yyVAL.nn=make_node(node{"var_types","",[]int{make_node(node{"VAR_TYPE",yyS[yypt-0].s,[]int{}})}})}
	case 120:
		//line ./src/cg_ir/ir_gen.y:657
		{yyVAL.nn=make_node(node{"var_types","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-0].s,[]int{}})}})}
	case 121:
		//line ./src/cg_ir/ir_gen.y:660
		{yyVAL.nn=make_node(node{"path","",[]int{yyS[yypt-0].nn}})}
	case 122:
		//line ./src/cg_ir/ir_gen.y:661
		{yyVAL.nn=make_node(node{"path","",[]int{make_node(node{"SYM_OPEN_SQ","{",[]int{}}),yyS[yypt-2].nn,yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_SQ","}",[]int{}})}})}
	case 123:
		//line ./src/cg_ir/ir_gen.y:665
		{yyVAL.nn=make_node(node{"maybe_size","",[]int{make_node(node{";","",[]int{}}),make_node(node{"LIT_INT",yyS[yypt-0].s,[]int{}})}})}
	case 124:
		//line ./src/cg_ir/ir_gen.y:666
		{yyVAL.nn=make_node(node{"maybe_size","",[]int{}})}
	case 125:
		//line ./src/cg_ir/ir_gen.y:670
		{yyVAL.nn=make_node(node{"maybe_exprs","",[]int{yyS[yypt-0].nn}})}
	case 126:
		//line ./src/cg_ir/ir_gen.y:671
		{yyVAL.nn=make_node(node{"maybe_exprs","",[]int{}})}
	case 127:
		//line ./src/cg_ir/ir_gen.y:675
		{yyVAL.nn=make_node(node{"exprs","",[]int{yyS[yypt-0].nn}})}
	case 128:
		//line ./src/cg_ir/ir_gen.y:676
		{yyVAL.nn=make_node(node{"exprs","",[]int{yyS[yypt-2].nn,make_node(node{",","",[]int{}}),yyS[yypt-0].nn}})}
	case 129:
		//line ./src/cg_ir/ir_gen.y:682
		{yyVAL.nn=make_node(node{"maybe_assignment","",[]int{yyS[yypt-0].nn}})}
	case 130:
		//line ./src/cg_ir/ir_gen.y:683
		{yyVAL.nn=make_node(node{"maybe_assignment","",[]int{}})}
	case 131:
		//line ./src/cg_ir/ir_gen.y:687
		{yyVAL.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-0].s,[]int{}})}})}
	case 132:
		//line ./src/cg_ir/ir_gen.y:688
		{yyVAL.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-3].s,[]int{}}),make_node(node{"SYM_OPEN_SQ","[",[]int{}}),yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_SQ","]",[]int{}})}})}
	case 133:
		//line ./src/cg_ir/ir_gen.y:689
		{yyVAL.nn=make_node(node{"hole","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),make_node(node{".","",[]int{}}),yyS[yypt-0].nn}})}
	case 134:
		//line ./src/cg_ir/ir_gen.y:692
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"=","",[]int{}}),yyS[yypt-0].nn}})}
	case 135:
		//line ./src/cg_ir/ir_gen.y:693
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"+=","OP_ADDEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 136:
		//line ./src/cg_ir/ir_gen.y:694
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"-=","OP_SUBEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 137:
		//line ./src/cg_ir/ir_gen.y:695
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"<=","OP_LEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 138:
		//line ./src/cg_ir/ir_gen.y:696
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{">=","OP_GEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 139:
		//line ./src/cg_ir/ir_gen.y:697
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"*=","OP_MULEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 140:
		//line ./src/cg_ir/ir_gen.y:698
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"/=","OP_DIVEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 141:
		//line ./src/cg_ir/ir_gen.y:699
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"%=","OP_MODEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 142:
		//line ./src/cg_ir/ir_gen.y:700
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"&=","OP_ANDEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 143:
		//line ./src/cg_ir/ir_gen.y:701
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"<<=","OP_SHLEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 144:
		//line ./src/cg_ir/ir_gen.y:702
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{">>=","OP_SHREQ",[]int{}}),yyS[yypt-0].nn}})}
	case 145:
		//line ./src/cg_ir/ir_gen.y:703
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"|=","OP_OREQ",[]int{}}),yyS[yypt-0].nn}})}
	case 146:
		//line ./src/cg_ir/ir_gen.y:704
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"^=","OP_XOREQ",[]int{}}),yyS[yypt-0].nn}})}
	case 147:
		//line ./src/cg_ir/ir_gen.y:705
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"==","OP_EQEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 148:
		//line ./src/cg_ir/ir_gen.y:706
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"!=","OP_NOTEQ",[]int{}}),yyS[yypt-0].nn}})}
	case 149:
		//line ./src/cg_ir/ir_gen.y:707
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"=int","OPEQ_INT",[]int{}}),yyS[yypt-0].nn}})}
	case 150:
		//line ./src/cg_ir/ir_gen.y:708
		{yyVAL.nn=make_node(node{"assignment","",[]int{yyS[yypt-2].nn,make_node(node{"=float","OPEQ_FLOAT",[]int{}}),yyS[yypt-0].nn}})}
	case 151:
		//line ./src/cg_ir/ir_gen.y:716
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"+","",[]int{}}),yyS[yypt-0].nn}})}
	case 152:
		//line ./src/cg_ir/ir_gen.y:717
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"-","",[]int{}}),yyS[yypt-0].nn}})}
	case 153:
		//line ./src/cg_ir/ir_gen.y:718
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"&","",[]int{}}),yyS[yypt-0].nn}})}
	case 154:
		//line ./src/cg_ir/ir_gen.y:719
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"|","",[]int{}}),yyS[yypt-0].nn}})}
	case 155:
		//line ./src/cg_ir/ir_gen.y:720
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"^","",[]int{}}),yyS[yypt-0].nn}})}
	case 156:
		//line ./src/cg_ir/ir_gen.y:721
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"/","",[]int{}}),yyS[yypt-0].nn}})}
	case 157:
		//line ./src/cg_ir/ir_gen.y:722
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"*","",[]int{}}),yyS[yypt-0].nn}})}
	case 158:
		//line ./src/cg_ir/ir_gen.y:723
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{">","",[]int{}}),yyS[yypt-0].nn}})}
	case 159:
		//line ./src/cg_ir/ir_gen.y:724
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"<","",[]int{}}),yyS[yypt-0].nn}})}
	case 160:
		//line ./src/cg_ir/ir_gen.y:725
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"%","",[]int{}}),yyS[yypt-0].nn}})}
	case 161:
		//line ./src/cg_ir/ir_gen.y:726
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{".","",[]int{}}),yyS[yypt-0].nn}})}
	case 162:
		//line ./src/cg_ir/ir_gen.y:727
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_RSHIFT",">>",[]int{}}),yyS[yypt-0].nn}})}
	case 163:
		//line ./src/cg_ir/ir_gen.y:728
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_LSHIFT","<<",[]int{}}),yyS[yypt-0].nn}})}
	case 164:
		//line ./src/cg_ir/ir_gen.y:729
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_ANDAND","&&",[]int{}}),yyS[yypt-0].nn}})}
	case 165:
		//line ./src/cg_ir/ir_gen.y:730
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_OROR","||",[]int{}}),yyS[yypt-0].nn}})}
	case 166:
		//line ./src/cg_ir/ir_gen.y:731
		{yyVAL.nn=make_node(node{"opeq_ops","",[]int{make_node(node{"OP_POWER","**",[]int{}}),yyS[yypt-0].nn}})}
	case 168:
		//line ./src/cg_ir/ir_gen.y:736
		{yyVAL.nn=make_node(node{"expr","",[]int{yyS[yypt-0].nn}})}
	case 169:
		//line ./src/cg_ir/ir_gen.y:737
		{yyVAL.nn=make_node(node{"expr","",[]int{yyS[yypt-0].nn}})}
	case 170:
		//line ./src/cg_ir/ir_gen.y:743
		{yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
	  yyVAL.map["value"]=yyS[yypt-0].map["value"];yyVAL.map["place"]=yyS[yypt-0].map["place"];yyVAL.map["type"]=yyS[yypt-0].map["type"];}
	case 171:
		//line ./src/cg_ir/ir_gen.y:746
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;
	  yyVAL.map["value"]=yyS[yypt-0].map["value"];yyVAL.map["place"]=yyS[yypt-0].map["place"];yyVAL.map["type"]=yyS[yypt-0].map["type"];yyVAL.code=yyS[yypt-0].code;}
	case 172:
		//line ./src/cg_ir/ir_gen.y:749
		{}
	case 173:
		//line ./src/cg_ir/ir_gen.y:750
		{yyVAL.nn=make_node(node{"exp","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),make_node(node{":","",[]int{}}),yyS[yypt-0].nn}})}
	case 174:
		//line ./src/cg_ir/ir_gen.y:751
		{}
	case 175:
		//line ./src/cg_ir/ir_gen.y:752
		{yyVAL.nn=make_node(node{"exp","",[]int{make_node(node{"&","",[]int{}}),yyS[yypt-0].nn}})}
	case 176:
		//line ./src/cg_ir/ir_gen.y:753
		{yyVAL.nn=make_node(node{"exp","",[]int{make_node(node{"OP_ANDMUT","&mut",[]int{}}),yyS[yypt-0].nn}})}
	case 177:
		//line ./src/cg_ir/ir_gen.y:754
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="-, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 178:
		//line ./src/cg_ir/ir_gen.y:760
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="+, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 179:
		//line ./src/cg_ir/ir_gen.y:766
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="&, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 180:
		//line ./src/cg_ir/ir_gen.y:771
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="|, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 181:
		//line ./src/cg_ir/ir_gen.y:777
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="^, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 182:
		//line ./src/cg_ir/ir_gen.y:782
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="/, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 183:
		//line ./src/cg_ir/ir_gen.y:788
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="*, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 184:
		//line ./src/cg_ir/ir_gen.y:794
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"]; yyVAL.map["true"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);
	    yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="ifgoto, jg, $1.map["value"], $3.map["value"], $$.map["true"]";r=new(node);q.next.next=r;r.value="=, $$.map["value"], 0";r.next=new(node);r.next.value="jmp, $$.map["after"]";r.next.next=new(node);s=r.next.next;s.value="label, $$.map["true"]";s.next=new(node);s.next.value="=, $$.map["value"], 1";s.next.next=new(node);s.next.next.value="label, $$.map["after"]"; }
	case 185:
		//line ./src/cg_ir/ir_gen.y:800
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];yyVAL.map["true"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);
	    yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="ifgoto, jl, $1.map["value"], $3.map["value"], $$.map["true"]";r=new(node);q.next.next=r;r.value="=, $$.map["value"], 0";r.next=new(node);r.next.value="jmp, $$.map["after"]";r.next.next=new(node);s=r.next.next;s.value="label, $$.map["true"]";s.next=new(node);s.next.value="=, $$.map["value"], 1";s.next.next=new(node);s.next.next.value="label, $$.map["after"]"; }
	case 186:
		//line ./src/cg_ir/ir_gen.y:806
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"];
	  yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="%, $$.map["value"],$1.map["value"], $3.map["value"]"; }
	case 187:
		//line ./src/cg_ir/ir_gen.y:812
		{yyVAL.nn=make_node(node{"exp","",[]int{yyS[yypt-2].nn,make_node(node{".","",[]int{}}),yyS[yypt-0].nn}})}
	case 190:
		//line ./src/cg_ir/ir_gen.y:815
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"]; yyVAL.map["false"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);
	    yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};
	  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="ifgoto, je, $1.map["value"], 0, $$.map["false"]";r=new(node);q.next.next=r;r.value="ifgoto, je, $3.map["value"], 0, $$.map["false"]";r.next=new(node);rr=r.next; rr.value="=, $$.map["value"], 1";rr.next=new(node);rr.next.value="jmp, $$.map["after"]";rr.next.next=new(node);s=rr.next.next;s.value="label, $$.map["false"]";s.next=new(node);s.next.value="=, $$.map["value"], 0";s.next.next=new(node);s.next.next.value="label, $$.map["after"]"; }
	case 191:
		//line ./src/cg_ir/ir_gen.y:821
		{
	  yyVAL.map= make_entry("temp"+strconv.Itoa(temp_num));temp_num+=1;yyVAL.map["type"]=yyS[yypt-2].map["type"]; yyVAL.map["true"]="label"+strconv.Itoa(label_num);yyVAL.map["after"]="label"+strconv.Itoa(label_num);
	    yyVAL.code=new (node);p=copy_nodes(yyS[yypt-2].code,yyVAL.code); p.next=new(node);q=copy_nodes(p.next,yyS[yypt-0].code);q.next=new(node);if(yyS[yypt-2].map==NULL)||(yyS[yypt-0].map==NULL) {log.Fatal("variable not declared")};  if(yyS[yypt-0].map["type"]!=yyS[yypt-2].map["type"]) {log.Fatal("Type Mismatch")};
	  q.next.value="ifgoto, je, $1.map["value"], 1, $$.map["true"]";r=new(node);q.next.next=r;r.value="ifgoto, je, $3.map["value"], 1, $$.map["true"]";r.next=new(node);rr=r.next; rr.value="=, $$.map["value"], 0";rr.next=new(node);rr.next.value="jmp, $$.map["after"]";rr.next.next=new(node);s=rr.next.next;s.value="label, $$.map["true"]";s.next=new(node);s.next.value="=, $$.map["value"], 1";s.next.next=new(node);s.next.next.value="label, $$.map["after"]"; }
	case 195:
		//line ./src/cg_ir/ir_gen.y:829
		{yyVAL.nn=make_node(node{"exp","",[]int{make_node(node{"CONTINUE","",[]int{}}),make_node(node{"IDENTIFIER",yyS[yypt-0].s,[]int{}})}}) }
	case 196:
		//line ./src/cg_ir/ir_gen.y:830
		{yyVAL.nn=make_node(node{"exp","",[]int{make_node(node{"UNSAFE","",[]int{}}),yyS[yypt-0].nn}}) }
	case 197:
		//line ./src/cg_ir/ir_gen.y:831
		{yyVAL.map=yyS[yypt-0].map;yyVAL.code=yyS[yypt-0].code;}
	case 198:
		//line ./src/cg_ir/ir_gen.y:832
		{}
	case 199:
		//line ./src/cg_ir/ir_gen.y:833
		{}
	case 200:
		//line ./src/cg_ir/ir_gen.y:837
		{yyVAL.map=yyS[yypt-2].map;yyVAL.code=yyS[yypt-2].code;}
	case 201:
		//line ./src/cg_ir/ir_gen.y:838
		{yyVAL.map=yyS[yypt-0].map;yyVAL.code=yyS[yypt-0].code;}
	case 202:
		//line ./src/cg_ir/ir_gen.y:842
		{yyVAL.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-0].s,[]int{}})}})}
	case 203:
		//line ./src/cg_ir/ir_gen.y:843
		{yyVAL.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),make_node(node{"SYM_COLCOL","::",[]int{}}),yyS[yypt-0].nn}})}
	case 204:
		//line ./src/cg_ir/ir_gen.y:844
		{yyVAL.nn=make_node(node{"func_identifier","",[]int{make_node(node{"IDENTIFIER",yyS[yypt-1].s,[]int{}}),make_node(node{"!","",[]int{}})}})}
	case 205:
		//line ./src/cg_ir/ir_gen.y:848
		{yyVAL.nn=make_node(node{"struct_expr","",[]int{make_node(node{"SYM_OPEN_CURLY","{",[]int{}}),yyS[yypt-2].nn,yyS[yypt-1].nn,make_node(node{"SYM_CLOSE_CURLY","}",[]int{}})}})}
	case 206:
		//line ./src/cg_ir/ir_gen.y:852
		{yyVAL.nn=make_node(node{"field_inits","",[]int{yyS[yypt-0].nn}})}
	case 207:
		//line ./src/cg_ir/ir_gen.y:853
		{yyVAL.nn=make_node(node{"field_inits","",[]int{yyS[yypt-2].nn,make_node(node{",","",[]int{}}),yyS[yypt-0].nn}})}
	case 208:
		//line ./src/cg_ir/ir_gen.y:857
		{yyVAL.nn=make_node(node{"field_init","",[]int{yyS[yypt-3].nn,make_node(node{"IDENTIFIER",yyS[yypt-2].s,[]int{}}),make_node(node{":","",[]int{}}),yyS[yypt-0].nn}})}
	case 209:
		//line ./src/cg_ir/ir_gen.y:861
		{yyVAL.nn=make_node(node{"default_field_init","",[]int{make_node(node{",","",[]int{}})}})}
	case 210:
		//line ./src/cg_ir/ir_gen.y:862
		{yyVAL.nn=make_node(node{"default_field_init","",[]int{make_node(node{",","",[]int{}}),make_node(node{"OP_DOTDOT","..",[]int{}}),yyS[yypt-0].nn}})}
	case 211:
		//line ./src/cg_ir/ir_gen.y:863
		{yyVAL.nn=make_node(node{"default_field_init","",[]int{}})}
	}
	goto yystack /* stack new state and value */
}
