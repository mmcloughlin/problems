def split_parts(s, pred, num):
    """
    split_parts is a general function that splits s into num substrings that
    satisfy pred.
    """
    n = len(s)
    if n == 0:
        return []
    if num == 1:
        return [[s]] if pred(s) else []
    results = []
    for p in xrange(1, n+1):
        prefix = s[:p]
        if not pred(prefix):
            continue
        for r in split_parts(s[p:], pred, num-1):
            results.append([prefix] + r)
    return results


def is_ip_part(s):
    if s[0] == '0':
        return s == '0'
    return int(s) < 256


def valid_ips(digits):
    splits = split_parts(digits, is_ip_part, 4)
    return ['.'.join(parts) for parts in splits]


def print_results(digits):
    print ' in:', digits
    ips = valid_ips(digits)
    for i, ip in enumerate(ips):
        print 'out:', ip


def main():
    print_results('19216811')
    print_results('19210001')


if __name__ == '__main__':
    main()
