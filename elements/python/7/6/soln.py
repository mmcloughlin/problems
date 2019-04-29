def say(digits):
    cur = digits[0]
    run = 1
    result = ''
    for digit in digits[1:]:
        if digit == cur:
            run += 1
            continue
        result += str(run) + cur
        cur = digit
        run = 1
    result += str(run) + cur
    return result


def look_and_say(n):
    digits = '1'
    for _ in xrange(n):
        digits = say(digits)
    return digits


def test():
    expect = [
            1,
            11,
            21,
            1211,
            111221,
            312211,
            13112221,
            1113213211,
            ]
    for n, digits in enumerate(expect):
        assert look_and_say(n) == str(digits)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
