	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14	sdk_version 10, 14
	.globl	_main                   ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$64, %rsp
	leaq	-48(%rbp), %rsi
	movq	___stack_chk_guard@GOTPCREL(%rip), %rax
	movq	(%rax), %rax
	movq	%rax, -8(%rbp)
	movl	$0, -52(%rbp)
	movq	L___const.main.b1(%rip), %rax
	movq	%rax, -18(%rbp)
	movw	L___const.main.b1+8(%rip), %cx
	movw	%cx, -10(%rbp)
	movq	-18(%rbp), %rax
	movq	%rax, -48(%rbp)
	movw	-10(%rbp), %cx
	movw	%cx, -40(%rbp)
	movq	-18(%rbp), %rax
	movq	%rax, -39(%rbp)
	movw	-10(%rbp), %cx
	movw	%cx, -31(%rbp)
	leaq	L_.str(%rip), %rdi
	movb	$0, %al
	callq	_printf
	movq	___stack_chk_guard@GOTPCREL(%rip), %rsi
	movq	(%rsi), %rsi
	movq	-8(%rbp), %rdi
	cmpq	%rdi, %rsi
	movl	%eax, -56(%rbp)         ## 4-byte Spill
	jne	LBB0_2
## %bb.1:
	xorl	%eax, %eax
	addq	$64, %rsp
	popq	%rbp
	retq
LBB0_2:
	callq	___stack_chk_fail
	ud2
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__cstring,cstring_literals
L___const.main.b1:                      ## @__const.main.b1
	.asciz	"abcdefghi"

L_.str:                                 ## @.str
	.asciz	"b2: %s\n"


.subsections_via_symbols
