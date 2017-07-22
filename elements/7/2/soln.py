import random


def replace_as_remove_bs(s, n):
    # remove bs first
    wr = 0
    num_a = 0
    for rd in xrange(n):
        if s[rd] == 'b':
            n -= 1
            continue
        if s[rd] == 'a':
            num_a += 1
        s[wr] = s[rd]
        wr += 1
    # replace "a"s with "dd"s, backwards
    final_len = n + num_a
    wr = final_len - 1
    rd = n - 1
    while rd >= 0:
        c = s[rd]
        if c == 'a':
            s[wr] = 'd'
            s[wr-1] = 'd'
            wr -= 2
        else:
            s[wr] = c
            wr -= 1
        rd -= 1
    return ''.join(s[:final_len])


def test(trials=1000, n=100):
    for trial in xrange(trials):
        s = ''.join([random.choice('abcd') for _ in xrange(n)])
        expect = s.replace('a', 'dd').replace('b', '')
        expanded = list(s) + [' ']*n
        assert expect == replace_as_remove_bs(expanded, n)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
