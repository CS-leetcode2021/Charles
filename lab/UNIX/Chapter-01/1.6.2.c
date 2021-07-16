#include "apue.h"

// 打印进程ID
int main(int argc, char *argv[]) {
    printf("hello world from process ID %ld\n", (long)getpid());
    exit(0);
}