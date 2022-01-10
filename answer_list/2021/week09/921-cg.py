# week 9, 2021

def intersection(list1, list2, list3):
    maxNumSize = 200
    arr = [0] * maxNumSize
    combinedList = list1 + list2 + list3

    for i in combinedList:
        arr[i] += 1

    for i in range(len(arr)):
        if arr[i] == 3:
            return i


print(intersection([1, 2, 3, 4], [2, 4, 6, 8], [3, 4, 5]))
# [4]
