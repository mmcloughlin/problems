def closest_repeat(words):
    last_position = dict()
    d = float('inf')
    i, j = None, None
    for idx, word in enumerate(words):
        last_idx = last_position.get(word, None)
        if last_idx is not None and idx - last_idx < d:
            i, j, d = last_idx, idx, idx - last_idx
        last_position[word] = idx
    return i, j


def test():
    words = 'All work and no play makes for no work no fun and no results'.split()
    i, j = closest_repeat(words)
    assert i == 7 and j == 9
    assert 'no' == words[i]
    assert 'no' == words[j]
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
