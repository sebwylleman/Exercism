def find_anagrams(word: str, candidates: list) -> set:
    anagrams = []
    for candidate in candidates:
        if char_frequency(word) == char_frequency(candidate) and candidate.lower() != word.lower():
            anagrams.append(candidate)
    return anagrams

def char_frequency(word: str) -> dict:
    frequency = {}
    lower_word = word.lower()
    for char in lower_word:
        if char not in frequency:
            frequency[char] = 0
        frequency[char] += 1
    return frequency