

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
label_CG1:
movl a1, %eax
movl $1,%eax
movl a2, %edi
movl $2,%edi
movl a3, %ebx
movl $3,%ebx
movl a4, %ecx
movl $4,%ecx
movl a0, %edx
movl $0,%edx
movl $3,a(,%edx,4)
movl $1,a(,%eax,4)
movl $7,a(,%edi,4)
movl $8,a(,%ebx,4)
movl $2,a(,%ecx,4)
movl pivot, %esi
movl a(,%edx,4),%esi
movl %eax,a1
movl start, %eax
movl $0,%eax
movl %edx,a0
movl end, %edx
movl $4,%edx
movl %esi,pivot
movl i, %esi
movl $0,%esi
movl %esi, i
movl %eax, start
movl %ebx, a3
movl %ecx, a4
movl %edx, end
movl %edi, a2
sita:
movl i, %ecx
movl %ecx, i
cmpl $4,%ecx
jg ram
label_CG2:
movl i, %ebx
movl t1, %ecx
movl a(,%ebx,4),%ecx
movl pivot, %edx
movl %ebx, i
movl %ecx, t1
movl %edx, pivot
cmpl %ecx,%edx
jl sham
label_CG3:
movl i, %ecx
movl t2, %edx
movl a(,%ecx,4),%edx
movl start, %ebx
movl %edx,b(,%ebx,4)
movl %ebx,%ebx
addl $1,%ebx
movl %ecx, i
movl %edx, t2
movl %ebx, start
jmp gita
sham:
movl i, %ebx
movl t2, %esi
movl a(,%ebx,4),%esi
movl end, %eax
movl %esi,b(,%eax,4)
movl %eax,%eax
subl $1,%eax
movl %esi, t2
movl %eax, end
movl %ebx, i
gita:
movl i, %esi
movl %esi,%esi
addl $1,%esi
movl %esi, i
jmp sita
ram:
movl pivot, %eax
movl start, %ebx
movl %eax,b(,%ebx,4)
movl j, %edx
movl $0,%edx
movl %ebx, start
movl %edx, j
movl %eax, pivot
lll:
movl j, %esi
movl %esi, j
cmpl $4,%esi
jg mota
label_CG4:
movl i, %eax
movl t1, %ebx
movl b(,%eax,4),%ebx
pushl %ebx
pushl $asd
movl %eax, i
movl %ebx, t1
label_CG5:
call  printf
jmp lll
mota:
movl $1,%eax
movl $0,%ebx
int $0x80
label_CG6:


