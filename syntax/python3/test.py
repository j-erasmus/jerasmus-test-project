#!/usr/local/bin/python3

# Python 3

def test():
    raise Exception from foo

def exceptions():
    try:
        print("Hello")
    except Exception as e:
        print("Exception")
