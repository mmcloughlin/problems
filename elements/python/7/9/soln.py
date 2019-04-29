def sinusoidal(s):
    n = len(s)
    result = ''
    for i in xrange(1, n, 4): result += s[i]
    for i in xrange(0, n, 2): result += s[i]
    for i in xrange(3, n, 4): result += s[i]
    return result


def main():
    assert sinusoidal('Hello_World!') == 'e_lHloWrdlo!'
    print 'pass'


if __name__ == '__main__':
    main()
