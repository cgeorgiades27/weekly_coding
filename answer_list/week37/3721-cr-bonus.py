## Sieve of Eratosthenes generator
from itertools import count


def make_primes(limit, skip=0) -> object:
    bool_dict = {x:True for x in range(2, limit)}
    primes = []

    for x in bool_dict.keys():
        total = x*x

        if bool_dict[x]:
            primes.append(x)

        while total < limit:
            bool_dict[total] = False
            total += x

    return iter(primes)


def prime_generator():
    end = 30
    prime_bucket = make_primes(end)
    while prime_bucket:
        try:
            yield next(prime_bucket)
        except StopIteration:
            start = end
            end += 30
            prime_bucket = make_primes(end, start)
            continue


c = prime_generator()

for i in range(30):
    print(next(c))

