def factors(value):
    prime_factors = []
    remainder = value

    # `+ 1` accounts for iterations for `factors(2)`` and `factors(3)`
    for i in range(2, value + 1):
        if remainder == 1:
            break

        while remainder % i == 0:
            prime_factors.append(i)
            remainder //= i

    return prime_factors


print(factors(60))
