RESULTS = {1: True}

def check_collatz(x):
    if x in RESULTS:
        return RESULTS[x]
    nxt = x//2 if x%2 == 0 else 3*x+1
    r = check_collatz(nxt)
    RESULTS[x] = r
    return r


def verify_collatz(n):
    """
    Verifies the collatz conjecture for all integers less than n.
    """
    return all(map(check_collatz, xrange(1, n)))


def test(n):
    assert verify_collatz(n)
    print 'pass'


def main():
    test(1000000)


if __name__ == '__main__':
    main()
