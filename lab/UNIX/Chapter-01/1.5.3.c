#include "apue.h"

#define BUFFSIZE 4096
// 可用于复制任一UNIX普通文件
int main(int argc, char *argv[]) {
    int n;
    char buf[BUFFSIZE];

    // STDIN_FILENO 用户的标准输入
    while ((n = read(STDIN_FILENO, buf, BUFFSIZE)) > 0)
        // STDOUT_FILENO 用户的标准输出
        if (write(STDOUT_FILENO, buf, n) != n)
            err_sys("write error");

    if (n < 0)
        err_sys("read error");
    exit(0);
}