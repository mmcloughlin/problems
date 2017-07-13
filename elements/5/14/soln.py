def open_doors_naive(n):
    state = [0]*n
    for i in range(1, n+1):
        for door in range(i, n+1, i):
            state[door-1] ^= 1
    return [i+1 for i, s in enumerate(state) if s == 1]


def open_doors(n):
    squares = []
    x = 1
    s = 1
    while s <= n:
        squares.append(s)
        x += 1
        s = x*x
    return squares


def test():
    assert open_doors(500) == open_doors_naive(500)
    print 'pass'


if __name__ == '__main__':
    test()
