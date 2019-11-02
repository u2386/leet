# coding: utf-8


def count_bitset(nums):
    ret = [0, 1]

    for i in nums[2:]:
        if i & 1 == 1:
            ret.append(ret[i-1] + 1)
        else:
            ret.append(ret[i//2])
    return ret


def test1():
    assert count_bitset([0, 1, 2]) == [0, 1, 1]


def test2():
    assert count_bitset([0, 1, 2, 3, 4, 5]) == [0, 1, 1, 2, 1, 2]


def main():
    n = 10

    if n == 0:
        return [0]
    elif n == 1:
        return [0, 1]

    print(count_bitset([i for i in range(n+1)]))


if __name__ == '__main__':
    main()
