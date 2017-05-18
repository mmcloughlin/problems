import sys


def xormult(x, y):
    z = 0
    while y:
        if y&1:
            z ^= x
        x <<= 1
        y >>= 1
    return z


def main():
    x, y = map(int, sys.argv[1:])
    print '{x}@{y} = {z}'.format(x=x, y=y, z=xormult(x, y))


if __name__ == '__main__':
    main()
