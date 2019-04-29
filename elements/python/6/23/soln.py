def rotate_right_out_of_place(A):
    return map(list, zip(*reversed(A)))


def rotate_right_in_place(A):
    n = len(A)
    for X in range((n+1)/2):
        for Y in range(n/2):
            x, y = X, Y
            t = A[y][x]
            for _ in range(4):
                x, y = n-1-y, x
                A[y][x], t = t, A[y][x]


def test(n=4):
    print 'TEST n={n}'.format(n=n)
    A = [range(i, i+n) for i in range(0, n*n, n)]
    expect = rotate_right_out_of_place(A)
    print 'before:', A
    rotate_right_in_place(A)
    print 'after:', A
    print 'expect:', expect
    assert expect == A


def main():
    for n in range(1, 10):
        test(n)


if __name__ == '__main__':
    main()
