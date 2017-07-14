import random


def dutch_flag(a, i):
    n = len(a)
    v = a[i]
    lo, hi = 0, n
    j = 0
    while j < hi:
        if a[j] == v:
            j += 1
            continue
        if a[j] < v:
            a[lo], a[j] = a[j], a[lo]
            lo += 1
            j += 1
            continue
        if a[j] > v:
            hi -= 1
            a[hi], a[j] = a[j], a[hi]


def test(n=20):
    a = [random.randrange(3) for i in range(n)]
    print a
    dutch_flag(a, n/2)
    print a


def main():
    test()


if __name__ == '__main__':
    main()
