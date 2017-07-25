import sys


def tail_rolling_window(f, n, out=sys.stdout):
    """
    tail_rolling_window implements tail by reading the whole file and keeping
    the last n lines in a ring buffer.
    """
    lines = [None]*n
    i = 0
    for line in f:
        lines[i%n] = line
        i += 1
    for _ in xrange(n):
        line = lines[i%n]
        if line is not None:
            out.write(line)
        i += 1


def tail_seek(f, n, out=sys.stdout):
    """
    tail_seek implements seek by seeking to the end of the file.
    """
    BUFSIZE = 128
    numnl = 0
    trailing = ''

    f.seek(0, 2)
    offset = f.tell()

    while offset > 0:
        read_bytes = min(offset, BUFSIZE)
        offset -= read_bytes
        f.seek(offset, 0)
        buf = f.read(read_bytes)
        assert len(buf) == read_bytes
        trailing = buf + trailing
        for i in xrange(read_bytes-1, -1, -1):
            if trailing[i] == '\n':
                numnl += 1
            if numnl == n+1:
                out.write(trailing[i+1:])
                return

    out.write(trailing)


def main(args, impl=tail_seek):
    filename = args[1]
    n = int(args[2])
    with open(filename) as f:
        impl(f, n)


if __name__ == '__main__':
    main(sys.argv)
