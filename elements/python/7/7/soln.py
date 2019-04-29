VALUES = dict(
        I=1,
        V=5,
        X=10,
        L=50,
        C=100,
        D=500,
        M=1000,
)

EXCEPTIONS = dict(
        I=['V', 'X'],
        X=['L', 'C'],
        C=['D', 'M'],
        )


def roman_numeral_value(numeral):
    v = 0
    i = 0
    n = len(numeral)
    while i < n:
        ch = numeral[i]
        if i+1 < n and numeral[i+1] in EXCEPTIONS.get(ch, []):
            v += VALUES[numeral[i+1]] - VALUES[ch]
            i += 2
        else:
            v += VALUES[ch]
            i += 1
    return v


def test():
    numerals = ['XXXXXIIIIIIIII', 'LVIIII', 'LIX']
    for numeral in numerals:
        assert 59 == roman_numeral_value(numeral)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
