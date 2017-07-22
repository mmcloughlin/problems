import string


def isalnum(s):
    return set(s) <= set(string.letters + string.digits)


def chars_equal(c1, c2):
    return (ord(c1)^ord(c2))&(~32) == 0


def palindromic(s):
    n = len(s)
    l = 0
    r = n - 1
    while l <= r:
        while l < n and not isalnum(s[l]):
            l += 1
        while r >= 0 and not isalnum(s[r]):
            r -= 1
        if l >= n or r < 0:
            break
        if not chars_equal(s[l], s[r]):
            return False
        l += 1
        r -= 1
    return True


def test():
    assert palindromic('A man, a plan, a canal, Panama')
    assert palindromic('Able was I, ere I saw Elba!')
    assert not palindromic('Ray a Ray')
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
