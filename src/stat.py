def check(line):
    size = len(line)
    if size==0:
        return False
    for i in range(0, size-1):
        if (line[i]=="/" and line[i+1]=="/"):
            return True
    if line[0]=="*":
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
import sys
keywords = ["go","java","cpp","c","py","h"]
def check_keyword(item,file_name):
    l = len(item)
    l = l + 1
    name = file_name[-l:]
    if name=="."+item:
        return True
    return False
path = ""+sys.argv[1]
import os
total = 0
total1 = 0
for a, b, c in os.walk(path):
    for filename in c:
        file_name = a + "/" + filename
        flag = True
        for item in keywords:
            if check_keyword(item, file_name):
                flag = False
        if flag==True:
            continue
        tmp1, tmp= getLine(file_name)
        print file_name+ " : " + str(tmp1)+"/"+str(tmp)
        total = total + tmp
        total1 = total1 + tmp1
print "TOTAL : " + str(total1)+"/"+str(total)
print "PERCENTAGE : " + str(100 * float(total1)/total)+"%"

