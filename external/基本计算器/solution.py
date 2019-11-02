# coding: utf-8


def make_scanner(src):
    return {
        'src': src,
        'length': len(src),
        'offset': 0,
        'ch': None,
    }


def update_offset(scanner, offset):
    if offset >= 0 and offset < scanner['length']:
        scanner['offset'] = offset
        scanner['ch'] = scanner['src'][offset]
        return
    scanner['offset'] = scanner['ch'] = -1


def skip_whitespace(scanner):
    offset = scanner['offset']
    while offset < scanner['length'] and scanner['src'][offset] == ' ':
        offset += 1

    update_offset(scanner, offset)


def scan_number(scanner):
    skip_whitespace(scanner)

    offset = scanner['offset']
    while (
        offset < scanner['length']
        and '0' <= scanner['src'][offset] <= '9'
    ):
        offset += 1
    num = scanner['src'][scanner['offset']: offset]

    update_offset(scanner, offset)

    skip_whitespace(scanner)
    return int(num)


def scan_opt(scanner):
    skip_whitespace(scanner)

    offset = scanner['offset']
    if scanner['src'][offset] not in '+-*/':
        return
    opt = scanner['src'][offset]

    update_offset(scanner, offset+1)

    skip_whitespace(scanner)

    return opt


def scan(scanner):
    skip_whitespace(scanner)

    if '0' <= scanner['ch'] <= '9':
        return scan_number(scanner)
    elif scanner['ch'] in '+-*/':
        return scan_opt(scanner)
    raise TypeError('Error token {} on pos {}'.format(
        scanner['ch'], scanner['offset']))


def evaluate(src):
    scanner = make_scanner(src)
    stack = []

    while scanner['ch'] != -1:
        ev = scan(scanner)
        if isinstance(ev, int):
            stack.append(ev)
        elif isinstance(ev, str):
            right = scan(scanner)
            if not isinstance(right, int):
                raise TypeError('Pos {} must be number'.format(
                    scanner['offset']))
            if ev == '*':
                left = stack.pop()
                stack.append(left * right)
            elif ev == '/':
                left = stack.pop()
                stack.append(left // right)
            elif ev == '-':
                stack.append(right * -1)
            else:
                stack.append(right)

    return sum(stack)


def test_scan_number():
    scanner = make_scanner('1234#23')
    assert scan_number(scanner) == 1234


def test_skip_whitespace():
    scanner = make_scanner('  c')
    skip_whitespace(scanner)
    assert scanner['src'][scanner['offset']] == 'c'


def test_whitespace_wrap_number():
    src = '   123  '
    scanner = make_scanner(src)

    assert scan_number(scanner) == 123
    assert scanner['offset'] == -1
    assert scanner['ch'] == -1


def test1():
    assert evaluate('3+2*2') == 7


def test2():
    assert evaluate('3/2') == 1


def main():
    print(evaluate('3+2*2'))


if __name__ == '__main__':
    main()
