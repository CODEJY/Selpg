# -*- coding: UTF-8 -*-
'''
for i in range(1,101):
    if i == 50:
        print("\f")
    print(i)
'''
fw = open("./input.txt","w")
for i in range(1,101):
    if i == 50:
        fw.write("\f")
    print >>fw,str(i)
fw.close()