def column_index(ref):
    n = 0
    for ch in ref:
        x = ord(ch) - ord('A') + 1
        n = 26*n + x
    return n


def test():
    assert column_index('A') == 1
    assert column_index('Z') == 26
    assert column_index('AA') == 27
    assert column_index('AZ') == 52


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
