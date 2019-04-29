def pascal(n):
    """
    Computes the first n rows of Pascal's Triangle.
    """
    rows = [[1]]
    for i in range(1, n):
        row = [1] + [rows[i-1][j] + rows[i-1][j+1] for j in xrange(i-1)] + [1]
        rows.append(row)
    return rows


def print_pascal(n=10):
    rows = pascal(n)
    assert len(rows) == n
    for row in rows:
        print ','.join(map(str, row))


def main():
    print_pascal(10)


if __name__ == '__main__':
    main()
