

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
label1:
movl i, %esi
movl $0,%esi
movl sum, %eax
movl $0,%eax
movl %esi, i
movl %eax, sum
label2:
movl i, %eax
movl %eax, i
cmpl $10,%eax
jg label6
label3:
movl i, %ebx
movl %ebx,a(,%ebx,4)
movl mad, %esi
movl a(,%ebx,4),%esi
pushl %esi
pushl $val
movl %esi, mad
movl %ebx, i
label4:
call  printf
movl sum, %ecx
pushl %ecx
pushl $SUM
movl %ecx, sum
label5:
call  printf
movl sum, %eax
movl i, %ebx
movl %eax,%eax
addl %ebx,%eax
movl %ebx,%ebx
addl $1,%ebx
movl %eax, sum
movl %ebx, i
jmp label2
label6:
movl sum, %ebx
pushl %ebx
pushl $SUM
movl %ebx, sum
label7:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80
label8:


