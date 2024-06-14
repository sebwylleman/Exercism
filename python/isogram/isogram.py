def is_isogram(string: str) -> bool:
    unique_letters = set()
    string_lowered = string.lower()

    for char in string_lowered:
        if char in unique_letters:
            return False
        if char == " " or char == "-":
            continue
        unique_letters.add(char)

    return True