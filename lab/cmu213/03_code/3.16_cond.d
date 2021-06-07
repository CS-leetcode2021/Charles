
3.16_cond.o：     文件格式 elf64-x86-64


Disassembly of section .text:

0000000000000000 <cond>:
   0:	f3 0f 1e fa          	endbr64 
   4:	48 85 f6             	test   %rsi,%rsi
   7:	74 08                	je     11 <cond+0x11>
   9:	48 39 3e             	cmp    %rdi,(%rsi)
   c:	7d 03                	jge    11 <cond+0x11>
   e:	48 89 3e             	mov    %rdi,(%rsi)
  11:	c3                   	retq   
