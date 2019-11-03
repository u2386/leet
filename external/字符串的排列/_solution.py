# coding: utf-8


def _premut(s, index, prem, ret, used):
    if index == len(s):
        ret.append(prem.copy())
        return

    for i in range(len(s)):
        if used[i] or (s[i] == s[i-1] and used[i-1]):
            continue

        prem.append(s[i])
        used[i] = 1

        _premut(s, index+1, prem, ret, used)

        used[i] = 0
        prem.pop()


def premutation(s):
    if len(s) == 0:
        return []
    elif len(s) == 1:
        return [s]

    ret = []
    used = [0 for _ in range(len(s))]
    _premut(sorted(s), 0, [], ret, used)
    return ret


def get_next(pat):
    nxt = [0] * len(pat)
    nxt[0] = -1

    k = -1
    j = 0

    while j < len(pat) - 1:
        if k == -1 or pat[j] == pat[k]:
            k += 1
            j += 1
            nxt[j] = k
        else:
            k = nxt[k]
    return nxt


def test1():
    s = []
    assert premutation(s) == []


def test2():
    s = [1]
    assert premutation(s) == [[1]]


def test3():
    s = [1, 2, 3]
    assert premutation(s) == [
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1]
    ]


def test4():
    s = [1, 1, 2]
    assert premutation(s) == [
        [1, 1, 2],
        [1, 2, 1],
        [2, 1, 1],
    ]


def main():
    s = [1, 1, 2]
    print(premutation(s))


if __name__ == '__main__':
    main()
