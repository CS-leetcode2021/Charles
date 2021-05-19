#!/bin/bash
echo "Hello World!"

your_name="tom"
echo $your_name

your_name="tang"
echo $your_name


# myUrl="https://www.google.com"
# readonly myUrl 变量只读
# myUrl="https://www.runoob.com"


myUrl="https://www.runoob.com"
unset myUrl   # 取消变量
echo $myUrl

string="runoob is a great site"
echo ${string:1:4} # 输出 unoo

string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4

var=http://www.aaa.com/123.html # *// 表示从左边开始删除第一个 // 号及左边的所有字符
echo ${var#*123}

# 获取字符串长度的两种方式
string="hello,everyone my name is xiaoming"
expr length "$string"
echo ${#string}
