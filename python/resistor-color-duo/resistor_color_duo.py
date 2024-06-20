def value(colors):
    color_resistance_value = {
        "black": "0",
        "brown": "1",
        "red": "2",
        "orange": "3",
        "yellow": "4",
        "green": "5",
        "blue": "6",
        "violet": "7",
        "grey": "8",
        "white": "9",
    }

    return int(color_resistance_value[colors[0]] + color_resistance_value[colors[1]])
