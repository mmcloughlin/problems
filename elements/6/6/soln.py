def dedupe(a):
    if len(a) == 0:
        return 0
    rd = 1
    wr = 1
    while rd < len(a):
        if a[rd] != a[wr-1]:
            a[wr] = a[rd]
            wr += 1
        rd += 1
    return wr


def test():
    a = [2,3,5,5,7,11,11,11,13]
    n = dedupe(a)
    assert a[:n] == [2,3,5,7,11,13]
    print 'pass'


if __name__ == '__main__':
    test()
