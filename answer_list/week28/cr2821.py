# Naive

def last_man_standing(souls: int, skipafew: int) -> int:
    """Takes a number of people 'souls' to be executed and an increment 'skipafew', returns the last to be executed"""
    cnt = 0
    goner = 0
    soul_list = [x for x in range(1, souls+1)]
    while souls > 1:
        cnt += 1
        goner += 1
        if cnt == skipafew-1:
            soul_list.pop(goner)
            souls -= 1
            cnt = 0
        if goner >= souls-1:
            goner -= souls
    return soul_list.pop()


# Test
print(last_man_standing(5, 2))
print(last_man_standing(6, 2))
print(last_man_standing(8, 3))
