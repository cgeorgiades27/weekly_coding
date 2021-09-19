## Sieve of Eratosthenes O(log n log n)

def get_primes(limit: int) -> list:

    bool_dict = {x:True for x in range(2, limit)}
    prime_set = []

    if limit < 3:
        return list()

    for x in bool_dict.keys():
        if bool_dict[x]:
            prime_set.append(x)
        total = x*x
        while total < limit:
            bool_dict[total] = False
            total += x

    return prime_set


# Tests
small = 100
medium = 1000
large = 100000
dont_do_it = 100000000

import timeit
print(timeit.timeit('get_primes(small)', number=10, globals=globals()))
#print(get_primes(large))

