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
	movl	$10, %eax
	movl	%eax, %edi
	movq	___stack_chk_guard@GOTPCREL(%rip), %rcx
	movq	(%rcx), %rcx
	movq	%rcx, -8(%rbp)
	movl	$0, -36(%rbp)
	callq	_malloc
	movq	$-1, %rdx
	movq	%rax, -48(%rbp)
	movq	-48(%rbp), %rdi
	leaq	l_.str(%rip), %rsi
	callq	___strcpy_chk
	leaq	-32(%rbp), %rsi
	movq	-48(%rbp), %rcx
	movq	(%rcx), %rdx
	movq	%rdx, -32(%rbp)
	movw	8(%rcx), %r8w
	movw	%r8w, -24(%rbp)
	movq	-48(%rbp), %rcx
	movq	(%rcx), %rdx
	movq	%rdx, -23(%rbp)
	movw	8(%rcx), %r8w
	movw	%r8w, -15(%rbp)
	leaq	L_.str.1(%rip), %rdi
	movq	%rax, -56(%rbp)         ## 8-byte Spill
	movb	$0, %al
	callq	_printf
	movq	___stack_chk_guard@GOTPCREL(%rip), %rcx
	movq	(%rcx), %rcx
	movq	-8(%rbp), %rdx
	cmpq	%rdx, %rcx
	movl	%eax, -60(%rbp)         ## 4-byte Spill
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
	.section	__TEXT,__const
l_.str:                                 ## @.str
	.asciz	"abcdefghi\000"

	.section	__TEXT,__cstring,cstring_literals
L_.str.1:                               ## @.str.1
	.asciz	"b2: %s\n"


.subsections_via_symbols
