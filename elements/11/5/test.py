import random

import median


def main():
    s = median.Streaming()
    data = []
    for i in range(100):
        x = random.randrange(100000)
        s.add(x)
        data.append(x)
        m = s.median()
        assert m == median.naive(data)
    print 'pass'


if __name__ == '__main__':
    main()
