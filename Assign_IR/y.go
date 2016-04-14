//line ./src/cg_ir/ir_gen.y:2
package main

import __yyfmt__ "fmt"

//line ./src/cg_ir/ir_gen.y:2
import "../Assign_IR/src/symtable"
import "fmt"
import "log"

/* import "os" */
import "strconv"
import "strings"

var line = 0
var temp_num = 0
var label_num = 0

func print_ircode(a *node) {
	if a == nil {
		fmt.Println("No code to print")
		return
	}
	fmt.Println(a.value)
	for a.next != nil {
		a = a.next
		fmt.Println(a.value)
	}
}

func list_end(l **node) *node {
	if (*l) == nil {
		(*l) = new(node)
		return *l
	}
	p := *l
	for p.next != nil {
		p = p.next
	}
	return p
}

type node struct {
	value string
	next  *node
}

func copy_nodes(a *node, b *node) *node {
	if a == nil {
		return b
	}
	b.value = a.value
	for a.next != nil {
		b.next = new(node)
		b = b.next
		a = a.next
		b.value = a.value
	}
	return b
}

func space(a string, i int) int {
	for ; a[i] == ' '; i++ {
	}
	return i
}

func btoi(a bool) int64 {
	if a == false {
		return 0
	}
	return 1
}

func itob(a int64) bool {
	if a == 0 {
		return false
	}
	return true
}

//line ./src/cg_ir/ir_gen.y:71
type yySymType struct {
	yys  int
	code *node
	mp   map[string]string
	n    int
	s    string
	b    bool
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

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
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
	"'='",
	"'!'",
	"OP_SHLEQ",
	"OP_SHREQ",
	"OP_OREQ",
	"OP_ANDEQ",
	"OP_XOREQ",
	"'<'",
	"'>'",
	"'|'",
	"'^'",
	"'&'",
	"'+'",
	"'-'",
	"'.'",
	"'*'",
	"'/'",
	"'%'",
	"';'",
	"','",
	"':'",
	"'['",
	"']'",
	"'_'",
	"'~'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 58,
	15, 200,
	-2, 109,
	-1, 86,
	12, 193,
	41, 122,
	45, 122,
	50, 122,
	51, 122,
	52, 122,
	53, 122,
	54, 122,
	56, 122,
	57, 122,
	137, 122,
	139, 122,
	140, 122,
	141, 122,
	142, 122,
	143, 122,
	-2, 160,
	-1, 148,
	14, 61,
	-2, 191,
	-1, 149,
	12, 193,
	-2, 160,
	-1, 302,
	41, 123,
	45, 123,
	50, 123,
	51, 123,
	52, 123,
	53, 123,
	54, 123,
	56, 123,
	57, 123,
	137, 123,
	139, 123,
	140, 123,
	141, 123,
	142, 123,
	143, 123,
	-2, 161,
	-1, 335,
	12, 193,
	-2, 160,
	-1, 378,
	15, 206,
	156, 206,
	-2, 191,
}

const yyNprod = 207
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1097

var yyAct = [...]int{

	74, 65, 365, 66, 223, 75, 93, 318, 313, 227,
	186, 245, 69, 226, 82, 20, 143, 144, 267, 85,
	31, 108, 374, 274, 123, 56, 145, 146, 147, 352,
	362, 137, 136, 142, 140, 141, 330, 32, 132, 131,
	219, 137, 136, 142, 285, 15, 58, 55, 372, 348,
	338, 189, 346, 328, 283, 230, 229, 95, 99, 100,
	98, 101, 102, 97, 194, 128, 127, 126, 14, 41,
	222, 105, 96, 269, 7, 45, 278, 150, 150, 150,
	342, 150, 84, 218, 160, 354, 153, 351, 178, 179,
	180, 148, 151, 152, 3, 154, 343, 326, 325, 183,
	309, 158, 5, 16, 310, 60, 371, 344, 272, 40,
	106, 345, 139, 138, 134, 135, 133, 132, 131, 130,
	137, 136, 142, 363, 266, 341, 195, 124, 120, 111,
	104, 199, 200, 201, 202, 203, 204, 205, 206, 207,
	208, 209, 210, 211, 212, 213, 214, 215, 216, 107,
	197, 35, 37, 15, 218, 103, 224, 43, 220, 221,
	228, 225, 232, 233, 234, 235, 236, 237, 238, 239,
	240, 241, 242, 243, 244, 285, 113, 263, 176, 331,
	62, 86, 50, 80, 81, 228, 119, 262, 83, 373,
	106, 339, 317, 264, 188, 265, 308, 77, 268, 196,
	106, 95, 99, 100, 98, 101, 102, 97, 4, 122,
	121, 16, 185, 198, 275, 105, 96, 30, 21, 276,
	277, 25, 370, 150, 15, 303, 284, 54, 143, 144,
	177, 89, 181, 53, 117, 19, 376, 279, 145, 146,
	147, 271, 125, 27, 190, 191, 192, 64, 15, 366,
	286, 287, 288, 289, 290, 291, 292, 293, 294, 295,
	296, 297, 298, 299, 300, 301, 94, 52, 91, 187,
	64, 53, 32, 118, 104, 184, 51, 77, 112, 109,
	306, 79, 16, 76, 150, 321, 323, 324, 304, 329,
	182, 322, 59, 11, 90, 316, 315, 12, 327, 103,
	10, 177, 92, 53, 8, 52, 16, 78, 51, 332,
	23, 37, 364, 57, 87, 64, 36, 34, 270, 193,
	110, 336, 337, 333, 88, 177, 49, 135, 133, 132,
	131, 219, 137, 136, 142, 350, 307, 52, 228, 273,
	280, 320, 347, 367, 349, 340, 358, 361, 353, 357,
	314, 360, 225, 312, 217, 73, 369, 72, 368, 316,
	315, 356, 355, 359, 114, 115, 116, 71, 70, 68,
	63, 86, 61, 80, 81, 150, 377, 375, 83, 305,
	106, 39, 319, 38, 129, 29, 28, 26, 18, 378,
	24, 95, 99, 100, 98, 101, 102, 97, 302, 17,
	6, 13, 2, 1, 311, 105, 96, 9, 0, 0,
	0, 0, 0, 0, 48, 231, 0, 22, 47, 143,
	144, 89, 0, 33, 0, 0, 0, 46, 0, 145,
	146, 147, 0, 0, 0, 143, 144, 140, 141, 0,
	231, 0, 0, 0, 0, 145, 146, 147, 0, 0,
	143, 144, 0, 140, 141, 0, 94, 0, 91, 0,
	145, 146, 147, 0, 104, 0, 7, 77, 140, 141,
	67, 79, 0, 76, 0, 143, 144, 0, 0, 0,
	0, 0, 0, 0, 0, 145, 146, 147, 0, 103,
	143, 144, 92, 140, 141, 0, 0, 78, 0, 0,
	145, 146, 147, 0, 87, 0, 0, 0, 140, 141,
	0, 0, 0, 0, 88, 139, 138, 134, 135, 133,
	132, 131, 219, 137, 136, 142, 0, 0, 0, 0,
	0, 139, 138, 134, 135, 133, 132, 131, 219, 137,
	136, 142, 0, 0, 0, 0, 139, 138, 134, 135,
	133, 132, 131, 130, 137, 136, 142, 257, 258, 0,
	0, 0, 0, 0, 0, 0, 0, 259, 260, 261,
	0, 139, 138, 134, 135, 133, 132, 131, 219, 137,
	136, 142, 0, 0, 0, 0, 139, 138, 134, 135,
	133, 132, 131, 130, 137, 136, 142, 143, 144, 0,
	0, 0, 0, 0, 0, 0, 0, 145, 146, 147,
	0, 0, 53, 53, 159, 140, 141, 51, 51, 44,
	112, 156, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 95, 99, 100, 98, 101, 102,
	97, 0, 0, 0, 0, 0, 52, 52, 105, 96,
	0, 0, 0, 254, 253, 249, 250, 248, 246, 247,
	256, 252, 251, 255, 0, 0, 0, 49, 49, 0,
	143, 144, 0, 0, 0, 0, 0, 0, 0, 0,
	145, 146, 147, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 139, 138, 134, 135, 133, 132, 131,
	219, 137, 136, 142, 335, 0, 0, 104, 0, 334,
	0, 83, 0, 106, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 95, 99, 100, 98, 101, 102,
	97, 0, 103, 0, 0, 0, 0, 0, 105, 96,
	0, 0, 0, 0, 0, 42, 0, 0, 0, 0,
	0, 0, 0, 0, 89, 48, 48, 0, 149, 47,
	47, 0, 0, 0, 157, 83, 0, 106, 46, 46,
	133, 132, 131, 219, 137, 136, 142, 0, 95, 99,
	100, 98, 101, 102, 97, 0, 0, 0, 0, 94,
	0, 91, 105, 96, 0, 0, 0, 104, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 86, 89, 143,
	144, 0, 0, 0, 83, 0, 106, 0, 0, 145,
	146, 147, 103, 0, 0, 92, 0, 95, 99, 100,
	98, 101, 102, 97, 0, 0, 0, 87, 0, 149,
	0, 105, 96, 94, 0, 91, 83, 88, 106, 0,
	0, 104, 0, 0, 0, 0, 0, 89, 0, 95,
	99, 100, 98, 101, 282, 97, 0, 0, 0, 0,
	0, 0, 0, 105, 281, 0, 103, 0, 0, 92,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 89,
	0, 87, 94, 0, 91, 0, 149, 0, 0, 0,
	104, 88, 0, 155, 0, 106, 0, 134, 135, 133,
	132, 131, 219, 137, 136, 142, 95, 99, 100, 98,
	101, 102, 97, 0, 94, 103, 91, 0, 92, 0,
	105, 96, 104, 0, 0, 0, 0, 0, 0, 0,
	87, 143, 144, 0, 0, 0, 89, 0, 0, 0,
	88, 145, 146, 147, 174, 0, 0, 103, 175, 0,
	92, 0, 0, 162, 163, 164, 165, 166, 0, 172,
	173, 32, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 88, 91, 0, 0, 0, 0, 0, 104,
	0, 95, 99, 100, 98, 101, 102, 97, 95, 99,
	100, 98, 101, 102, 97, 105, 96, 0, 0, 0,
	0, 0, 105, 96, 103, 0, 0, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 88,
	0, 0, 132, 131, 219, 137, 136, 142, 0, 0,
	161, 0, 168, 169, 170, 167, 171, 0, 0, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 103,
	0, 0, 0, 0, 0, 0, 103,
}
var yyPact = [...]int{

	-26, -1000, -1000, 299, -26, 288, -1000, 292, -1000, -1000,
	-87, -35, 223, 204, -26, 288, -1000, 207, 237, 267,
	-26, 307, -1000, -1000, -1000, -66, -1000, 607, 214, -109,
	-1000, -132, -1000, -1000, -110, -1000, 287, -1000, 366, -66,
	-1000, -137, -1000, -1000, 266, -1000, 608, 307, 307, 307,
	-1000, 228, -1000, -1000, -1000, 267, 608, 113, 148, -133,
	112, 366, -1000, -88, -89, -1000, -90, 307, -1000, -1000,
	-1000, -1000, -1000, -1000, 442, -1000, 753, 753, 753, 186,
	891, 609, -1000, 753, 913, -1000, 168, 753, 753, 753,
	220, 285, 186, -1000, 270, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 366, -1000, 264, -1000,
	38, -1000, 608, -1000, 608, 608, 608, -91, -1000, -1000,
	-1000, -1000, 802, 298, -1000, -1000, -1000, -1000, -1000, 267,
	802, 753, 753, 753, 753, 753, 753, 753, 753, 753,
	753, 753, 753, 753, 753, 753, 753, 753, -1000, 144,
	549, 186, 186, -1000, -33, 802, 802, -1000, -99, -100,
	427, 753, 753, 753, 753, 753, 753, 753, 753, 753,
	753, 753, 753, 753, 509, 509, 753, 204, 549, 893,
	893, 802, -1000, -1000, -1000, 109, -141, 61, 608, -1000,
	-1000, -1000, -1000, 230, 68, -1000, -1000, -134, -1000, -121,
	-121, -121, 893, 180, 622, -1000, -1000, 761, 761, 761,
	761, -1000, -111, -111, -111, -111, -111, 200, 753, 753,
	-19, -1000, 834, -101, 402, -1000, 213, -112, -1000, -1000,
	-1000, -1000, 549, 549, 549, 549, 549, 549, 549, 549,
	549, 549, 549, 549, 549, -1000, 802, 802, 802, 802,
	802, 802, 802, 802, 802, 802, 802, 802, 802, 802,
	802, 802, -1000, 387, -1000, 212, -1000, -1000, 973, 264,
	-1000, -1000, -1000, 59, 608, 32, 371, -121, 96, 186,
	186, 37, 36, 753, -102, 802, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 23, -1000, -1000, 699, 509,
	509, -1000, 35, -1000, -21, -1000, 34, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 67, 81, -103, -1000, -1000,
	-1000, 264, 549, -106, 802, 73, -1000, -1000, -127, 32,
	6, 966, 802, 973, -1000, -1000, 802, -1000, -1000, 19,
	-32, 244, -1000, -1000, 176, -1000, 34, -1000, -1000, -1000,
	209, 442, -1000, 66, 33, -1000, -135, -1000, -1000, -1000,
	186, 225, -1000, 244, 753, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 403, 402, 401, 15, 208, 294, 400, 399, 390,
	388, 387, 386, 385, 217, 20, 129, 383, 105, 381,
	109, 10, 19, 379, 372, 180, 370, 1, 3, 369,
	12, 368, 367, 357, 355, 13, 14, 354, 353, 8,
	350, 345, 343, 6, 341, 340, 4, 316, 339, 336,
	0, 323, 9, 11, 7, 320, 75, 182, 319, 5,
	82, 317, 313, 151, 312, 2,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 3, 5, 7, 8,
	10, 12, 12, 13, 13, 14, 11, 11, 11, 11,
	9, 17, 17, 19, 19, 20, 21, 21, 21, 23,
	23, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 18, 18, 24, 24, 25, 25, 25, 25,
	27, 27, 27, 27, 27, 27, 34, 34, 34, 34,
	29, 37, 38, 38, 39, 42, 42, 41, 41, 30,
	30, 44, 44, 43, 31, 32, 33, 33, 33, 26,
	48, 48, 49, 49, 49, 49, 49, 49, 49, 40,
	40, 40, 40, 40, 40, 40, 54, 45, 45, 15,
	55, 55, 16, 16, 16, 16, 16, 16, 47, 47,
	57, 57, 56, 56, 58, 58, 35, 35, 52, 52,
	46, 46, 60, 60, 60, 59, 59, 59, 59, 59,
	59, 59, 59, 59, 59, 59, 59, 59, 59, 59,
	53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 28, 28, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 50, 50, 6, 6, 6, 4, 61, 61, 63,
	62, 62, 62, 51, 64, 64, 65,
}
var yyR2 = [...]int{

	0, 1, 5, 2, 4, 0, 0, 1, 4, 2,
	3, 1, 0, 1, 3, 3, 2, 2, 3, 0,
	4, 1, 0, 1, 2, 4, 1, 3, 4, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 0, 2, 1, 2, 2, 1, 2,
	1, 1, 1, 1, 1, 1, 5, 2, 3, 3,
	7, 0, 1, 3, 4, 1, 1, 2, 0, 3,
	5, 1, 1, 3, 3, 2, 5, 5, 9, 5,
	2, 0, 2, 3, 4, 6, 2, 2, 0, 1,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 1, 2, 3, 3, 3, 3, 1, 0,
	1, 1, 1, 4, 2, 0, 1, 0, 1, 3,
	1, 0, 1, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 0, 1, 1, 1,
	1, 4, 3, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 1, 2, 2, 1, 1,
	2, 1, 3, 1, 3, 2, 4, 1, 3, 4,
	1, 3, 0, 4, 1, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, 120, -5, 128, -7, 100, 5, -2,
	-6, 5, 5, -3, 155, 80, 138, -8, -10, 12,
	-4, 14, -2, -6, -9, 14, -11, 6, -12, -13,
	-14, -15, 5, -2, -61, -63, -47, 4, -17, -19,
	-20, 135, 138, -16, 12, -56, 161, 152, 148, 60,
	-57, 10, 39, 5, 13, 156, 157, -62, 156, 5,
	-18, -24, -25, -26, -5, -27, -28, 104, -29, -30,
	-31, -32, -33, -34, -50, -59, 107, 101, 131, 105,
	7, 8, -36, 12, -60, -22, 5, 138, 148, 55,
	-6, 92, 126, -43, 90, 25, 40, 31, 28, 26,
	27, 29, 30, 123, 98, 39, 14, -20, 158, 13,
	-55, -16, 12, -16, -47, -47, -47, -57, -14, -16,
	15, -63, 61, 157, 15, -25, 155, 155, 155, -47,
	151, 150, 149, 148, 146, 147, 153, 152, 145, 144,
	66, 67, 154, 48, 49, 58, 59, 60, -36, 5,
	-50, -36, -36, -43, -36, 12, 12, 155, -22, 5,
	-50, 137, 50, 51, 52, 53, 54, 142, 139, 140,
	141, 143, 56, 57, 41, 45, 10, 157, -50, -50,
	-50, 12, 5, -43, 5, -18, -21, 5, 156, 13,
	-16, -16, -16, -58, 155, -28, -56, -15, -60, -50,
	-50, -50, -50, -50, -50, -50, -50, -50, -50, -50,
	-50, -50, -50, -50, -50, -50, -50, -37, 10, 151,
	-43, -43, 103, -46, -50, -59, -35, -52, -28, 155,
	155, 13, -50, -50, -50, -50, -50, -50, -50, -50,
	-50, -50, -50, -50, -50, -53, 149, 150, 148, 146,
	147, 153, 152, 145, 144, 154, 151, 48, 49, 58,
	59, 60, -53, -50, -4, -35, 15, 159, 137, 12,
	-16, 11, 40, -48, 157, 14, -50, -50, 95, -36,
	-45, 40, 30, 155, 13, 156, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, 11, 13, -22, -23, -21, -49, 137, 41,
	45, -16, -38, -39, -40, -15, -22, 160, -54, 11,
	-44, -43, -30, -43, -43, 61, 61, -36, 155, -28,
	13, 156, -50, -51, 10, 5, -53, -53, 15, 156,
	-41, 146, 101, 62, 40, 30, 155, -21, 155, -52,
	-50, 14, 156, -39, 79, -15, -22, -54, -28, -22,
	-46, -50, 11, 155, -64, -65, 5, -42, -28, -27,
	13, 40, 15, 156, 157, -43, 11, -65, -36,
}
var yyDef = [...]int{

	5, -2, 1, 0, 5, 0, 7, 0, 6, 3,
	0, 193, 0, 0, 5, 0, 195, 0, 19, 12,
	5, 109, 4, 194, 8, 22, 9, 0, 0, 11,
	13, 0, 99, 2, 202, 197, 0, 108, 43, 21,
	23, 0, 16, 17, 0, 102, 0, 109, 109, 109,
	112, 0, 110, 111, 10, 0, 0, 0, -2, 0,
	0, 42, 45, 0, 0, 48, 0, 109, 50, 51,
	52, 53, 54, 55, 157, 158, 0, 0, 0, 0,
	0, 0, 191, 0, 0, 159, -2, 0, 0, 0,
	0, 185, 0, 188, 189, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 43, 24, 0, 18,
	0, 100, 0, 103, 0, 0, 0, 115, 14, 15,
	196, 198, 0, 0, 20, 44, 46, 47, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, -2, -2,
	0, 191, 191, 75, 191, 121, 117, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 156, 156, 0, 0, 163, 164,
	165, 117, 186, 187, 190, 0, 0, 26, 0, 107,
	104, 105, 106, 0, 0, 201, 199, 81, 124, 178,
	166, 167, 168, 169, 170, 171, 172, 173, 174, 175,
	176, 177, 179, 180, 181, 182, 183, 0, 0, 0,
	69, 74, 0, 0, 0, 120, 0, 116, 118, 58,
	59, 192, 125, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 139, 0, 162, 0, 73, 25, 0, 0,
	101, 113, 114, 88, 0, 0, 0, 178, 0, 191,
	0, 32, 38, 0, 0, 0, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 149, 150, 151, 152, 153,
	154, 155, -2, 184, 27, 0, 29, 79, 0, 156,
	156, 80, 0, 62, 68, 89, 90, 91, 92, 161,
	70, 71, 72, 76, 77, 0, 0, 191, 56, 119,
	28, 0, 82, 0, 0, -2, 86, 87, 0, 0,
	0, 0, 0, 0, 97, 98, 121, 30, 83, 0,
	157, 0, 60, 63, 0, 93, 94, 95, 67, 96,
	0, 0, 84, 0, 0, 204, 0, 64, 65, 66,
	0, 0, 203, 0, 0, 78, 85, 205, -2,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 138, 3, 3, 3, 154, 148, 3,
	3, 3, 152, 149, 156, 150, 151, 153, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 157, 155,
	144, 137, 145, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 158, 3, 159, 147, 160, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 146, 3, 161,
}
var yyTok2 = [...]int{

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
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:278
		{
			print_ircode(yyDollar[1].code)
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:282
		{
			yyVAL.code = yyDollar[4].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[5].code
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:283
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[2].code
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:288
		{
			fmt.Println("in rust _marker " + yyDollar[0].s)
			yyVAL.mp = symtab.Make_entry(yyDollar[0].s)
			yyVAL.mp["type"] = "struct"
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:291
		{
			yyVAL.code = yyDollar[1].code
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:296
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "func" + yyDollar[2].s
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = new(node)
			yyVAL.code.value = "jmp, " + yyVAL.mp["after"]
			yyVAL.code.next = new(node)
			yyVAL.code.next.value = "label, " + yyVAL.mp["begin"] + ", " + yyVAL.mp["funargs"]
			yyVAL.code.next.next = new(node)
			yyVAL.code.next.next = yyDollar[4].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "label, " + yyVAL.mp["after"]
			if yyDollar[2].s == "main" {
				pp := list_end(&yyVAL.code)
				pp.next = new(node)
				pp.next.value = "exit"
			}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:332
		{
			yyVAL.code = yyDollar[3].code
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:361
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:362
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:363
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:367
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "string, " + yyVAL.mp["value"] + ", " + (yyDollar[1].s[1:])[0:len(yyDollar[1].s)-2] + "\\n\\0"
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:368
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:369
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", 1"
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:370
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", 0"
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:371
		{
			if (yyDollar[1].s == "i8") || (yyDollar[1].s == "i16") || (yyDollar[1].s == "i32") || (yyDollar[1].s == "i64") || (yyDollar[1].s == "isize") || (yyDollar[1].s == "u8") || (yyDollar[1].s == "u16") || (yyDollar[1].s == "u32") || (yyDollar[1].s == "u64") || (yyDollar[1].s == "usize") {
				yyVAL.s = "int"
			} else {
				yyVAL.s = "str"
			}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:376
		{
			yyVAL.code = yyDollar[1].code
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:381
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[2].code
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:382
		{
			yyVAL.code = yyDollar[1].code
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:386
		{
			yyVAL.code = yyDollar[1].code
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:388
		{
			yyVAL.code = yyDollar[1].code
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:389
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:394
		{
			yyVAL.code = yyDollar[1].code
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:395
		{
			yyVAL.code = yyDollar[1].code
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:396
		{
			yyVAL.code = yyDollar[1].code
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:397
		{
			yyVAL.code = yyDollar[1].code
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:398
		{
			yyVAL.code = yyDollar[1].code
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:399
		{
			yyVAL.code = yyDollar[1].code
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:404
		{
			yyVAL.code = new(node)
			yyVAL.code.value = "ret"
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:405
		{
			yyVAL.code = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "ret, " + yyDollar[2].mp["value"]
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:406
		{
			yyVAL.code = new(node)
			yyDollar[2].mp = symtab.Find_id(yyDollar[2].s)
			if yyDollar[2].mp == nil {
				log.Fatal("Returning undefined identifier; ")
			}
			yyVAL.code.value = "ret, " + yyDollar[2].mp["value"]
		}
	case 60:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:413
		{
			yyVAL.code = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = yyDollar[5].code

		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:420
		{
			yyVAL.mp = symtab.Make_entry("case_exp")
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["cae_exp"] + ", " + yyDollar[0].mp["value"]
			yyVAL.mp["after_match"] = "label" + strconv.Itoa(label_num)
			label_num += 1
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:427
		{
			yyVAL.code = yyDollar[1].code
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:428
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:432
		{

			temp := symtab.Find_id("case_exp")
			if yyVAL.s == "_" {
				yyVAL.code = yyDollar[4].code
				p := list_end(&yyVAL.code)
				p.next = new(node)
				p.next.value = "lable" + yyVAL.mp["after_match"]
			}

			if yyVAL.code == nil {
				yyVAL.code = new(node)
				yyVAL.code.value = "ifgoto, jne, " + temp["value"] + ", " + yyDollar[1].mp["value"] + ", label" + strconv.Itoa(label_num)
				yyVAL.code.next = yyDollar[4].code
				p := list_end(&yyVAL.code)
				p.next = new(node)
				p.next.value = "jmp, " + temp["after_match"]
				p.next.next = new(node)
				p.next.next.value = "label, " + strconv.Itoa(label_num)
			}

		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:450
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:451
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:461
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			fmt.Println("printing exp $2.mp", yyDollar[2].mp)
			if yyDollar[2].mp == nil {
				log.Fatal("Bad If   block;;;")
			}

			yyVAL.code = new(node)
			p := yyVAL.code
			if yyDollar[2].code != nil {
				o := copy_nodes(yyDollar[2].code, yyVAL.code)
				o.next = new(node)
				p = o.next
			} else {
				p = yyVAL.code
			}
			p.value = "ifgoto, je, " + yyDollar[2].mp["value"] + ", 0, " + yyVAL.mp["after"]
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			q.next.value = "label, " + yyVAL.mp["after"]

		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:473
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			if yyDollar[2].mp == nil {
				log.Fatal("Expression or block  not declared in IF statement")
			}
			yyVAL.code = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "ifgoto, je, " + yyDollar[2].mp["value"] + ", 1, " + yyVAL.mp["true"]
			p.next.next = new(node)
			p.next.next = yyDollar[5].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			q.next.value = "jmp, " + yyVAL.mp["after"]
			q.next.next = new(node)
			q.next.next.value = "label, " + yyVAL.mp["true"]
			q.next.next.next = new(node)
			q.next.next.next = yyDollar[3].code
			r := list_end(&yyVAL.code)
			r.next = new(node)
			r.next.value = "label, " + yyVAL.mp["after"]
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:487
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:488
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:492
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:496
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.code = new(node)
			yyVAL.code.value = "label, " + yyVAL.mp["begin"]
			yyVAL.code.next = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "ifgoto, je, " + yyDollar[2].mp["value"] + ", 0, " + yyVAL.mp["after"]
			p.next.next = yyDollar[3].code
			r := list_end(&yyVAL.code)
			r.next = new(node)
			r.next.value = "jmp, " + yyVAL.mp["begin"]
			r.next.next = new(node)
			r.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:501
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			if yyDollar[2].code == nil {
				log.Fatal("variable not declared")
			}
			yyVAL.code = new(node)
			yyVAL.code.value = "label, " + yyVAL.mp["begin"]
			yyVAL.code.next = new(node)
			p := copy_nodes(yyDollar[2].code, yyVAL.code.next)
			p.next = new(node)
			p.next.value = "jmp, " + yyVAL.mp["begin"]
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:506
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.code = new(node)
			p := copy_nodes(yyDollar[2].code, yyVAL.code)
			p.next = new(node)
			q := copy_nodes(yyDollar[4].code, p.next)
			tmp := symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1

			q.next = new(node)
			q.next.value = "=, " + tmp["value"] + ", " + "0"
			q.next.next = new(node)
			q.next.next.value = "label, " + yyVAL.mp["begin"]
			r := q.next.next
			r.next = new(node)
			r.next.value = "ifgoto, jg, " + tmp["value"] + ", " + yyDollar[4].mp["size"] + ", " + yyVAL.mp["after"]
			r.next.next = new(node)
			r.next.next.value = "=, " + yyDollar[2].mp["value"] + ", " + yyDollar[4].mp[tmp["value"]]
			r.next.next.next = new(node)
			s := copy_nodes(yyDollar[5].code, r.next.next.next)
			s.next = new(node)
			s.next.value = "+, " + tmp["value"] + ", " + tmp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "jmp, " + yyVAL.mp["begin"]
			t := s.next.next
			t.next = new(node)
			t.next.value = "label, " + yyVAL.mp["after"]
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:520
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.code = yyDollar[2].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			q.next.value = "=, " + yyDollar[2].mp["value"] + ", " + yyDollar[4].mp["start"]
			q.next.next = new(node)
			q.next.next.value = "label, " + yyVAL.mp["begin"]
			r := q.next.next
			r.next = new(node)
			r.next.value = "ifgoto, jg, " + yyDollar[2].mp["value"] + ", " + yyDollar[4].mp["end"] + ", " + yyVAL.mp["after"]

			//r.next.next=$5.code;
			s := list_end(&r)
			s.next = new(node)
			s.next = yyDollar[5].code
			s = list_end(&s)
			s.next = new(node)
			s.next.value = "+, " + yyDollar[2].mp["value"] + ", " + yyDollar[2].mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "jmp, " + yyVAL.mp["begin"]
			t := s.next.next
			t.next = new(node)
			t.next.value = "label, " + yyVAL.mp["after"]

		}
	case 78:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:537
		{

			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.code = yyDollar[3].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "label, " + yyVAL.mp["begin"]
			p.next.next = yyDollar[5].code
			q := list_end(&p)
			q.next = new(node)
			q.next.value = "ifgoto, je, " + yyDollar[5].mp["value"] + ", 0, " + yyVAL.mp["after"]

			q.next.next = yyDollar[9].code
			s := list_end(&q)
			s.next = yyDollar[7].code
			t := list_end(&s)
			t.next = new(node)
			t.next.value = "jmp, " + yyVAL.mp["begin"]
			u := t.next
			u.next = new(node)
			u.next.value = "label, " + yyVAL.mp["after"]

		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:552
		{
			fmt.Println("in let", yyDollar[4].mp, yyDollar[5].s)
			fmt.Println("OOOOOOOOOO", yyDollar[5])
			if yyDollar[3].mp == nil {
				log.Fatal("Variable name not present in let")
			}
			if yyDollar[4].mp == nil {
				if yyDollar[4].s != "" {
					if yyDollar[5].mp != nil {
						/*let mut y:i32 = expr */
						fmt.Println(yyDollar[5].mp["type"], yyDollar[4].s)
						if yyDollar[5].mp["type"] != yyDollar[4].s {
							log.Fatal("Type mismatch in let ;;")
						}
						yyDollar[3].mp["type"] = yyDollar[2].s + yyDollar[5].mp["type"]
						fmt.Println("MMMMMMMMMMMMMMMMMMMM", yyDollar[4].s)
						yyVAL.code = new(node)
						if yyDollar[5].code != nil {
							p := copy_nodes(yyDollar[5].code, yyVAL.code)
							p.next = new(node)
							if yyDollar[5].mp["Array"] == "true" {
								p2 := &p.next
								if yyDollar[5].mp["args"] != "" {
									s2 := strings.Split(yyDollar[5].mp["args"], ", ")
									for i := 0; i < yyDollar[5].n; i++ {
										(*p2).value = "[]=, " + strconv.Itoa(i) + ", " + yyDollar[3].mp["value"] + ", " + s2[i]
										(*p2).next = new(node)
										p2 = &((*p2).next)
									}
								} else {
									for i := 0; i < yyDollar[5].n; i++ {
										(*p2).value = "[]=, " + strconv.Itoa(i) + ", " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
										(*p2).next = new(node)
										p2 = &((*p2).next)
									}
								}
							} else {
								if yyDollar[5].mp["isfunc"] == "true" {
									p.next.value = "=ret, " + yyDollar[3].mp["value"] + ", "
								} else {
									p.next.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
								}
							}

						} else {
							if yyDollar[5].mp["Array"] == "true" {
								p2 := &yyVAL.code
								if yyDollar[5].mp["args"] != "" {
									s2 := strings.Split(yyDollar[5].mp["args"], ", ")
									for i := 0; i < yyDollar[5].n; i++ {
										(*p2).value = "[]=, " + strconv.Itoa(i) + ", " + yyDollar[3].mp["value"] + ", " + s2[i]
										(*p2).next = new(node)
										p2 = &((*p2).next)
									}
								} else {
									for i := 0; i < yyDollar[5].n; i++ {
										(*p2).value = "[]=, " + strconv.Itoa(i) + ", " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
										(*p2).next = new(node)
										p2 = &((*p2).next)
									}
								}
							} else {
								if yyDollar[5].mp["isfunc"] == "true" {
									yyVAL.code.value = "=ret, " + yyDollar[3].mp["value"] + ", "
								} else {
									yyVAL.code.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
								}
							}

						}
					} else { /*let  y:i32 */
						yyDollar[3].mp["type"] = yyDollar[2].s + yyDollar[4].s

					}
				} else { /* let y = 5 */
					fmt.Println("FFFFFFFFFFFFFFFFFFF")
					print_ircode(yyDollar[5].code)
					fmt.Println("FFFFFFFFFFFFFFFFFFF")
					if yyDollar[5].mp == nil {
						log.Fatal("incomplete let expression  ;")
					}
					yyDollar[3].mp["type"] = yyDollar[2].s + yyDollar[5].mp["type"]
					yyVAL.code = new(node)
					yyVAL.code = yyDollar[5].code
					p := list_end(&yyVAL.code)
					p.next = new(node)
					p.next.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
					print_ircode(yyDollar[5].code)
					fmt.Println("FFFFFFFFFFFFFFFFFFF")
				}
			} else {

				if yyDollar[4].mp["type"] != "struct" {
					log.Fatal("struct " + yyDollar[4].mp["value"] + "not defined prior to use;")
				}
				str_slice := strings.Split(yyDollar[5].s, ",")
				yyVAL.code = yyDollar[5].code
				p := list_end(&yyVAL.code)
				temp := symtab.Make_entry(yyDollar[3].mp["value"] + "_" + str_slice[0])
				for i := 0; i < len(str_slice); i += 2 {

					temp = symtab.Make_entry(yyDollar[3].mp["value"] + "_" + str_slice[i])
					p.next = new(node)
					p.next.value = "=, " + temp["value"] + ", " + str_slice[i+1]
					p = p.next
				}
				fmt.Println("in let, elssssss", yyDollar[5].s)
				print_ircode(yyVAL.code)
				fmt.Println("in let, elssssss", yyDollar[5].s)
			}

		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:645
		{
			if yyDollar[2].mp == nil {
				yyVAL.s = yyDollar[2].s
				yyVAL.mp = nil
				yyVAL.code = nil
			} else {
				yyVAL.code = yyDollar[2].code
				yyVAL.mp = yyDollar[2].mp
			}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:646
		{
			yyVAL.s = ""
			yyVAL.mp = nil
			yyVAL.code = nil
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:651
		{
			fmt.Println("jjdddlsddd")
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:652
		{
			fmt.Println("jjdddlsddddqqqqqq")
			yyVAL.code = yyDollar[2].code
			yyVAL.s = yyDollar[2].s
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:654
		{
			fmt.Println("jjdddlsdddww")
			yyVAL.code = yyDollar[3].code
			yyVAL.mp = yyDollar[3].mp
			yyVAL.n = yyDollar[3].n
		}
	case 85:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:656
		{
			fmt.Println("jjdddlsdddeeeeeee")
			yyVAL.code = yyDollar[3].code
			yyVAL.mp = yyDollar[3].mp
			yyVAL.n = yyDollar[5].n
			yyVAL.mp["Array"] = "true"
			yyVAL.mp["type"] = "Array_" + yyVAL.mp["type"] + "_" + strconv.Itoa(yyVAL.n)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:658
		{
			fmt.Println("jjdddlsdddyyyyyyyy")
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.code = new(node)
			yyVAL.code = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			if yyDollar[2].mp["op"] == "" {
				p.next.value = "=" + ", " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n)
			} else {
				p.next.value = yyDollar[2].mp["op"] + ", " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n) + ", " + yyDollar[2].mp["value"]
			}
			yyVAL.mp["type"] = "int"
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:669
		{
			fmt.Println("jjdddlsdddiiiii")
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:670
		{
			yyVAL.s = ""
			fmt.Println("jjdddlsdddmmmmm")
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:674
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = nil
			yyVAL.s = ""
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:675
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
			yyVAL.s = ""
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:676
		{
			yyVAL.s = "_"
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:677
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:678
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:684
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:688
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["start"] = strconv.Itoa(yyDollar[1].n)
			yyVAL.mp["end"] = strconv.Itoa(yyDollar[3].n)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:689
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["start"] = strconv.Itoa((int)(([]rune(yyDollar[1].s))[0]))
			yyVAL.mp["end"] = strconv.Itoa((int)(([]rune(yyDollar[3].s))[0]))
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:693
		{

			yyDollar[1].mp = symtab.Find_id(yyDollar[1].s)
			if yyDollar[1].mp == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
			}
			yyVAL.mp = yyDollar[1].mp
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:711
		{
			yyVAL.s = yyDollar[1].s
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:712
		{
			yyVAL.s = "~" + yyDollar[2].s
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:713
		{
			yyVAL.s = "*" + yyDollar[2].s + yyDollar[3].s
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:714
		{
			yyVAL.s = "&" + yyDollar[2].s + yyDollar[3].s
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:715
		{
			yyVAL.s = "**" + yyDollar[2].s + yyDollar[3].s
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:720
		{
			yyVAL.s = yyDollar[1].s + "_"
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:721
		{
			yyVAL.s = ""
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:726
		{
			yyVAL.mp = nil
			if (yyDollar[1].s == "i8") || (yyDollar[1].s == "i16") || (yyDollar[1].s == "i32") || (yyDollar[1].s == "i64") || (yyDollar[1].s == "isize") || (yyDollar[1].s == "u8") || (yyDollar[1].s == "u16") || (yyDollar[1].s == "u32") || (yyDollar[1].s == "u64") || (yyDollar[1].s == "usize") {
				yyVAL.s = "int"
			} else {
				yyVAL.s = "str"
			}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:728
		{

			yyVAL.mp = symtab.Find_id(yyDollar[1].s)
			fmt.Println("in var_type", yyVAL.mp)
			if yyVAL.mp == nil {
				log.Fatal("var_type not defined,")
			}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:737
		{
			yyVAL.s = yyDollar[1].s
			yyVAL.mp = yyDollar[1].mp
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:738
		{
			yyVAL.s = "Array_" + yyDollar[2].s + "_" + yyDollar[3].s
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:742
		{
			yyVAL.s = strconv.Itoa(yyDollar[2].n)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:747
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:748
		{
			yyVAL.code = nil
			yyVAL.mp = nil
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:752
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp["args"] = yyDollar[1].mp["value"] + ", "
			yyVAL.n = 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:753
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			yyVAL.mp["args"] = yyDollar[1].mp["args"] + yyDollar[3].mp["value"] + ", "
			if len(yyDollar[1].mp["type"]) > 5 && (yyDollar[1].mp["type"])[0:5] == "Array" {
				sss := strings.Split(yyDollar[1].mp["type"], "_")
				yyDollar[1].mp["type"] = sss[1]

			}
			yyVAL.n = yyDollar[1].n + 1
			yyVAL.mp["type"] = "Array_" + yyDollar[1].mp["type"] + "_" + strconv.Itoa(yyVAL.n)
			fmt.Println("LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL", yyVAL.mp["args"])
			yyVAL.mp["Array"] = "true"

		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:767
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:772
		{
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:779
		{
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyDollar[1].mp["value2"] = yyDollar[3].mp["value"]
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:789
		{
			p := symtab.Find_id(yyDollar[1].mp["value"] + "_" + yyDollar[3].mp["value"])
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].mp["value"] + "_" + yyDollar[3].mp["value"])
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}

		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:801
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p)
			q.next = new(node)
			if yyDollar[1].mp["value2"] == "" {
				if yyDollar[3].mp["isfunc"] == "true" {
					q.next.value = "=ret, " + yyDollar[1].mp["value"] + ", "
				} else {
					q.next.value = "=, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
				}
			} else {
				q.next.value = "[]=, " + yyDollar[1].mp["value2"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:803
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "+, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:804
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "-, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:806
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "*, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:807
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "/, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:808
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "%, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:809
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "&, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:814
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "|, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:815
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "^, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:820
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.code = new(node)
			yyVAL.code = yyDollar[3].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			if yyDollar[3].mp["op"] == "" {
				p.next.value = "=" + ", " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[2].n)
			} else {
				p.next.value = yyDollar[3].mp["op"] + ", " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[2].n) + ", " + yyDollar[3].mp["value"]
			}
			p.next.next = new(node)
			p.next.next.value = "=, " + yyDollar[1].mp["value"] + ", " + yyVAL.mp["value"]
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:841
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "+"
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:842
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "-"
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:843
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "&"
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:844
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "|"
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:845
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "^"
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:846
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "/"
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:847
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "*"
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:848
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = ">"
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:849
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "<"
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:850
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "%"
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:851
		{
			fmt.Println("LLLLLLLLLLLLLLLL")
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "."
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:854
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "&&"
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:855
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
			yyVAL.mp["op"] = "||"
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:857
		{
			yyVAL.s = ""
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:861
		{
			fmt.Println("hello in expr")
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:862
		{
			fmt.Println("sadsad")
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:868
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:870
		{
			fmt.Println("jjdddlsvvvvvvv")
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}

		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:880
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = new(node)
			yyVAL.code.value = "=[], " + yyVAL.mp["value"] + ", " + yyDollar[1].s + ", " + yyDollar[3].mp["value"]
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:887
		{
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:891
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "-, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:896
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "+, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:901
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "&, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:905
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "|, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:910
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "^, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:914
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "/, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:919
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "*, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:924
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)

			//q:=copy_nodes(p.next,$3.code);
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}
			q.next.value = "ifgoto, jg, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"] + ", " + yyVAL.mp["true"]
			r := new(node)
			q.next.next = r
			r.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			r.next = new(node)
			r.next.value = "jmp, " + yyVAL.mp["after"]
			r.next.next = new(node)
			s := r.next.next
			s.value = "label, " + yyVAL.mp["true"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:950
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)

			//q:=copy_nodes(p.next,$3.code);
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}
			q.next.value = "ifgoto, jl, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"] + ", " + yyVAL.mp["true"]
			r := new(node)
			q.next.next = r
			r.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			r.next = new(node)
			r.next.value = "jmp, " + yyVAL.mp["after"]
			r.next.next = new(node)
			s := r.next.next
			s.value = "label, " + yyVAL.mp["true"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:975
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)

			//q:=copy_nodes(p.next,$3.code);
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}
			q.next.value = "ifgoto, jle, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"] + ", " + yyVAL.mp["true"]
			r := new(node)
			q.next.next = r
			r.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			r.next = new(node)
			r.next.value = "jmp, " + yyVAL.mp["after"]
			r.next.next = new(node)
			s := r.next.next
			s.value = "label, " + yyVAL.mp["true"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:999
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)

			//q:=copy_nodes(p.next,$3.code);
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}
			q.next.value = "ifgoto, jge, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"] + ", " + yyVAL.mp["true"]
			r := new(node)
			q.next.next = r
			r.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			r.next = new(node)
			r.next.value = "jmp, " + yyVAL.mp["after"]
			r.next.next = new(node)
			s := r.next.next
			s.value = "label, " + yyVAL.mp["true"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1024
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "%, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1030
		{
			fmt.Println("in a.b")
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			q.next.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + "_" + yyDollar[3].mp["value"]

		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1036
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = ">>, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1040
		{ //incorrect
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "<<, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1044
		{

			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["false"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyDollar[1].code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}
			q.next.value = "ifgoto, je, " + yyDollar[1].mp["value"] + ", 0, " + yyVAL.mp["false"]
			r := new(node)
			q.next.next = r
			r.value = "ifgoto, je, " + yyDollar[3].mp["value"] + ", 0, " + yyVAL.mp["false"]
			r.next = new(node)
			rr := r.next
			rr.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			rr.next = new(node)
			rr.next.value = "jmp, " + yyVAL.mp["after"]
			rr.next.next = new(node)
			s := rr.next.next
			s.value = "label, " + yyVAL.mp["false"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1063
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			label_num += 1

			yyVAL.code = yyDollar[1].code
			p := list_end(&yyDollar[1].code)
			p.next = new(node)
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			if (yyDollar[1].mp == nil) || (yyDollar[3].mp == nil) {
				log.Fatal("variable not declared")
			}
			if yyDollar[3].mp["type"] != yyDollar[1].mp["type"] {
				log.Fatal("Type Mismatch")
			}

			q.next.value = "ifgoto, je, " + yyDollar[1].mp["value"] + ", 1, " + yyVAL.mp["true"]
			r := new(node)
			q.next.next = r
			r.value = "ifgoto, je, " + yyDollar[3].mp["value"] + ", 1, " + yyVAL.mp["true"]
			r.next = new(node)
			rr := r.next
			rr.value = "=, " + yyVAL.mp["value"] + ", " + "0"
			rr.next = new(node)
			rr.next.value = "jmp, " + yyVAL.mp["after"]
			rr.next.next = new(node)
			s := rr.next.next
			s.value = "label, " + yyVAL.mp["true"]
			s.next = new(node)
			s.next.value = "=, " + yyVAL.mp["value"] + ", " + "1"
			s.next.next = new(node)
			s.next.next.value = "label, " + yyVAL.mp["after"]
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1081
		{ //incorrect

		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1084
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "arg, " + yyDollar[3].mp["args"] + ", , , "
			q.next.next = new(node)
			q.next.next.value = "call, " + yyDollar[1].mp["value"] + ", " + ", "
			yyVAL.mp["isfunc"] = "true"
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1093
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1094
		{
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1095
		{
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1099
		{
			fmt.Println("jjdddlsdddddcccccc")
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1100
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1104
		{
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}

		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1112
		{
			p := symtab.Find_id(yyDollar[1].s + "_" + yyDollar[3].s)
			if p == nil {
				yyVAL.mp = symtab.Make_entry(yyDollar[1].s + "_" + yyDollar[3].s)
			} else {
				yyVAL.mp = p
			}

		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1119
		{
			p := symtab.Find_id(yyDollar[1].s + "_")
			if p == nil {
				yyVAL.mp = symtab.Make_entry(yyDollar[1].s + "_")
			} else {
				yyVAL.mp = p
			}

		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1129
		{

		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1141
		{
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1151
		{
			yyVAL.s = yyDollar[3].s
			yyVAL.code = yyDollar[3].code

		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1158
		{
			yyVAL.s = yyDollar[1].s
			yyVAL.code = yyDollar[1].code
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1159
		{
			yyVAL.s = yyDollar[1].s + "," + yyDollar[3].s
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:1163
		{
			yyVAL.s = yyDollar[1].s + "," + yyDollar[3].mp["value"]
			yyVAL.code = yyDollar[3].code
		}
	}
	goto yystack /* stack new state and value */
}
