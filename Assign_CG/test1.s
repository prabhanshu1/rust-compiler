

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
movl n, %esi
movl $32,%esi
movl i, %ecx
movl $1,%ecx
movl count, %edi
movl $0,%edi
movl %ecx, i
movl %edi, count
movl %esi, n
ram:
movl i, %eax
movl n, %ebx
movl %eax, i
movl %ebx, n
cmpl %eax,%ebx
jl label1
label_CG2:
movl $0,%edx
movl d, %eax
movl n, %ecx
movl i, %edi
movl %ecx,%eax
divl %edi
movl %edx,%eax
movl %eax, d
movl %ecx, n
movl %edi, i
cmpl $0,%eax
jne label0
label_CG3:
movl count, %edi
movl %edi,%edi
addl $1,%edi
movl %edi, count
label0:
movl i, %esi
movl %esi,%esi
addl $1,%esi
movl %esi, i
jmp ram
label1:
movl count, %eax
movl %eax, count
cmpl $2,%eax
je label2
label2:
movl count, %eax
pushl %eax
pushl $asd
movl %eax, count
label_CG4:
call  printf
pushl count
movl n, %edx
pushl %edx
movl d, %ebx
pushl %ebx
pushl $asd
movl %edx, n
movl %ebx, d
label_CG5:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG6:


