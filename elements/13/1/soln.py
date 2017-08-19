def group_by_hash(words, h):
    groups = dict()
    for word in words:
        groups.setdefault(h(word), []).append(word)
    return groups.values()


def anagram_hash(word):
    return ''.join(sorted(word))


def anagram_groups(words):
    return group_by_hash(words, anagram_hash)


def main(wordlist='/usr/share/dict/words'):
    words = [l.strip() for l in open(wordlist)]
    for group in anagram_groups(words):
        if len(group) > 1:
            print group


if __name__ == '__main__':
    main()
