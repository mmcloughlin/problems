import collections


class MaxQueue(object):
    """
    MaxQueue provides a base queue API and tracks the largest item contained
    in the queue.
    """
    def __init__(self):
        self.items = []
        self.peaks = []

    def enqueue(self, x):
        self.items.append(x)
        while len(self.peaks) > 0 and self.peaks[-1] < x:
            self.peaks.pop(-1)
        self.peaks.append(x)

    def dequeue(self):
        if self.empty():
            return None
        x = self.items.pop(0)
        if x == self.peaks[0]:
            self.peaks.pop(0)
        return x

    def max(self):
        return self.peaks[0] if not self.empty() else None

    def empty(self):
        return len(self.items) == 0


TimestampValue = collections.namedtuple('TimestampValue', ['timestamp', 'value'])


def max_rolling_window(points, window_length):
    q = MaxQueue()
    maxima = []
    t = 0
    tail = 0
    head = 0
    while t <= points[-1].timestamp:
        while head < len(points) and points[head].timestamp <= t:
            q.enqueue(points[head].value)
            head += 1
        while points[tail].timestamp < t - window_length:
            q.dequeue()
            tail += 1
        maxima.append(TimestampValue(
            timestamp=t,
            value=q.max(),
            ))
        t += 1
    return maxima


def test():
    """
    test example from figure 9.4
    """
    points = [
            1.3, None, 2.5, 3.7,
            None, 1.4, 2.6, None,
            2.2, 1.7, None, None,
            None, None, 1.7
            ]
    maxima = [
            1.3, 1.3, 2.5, 3.7,
            3.7, 3.7, 3.7, 2.6,
            2.6, 2.6, 2.2, 2.2,
            1.7, None, 1.7,
            ]
    timestamped_points = []
    for t, p in enumerate(points):
        if p is not None:
            timestamped_points.append(TimestampValue(
                timestamp=t,
                value=p,
                ))
    results = max_rolling_window(timestamped_points, 3)
    for t, r in enumerate(results):
        assert r.timestamp == t
        assert maxima[t] == r.value
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
