def find_upper_bound(x):
    """
    find_upper_bound returns an integer n with n >= len(x), without using the
    len() function. The complexity is O(log(len(x)).
    """
    n = 1
    while True:
        try:
            v = x[n]
        except IndexError:
            return n
        n *= 2


def binary_search(x, target):
    """
    binary_search looks for target in the list x, without assuming the length
    of x.
    """

    left = 0
    right = find_upper_bound(x) - 1

    while left <= right:
        mid = (left + right) // 2

        try:
            x[mid]
        except IndexError:
            right = mid - 1
            continue

        if x[mid] == target:
            return mid
        if x[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    return None


def test():
    x = list(range(6))

    i = binary_search(x, 4)
    assert i == 4

    i = binary_search(x, 1)
    assert i == 1

    i = binary_search(x, 42)
    assert i is None


if __name__ == '__main__':
    test()
