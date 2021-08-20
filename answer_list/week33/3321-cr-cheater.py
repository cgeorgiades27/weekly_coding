# Cheat Solution and Cheat Bonus
import heapq


def merge(list1, list2):
    return list(heapq.merge(list1, list2))


def multi_merge(list_2d):
    list_count = len(list_2d)

    for i in range(list_count):
        list_2d.extend(list_2d.pop(0))
    heapq.heapify(list_2d)

    return list_2d


# Tests
my_list = [3, 4, 6, 10, 11, 15]
alices_list = [1, 5, 8, 12, 14, 19]
bonus_list = [[3, 4, 6, 10, 11, 15], [1, 5, 8, 12, 14, 19]]

print(merge(my_list, alices_list))
print(multi_merge(bonus_list))
