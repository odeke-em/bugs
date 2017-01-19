"".f1 t=1 size=176 args=0x80 locals=0x108
	0x0000 00000 (main.go:5)	TEXT	"".f1(SB), $264-128
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	LEAQ	-136(SP), AX
	0x0011 00017 (main.go:5)	CMPQ	AX, 16(CX)
	0x0015 00021 (main.go:5)	JLS	166
	0x001b 00027 (main.go:5)	SUBQ	$264, SP
	0x0022 00034 (main.go:5)	MOVQ	BP, 256(SP)
	0x002a 00042 (main.go:5)	LEAQ	256(SP), BP
	0x0032 00050 (main.go:5)	FUNCDATA	$0, gclocals·8658ec02c587fb17d31955e2d572c2ff(SB)
	0x0032 00050 (main.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0032 00050 (main.go:5)	LEAQ	"".t+272(FP), DI
	0x003a 00058 (main.go:5)	XORPS	X0, X0
	0x003d 00061 (main.go:5)	DUFFZERO	$266
	0x0050 00080 (main.go:6)	PCDATA	$0, $0
	0x0050 00080 (main.go:6)	CALL	"".g(SB)
	0x0055 00085 (main.go:6)	LEAQ	""..autotmp_0+128(SP), DI
	0x005d 00093 (main.go:6)	MOVQ	SP, SI
	0x0060 00096 (main.go:6)	DUFFCOPY	$784
	0x0073 00115 (main.go:5)	LEAQ	"".t+272(FP), DI
	0x007b 00123 (main.go:6)	LEAQ	""..autotmp_0+128(SP), SI
	0x0083 00131 (main.go:6)	DUFFCOPY	$784
	0x0096 00150 (main.go:6)	MOVQ	256(SP), BP
	0x009e 00158 (main.go:6)	ADDQ	$264, SP
	0x00a5 00165 (main.go:6)	RET
	0x00a6 00166 (main.go:6)	NOP
	0x00a6 00166 (main.go:5)	PCDATA	$0, $-1
	0x00a6 00166 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00ab 00171 (main.go:5)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 84 24 78 ff ff  eH..%....H..$x..
	0x0010 ff 48 3b 41 10 0f 86 8b 00 00 00 48 81 ec 08 01  .H;A.......H....
	0x0020 00 00 48 89 ac 24 00 01 00 00 48 8d ac 24 00 01  ..H..$....H..$..
	0x0030 00 00 48 8d bc 24 10 01 00 00 0f 57 c0 48 89 6c  ..H..$.....W.H.l
	0x0040 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48 8b 6d 00  $.H.l$......H.m.
	0x0050 e8 00 00 00 00 48 8d bc 24 80 00 00 00 48 89 e6  .....H..$....H..
	0x0060 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48  H.l$.H.l$......H
	0x0070 8b 6d 00 48 8d bc 24 10 01 00 00 48 8d b4 24 80  .m.H..$....H..$.
	0x0080 00 00 00 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00  ...H.l$.H.l$....
	0x0090 00 00 48 8b 6d 00 48 8b ac 24 00 01 00 00 48 81  ..H.m.H..$....H.
	0x00a0 c4 08 01 00 00 c3 e8 00 00 00 00 e9 50 ff ff ff  ............P...
	rel 5+4 t=16 TLS+0
	rel 72+4 t=8 runtime.duffzero+266
	rel 81+4 t=8 "".g+0
	rel 107+4 t=8 runtime.duffcopy+784
	rel 142+4 t=8 runtime.duffcopy+784
	rel 167+4 t=8 runtime.morestack_noctxt+0
"".f2 t=1 size=176 args=0x80 locals=0x108
	0x0000 00000 (main.go:9)	TEXT	"".f2(SB), $264-128
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	LEAQ	-136(SP), AX
	0x0011 00017 (main.go:9)	CMPQ	AX, 16(CX)
	0x0015 00021 (main.go:9)	JLS	166
	0x001b 00027 (main.go:9)	SUBQ	$264, SP
	0x0022 00034 (main.go:9)	MOVQ	BP, 256(SP)
	0x002a 00042 (main.go:9)	LEAQ	256(SP), BP
	0x0032 00050 (main.go:9)	FUNCDATA	$0, gclocals·8658ec02c587fb17d31955e2d572c2ff(SB)
	0x0032 00050 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0032 00050 (main.go:9)	LEAQ	"".~r0+272(FP), DI
	0x003a 00058 (main.go:9)	XORPS	X0, X0
	0x003d 00061 (main.go:9)	DUFFZERO	$266
	0x0050 00080 (main.go:10)	PCDATA	$0, $0
	0x0050 00080 (main.go:10)	CALL	"".g(SB)
	0x0055 00085 (main.go:10)	LEAQ	""..autotmp_1+128(SP), DI
	0x005d 00093 (main.go:10)	MOVQ	SP, SI
	0x0060 00096 (main.go:10)	DUFFCOPY	$784
	0x0073 00115 (main.go:9)	LEAQ	"".~r0+272(FP), DI
	0x007b 00123 (main.go:10)	LEAQ	""..autotmp_1+128(SP), SI
	0x0083 00131 (main.go:10)	DUFFCOPY	$784
	0x0096 00150 (main.go:10)	MOVQ	256(SP), BP
	0x009e 00158 (main.go:10)	ADDQ	$264, SP
	0x00a5 00165 (main.go:10)	RET
	0x00a6 00166 (main.go:10)	NOP
	0x00a6 00166 (main.go:9)	PCDATA	$0, $-1
	0x00a6 00166 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x00ab 00171 (main.go:9)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 84 24 78 ff ff  eH..%....H..$x..
	0x0010 ff 48 3b 41 10 0f 86 8b 00 00 00 48 81 ec 08 01  .H;A.......H....
	0x0020 00 00 48 89 ac 24 00 01 00 00 48 8d ac 24 00 01  ..H..$....H..$..
	0x0030 00 00 48 8d bc 24 10 01 00 00 0f 57 c0 48 89 6c  ..H..$.....W.H.l
	0x0040 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48 8b 6d 00  $.H.l$......H.m.
	0x0050 e8 00 00 00 00 48 8d bc 24 80 00 00 00 48 89 e6  .....H..$....H..
	0x0060 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48  H.l$.H.l$......H
	0x0070 8b 6d 00 48 8d bc 24 10 01 00 00 48 8d b4 24 80  .m.H..$....H..$.
	0x0080 00 00 00 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00  ...H.l$.H.l$....
	0x0090 00 00 48 8b 6d 00 48 8b ac 24 00 01 00 00 48 81  ..H.m.H..$....H.
	0x00a0 c4 08 01 00 00 c3 e8 00 00 00 00 e9 50 ff ff ff  ............P...
	rel 5+4 t=16 TLS+0
	rel 72+4 t=8 runtime.duffzero+266
	rel 81+4 t=8 "".g+0
	rel 107+4 t=8 runtime.duffcopy+784
	rel 142+4 t=8 runtime.duffcopy+784
	rel 167+4 t=8 runtime.morestack_noctxt+0
"".init t=1 size=79 args=0x0 locals=0x8
	0x0000 00000 (main.go:13)	TEXT	"".init(SB), $8-0
	0x0000 00000 (main.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:13)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:13)	JLS	72
	0x000f 00015 (main.go:13)	SUBQ	$8, SP
	0x0013 00019 (main.go:13)	MOVQ	BP, (SP)
	0x0017 00023 (main.go:13)	LEAQ	(SP), BP
	0x001b 00027 (main.go:13)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (main.go:13)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (main.go:13)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (main.go:13)	CMPB	AL, $1
	0x0024 00036 (main.go:13)	JLS	$0, 47
	0x0026 00038 (main.go:13)	MOVQ	(SP), BP
	0x002a 00042 (main.go:13)	ADDQ	$8, SP
	0x002e 00046 (main.go:13)	RET
	0x002f 00047 (main.go:13)	JNE	$0, 56
	0x0031 00049 (main.go:13)	PCDATA	$0, $0
	0x0031 00049 (main.go:13)	CALL	runtime.throwinit(SB)
	0x0036 00054 (main.go:13)	UNDEF
	0x0038 00056 (main.go:13)	MOVB	$2, "".initdone·(SB)
	0x003f 00063 (main.go:13)	MOVQ	(SP), BP
	0x0043 00067 (main.go:13)	ADDQ	$8, SP
	0x0047 00071 (main.go:13)	RET
	0x0048 00072 (main.go:13)	NOP
	0x0048 00072 (main.go:13)	PCDATA	$0, $-1
	0x0048 00072 (main.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x004d 00077 (main.go:13)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 39 48  eH..%....H;a.v9H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 02 48  ...............H
	0x0040 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb b1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 73+4 t=8 runtime.morestack_noctxt+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb t=8 dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·8658ec02c587fb17d31955e2d572c2ff t=8 dupok size=10
	0x0000 01 00 00 00 10 00 00 00 00 00                    ..........
go.info."".f1 t=45 size=38
	0x0000 02 22 22 2e 66 31 00 00 00 00 00 00 00 00 00 00  ."".f1..........
	0x0010 00 00 00 00 00 00 00 01 05 74 00 01 9c 00 00 00  .........t......
	0x0020 00 00 00 00 00 00                                ......
	rel 7+8 t=1 "".f1+0
	rel 15+8 t=1 "".f1+176
	rel 29+8 t=28 go.info."".T+0
go.info."".f2 t=45 size=40
	0x0000 02 22 22 2e 66 32 00 00 00 00 00 00 00 00 00 00  ."".f2..........
	0x0010 00 00 00 00 00 00 00 01 05 7e 72 30 00 01 9c 00  .........~r0....
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 7+8 t=1 "".f2+0
	rel 15+8 t=1 "".f2+176
	rel 31+8 t=28 go.info."".T+0
"".g.args_stackmap t=8 size=16
	0x0000 02 00 00 00 20 00 00 00 00 00 00 00 00 00 00 00  .... ...........
go.info."".init t=45 size=27
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 00                 ...........
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+79
"".initdone· t=32 size=1
runtime.gcbits.01 t=8 dupok size=1
	0x0000 01                                               .
type..namedata.*[]int. t=8 dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 00 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int.+0
	rel 48+8 t=1 type.[]int+0
type.[]int t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int.+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..hashfunc128 t=8 dupok size=16
	0x0000 00 00 00 00 00 00 00 00 80 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc128 t=8 dupok size=16
	0x0000 00 00 00 00 00 00 00 00 80 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
type..alg128 t=8 dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc128+0
	rel 8+8 t=1 type..eqfunc128+0
type..namedata.*main.T. t=8 dupok size=10
	0x0000 01 00 07 2a 6d 61 69 6e 2e 54                    ...*main.T
type.*"".T t=8 size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8b 54 ca 19 00 08 08 36 00 00 00 00 00 00 00 00  .T.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.T.+0
	rel 48+8 t=1 type."".T+0
runtime.gcbits. t=8 dupok size=0
type."".T t=8 size=88
	0x0000 80 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 b5 a4 ef 44 07 08 08 91 00 00 00 00 00 00 00 00  ...D............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 10 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg128+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.T.+0
	rel 44+4 t=5 type.*"".T+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
	rel 72+4 t=5 type..importpath."".+0
