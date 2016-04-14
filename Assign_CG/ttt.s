

.globl main


.section .data
temp1:
.long 69
temp1end:
temp2:
.long 69
temp2end:
y:
.long 69
yend:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:


.section .text


label_CG1:
jmp label4
main:
pushl %ebp
movl %esp,%ebp
label0:
cmpl $0,$
je label2
label_CG2:
movl temp1, %edx
movl $1,%edx
movl y, %ebx
movl temp2, %esi
movl %ebx,%esi
addl %edx,%esi
movl %esi,%ebx
movl %edx, temp1
movl %esi, temp2
movl %ebx, y
jmp label0
label2:
label4:
movl $1,%eax
movl $0,%ebx
int $0x80


