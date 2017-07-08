def max_subset_sum(a, m=1):
    cur = 0
    best = 0
    for x in a:
        t = m*x
        cur = max([cur + t, t, 0])
        if cur > best:
            best = cur
    return best


def max_circular_subset_sum(a):
    best_regular = max_subset_sum(a)

    flipped = max_subset_sum(a, m=-1)
    total = sum(a)
    best_flipped = total + flipped

    return max([best_regular, best_flipped])


def main():
    a = [904, 40, 523, 12, -335, -385, -124, 481, -31]
    assert max_subset_sum(a) == 1479
    print max_circular_subset_sum(a)


if __name__ == '__main__':
    main()
