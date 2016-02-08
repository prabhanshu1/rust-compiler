

.globl main


.section .data
a:
.long 69
aend:
c:
.long 69
cend:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
str:
.ascii "%d %d \n\0"
strend:


.section .text


main:
label1:
movl a, %eax
movl $32,%eax
movl c, %ebx
movl $0,%ebx
subl %eax,%ebx
movl %eax,%eax
xorl %ebx,%eax
pushl %ebx
pushl %eax
pushl $str
movl %eax, a
movl %ebx, c
label2:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80


