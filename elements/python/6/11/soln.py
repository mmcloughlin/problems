def longest_increasing_subarray(a):
    A = a + [float('-inf')]
    best_s, best_n = 0, 1
    s, n = 0, 1
    for i in range(len(a)):
        if A[i] <= A[i+1]:
            n += 1
            continue
        if n > best_n:
            best_s, best_n = s, n
        s = i+1
        n = 1
    return best_s, best_s + best_n


def test():
    # example in the book
    a = [2, 11, 3, 5, 13, 7, 19, 17, 23]
    r = longest_increasing_subarray(a)
    assert r == (2, 5)

    # entire thing is increasing
    a = range(13)
    r = longest_increasing_subarray(a)
    assert r == (0, 13)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
