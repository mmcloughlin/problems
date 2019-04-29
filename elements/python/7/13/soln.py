import sys


def format_line(A, L, space=' '):
    if len(A) == 1:
        w = A[0]
        return w + space*(L - len(w))
    total_space = L - sum(map(len, A))
    num_spaces = len(A)-1
    line = A[0]
    for w in A[1:]:
        n = (total_space + num_spaces - 1) // num_spaces
        line += space*n + w
        num_spaces -= 1
        total_space -= n
    return line


def justify(A, L=78, space=' '):
    current_line = [A[0]]
    current_chars_min = len(A[0])
    lines = []
    for w in A[1:]:
        next_chars_min = current_chars_min + 1 + len(w)
        if next_chars_min > L:
            lines.append(format_line(current_line, L, space=space))
            current_line = []
            current_chars_min = -1
        current_line.append(w)
        current_chars_min += 1 + len(w)
    if len(current_line) > 0:
        last = space.join(current_line)
        last += space * (L - len(last))
        lines.append(last)
    return lines


def display_lines(lines):
    for line in lines:
        print '|{}|'.format(line)


def test():
    A = 'The quick brown fox jumped over the lazy dogs.'.split()
    L = 11
    lines = justify(A, L, space='_')
    display_lines(lines)


def main(args):
    filename = args[1]
    L = int(args[2])
    A = []
    with open(filename) as f:
        for line in f:
            words = line.strip().split()
            A.extend(words)
    lines = justify(A, L, space='_')
    display_lines(lines)


if __name__ == '__main__':
    main(sys.argv)
