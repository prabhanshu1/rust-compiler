0     &{label   label1  -3}
1     &{= b 10  0}
2     &{= n 10  0}
3     &{= mod 2713  0}
4     &{= rand1 31  0}
5     &{= rand2 101  0}
6     &{= cur 23  0}
7     &{= i 0  0}
8     &{label   label2  -3}
9     &{ifgoto je i n label4}
10     &{label   label3  -3}
11     &{* temp rand1 cur 0}
12     &{+ temp temp rand2 0}
13     &{% cur temp mod 0}
14     &{print  cur  -2}
15     &{[]= i a cur 0}
16     &{+ i i 1 0}
17     &{ifgoto je 1 aa label2}
18     &{label   label4  -3}
19     &{= i 0  0}
20     &{label   label5  -3}
21     &{ifgoto je i n label12}
22     &{label   label6  -3}
23     &{+ j i 1 0}
24     &{label   label7  -3}
25     &{ifgoto je j n label11}
26     &{label   label8  -3}
27     &{=[] t1 a i 0}
28     &{=[] t2 a j 0}
29     &{ifgoto jle t1 t2 label10}
30     &{label   label9  -3}
31     &{+ t1 t1 t2 0}
32     &{- t2 t1 t2 0}
33     &{- t1 t1 t2 0}
34     &{label   label10  -3}
35     &{[]= i a t1 0}
36     &{[]= j a t2 0}
37     &{+ j j 1 0}
38     &{ifgoto je 1 aa label7}
39     &{label   label11  -3}
40     &{+ i i 1 0}
41     &{ifgoto je 1 aa label5}
42     &{label   label12  -3}
43     &{= i 0  0}
44     &{label   label13  -3}
45     &{ifgoto je i n label6}
46     &{label   label14  -3}
47     &{=[] t1 a i 0}
48     &{print  t1  -2}
49     &{ifgoto je 1 aa label13}
50     &{label   end  -3}


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
a:
.rept 100
.long 69
.endr
aend:


.section .text


_start:
label1:
movl $b, %ebx
movl $10,%ebx
movl $n, %esp
movl $10,%esp
movl $mod, %ecx
movl $2713,%ecx
movl $rand1, %edx
movl $31,%edx
movl $rand2, %edi
movl $101,%edi
movl $cur, %eax
movl $23,%eax
movl %eax,cur
movl $i, %eax
movl $0,%eax
movl %edx, rand1
movl %edi, rand2
movl %esp, n
movl %eax, i
movl %ebx, b
movl %ecx, mod
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
movl $rand1, %edi
movl $cur, %ebx
movl %edi,%eax
mull %ebx
movl $rand2, %ecx
movl %eax,%eax
addl %ecx,%eax
movl $cur ,%ebx
movl $0,%edx
movl $cur ,%ebx
movl %eax,temp
movl $cur, %eax
movl $temp, %esp
movl $mod, %ebx
movl %esp,%eax
divl %ebx
movl %edx,%eax
movl $i, %edx
movl %eax,a(,%edx,4)
movl %edx,%edx
addl $1,%edx
movl %ecx,rand2
movl $aa, %ecx
cmpl $1,%ecx
je label2
movl %edx, i
movl %edi, rand1
movl %esp, temp
movl %eax, cur
movl %ebx, mod
movl %ecx, aa
label4:
movl $i, %eax
movl $0,%eax
movl %eax, i
label5:
movl $i, %esp
movl $n, %eax
cmpl %esp,%eax
je label12
movl %esp, i
movl %eax, n
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
movl $i, %ebx
movl $t1, %ecx
movl a(,%ebx,4),%ecx
movl $j, %edx
movl $t2, %eax
movl a(,%edx,4),%eax
cmpl %ecx,%eax
jle label10
movl %edx, j
movl %eax, t2
movl %ebx, i
movl %ecx, t1
label9:
movl $t1, %ebx
movl $t2, %esp
movl %ebx,%ebx
addl %esp,%ebx
movl $temporary_compiler_variable, %eax
movl %esp,%eax
movl %ebx,%esp
subl %eax,%esp
movl %ebx,%ebx
subl %esp,%ebx
movl %esp, t2
movl %eax, temporary_compiler_variable
movl %ebx, t1
label10:
movl $t1, %ebx
movl $i, %ecx
movl %ebx,a(,%ecx,4)
movl $t2, %edx
movl $j, %eax
movl %edx,a(,%eax,4)
movl %eax,%eax
addl $1,%eax
movl $aa, %edi
cmpl $1,%edi
je label7
movl %eax, j
movl %ebx, t1
movl %ecx, i
movl %edx, t2
movl %edi, aa
label11:
movl $i, %edx
movl %edx,%edx
addl $1,%edx
movl $aa, %ebx
cmpl $1,%ebx
je label5
movl %edx, i
movl %ebx, aa
label12:
movl $i, %edx
movl $0,%edx
movl %edx, i
label13:
movl $i, %esp
movl $n, %ecx
cmpl %esp,%ecx
je label6
movl %esp, i
movl %ecx, n
label14:
movl $i, %ebx
movl $t1, %esp
movl a(,%ebx,4),%esp
movl $aa, %ecx
cmpl $1,%ecx
je label13
movl %esp, t1
movl %ebx, i
movl %ecx, aa


