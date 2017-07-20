def verify_sudoku(S):
    rows = [0]*9
    cols = [0]*9
    blocks = [0]*9

    for y, row in enumerate(S):
        for x, digit in enumerate(row):
            flag = 1 << digit
            rows[y] |= flag
            cols[x] |= flag
            block_idx = (x//3) + 3*(y//3)
            blocks[block_idx] |= flag

    return all([mask == 0b1111111110 for mask in rows + cols + blocks])


def generate_valid_sudoku():
    return [[1+(i + j) % 9 for j in xrange(9)] for i in [0,3,6,1,4,7,2,5,8]]


def test():
    S = generate_valid_sudoku()
    assert verify_sudoku(S)

    S = [range(1,10) for _ in xrange(9)]
    assert not verify_sudoku(S)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
