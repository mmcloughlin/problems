import sys


def rotate(s, n):
    """
    Circularly rotate s by n positions.
    """
    return s[n:] + s[:n]


def first_lexicographic_rotation(s):
    """
    Sub-optimal but "obviously" correct solution.
    """
    n = len(s)
    ordering = sorted(range(n), key=lambda i: rotate(s, i))
    best = ordering[0]
    return best, rotate(s, best)


def main():
    word = sys.argv[1]
    n, rot = first_lexicographic_rotation(word)
    print n, rot


if __name__ == '__main__':
    main()
