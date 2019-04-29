import random


KEYPAD = {
        '2': 'ABC',
        '3': 'DEF',
        '4': 'GHI',
        '5': 'JKL',
        '6': 'MNO',
        '7': 'PQRS',
        '8': 'TUV',
        '9': 'WXYZ',
        }


def mnemonics(number):
    if number == '':
        return ['']
    digit = number[0]
    results = []
    for char in KEYPAD.get(digit, []):
        for sub in mnemonics(number[1:]):
            results.append(char + sub)
    return results


def test(n=10):
    number = ''
    expect = 1
    for _ in xrange(n):
        digit = str(random.randrange(2, 10))
        expect *= len(KEYPAD[digit])
        number += digit
    results = mnemonics(number)
    assert expect == len(results)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
