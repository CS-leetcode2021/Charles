
exchange.o：     文件格式 elf64-x86-64


Disassembly of section .text:

0000000000000000 <exchange>:
   0:	f3 0f 1e fa          	endbr64 
   4:	48 8b 07             	mov    (%rdi),%rax      // rax 是返回值的寄存器 rdi是第一个参数的寄存器，rsi是第二个参数的寄存器 
                                                      // 将rdi的地址的数据mov到rax中作为返回值
   7:	48 89 37             	mov    %rsi,(%rdi)      // 将第二个数据mov到第一个数据中去
   a:	c3                   	retq                    // 返回一个四字长的类型值，也就是8个字节
                                                      // 从内存读取数据到寄存器再从寄存器到内存
