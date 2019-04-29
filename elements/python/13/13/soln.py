def find_substrings(sentence, words):
    n = len(words[0])
    total = n*len(words)
    return [sentence[i:i+total] for i in xrange(len(sentence)) if
            match_words_in_dict(sentence[i:], set(words), n)]


def match_words_in_dict(s, words, n):
    if len(words) == 0:
        return True
    if len(s) < n:
        return False
    pre = s[:n]
    if pre not in words:
        return False
    return match_words_in_dict(s[n:], words - set([pre]), n)


def main():
    sentence = 'amanaplanacanal'
    words = ['can', 'apl', 'ana']
    print find_substrings(sentence, words)


if __name__ == '__main__':
    main()
