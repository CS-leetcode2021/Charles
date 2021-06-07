
3.18_test.o：     文件格式 elf64-x86-64


Disassembly of section .text:

0000000000000000 <test>:
   0:	f3 0f 1e fa          	endbr64 
   4:	48 8d 04 37          	lea    (%rdi,%rsi,1),%rax
   8:	48 01 d0             	add    %rdx,%rax
   b:	48 83 ff fd          	cmp    $0xfffffffffffffffd,%rdi
   f:	7d 15                	jge    26 <test+0x26>
  11:	48 39 d6             	cmp    %rdx,%rsi
  14:	7d 08                	jge    1e <test+0x1e>
  16:	48 89 f8             	mov    %rdi,%rax
  19:	48 0f af c6          	imul   %rsi,%rax
  1d:	c3                   	retq   
  1e:	48 89 f0             	mov    %rsi,%rax
  21:	48 0f af c2          	imul   %rdx,%rax
  25:	c3                   	retq   
  26:	48 83 ff 02          	cmp    $0x2,%rdi
  2a:	7e 07                	jle    33 <test+0x33>
  2c:	48 89 f8             	mov    %rdi,%rax
  2f:	48 0f af c2          	imul   %rdx,%rax
  33:	c3                   	retq   
