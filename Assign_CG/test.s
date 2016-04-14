

.globl main


.section .data
n:
.long 69
nend:
i:
.long 69
iend:
count:
.long 69
countend:
d:
.long 69
dend:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
asd:
.ascii "dekh prabhanshu print hota hai %d %d %d\n\0"
asdend:


.section .text


main:
label_CG1:
movl n, %ecx
movl $32,%ecx
movl i, %edi
movl $1,%edi
movl count, %eax
movl $0,%eax
movl %ecx, n
movl %edi, i
movl %eax, count
cmpl %edi,%ecx
jl label1
label_CG2:
movl $0,%edx
movl d, %eax
movl n, %ebx
movl i, %ecx
movl %ebx,%eax
divl %ecx
movl %edx,%eax
movl %eax, d
movl %ebx, n
movl %ecx, i
cmpl $0,%eax
jne label0
label_CG3:
movl count, %ecx
movl %ecx,%ecx
addl $1,%ecx
movl %ecx, count
label0:
movl i, %eax
movl %eax,%eax
addl $1,%eax
movl %eax, i
jmp 4
label1:
movl count, %ebx
movl %ebx, count
cmpl $2,%ebx
je label2
label2:
movl count, %eax
pushl %eax
pushl $asd
movl %eax, count
label_CG4:
call  printf
pushl count
movl n, %eax
pushl %eax
movl d, %edi
pushl %edi
pushl $asd
movl %eax, n
movl %edi, d
label_CG5:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG6:


