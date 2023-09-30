"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


# TODO: define the 'EXPECTED_BAKE_TIME' constant.
EXPECTED_BAKE_TIME = 40


# TODO: Remove 'pass' and complete the 'bake_time_remaining()' function below.
def bake_time_remaining(time_in_oven):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - time_in_oven


def preparation_time_in_minutes(number_of_layers):
    """Implement the preparation_time_in_minutes(number_of_layers) function that takes the number of layers you want to add to the lasagna as an argument and returns how many minutes you would spend making them. Assume each layer takes 2 minutes to prepare."""
    return 2 * number_of_layers


# TODO: define the 'elapsed_time_in_minutes()' function below.
# Remember to add a docstring (you can copy and then alter the one from bake_time_remaining.)
def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Implement the elapsed_time_in_minutes(number_of_layers, elapsed_bake_time) function that has two parameters: number_of_layers (the number of layers added to the lasagna) and elapsed_bake_time (the number of minutes the lasagna has been baking in the oven). This function should return the total number of minutes you've been cooking, or the sum of your preparation time and the time the lasagna has already spent baking in the oven."""
    return 2 * number_of_layers + elapsed_bake_time
