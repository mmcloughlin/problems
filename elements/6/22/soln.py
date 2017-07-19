def rotate_left(a):
    """
    Returns the result of rotating the matrix a 90 degrees left.
    """
    return list(zip(*[reversed(r) for r in a]))


def print_spiral_order(a):
    """
    Print a multi-dimensional array in spiral order.
    """
    while len(a) > 0:
        print ' '.join(map(str, a.pop(0))),
        a = rotate_left(a)
    print # final new line


def main():
    a = [
            [1, 2, 3],
            [8, 9, 4],
            [7, 6, 5],
        ]
    print_spiral_order(a)


if __name__ == '__main__':
    main()
