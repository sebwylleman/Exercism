def equilateral(sides):
    return is_triangle(sides) and len(set(sides)) == 1


def isosceles(sides):
    return is_triangle(sides) and len(set(sides)) == 2 or equilateral(sides)


def scalene(sides):
    return is_triangle(sides) and not equilateral(sides) and not (isosceles(sides))


def is_triangle(sides):
    a, b, c = sides
    if a * b * c <= 0:
        return False
    return a + b >= c and b + c >= a and a + c >= b
