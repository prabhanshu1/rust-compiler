

.globl main


.section .data
i:
.long 69
iend:
sum:
.long 69
sumend:
chap:
.long 69
chapend:
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
scan:
.ascii "%d\0"
scanend:
a:
.rept 100
.long 69
.endr
aend:


.section .text


main:
label_CG1:
movl i, %ecx
movl $0,%ecx
movl sum, %edx
movl $0,%edx
movl %ecx, i
movl %edx, sum
sham:
movl i, %eax
movl %eax, i
cmpl $10,%eax
jg ram
label_CG2:
movl i, %ecx
movl %ecx,a(,%ecx,4)
movl chap, %edi
movl $mad,%edi
pushl %edi
pushl $scan
movl %ecx, i
movl %edi, chap
label_CG3:
call  scanf
movl sum, %eax
pushl %eax
pushl $val
movl %eax, sum
label_CG4:
call  printf
movl sum, %eax
movl mad, %ebx
movl %eax,%eax
addl %ebx,%eax
movl i, %ecx
movl %ecx,%ecx
addl $1,%ecx
movl %eax, sum
movl %ebx, mad
movl %ecx, i
jmp sham
ram:
movl sum, %edx
pushl %edx
pushl $SUM
movl %edx, sum
label_CG5:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG6:


