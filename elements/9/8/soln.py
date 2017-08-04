class Stack(object):
    def __init__(self):
        self.stack = []

    def push(self, x):
        self.stack.append(x)

    def pop(self):
        if self.size() == 0:
            return None
        return self.stack.pop(-1)

    def peek(self):
        if self.size() == 0:
            return None
        return self.stack[-1]

    def size(self):
        return len(self.stack)


def sunset_views(apartment_heights):
    s = Stack()
    for h in apartment_heights:
        while s.peek() is not None and s.peek() <= h:
            s.pop()
        s.push(h)
    return s.stack


def test():
    apartments = [70, 33, 1, 45, 50, 12, 24, 23, 25, 25, 6, 12, 3, 7, 12, 5, 6]
    for h in apartments:
        print '{bar} {h}'.format(bar='#'*h, h=h)
    views = sunset_views(apartments)
    assert views == [70, 50, 25, 12, 6]


def main():
    test()


if __name__ == '__main__':
    main()
