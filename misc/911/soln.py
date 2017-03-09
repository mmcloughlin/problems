def count_non_emergency_numbers(n):
    """
    Compute the number of n-digit numbers (not starting with 0) that do not
    contain "911".
    """
    a = []
    for m in range(n+1):
        max_k = m // 3
        c = [10**m] + [0]*n
        for k in range(1, max_k+1):
            cnt = 0
            for i in range(m-3+1):
                j = m-3-i
                cnt += a[i][0] * a[j][k-1]
            c[k] = cnt
            c[0] -= cnt
        a.append(c)
    return a[n][0] - a[n-1][0]



def naive(n):
    """
    Uses a naive (obviously correct) method to compute the number of n-digit
    numbers don't start with 0 and that don't contain "911".
    """
    count = 0
    for x in range(10**n):
        s = str(x)
        if len(s) == n and s.find('911') == -1:
            count += 1
    return count


def verify(n):
    """
    Confirm that count_non_emergency_numbers gives the same answer as a naive
    computation of the same thing.
    """
    expect = naive(n)
    got = count_non_emergency_numbers(n)
    print '{result} expect={expect} got={got}'.format(
            result='ok' if expect == got else 'FAIL',
            expect=expect,
            got=got,
            )


def main():
    print count_non_emergency_numbers(10)


if __name__ == '__main__':
    main()
