def swap(x, i, j):
    """
    swap bits i and j of x.
    """
    mask = (1<<i) | (1<<j)
    m = x&mask
    if m == 0 or m == mask:
        return x
    return x^mask
