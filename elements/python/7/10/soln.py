import string
import random


def decode_rle(s):
    n = len(s)
    i = 0
    result = ''
    while i < n:
        run = 0
        while ord(s[i]) >= ord('0') and ord(s[i]) <= ord('9'):
            run = 10*run + (ord(s[i]) - ord('0'))
            i += 1
        result += s[i]*run
        i += 1
    return result


def encode_rle(s):
    n = len(s)
    i = 0
    result = ''
    while i < n:
        ch = s[i]
        run = 0
        while i < n and s[i] == ch:
            run += 1
            i += 1
        result += str(run) + ch
    return result


def random_test_case(n, max_run_length=200):
    """
    Returns random run-length-encoded data with n runs.
    """
    result = ''
    chars = random.sample(string.lowercase, n)
    for i in xrange(n):
        result += str(random.randrange(1, max_run_length+1))
        result += chars[i]
    return result


def test(trials=1000):
    # some concrete test cases
    assert decode_rle('3e4f2e') == 'eeeffffee'
    assert decode_rle('3e11f2e') == 'eeefffffffffffee'
    assert len(decode_rle('1a11b111c')) == 123
    assert encode_rle('aaaabcccaa') == '4a1b3c2a'

    # roundtrips
    for _ in xrange(trials):
        enc = random_test_case(10)
        assert encode_rle(decode_rle(enc)) == enc

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
