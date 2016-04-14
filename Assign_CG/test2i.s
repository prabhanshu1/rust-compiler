

.globl main


.section .data
i:
.long 69
iend:
sum:
.long 69
sumend:
mad:
.long 69
madend:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
val:
.ascii "Value=%d \n\0"
valend:
SUM:
.ascii "SUM=%d \n\0"
SUMend:
a:
.rept 100
.long 69
.endr
aend:


.section .text


main:
label_CG1:
movl i, %edx
movl $0,%edx
movl sum, %eax
movl $0,%eax
movl %edx, i
movl %eax, sum
sham:
movl i, %eax
movl %eax, i
cmpl $10,%eax
jg ram
label_CG2:
movl i, %edx
movl %edx,a(,%edx,4)
movl mad, %eax
movl a(,%edx,4),%eax
pushl %eax
pushl $val
movl %eax, mad
movl %edx, i
label_CG3:
call  printf
movl sum, %eax
pushl %eax
pushl $SUM
movl %eax, sum
label_CG4:
call  printf
movl sum, %eax
movl i, %edi
movl %eax,%eax
addl %edi,%eax
movl %edi,%edi
addl $1,%edi
movl %eax, sum
movl %edi, i
jmp sham
ram:
movl sum, %eax
pushl %eax
pushl $SUM
movl %eax, sum
label_CG5:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG6:


