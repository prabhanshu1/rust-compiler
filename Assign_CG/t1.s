

.globl main


.section .data
b:
.long 69
bend:
n:
.long 69
nend:
mod:
.long 69
modend:
rand1:
.long 69
rand1end:
rand2:
.long 69
rand2end:
cur:
.long 69
curend:
i:
.long 69
iend:
temp:
.long 69
tempend:
aa:
.long 69
aaend:
j:
.long 69
jend:
t1:
.long 69
t1end:
t2:
.long 69
t2end:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
a:
.rept 100
.long 69
.endr
aend:


.section .text


main:
label1:
movl b, %edi
movl $10,%edi
movl n, %eax
movl $10,%eax
movl mod, %ebx
movl $2713,%ebx
movl rand1, %ecx
movl $31,%ecx
movl rand2, %esp
movl $101,%esp
movl cur, %edx
movl $23,%edx
movl %eax,n
movl i, %eax
movl $0,%eax
movl %eax, i
movl %ebx, mod
movl %ecx, rand1
movl %edx, cur
movl %edi, b
movl %esp, rand2
label2:
movl i, %eax
movl n, %ebx
movl %eax, i
movl %ebx, n
cmpl %eax,%ebx
je label4
label3:
movl $0,%edx
movl temp, %eax
movl rand1, %edi
movl cur, %ebx
movl %edi,%eax
mull %ebx
movl rand2, %ecx
movl %eax,%eax
addl %ecx,%eax
movl cur ,%ebx
movl $0,%edx
movl cur ,%ebx
movl %eax,temp
movl cur, %eax
movl temp, %ebx
movl mod, %esp
movl %ebx,%eax
divl %esp
movl %edx,%eax
movl i, %edx
movl %eax,a(,%edx,4)
movl %edx,%edx
addl $1,%edx
movl %eax,cur
movl aa, %eax
movl %edi, rand1
movl %esp, mod
movl %eax, aa
movl %ebx, temp
movl %ecx, rand2
movl %edx, i
cmpl $1,%eax
je label2
label4:
movl i, %eax
movl $0,%eax
movl %eax, i
label5:
movl i, %ebx
movl n, %ecx
movl %ebx, i
movl %ecx, n
cmpl %ebx,%ecx
je label12
label6:
movl i, %edx
movl j, %eax
movl %edx,%eax
addl $1,%eax
movl %edx, i
movl %eax, j
label7:
movl j, %ecx
movl n, %edx
movl %ecx, j
movl %edx, n
cmpl %ecx,%edx
je label11
label8:
movl i, %eax
movl t1, %ebx
movl a(,%eax,4),%ebx
movl j, %ecx
movl t2, %edx
movl a(,%ecx,4),%edx
movl %eax, i
movl %ebx, t1
movl %ecx, j
movl %edx, t2
cmpl %ebx,%edx
jle label10
label9:
movl t1, %edx
movl t2, %eax
movl %edx,%edx
addl %eax,%edx
movl temporary_compiler_variable, %edi
movl %eax,%edi
movl %edx,%eax
subl %edi,%eax
movl %edx,%edx
subl %eax,%edx
movl %eax, t2
movl %edx, t1
movl %edi, temporary_compiler_variable
label10:
movl t1, %edx
movl i, %eax
movl %edx,a(,%eax,4)
movl t2, %edi
movl j, %ebx
movl %edi,a(,%ebx,4)
movl %ebx,%ebx
addl $1,%ebx
movl aa, %esp
movl %eax, i
movl %ebx, j
movl %edx, t1
movl %edi, t2
movl %esp, aa
cmpl $1,%esp
je label7
label11:
movl i, %edi
movl %edi,%edi
addl $1,%edi
movl aa, %eax
movl %eax, aa
movl %edi, i
cmpl $1,%eax
je label5
label12:
movl i, %esp
movl $0,%esp
movl %esp, i
label13:
movl i, %eax
movl n, %ebx
movl %eax, i
movl %ebx, n
cmpl %eax,%ebx
je label6
label14:
movl i, %eax
movl t1, %edx
movl a(,%eax,4),%edx
movl aa, %ebx
movl %eax, i
movl %ebx, aa
movl %edx, t1
cmpl $1,%ebx
je label13
label15:
movl $1,%eax
movl $0,%ebx
int $0x80


