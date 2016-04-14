

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
label_CG1:
movl a, %edx
movl $32,%edx
movl c, %ebx
movl $0,%ebx
subl %edx,%ebx
movl %edx,%edx
xorl %ebx,%edx
pushl %ebx
pushl %edx
pushl $str
movl %ebx, c
movl %edx, a
label_CG2:
call  printf
movl $1,%eax
movl $0,%ebx
int $0x80


