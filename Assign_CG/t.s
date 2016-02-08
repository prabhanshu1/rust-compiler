

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
asd:
.ascii "asdasdsad%d"
asdend:
a:
.rept 100
.long 69
.endr
aend:


.section .text


main:
label1:
movl $b, %eax
movl $10,%eax
movl $n, %edi
movl $10,%edi
movl $mod, %ebx
movl $2713,%ebx
movl $rand1, %ecx
movl $31,%ecx
movl $rand2, %edx
movl $101,%edx
movl $cur, %esp
movl $23,%esp
movl %ecx,rand1
movl $i, %ecx
movl $0,%ecx
movl %edx, rand2
movl %edi, n
movl %esp, cur
movl %eax, b
movl %ebx, mod
movl %ecx, i
label2:
movl $i, %edi
movl $n, %eax
cmpl %edi,%eax
je label4
movl %eax, n
movl %edi, i
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
cmpl $1,%esp
je label2
movl %edi, mod
movl %esp, aa
movl %eax, cur
movl %ebx, rand1
movl %ecx, temp
movl %edx, i
label4:
movl $i, %eax
movl $0,%eax
movl %eax, i
label5:
movl $i, %ebx
movl $n, %esp
cmpl %ebx,%esp
je label12
movl %esp, n
movl %ebx, i
label6:
movl $i, %eax
movl $j, %edx
movl %eax,%edx
addl $1,%edx
movl %eax, i
movl %edx, j
label7:
movl $j, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label11
movl %eax, j
movl %ebx, n
label8:
movl $i, %edx
movl $t1, %ebx
movl a(,%edx,4),%ebx
movl $j, %ecx
movl $t2, %edi
movl a(,%ecx,4),%edi
cmpl %ebx,%edi
jle label10
movl %ebx, t1
movl %ecx, j
movl %edx, i
movl %edi, t2
label9:
movl $t1, %ebx
movl $t2, %ecx
movl %ebx,%ebx
addl %ecx,%ebx
movl $temporary_compiler_variable, %edi
movl %ecx,%edi
movl %ebx,%ecx
subl %edi,%ecx
movl %ebx,%ebx
subl %ecx,%ebx
movl %ebx, t1
movl %ecx, t2
movl %edi, temporary_compiler_variable
label10:
movl $t1, %ecx
movl $i, %edx
movl %ecx,a(,%edx,4)
movl $t2, %ebx
movl $j, %esp
movl %ebx,a(,%esp,4)
movl %esp,%esp
addl $1,%esp
movl $aa, %edi
cmpl $1,%edi
je label7
movl %edi, aa
movl %esp, j
movl %ebx, t2
movl %ecx, t1
movl %edx, i
label11:
movl $i, %eax
movl %eax,%eax
addl $1,%eax
movl $aa, %ebx
cmpl $1,%ebx
je label5
movl %eax, i
movl %ebx, aa
label12:
movl $i, %esp
movl $0,%esp
movl %esp, i
label13:
movl $i, %eax
movl $n, %ebx
cmpl %eax,%ebx
je label6
movl %ebx, n
movl %eax, i
label14:
movl $i, %eax
movl $t1, %ebx
movl a(,%eax,4),%ebx
movl $aa, %ecx
cmpl $1,%ecx
je label13
movl %eax, i
movl %ebx, t1
movl %ecx, aa
label15:


