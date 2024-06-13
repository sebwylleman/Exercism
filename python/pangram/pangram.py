import string


def is_pangram(sentence: string) -> bool:

    lowercase_sentence = sentence.lower()
    
    return all(char in lowercase_sentence for char in string.ascii_lowercase)