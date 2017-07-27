A = 0x5062146fc7303a6b
MASK = 0xffffffffffffffff


def h(s):
    """
    Computes a 64-bit hash of the string s.
    """
    H = 0
    for c in s:
        H = (A*H + ord(c)) & MASK
    return H


def strings_equal(S, T):
    return all([s == t for s, t in map(None, S, T)])


def find(needle, haystack):
    n = len(needle)
    B = A**n & MASK
    if len(haystack) < n:
        return -1
    target_hash = h(needle)
    current_hash = h(haystack[:n])
    for i in xrange(len(haystack)-n):
        if current_hash == target_hash and strings_equal(needle, haystack[i:i+n]):
            return i
        current_hash *= A
        current_hash -= B * ord(haystack[i])
        current_hash += ord(haystack[i+n])
        current_hash &= MASK
    return -1


def test():
    haystack = 'The Fat Cat Sat on the Mat'
    needle = 'Sat'
    assert 12 == find(needle, haystack)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
