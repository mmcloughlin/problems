SEP = ':'


def pair_key(p1, p2):
    return p1 + SEP + p2 if p1 <= p2 else p2 + SEP + p1


def highest_affinity_pair(views):
    user_views = dict()
    affinity = dict()
    mx = float('-inf')
    best = None
    for page, user in views:
        if page in user_views.setdefault(user, {}):
            continue
        for p2 in user_views[user].keys():
            k = pair_key(page, p2)
            affinity[k] = affinity.get(k, 0) + 1
            if affinity[k] > mx:
                mx = affinity[k]
                best = k
        user_views[user][page] = True
    return best


def test():
    views = [
            ('ayhoo', 'ap42'),
            ('oogleg', 'ap2'),
            ('tweeter', 'thl'),
            ('oogleg', 'aa314'),
            ('oogleg', 'aa314'),
            ('oogleg', 'thl'),
            ('tweeter', 'aa314'),
            ('tweeter', 'ap42'),
            ('ayhoo', 'aa314'),
            ]
    print highest_affinity_pair(views)


def main():
    test()


if __name__ == '__main__':
    main()
