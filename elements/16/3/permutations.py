def _print_permutations_final_fixed(x, n):
    """
    Print all permutations of the array x, keeping some of the final positions
    fixed.  The parameter n means how many of the initial positions of x may
    be changed. For example if len(x)=5 and n=3 this means the last 2
    positions of x must be left alone.
    """
    assert n >= 0
    assert n <= len(x)

    if n == 0:
        print x

    for i in range(n):
        y = list(x)
        y[i], y[n-1] = y[n-1], y[i]
        _print_permutations_final_fixed(y, n-1)


def print_permutations(x):
    """
    Print all permutations of the array x.
    """
    _print_permutations_final_fixed(x, len(x))
