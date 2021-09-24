## Sieve of Eratosthenes generator


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

    return iter(primes[skip:])


def prime_generator():
    indx = 1
    end = 30
    prime_bucket = make_primes(end)
    while prime_bucket:
        try:
            yield next(prime_bucket)
            indx += 1
        except StopIteration:
            end += 30
            prime_bucket = make_primes(end, indx)
            continue

# Tests
import timeit
prime = prime_generator()
print(timeit.timeit('next(prime)', number=100, globals=globals()))

#for i in range(30):
#    print(next(prime))

