#!/usr/bin/env python
# -*- coding:utf-8 -*-
a = "hello"


def main():
    global a
    print(f"global a is {a}, address: {id(a)}")
    a = "hello1"
    print(f"local a is {a}, address: {id(a)}")
    for i in range(2):
        a = "hello2"
        print(f"for a is {a}, address: {id(a)}")
        if i % 2 != 0:
            a = "hello3"
            print(f"if a is {a}, address: {id(a)}")


if __name__ == "__main__":
    main()



"""
执行结果 Python
global a is hello, address: 2056553020144
local a is hello1, address: 2056553020912
for a is hello2, address: 2056554223984
for a is hello2, address: 2056554223984
if a is hello3, address: 2056554224304

"""