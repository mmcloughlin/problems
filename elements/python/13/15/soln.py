SEP = ':'


def attr_key(attrs):
    return SEP.join(sorted(attrs))


def pair_users(users):
    attr_to_user = dict()
    pairs = []
    for user_id, attrs in users:
        k = attr_key(attrs)
        if k not in attr_to_user:
            attr_to_user[k] = user_id
            continue
        pairs.append((attr_to_user[k], user_id))
        del attr_to_user[k]
    return pairs, list(attr_to_user.values())


def main():
    users = [
            (12345, ['a', 'b']),
            (43252345, ['b', 'a']),
            (2323423, ['c', 'a']),
            (3423, ['c', 'a', 'd']),
            (34239, ['a', 'c']),
            ]
    print pair_users(users)


if __name__ == '__main__':
    main()
