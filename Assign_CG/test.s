

.globl _start


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
asd:
.ascii "asdasdsad%d"
asdend:
a:
.rept 100
.long 69
.endr
aend:


.section .text


_start:
label1:
movl $b, %eax
movl $10,%eax
movl $n, %ebx
movl $10,%ebx
movl $mod, %ecx
movl $2713,%ecx
movl $rand1, %edx
movl $31,%edx
movl $rand2, %edi
movl $101,%edi
movl $cur, %esp
movl $23,%esp
movl %eax,b
movl $i, %eax
movl $0,%eax
movl %eax, i
movl %ebx, n
movl %ecx, mod
movl %edx, rand1
movl %edi, rand2
movl %esp, cur
label2:
movl $i, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label4
movl %eax, i
movl %ebx, n
label3:
movl $0,%edx
movl $temp, %eax
movl $rand1, %ebx
movl $cur, %ecx
movl %ebx,%eax
mull %ecx
movl $rand2, %edx
movl %eax,%eax
addl %edx,%eax
movl $cur ,%ecx
movl %edx,rand2
movl $0,%edx
movl $cur ,%ecx
movl %eax,temp
movl $cur, %eax
movl $temp, %ecx
movl $mod, %edi
movl %ecx,%eax
divl %edi
movl %edx,%eax
movl $i, %edx
movl %eax,a(,%edx,4)
movl %edx,%edx
addl $1,%edx
movl $aa, %esp
cmpl %esp,$1
je label2
movl %eax, cur
movl %ebx, rand1
movl %ecx, temp
movl %edx, i
movl %edi, mod
movl %esp, aa
label4:
movl $i, %eax
movl $0,%eax
movl %eax, i
label5:
movl $i, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label12
movl %eax, i
movl %ebx, n
label6:
movl $i, %eax
movl $j, %ebx
movl %eax,%ebx
addl $1,%ebx
movl %eax, i
movl %ebx, j
label7:
movl $j, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label11
movl %eax, j
movl %ebx, n
label8:
movl $i, %eax
movl $t1, %ebx
movl a(,%eax,4),%ebx
movl $j, %ecx
movl $t2, %edx
movl a(,%ecx,4),%edx
cmpl %ebx,%edx
jle label10
movl %eax, i
movl %ebx, t1
movl %ecx, j
movl %edx, t2
label9:
movl $t1, %eax
movl $t2, %ebx
movl %eax,%eax
addl %ebx,%eax
movl $temporary_compiler_variable, %ecx
movl %ebx,%ecx
movl %eax,%ebx
subl %ecx,%ebx
movl %eax,%eax
subl %ebx,%eax
movl %eax, t1
movl %ebx, t2
movl %ecx, temporary_compiler_variable
label10:
movl $t1, %eax
movl $i, %ebx
movl %eax,a(,%ebx,4)
movl $t2, %ecx
movl $j, %edx
movl %ecx,a(,%edx,4)
movl %edx,%edx
addl $1,%edx
movl $aa, %edi
cmpl %edi,$1
je label7
movl %eax, t1
movl %ebx, i
movl %ecx, t2
movl %edx, j
movl %edi, aa
label11:
movl $i, %eax
movl %eax,%eax
addl $1,%eax
movl $aa, %ebx
cmpl %ebx,$1
je label5
movl %eax, i
movl %ebx, aa
label12:
movl $i, %eax
movl $0,%eax
movl %eax, i
label13:
movl $i, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label6
movl %eax, i
movl %ebx, n
label14:
movl $i, %eax
movl $t1, %ebx
movl a(,%eax,4),%ebx
movl $aa, %ecx
cmpl %ecx,$1
je label13
movl %eax, i
movl %ebx, t1
movl %ecx, aa
label15:


