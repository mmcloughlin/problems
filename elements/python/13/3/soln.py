import string


def char_counts(text):
    counts = {}
    for ch in text:
        if ch in string.whitespace:
            continue
        counts[ch] = counts.get(ch, 0) + 1
    return counts


def can_construct_letter(letter, magazine):
    letter_counts = char_counts(letter)
    magazine_counts = char_counts(magazine)
    for ch, required in letter_counts.items():
        if magazine_counts.get(ch, 0) < required:
            return False
    return True


def test():
    assert can_construct_letter('brown fox', 'the quick brown fox jumped ...')
    assert not can_construct_letter('blue fox', 'the quick brown fox jumped ...')
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
