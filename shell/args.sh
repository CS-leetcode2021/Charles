#!/bin/bash
# author:菜鸟教程
# url:www.runoob.com

# /args.sh 1 2 3
echo "Shell 传递参数实例！";
echo "执行的文件名：$0";
echo "第一个参数为：$1";
echo "第二个参数为：$2";
echo "第三个参数为：$3";

echo "传递的参数作为一个字符串显示：$*";

# $* 等价于传递了一个参数
# $@ 等价于传递了多个参数

echo "---$*演示---"
# shellcheck disable=SC2066
for i in "$*"; do
  echo %i
  done