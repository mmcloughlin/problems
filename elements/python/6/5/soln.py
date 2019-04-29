import random


def remove(a, k):
    """
    Remove all occurances of k from a.
    """
    n = len(a)
    rd = 0
    wr = 0
    while rd < n:
        if a[rd] != k:
            a[wr] = a[rd]
            wr += 1
        rd += 1
    del a[wr:]
    return wr


def main():
    a = [random.randrange(3) for _ in range(16)]
    print a
    m = remove(a, 2)
    print m, a


if __name__ == '__main__':
    main()
