def rotate(x, k):
    n = len(x)
    total_moved = 0
    cycle_start = 0
    while total_moved < n:
        last = x[cycle_start]
        i = (cycle_start + k) % n
        while i != cycle_start:
            x[i], last = last, x[i]
            i = (i + k) % n
            total_moved += 1
        x[cycle_start] = last
        total_moved += 1
        cycle_start += 1


def test(n, k):
    x = list(range(n))
    r = list(x)
    rotate(r, k)
    for i in range(n):
        assert x[i] == r[(i+k)%n]


if __name__ == '__main__':
    n = 24
    for k in range(n):
        test(n, k)
