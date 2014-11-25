#!/usr/bin/env python
#-*- encoding:UTF-8 -*-
import sys, os, time
def check(line):
    size = len(line)
    for i in range(0, size-1):
        if (line[i]=="/" and line[i+1]=="/"):
            return True
    return False
def getLine(file_name):
    f = open(file_name)
    total = 0
    com = 0
    while True:
        line = f.readline()
        if not line:
            break
        total = total + 1
        if check(line):
            com = com + 1
    return com, total
def getname(name):
    words = name.split("/")
    size = len(words)
    return words[size-1]

def Update(file_name):
    print file_name

    filename = getname(file_name)
    f = open(file_name)
    lines = f.read().split("\n")
    f.close()
    package_name = lines[0].split(" ")[1]

    stat_info = time.localtime(os.stat(file_name).st_ctime)
    date = str(stat_info.tm_year) + "-"+str(stat_info.tm_mon)+"-"+str(stat_info.tm_mday)


    content = '''/***************************************************************************************
**文件名：'''+filename+'''
**包名称：'''+package_name+'''
**创建日期：'''+date+'''
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/\n'''


    page = lines[0] + "\n"+"\n" +content+"\n"
    size = len(lines)
    for i in range(1, size):
        page = page + lines[i] + "\n"

    f_out =open(file_name,"w")
    f_out.write(page)
    f_out.close()

path = ""+sys.argv[1]

for a, b, c in os.walk(path):
    for filename in c:
        file_name = a + "/" + filename
        if file_name[-2:]!='go':
            continue
        Update(file_name)



