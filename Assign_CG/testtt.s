

.globl main


.section .data
temp0:
.long 69
temp0end:
y:
.long 69
yend:
temp1:
.long 69
temp1end:
temp2:
.long 69
temp2end:
temp10:
.long 69
temp10end:
temp11:
.long 69
temp11end:
temp20:
.long 69
temp20end:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
temp3:
.ascii "mtoa  %d\n\n\0"
temp3end:
temp7:
.ascii "mtoa %d\n\0"
temp7end:
temp14:
.ascii "FAIL\n\0"
temp14end:
temp12:
.ascii "SUCEESS\n\0"
temp12end:
temp18:
.ascii "Abhimanyu\n\0"
temp18end:


.section .text


label_CG1:
jmp label6
main:
pushl %ebp
movl %esp,%ebp
movl temp0, %eax
movl $5,%eax
movl y, %ebx
movl %eax,%ebx
movl temp1, %ecx
movl $1,%ecx
movl temp2, %edi
movl %ebx,%edi
addl %ecx,%edi
movl %edi,%ebx
pushl %ebx
pushl $temp3
movl %eax, temp0
movl %ebx, y
movl %ecx, temp1
movl %edi, temp2
label_CG2:
call  printf
label_CG3:
call  shit
movl %eax,y
movl y, %eax
pushl %eax
pushl $temp7
movl %eax, y
label_CG4:
call  printf
movl temp10, %edi
movl $2,%edi
movl y, %eax
movl %eax, y
movl %edi, temp10
cmpl %eax,%edi
jg label0
label_CG5:
movl temp11, %eax
movl $0,%eax
movl %eax, temp11
jmp label1
label0:
movl temp11, %esi
movl $1,%esi
movl %esi, temp11
label1:
movl temp11, %ecx
movl %ecx, temp11
cmpl $1,%ecx
je label4
label_CG6:
pushl $temp14
label_CG7:
call  printf
jmp label2
label4:
pushl $temp12
label_CG8:
call  printf
label2:
label6:
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG9:
jmp label7
shit:
pushl %ebp
movl %esp,%ebp
pushl $temp18
label_CG10:
call  printf
movl temp20, %esi
movl $1,%esi
movl %esi,temp20
movl temp20, %eax
movl %ebp,%esp
popl %ebp
ret
movl %eax, temp20
label7:


