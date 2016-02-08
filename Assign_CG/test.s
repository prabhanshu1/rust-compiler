

.globl main


.section .data
a1:
.long 69
a1end:
a2:
.long 69
a2end:
a3:
.long 69
a3end:
a4:
.long 69
a4end:
a0:
.long 69
a0end:
pivot:
.long 69
pivotend:
start:
.long 69
startend:
end:
.long 69
endend:
i:
.long 69
iend:
t1:
.long 69
t1end:
t2:
.long 69
t2end:
j:
.long 69
jend:
temporary_compiler_variable:
.long 69
temporary_compiler_variableend:
asd:
.ascii "%d\n\0"
asdend:
a:
.rept 100
.long 69
.endr
aend:
b:
.rept 100
.long 69
.endr
bend:


.section .text


main:
label1:
movl a1, %eax
movl $1,%eax
movl a2, %ebx
movl $2,%ebx
movl a3, %ecx
movl $3,%ecx
movl a4, %edx
movl $4,%edx
movl a0, %edi
movl $0,%edi
movl $3,a(,%edi,4)
movl $1,a(,%eax,4)
movl $7,a(,%ebx,4)
movl $8,a(,%ecx,4)
movl $2,a(,%edx,4)
movl pivot, %esi
movl a(,%edi,4),%esi
movl %eax,a1
movl start, %eax
movl $0,%eax
movl %eax,start
movl end, %eax
movl $4,%eax
movl %eax,end
movl i, %eax
movl $0,%eax
movl %eax, i
movl %ebx, a2
movl %ecx, a3
movl %edx, a4
movl %edi, a0
movl %esi, pivot
label2:
movl i, %eax
movl %eax, i
cmpl $4,%eax
jg label7
label3:
movl i, %eax
movl t1, %ebx
movl a(,%eax,4),%ebx
movl pivot, %ecx
movl %eax, i
movl %ebx, t1
movl %ecx, pivot
cmpl %ebx,%ecx
jl label5
label4:
movl i, %eax
movl t2, %ebx
movl a(,%eax,4),%ebx
movl start, %ecx
movl %ebx,b(,%ecx,4)
movl %ecx,%ecx
addl $1,%ecx
movl %eax, i
movl %ebx, t2
movl %ecx, start
jmp label6
label5:
movl i, %eax
movl t2, %ebx
movl a(,%eax,4),%ebx
movl end, %ecx
movl %ebx,b(,%ecx,4)
movl %ecx,%ecx
subl $1,%ecx
movl %eax, i
movl %ebx, t2
movl %ecx, end
label6:
movl i, %eax
movl %eax,%eax
addl $1,%eax
movl %eax, i
jmp label2
label7:
movl pivot, %eax
movl start, %ebx
movl %eax,b(,%ebx,4)
movl j, %ecx
movl $0,%ecx
movl %eax, pivot
movl %ebx, start
movl %ecx, j
label8:
movl j, %eax
movl %eax, j
cmpl $4,%eax
jg label11
label9:
movl i, %eax
movl t1, %ebx
movl b(,%eax,4),%ebx
pushl %ebx
pushl $asd
movl %eax, i
movl %ebx, t1
label10:
call  printf
jmp label8
label11:
movl $1,%eax
movl $0,%ebx
int $0x80
label12:


