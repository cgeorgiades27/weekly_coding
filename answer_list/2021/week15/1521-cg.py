'''
This problem was asked by Google.

Given two strings A and B, return whether or not A can be shifted some number of times to get B.

For example, if A is abcde and B is cdeab, return true. If A is abc and B is acb, return false
'''


import sys

# O(n^2)
def rollover1(original, rolled):

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

# O(n)
def rollover2(original, rolled):
    arr = [ char for char in original]
    for i in range(len(original)):
        arr.insert(0,arr.pop())
        print("roll {0}: {1}".format(i,''.join(arr)))
        if ''.join(arr) == rolled:
            print("matched on roll: {0}".format(i))
            return True
    return False



if len(sys.argv) > 1:
    print("Rollover method 1:")
    print(rollover1(sys.argv[1],sys.argv[2]))
    print("Rollover method 2:")
    print(rollover2(sys.argv[1],sys.argv[2]))

else:
    original1 = "abcde"
    rolled1 = "cdeab"
    print("Rollover method 1:")
    print(rollover1(original1,rolled1))
    print("Rollover method 2:")
    print(rollover2(original1,rolled1))
    original2 = "abc"
    rolled2 = "acb"
    print("Rollover method 1:")
    print(rollover1(original2,rolled2))
    print("Rollover method 2:")
    print(rollover2(original2,rolled2))