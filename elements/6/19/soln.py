import random


class Sampler(object):
    """
    Sampler maintains a uniform sample of size k from streaming data.
    """

    def __init__(self, k):
        self.n = 0
        self.k = k
        self.sampled = []

    def add(self, x):
        self.n += 1
        if self.n <= self.k:
            self.sampled.append(x)
            return
        p = self.k / float(self.n)
        if random.random() > p:
            return
        i = random.randrange(self.k)
        self.sampled[i] = x

    def sample(self):
        return self.sampled


def test(k, n):
    s = Sampler(k)
    for i in xrange(k):
        assert i == len(s.sample())
        s.add(i)
    for i in xrange(k, n):
        assert k == len(s.sample())
        s.add(i)
    print s.sample()
    print 'pass'


def main():
    test(10, 10000)


if __name__ == '__main__':
    main()
