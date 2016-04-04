//line ./src/cg_ir/ir_gen.y:2
package main

import __yyfmt__ "fmt"

//line ./src/cg_ir/ir_gen.y:2
import "../Assign_IR/src/symtable"
import "fmt"
import "log"

/* import "os" */
import "strconv"

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
	fmt.Println(a, b)
	b.value = a.value
	for a.next != nil {
		fmt.Println(a.value)
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

//line ./src/cg_ir/ir_gen.y:72
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
	-1, 38,
	15, 199,
	-2, 108,
	-1, 89,
	12, 192,
	41, 121,
	45, 121,
	50, 121,
	51, 121,
	52, 121,
	53, 121,
	54, 121,
	56, 121,
	57, 121,
	66, 121,
	67, 121,
	137, 121,
	139, 121,
	140, 121,
	141, 121,
	142, 121,
	143, 121,
	-2, 161,
	-1, 149,
	12, 192,
	-2, 161,
	-1, 308,
	41, 122,
	45, 122,
	50, 122,
	51, 122,
	52, 122,
	53, 122,
	54, 122,
	56, 122,
	57, 122,
	66, 122,
	67, 122,
	137, 122,
	139, 122,
	140, 122,
	141, 122,
	142, 122,
	143, 122,
	-2, 162,
}

const yyNprod = 202
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1052

var yyAct = [...]int{

	69, 96, 78, 219, 68, 188, 222, 88, 275, 280,
	35, 72, 114, 266, 86, 111, 273, 221, 142, 143,
	62, 58, 13, 291, 57, 87, 38, 355, 144, 145,
	146, 216, 242, 132, 131, 141, 137, 136, 140, 338,
	319, 137, 136, 140, 45, 335, 289, 225, 224, 196,
	129, 128, 127, 77, 15, 316, 43, 323, 218, 317,
	364, 116, 123, 124, 283, 36, 63, 109, 191, 325,
	16, 122, 55, 309, 346, 344, 333, 53, 65, 115,
	112, 324, 177, 7, 152, 98, 102, 103, 101, 104,
	282, 100, 157, 332, 52, 148, 151, 185, 153, 108,
	281, 16, 322, 3, 367, 268, 54, 42, 353, 351,
	271, 5, 354, 55, 139, 138, 134, 135, 133, 132,
	131, 141, 137, 136, 140, 365, 352, 51, 17, 265,
	192, 193, 194, 125, 150, 150, 59, 150, 109, 159,
	23, 197, 14, 126, 180, 181, 182, 54, 120, 110,
	215, 315, 16, 217, 80, 25, 223, 220, 107, 17,
	214, 227, 228, 229, 230, 231, 232, 233, 234, 235,
	236, 237, 238, 239, 240, 241, 187, 34, 179, 60,
	29, 320, 339, 106, 223, 198, 199, 200, 201, 202,
	203, 204, 205, 206, 207, 208, 209, 210, 211, 212,
	213, 264, 263, 269, 261, 291, 142, 143, 159, 259,
	17, 190, 61, 366, 178, 50, 144, 145, 146, 49,
	279, 93, 278, 178, 310, 277, 290, 10, 48, 179,
	267, 260, 24, 285, 56, 121, 183, 20, 27, 369,
	270, 31, 262, 36, 292, 293, 294, 295, 296, 297,
	298, 299, 300, 301, 302, 303, 304, 305, 306, 307,
	4, 189, 36, 98, 102, 103, 101, 104, 282, 100,
	284, 186, 150, 184, 313, 311, 147, 108, 281, 39,
	11, 12, 117, 118, 119, 327, 318, 330, 331, 55,
	8, 25, 336, 37, 53, 328, 46, 22, 195, 47,
	113, 67, 314, 130, 334, 135, 133, 132, 131, 141,
	137, 136, 140, 272, 286, 326, 340, 360, 321, 276,
	274, 2, 76, 54, 350, 67, 9, 75, 278, 345,
	348, 277, 349, 347, 74, 21, 107, 26, 73, 71,
	66, 64, 223, 150, 51, 357, 312, 361, 358, 342,
	343, 362, 41, 89, 40, 83, 84, 33, 220, 363,
	85, 106, 109, 337, 32, 30, 19, 28, 368, 18,
	67, 6, 1, 98, 102, 103, 101, 104, 105, 100,
	0, 0, 0, 0, 36, 0, 0, 108, 99, 0,
	0, 89, 0, 83, 84, 359, 0, 0, 85, 0,
	109, 0, 0, 92, 98, 102, 103, 101, 104, 282,
	100, 98, 102, 103, 101, 104, 105, 100, 108, 281,
	356, 0, 44, 0, 0, 108, 99, 0, 0, 0,
	0, 0, 50, 0, 0, 0, 49, 0, 97, 0,
	94, 92, 0, 0, 0, 48, 107, 0, 7, 80,
	0, 329, 70, 82, 0, 79, 0, 142, 143, 0,
	0, 0, 0, 0, 0, 0, 308, 144, 145, 146,
	0, 106, 0, 0, 95, 0, 97, 107, 94, 81,
	0, 0, 0, 226, 107, 0, 90, 80, 142, 143,
	0, 82, 0, 79, 0, 0, 91, 0, 144, 145,
	146, 0, 106, 142, 143, 0, 0, 0, 0, 106,
	0, 0, 95, 144, 145, 146, 0, 81, 142, 143,
	0, 0, 0, 0, 90, 0, 0, 0, 144, 145,
	146, 254, 255, 0, 91, 0, 0, 0, 0, 279,
	0, 256, 257, 258, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 139, 138, 134, 135, 133, 132, 131,
	141, 137, 136, 140, 142, 143, 0, 0, 0, 0,
	0, 0, 0, 0, 144, 145, 146, 0, 0, 0,
	0, 0, 0, 0, 139, 138, 134, 135, 133, 132,
	131, 141, 137, 136, 140, 0, 0, 0, 0, 139,
	138, 134, 135, 133, 132, 131, 141, 137, 136, 140,
	0, 0, 0, 0, 139, 138, 134, 135, 133, 132,
	131, 141, 137, 136, 140, 0, 0, 251, 250, 246,
	247, 245, 243, 244, 253, 249, 248, 252, 0, 0,
	0, 0, 0, 55, 0, 0, 142, 143, 53, 0,
	115, 0, 0, 0, 0, 0, 144, 145, 146, 0,
	139, 138, 134, 135, 133, 132, 131, 141, 137, 136,
	140, 89, 0, 142, 143, 0, 341, 54, 85, 0,
	109, 0, 0, 144, 145, 146, 0, 0, 0, 0,
	0, 98, 102, 103, 101, 104, 105, 100, 51, 0,
	0, 0, 158, 0, 0, 108, 99, 0, 0, 155,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 92, 98, 102, 103, 101, 104, 105, 100, 0,
	0, 0, 0, 0, 0, 0, 108, 99, 0, 0,
	0, 0, 0, 0, 0, 0, 89, 132, 131, 141,
	137, 136, 140, 85, 0, 109, 97, 0, 94, 0,
	0, 0, 0, 0, 107, 0, 98, 102, 103, 101,
	104, 105, 100, 133, 132, 131, 141, 137, 136, 140,
	108, 99, 0, 0, 0, 0, 50, 0, 0, 106,
	49, 0, 95, 0, 0, 107, 92, 0, 0, 48,
	0, 149, 0, 0, 90, 142, 143, 0, 85, 0,
	109, 0, 0, 0, 91, 144, 145, 146, 0, 0,
	106, 98, 102, 103, 101, 104, 105, 100, 0, 0,
	0, 97, 0, 94, 0, 108, 99, 0, 0, 107,
	0, 0, 0, 0, 0, 0, 0, 0, 149, 0,
	175, 92, 156, 0, 176, 85, 0, 109, 0, 161,
	162, 165, 166, 167, 106, 173, 174, 95, 98, 102,
	103, 101, 104, 288, 100, 163, 164, 0, 0, 90,
	0, 0, 108, 287, 0, 0, 97, 0, 94, 91,
	0, 0, 149, 0, 107, 0, 0, 0, 92, 154,
	0, 109, 0, 134, 135, 133, 132, 131, 141, 137,
	136, 140, 98, 102, 103, 101, 104, 105, 100, 106,
	0, 0, 95, 0, 0, 0, 108, 99, 0, 0,
	0, 0, 0, 97, 90, 94, 0, 0, 0, 0,
	0, 107, 92, 0, 91, 0, 160, 0, 169, 170,
	171, 168, 172, 98, 102, 103, 101, 104, 105, 100,
	0, 0, 0, 0, 0, 0, 106, 108, 99, 95,
	0, 0, 0, 0, 0, 0, 0, 97, 0, 94,
	0, 90, 0, 0, 0, 107, 0, 0, 0, 0,
	0, 91, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	106, 0, 0, 95, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 90, 107, 0, 0, 0,
	0, 0, 0, 0, 0, 91, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 106,
}
var yyPact = [...]int{

	-17, -1000, -1000, 285, -17, 275, -1000, 276, 128, -1000,
	-101, -10, 225, -17, 287, -17, 275, -1000, 166, 235,
	257, -1000, -130, -1000, 274, -1000, -1000, -1000, -1000, -79,
	-1000, 284, 221, -132, -1000, -136, -1000, 121, 151, -137,
	348, -79, -1000, -143, -1000, -1000, 67, -1000, 638, 287,
	287, 287, -1000, 108, -1000, -1000, -1000, 257, 638, -1000,
	-1000, 741, 741, 118, 348, -1000, -103, -104, -1000, -105,
	287, -1000, -1000, -1000, -1000, -1000, -1000, 516, -1000, 271,
	796, 796, 124, 887, 697, 796, -1000, 809, -1000, 72,
	796, 796, 796, 224, 268, 124, -1000, 266, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 348,
	-1000, 256, -1000, 55, -1000, 638, -1000, 638, 638, 638,
	-106, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	257, 796, 796, 796, 796, 796, 796, 796, 796, 796,
	796, 796, 796, 796, 796, 796, 796, 146, 124, 21,
	516, 124, -1000, -45, 741, 741, -1000, -107, -108, 470,
	741, 741, 741, 741, 741, 741, 741, 741, 741, 741,
	741, 741, 741, 741, 741, 483, 483, 796, 237, 128,
	516, 598, 598, 741, -1000, -1000, -1000, 114, -146, 93,
	638, -1000, -1000, -1000, -1000, 229, 70, -141, -111, -111,
	598, 158, 625, -1000, -1000, 757, 757, -1000, -111, -116,
	-116, -116, -116, -116, 379, -31, 796, -1000, 843, -109,
	-1000, 213, -133, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 741, 741, 741, 741, 741, 741, 741,
	741, 741, 741, 741, 741, 741, 741, 741, 741, -1000,
	455, -1000, 63, -1000, 211, -1000, -1000, 928, 256, -1000,
	-1000, -1000, 14, 638, 25, -1000, -44, -1000, -1000, -1000,
	-1000, 19, 7, 53, 440, 124, 124, 32, 15, 796,
	-110, 741, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 796,
	-1000, -1000, 26, -1000, -1000, 666, 483, 483, -1000, -1000,
	60, -5, 238, 741, 69, 96, -1000, -1000, -1000, -1000,
	-1000, -1000, 68, 82, -128, -1000, -1000, 409, -1000, 256,
	-1000, 741, -1000, -1000, -1000, -1000, 386, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 237, -1000, -1000, 49, -30,
	-1000, -1000, -1000, 200, -1000, 64, 124, 228, -1000, -1000,
}
var yyPgo = [...]int{

	0, 372, 321, 22, 260, 221, 371, 369, 367, 366,
	365, 364, 357, 177, 10, 12, 354, 66, 352, 107,
	5, 7, 346, 341, 78, 340, 4, 0, 339, 11,
	338, 334, 327, 322, 17, 320, 8, 319, 318, 317,
	14, 1, 315, 314, 3, 232, 313, 302, 6, 53,
	32, 9, 300, 299, 94, 298, 2, 25, 297, 293,
	140,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 4, 6, 7, 9,
	11, 11, 12, 12, 13, 10, 10, 10, 10, 8,
	16, 16, 18, 18, 19, 20, 20, 20, 22, 22,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 17, 17, 23, 23, 24, 24, 24, 24, 26,
	26, 26, 26, 26, 26, 33, 33, 33, 33, 28,
	28, 35, 35, 36, 39, 39, 38, 38, 29, 29,
	42, 42, 41, 30, 31, 32, 32, 32, 25, 46,
	46, 47, 47, 47, 47, 47, 47, 37, 37, 37,
	37, 37, 37, 37, 51, 51, 43, 43, 14, 52,
	52, 15, 15, 15, 15, 15, 15, 45, 45, 54,
	54, 53, 53, 55, 55, 34, 34, 48, 48, 44,
	44, 57, 57, 57, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 27, 27,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	49, 49, 5, 5, 5, 3, 58, 58, 60, 59,
	59, 59,
}
var yyR2 = [...]int{

	0, 1, 4, 2, 4, 0, 1, 4, 2, 3,
	1, 0, 1, 3, 3, 2, 2, 3, 0, 4,
	1, 0, 1, 2, 4, 1, 3, 4, 1, 3,
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
var yyChk = [...]int{

	-1000, -1, -2, 120, -4, 128, -6, 100, 5, -2,
	-5, 5, 5, -3, 14, 155, 80, 138, -7, -9,
	12, -2, -58, -60, -45, 4, -2, -5, -8, 14,
	-10, 6, -11, -12, -13, -14, 5, -59, 156, 5,
	-16, -18, -19, 135, 138, -15, 12, -53, 161, 152,
	148, 60, -54, 10, 39, 5, 13, 156, 157, 15,
	-60, 61, 157, -17, -23, -24, -25, -4, -26, -27,
	104, -28, -29, -30, -31, -32, -33, -49, -56, 107,
	101, 131, 105, 7, 8, 12, -40, -57, -21, 5,
	138, 148, 55, -5, 92, 126, -41, 90, 25, 40,
	31, 28, 26, 27, 29, 30, 123, 98, 39, 14,
	-19, 158, 13, -52, -15, 12, -15, -45, -45, -45,
	-54, -13, -15, -27, -27, 15, -24, 155, 155, 155,
	-45, 150, 149, 148, 146, 147, 153, 152, 145, 144,
	154, 151, 48, 49, 58, 59, 60, 5, -40, 5,
	-49, -40, -41, -40, 12, 12, 155, -21, 5, -49,
	137, 50, 51, 66, 67, 52, 53, 54, 142, 139,
	140, 141, 143, 56, 57, 41, 45, 10, 151, 157,
	-49, -49, -49, 12, 5, -41, 5, -17, -20, 5,
	156, 13, -15, -15, -15, -55, 155, -14, -49, -49,
	-49, -49, -49, -49, -49, -49, -49, -49, -49, -49,
	-49, -49, -49, -49, 14, -41, 10, -41, 103, -44,
	-56, -34, -48, -27, 155, 155, 13, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -50, 149, 150, 148, 146, 147, 153, 152,
	145, 144, 154, 151, 48, 49, 58, 59, 60, -50,
	-49, -57, 5, -3, -34, 15, 159, 137, 12, -15,
	11, 40, -46, 157, -35, -36, -37, -14, -21, 160,
	-51, 40, 30, 95, -49, -40, -43, 40, 30, 155,
	13, 156, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, 11, 10,
	13, -21, -22, -20, -47, 137, 41, 45, -15, 15,
	156, -38, 146, 101, 62, 62, -42, -41, -29, 11,
	-41, -41, 61, 61, -40, 155, -27, -49, 13, 156,
	-27, 10, -50, -50, 15, -36, 79, -14, -21, -51,
	-27, 40, 30, 40, 30, 155, 11, -20, -48, -49,
	-39, -27, -26, -44, 11, 155, 13, 40, -41, 11,
}
var yyDef = [...]int{

	5, -2, 1, 0, 5, 0, 6, 0, 0, 3,
	0, 192, 0, 5, 108, 5, 0, 194, 0, 18,
	11, 2, 201, 196, 0, 107, 4, 193, 7, 21,
	8, 0, 0, 10, 12, 0, 98, 0, -2, 0,
	42, 20, 22, 0, 15, 16, 0, 101, 0, 108,
	108, 108, 111, 0, 109, 110, 9, 0, 0, 195,
	197, 0, 0, 0, 41, 44, 0, 0, 47, 0,
	108, 49, 50, 51, 52, 53, 54, 158, 159, 0,
	0, 0, 0, 0, 0, 0, 191, 0, 160, -2,
	0, 0, 0, 0, 184, 0, 187, 188, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 42,
	23, 0, 17, 0, 99, 0, 102, 0, 0, 0,
	114, 13, 14, 200, 198, 19, 43, 45, 46, 48,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 191, -2,
	0, 191, 74, 191, 120, 116, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 157, 157, 0, 0, 0,
	164, 165, 166, 116, 185, 186, 189, 0, 0, 25,
	0, 106, 103, 104, 105, 0, 0, 80, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 178,
	179, 180, 181, 182, 0, 68, 0, 73, 0, 0,
	119, 0, 115, 117, 57, 58, 190, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 140,
	0, 123, 121, 163, 0, 72, 24, 0, 0, 100,
	112, 113, 86, 0, 0, 61, 67, 87, 88, 89,
	90, 31, 37, 0, 0, 191, 0, 31, 37, 0,
	0, 0, 141, 142, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, -2, 0,
	183, 26, 0, 28, 78, 0, 157, 157, 79, 59,
	0, 0, 0, 0, 0, 0, 69, 70, 71, 162,
	75, 76, 0, 0, 191, 55, 118, 0, 27, 0,
	81, 0, 84, 85, 60, 62, 0, 91, 92, 93,
	66, 94, 95, 96, 97, 120, 122, 29, 0, 158,
	63, 64, 65, 0, 82, 0, 0, 0, 77, 83,
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
		//line ./src/cg_ir/ir_gen.y:280
		{
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:297
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = yyDollar[2].s
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			yyVAL.code = new(node)
			yyVAL.code.value = "jmp, " + yyVAL.mp["after"]
			yyVAL.code.next = new(node)
			yyVAL.code.next.value = "label, " + yyVAL.mp["begin"]
			yyVAL.code.next.next = new(node)
			p := copy_nodes(yyDollar[4].code, yyVAL.code.next.next)
			p.next = new(node)
			p.next.value = "label, " + yyVAL.mp["after"]
			print_ircode(yyVAL.code)
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:331
		{
			yyVAL.code = yyDollar[3].code
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:360
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:361
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n)
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
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:366
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:367
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "str"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:368
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", 1"
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:369
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = "int"
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", 0"
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:370
		{
			if (yyDollar[1].s == "i8") || (yyDollar[1].s == "i16") || (yyDollar[1].s == "i32") || (yyDollar[1].s == "i64") || (yyDollar[1].s == "isize") || (yyDollar[1].s == "u8") || (yyDollar[1].s == "u16") || (yyDollar[1].s == "u32") || (yyDollar[1].s == "u64") || (yyDollar[1].s == "usize") {
				yyVAL.s = "int"
			} else {
				yyVAL.s = "str"
			}
			{
				yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
				temp_num += 1
				yyVAL.mp["type"] = "str"
				yyVAL.mp["value"] = yyDollar[1].s
				yyVAL.code = new(node)
				yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].s
			}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:376
		{
			yyVAL.code = yyDollar[1].code
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:381
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyDollar[1].code)
			p.next = yyDollar[2].code
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:382
		{
			yyVAL.code = yyDollar[1].code
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:386
		{
			if yyDollar[1].code == nil {
				fmt.Println("in stmt let code is nil")
			}
			yyVAL.code = yyDollar[1].code
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:389
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:435
		{
			fmt.Println("iff0")
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			fmt.Println("iff1", yyDollar[2], yyDollar[3])
			if yyDollar[2].mp == nil {
				log.Fatal("Bad If   block;;;")
			}
			fmt.Println("iff2")
			yyVAL.code = new(node)
			fmt.Println("ssdd")
			p := yyVAL.code
			if yyDollar[2].code != nil {
				o := copy_nodes(yyDollar[2].code, yyVAL.code)
				o.next = new(node)
				p = o.next
				fmt.Println("sss00")
			} else {
				p = yyVAL.code
			}
			p.value = "ifgoto, je, " + yyDollar[2].mp["value"] + ", 0, " + yyVAL.mp["after"]
			fmt.Println("iff2-1")
			p.next = new(node)
			if yyDollar[3].code == nil {
				fmt.Println("$3.code nil in expr_if")
			}
			p.next = yyDollar[3].code
			q := list_end(&yyVAL.code)
			q.next = new(node)
			fmt.Println("iff2-3")
			q.next.value = "label, " + yyVAL.mp["after"]
			fmt.Println("iff3")
			print_ircode(yyVAL.code)
			fmt.Println("iff4")
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:448
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			if yyDollar[2].mp == nil {
				log.Fatal("Expression or block  not declared in IF statement")
			}
			yyVAL.code = yyDollar[2].code
			p := list_end(&yyVAL.code)
			p.next = new(node)
			p.next.value = "ifgoto, je, " + yyDollar[2].mp["value"] + ", 1, " + yyVAL.mp["true"]
			p.next.next = new(node)
			print_ircode(yyVAL.code)
			fmt.Println(yyDollar[5].code, "sdfassaaasssssssssssssssssssssssx")
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
			fmt.Println("printing in if else block")
			print_ircode(yyVAL.code)
			fmt.Println("printed in if else block")
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:467
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
			fmt.Println("printing in  block or if")
			print_ircode(yyVAL.code)
			fmt.Println("printed in  block or if")
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:471
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:475
		{
			yyVAL.code = yyDollar[2].code
			fmt.Println("printing in  block")
			print_ircode(yyVAL.code)
			fmt.Println("printed in  block")
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:482
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
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
			fmt.Println("Printing WHILE")
			print_ircode(yyVAL.code)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:487
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
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
			print_ircode(yyVAL.code)
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:492
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
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
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:506
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
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
			fmt.Println("AAAAAAAAAAAA", yyDollar[5], yyDollar[2])
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
			fmt.Println("FFFFFFFFFFF")
			print_ircode(yyVAL.code)

		}
	case 77:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:523
		{
			fmt.Println("AT for inside start")
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["begin"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[3].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[5].code
			q := list_end(&p)
			q.next = new(node)
			q.next.value = "ifgoto, je, " + yyDollar[5].mp["value"] + ", 0, " + yyVAL.mp["after"]
			fmt.Println("AT for inside mid")
			q.next.next = yyDollar[9].code
			s := list_end(&q)
			s.next = yyDollar[7].code
			t := list_end(&s)
			t.next = new(node)
			t.next.value = "jmp, " + yyVAL.mp["begin"]
			u := t.next
			u.next = new(node)
			u.next.value = "label, " + yyVAL.mp["after"]
			fmt.Println("AT for inside end")
			print_ircode(yyVAL.code)
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:539
		{
			if yyDollar[3].mp == nil {
				log.Fatal("Variable name not present in let")
			}
			if yyDollar[4].mp == nil {
				if yyDollar[4].s != "" {
					if yyDollar[5].mp != nil {
						/*let mut y:i32 = expr */
						if yyDollar[5].mp["type"] != yyDollar[4].s {
							log.Fatal("Type mismatch in let ;;")
						}
						yyDollar[3].mp["type"] = yyDollar[5].mp["type"]
						yyVAL.code = new(node)
						if yyDollar[5].code != nil {
							p := copy_nodes(yyDollar[5].code, yyVAL.code)
							p.next = new(node)
							p.next.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
						} else {
							yyVAL.code.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
							fmt.Println("let 2-2")
						}
					} else { /*let  y:i32 */
						yyDollar[3].mp["type"] = yyDollar[4].s
						fmt.Println("let 3")
					}
				} else { /* let y = 5 */
					if yyDollar[5].mp == nil {
						log.Fatal("incomplete let expression  ;")
					}
					yyDollar[3].mp["type"] = yyDollar[5].mp["type"]
					yyVAL.code = new(node)
					yyVAL.code = yyDollar[5].code
					p := list_end(&yyVAL.code)
					p.next = new(node)
					p.next.value = "=, " + yyDollar[3].mp["value"] + ", " + yyDollar[5].mp["value"]
				}
			}
			print_ircode(yyDollar[5].code)
			fmt.Println("let after")
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:567
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
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:568
		{
			yyVAL.s = ""
			yyVAL.mp = nil
			yyVAL.code = nil
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:572
		{
			yyVAL.code = yyDollar[2].code
			yyVAL.mp = yyDollar[2].mp
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:575
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n)
			fmt.Println(yyDollar[1].n)
			yyVAL.mp["type"] = "int"
			if yyDollar[2].code != nil {
				yyVAL.code = yyDollar[2].code
				p := list_end(&yyDollar[2].code)
				p.next = new(node)
				p.next.value = "+, " + yyVAL.mp["value"] + ", " + strconv.Itoa(yyDollar[1].n) + ", " + yyDollar[2].mp["value"]
			} else {
				yyVAL.s = ""
			}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:584
		{
			yyVAL.s = ""
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:602
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["start"] = strconv.Itoa(yyDollar[1].n)
			yyVAL.mp["end"] = strconv.Itoa(yyDollar[3].n)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:603
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["start"] = strconv.Itoa((int)(([]rune(yyDollar[1].s))[0]))
			yyVAL.mp["end"] = strconv.Itoa((int)(([]rune(yyDollar[3].s))[0]))
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:607
		{
			yyDollar[1].mp = symtab.Find_id(yyDollar[1].s)
			if yyDollar[1].mp == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
			}
			yyVAL.mp = yyDollar[1].mp
			fmt.Println(yyDollar[1].mp, "sdaaa", yyVAL.mp)
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:621
		{
			if yyDollar[1].code == nil {
				yyVAL.s = yyDollar[1].s
			} else {
				yyVAL.code = yyDollar[1].code
				yyVAL.mp = yyDollar[1].mp
			}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:630
		{
			yyVAL.s = yyDollar[1].s
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:631
		{
			yyVAL.s = ""
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:636
		{
			if (yyDollar[1].s == "i8") || (yyDollar[1].s == "i16") || (yyDollar[1].s == "i32") || (yyDollar[1].s == "i64") || (yyDollar[1].s == "isize") || (yyDollar[1].s == "u8") || (yyDollar[1].s == "u16") || (yyDollar[1].s == "u32") || (yyDollar[1].s == "u64") || (yyDollar[1].s == "usize") {
				yyVAL.s = "int"
			} else {
				yyVAL.s = "str"
			}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:637
		{
			yyVAL.s = yyDollar[1].s
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:640
		{
			yyVAL.s = yyDollar[1].s
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:641
		{
			yyVAL.s = "Array_" + yyDollar[2].s + "_" + yyDollar[3].s
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:645
		{
			yyVAL.s = strconv.Itoa(yyDollar[2].n)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:650
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:655
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp["args"] = yyDollar[1].mp["value"] + ", "
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:656
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			yyVAL.mp["args"] = yyDollar[1].mp["args"] + yyDollar[3].mp["value"] + ", "
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:664
		{
			fmt.Println("In maybe assignment")
			yyVAL.code = yyDollar[1].code
			print_ircode(yyDollar[1].code)
			fmt.Println("In maybe assignment")
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:669
		{
			p := symtab.Find_id(yyDollar[1].s)
			fmt.Println("(in exp ) Identifier =", p["value"])
			if p == nil {
				fmt.Println("(in exp )new Identifier ", yyDollar[1].s)
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:680
		{
			fmt.Println("After func call")
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			fmt.Println("hole")
			print_ircode(yyVAL.code)
			fmt.Println("hole2")
			q := list_end(&p)
			q.next = new(node)
			q.next.value = "=, " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			fmt.Println("hole3")
			print_ircode(yyVAL.code)
			fmt.Println("hole4")
			fmt.Println("DONE func call")
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:682
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "+, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:683
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "-, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:686
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "*, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:687
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "/, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:688
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "%, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:689
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "&, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:692
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "|, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:693
		{
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "^, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			print_ircode(yyVAL.code)
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:696
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyDollar[1].mp["value"] + ", " + strconv.Itoa(yyDollar[2].n)
			fmt.Println(yyDollar[1].n)
			if yyDollar[3].code != nil {
				yyVAL.code.next = new(node)
				yyVAL.code.next = yyDollar[2].code
				p := list_end(&yyVAL.code.next)
				p.next = new(node)
				p.next.value = "+, " + yyDollar[1].mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
			} else {
				yyVAL.s = ""
			}
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:728
		{
			yyVAL.s = ""
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:732
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:733
		{
			yyVAL.code = yyDollar[1].code
			yyVAL.mp = yyDollar[1].mp
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:739
		{
			yyVAL.mp = yyDollar[1].mp
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:741
		{
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}

		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:749
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = new(node)
			yyVAL.code.value = "=, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp[yyDollar[3].mp["value"]]
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:756
		{
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:759
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
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:764
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
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:769
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
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:773
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
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:778
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
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:782
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
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:787
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
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:792
		{

			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = ">, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:798
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "<, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:804
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
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:810
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "., " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:815
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
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:819
		{
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
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:823
		{

			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["false"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyDollar[1].code)
			p.next = new(node)
			q := copy_nodes(p.next, yyDollar[3].code)
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
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:830
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.mp["true"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.mp["after"] = "label" + strconv.Itoa(label_num)
			label_num += 1
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyDollar[1].code)
			p.next = new(node)
			q := copy_nodes(p.next, yyDollar[3].code)
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
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:835
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "**, " + yyVAL.mp["value"] + ", " + yyDollar[1].mp["value"] + ", " + yyDollar[3].mp["value"]
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:839
		{
			yyVAL.mp = symtab.Make_entry("temp" + strconv.Itoa(temp_num))
			temp_num += 1
			yyVAL.mp["type"] = yyDollar[1].mp["type"]
			yyVAL.code = yyDollar[1].code
			p := list_end(&yyVAL.code)
			p.next = yyDollar[3].code
			q := list_end(&p.next)
			q.next = new(node)
			q.next.value = "push, " + yyDollar[3].mp["args"]
			q.next.next = new(node)
			q.next.next.value = "call, " + yyDollar[1].mp["value"] + ", "
			print_ircode(yyVAL.code)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:847
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:848
		{
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:849
		{
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:853
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:854
		{
			yyVAL.mp = yyDollar[1].mp
			yyVAL.code = yyDollar[1].code
			fmt.Println("DODODIDODODO ")
			print_ircode(yyVAL.code)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:858
		{
			p := symtab.Find_id(yyDollar[1].s)
			if p == nil {
				yyDollar[1].mp = symtab.Make_entry(yyDollar[1].s)
				yyVAL.mp = yyDollar[1].mp
			} else {
				yyVAL.mp = p
			}

		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:866
		{
			p := symtab.Find_id(yyDollar[1].s + "_" + yyDollar[3].s)
			if p == nil {
				yyVAL.mp = symtab.Make_entry(yyDollar[1].s + "_" + yyDollar[3].s)
			} else {
				yyVAL.mp = p
			}

		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./src/cg_ir/ir_gen.y:873
		{
			p := symtab.Find_id(yyDollar[1].s + "_")
			if p == nil {
				yyVAL.mp = symtab.Make_entry(yyDollar[1].s + "_")
			} else {
				yyVAL.mp = p
			}

		}
	}
	goto yystack /* stack new state and value */
}
