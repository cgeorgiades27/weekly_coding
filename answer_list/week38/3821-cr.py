# Subtraction method

def sub_division(numerator: int, denominator: int) -> int:
    quotient = 0
    top = abs(numerator)
    bottom = abs(denominator)
    # check zero division
    if bottom == 0:
        raise ValueError(f"Division by zero ({numerator} ÷ {denominator})")

    if numerator < 0 and denominator < 0 or numerator > 0 and denominator > 0:
        sign = 1
    else:
        sign = -1

    while top >= bottom:
        top -= bottom
        quotient += 1
    return quotient * sign


# Bitwise

def bit_division(numerator: int, denominator: int) -> int:
    quotient = 0
    top = abs(numerator)
    bottom = abs(denominator)

    if bottom == 0:
        raise ValueError(f"Division by zero ({numerator} ÷ {denominator})")

    if numerator < 0 and denominator < 0 or numerator > 0 and denominator > 0:
        sign = 1
    else:
        sign = -1

    while top >= bottom:
        count = 0
        while top >= (bottom << count):
            count += 1
        count -= 1
        top = top - (bottom << count)
        quotient += 1 << count

    return quotient * sign


# Tests
from random import randint

for _ in range(10):
    dividend = randint(-1000, 1000)
    divisor = randint(-100, 100)
    try:
        print(f"sub: {sub_division(dividend, divisor)}",
              f" bit: {bit_division(dividend, divisor)}", sep=',')
    except Exception as e:
        print(e)

# Static cases
# print(sub_division(440, 8))
# print(sub_division(183, 3))
# print(sub_division(-2040, 24))
# print(sub_division(200, 0))
# print(bit_division(440, 8))
# print(bit_division(-2040, 24))
# print(bit_division(183, 3))
# print(bit_division(200, 0))

# Timeit
import timeit

print("sub division time:",
      timeit.timeit('sub_division(randint(-1000, 1000), randint(1, 100))',
                    number=100, globals=globals()))
print("bit division time:",
      timeit.timeit('bit_division(randint(-1000, 1000), randint(1, 100))',
                    number=100, globals=globals()))
