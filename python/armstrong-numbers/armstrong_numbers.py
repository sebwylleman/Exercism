import functools


def is_armstrong_number(number):
    number_str = str(number)
    result = sum([int(n) ** len(number_str) for n in number_str])
    return result == number
