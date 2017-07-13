import collections


Rect = collections.namedtuple('Rect', ['x', 'y', 'w', 'h'])


def intersect_intervals(l0, s0, l1, s1):
    """
    Compute the intersection of intervals (l0, l0+s0) and (l1, l1+s1).
    """
    l = max(l0, l1)
    r = min(l0+s0, l1+s1)
    if l > r:
        return None
    return (l, r-l)


def intersect_rectangles(r0, r1):
    xint = intersect_intervals(r0.x, r0.w, r1.x, r1.w)
    if xint is None:
        return None
    yint = intersect_intervals(r0.y, r0.h, r1.y, r1.h)
    if yint is None:
        return None
    return Rect(x=xint[0], y=yint[0], w=xint[1], h=yint[1])


def test():
    # no intersection
    r0 = Rect(0, 0, 1, 1)
    r1 = Rect(9, 9, 1, 1)
    r = intersect_rectangles(r0, r1)
    assert r is None

    # complete containment
    r0 = Rect(0, 0, 10, 10)
    r1 = Rect(5, 5, 1, 1)
    r = intersect_rectangles(r0, r1)
    assert r == r1

    # intersection
    r0 = Rect(0, 0, 2, 2)
    r1 = Rect(1, 1, 2, 2)
    r = intersect_rectangles(r0, r1)
    assert r == Rect(1, 1, 1, 1)



def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
