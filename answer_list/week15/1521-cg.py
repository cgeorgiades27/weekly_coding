'''
This problem was asked by Google.

Given two strings A and B, return whether or not A can be shifted some number of times to get B.

For example, if A is abcde and B is cdeab, return true. If A is abc and B is acb, return false
'''


import sys

def rollover(original, rolled):

    placeholder = []
    for i in range(len(original)):
        placeholder.append('-')
    for i in range(len(original)):
        for j in range(len(original)):
            placeholder[j] = original[(i + j) % len(original)]
        print("roll {0}: {1}".format(i,''.join(placeholder)))
        if ''.join(placeholder) == rolled:
            print("matched on roll: {0}".format(i))
            return True
    return False
        


if len(sys.argv) > 1:
    print(rollover(sys.argv[1],sys.argv[2]))
else:
    original1 = "abcde"
    rolled1 = "cdeab"
    print(rollover(original1,rolled1))
    original2 = "abc"
    rolled2 = "acb"
    print(rollover(original2,rolled2))