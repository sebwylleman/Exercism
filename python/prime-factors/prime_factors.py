def factors(value: int) -> list[int]:
    prime_factors = []
    remainder = value

    i = 2
    while remainder != 1:
        while remainder % i == 0:
            prime_factors.append(i)
            remainder //= i
        i += 1

    return prime_factors
