def has_palindromic_permutation(word):
    counts = dict()
    for c in word:
        counts[c] = counts.get(c, 0) + 1
    num_odd = sum([n%2 for n in counts.values()])
    return num_odd < 2


def test():
    assert has_palindromic_permutation('edified')
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
