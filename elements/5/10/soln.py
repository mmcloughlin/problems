def reverse_digits(x):
    if x < 0:
        return -reverse_digits(-x)
    result = 0
    while x != 0:
        result = 10*result + (x % 10)
        x /= 10
    return result


def test():
    assert reverse_digits(42) == 24
    assert reverse_digits(-314) == -413
    assert reverse_digits(0) == 0


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
