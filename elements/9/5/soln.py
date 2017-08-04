def normalize_path(path):
    parts = path.split('/')
    if len(parts) == 0:
        return ''
    is_absolute = parts[0] == ''
    stack = []
    for part in parts:
        if part == '' or part == '.':
            continue
        if part == '..':
            if len(stack) > 0 and stack[-1] != '..':
                stack.pop(-1)
                continue
        stack.append(part)

    path = '/'.join(stack)
    if not is_absolute:
        return path
    if path.startswith('..'):
        return None
    return '/' + path


def test():
    assert normalize_path('/usr/lib/../bin/gcc') == '/usr/bin/gcc'
    assert normalize_path('scripts//./../scripts/awkscripts/./.') == 'scripts/awkscripts'
    assert normalize_path('/usr/../../bin/gcc') is None
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
