class Trie(object):
    def __init__(self):
        self.children = {}

    def add(self, s):
        if len(s) == 0:
            return
        self.children.setdefault(s[0], Trie()).add(s[1:])

    def unique_prefix(self, s):
        if len(s) == 0:
            return ''
        sub = self.children.get(s[0])
        rest = sub.unique_prefix(s[1:]) if sub is not None else ''
        return s[0] + rest


def test():
    t = Trie()
    t.add('dog')
    t.add('be')
    t.add('cut')
    assert t.unique_prefix('cat') == 'ca'

    t.add('car')
    assert t.unique_prefix('cat') == 'cat'

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
