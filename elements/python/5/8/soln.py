import sys


DIGIT_FOR_VALUE = '0123456789ABCDEF'
VALUE_FOR_DIGIT = {d: i for i, d in enumerate(DIGIT_FOR_VALUE)}


def from_base(s, b=10):
    x = 0
    for d in s:
        if d not in VALUE_FOR_DIGIT:
            raise Exception('unknown digit')
        v = VALUE_FOR_DIGIT[d]
        if v >= b:
            raise Exception('invalid digit for base')
        x = b*x + v
    return x


def to_base(x, b=10):
    s = ''
    while x != 0:
        d = DIGIT_FOR_VALUE[x % b]
        s = d + s
        x = x//b
    return s


def convert_base(s, b1, b2):
    return to_base(from_base(s, b1), b2)


def main():
    s, b1, b2 = sys.argv[1:]
    converted = convert_base(s, int(b1), int(b2))
    print converted


if __name__ == '__main__':
    main()
