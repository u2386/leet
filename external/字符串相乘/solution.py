# coding: utf-8


def add(a0, a1):
    if a0 == '0':
        return a1
    elif a1 == '0':
        return a0

    ret = []

    i = len(a0) - 1
    j = len(a1) - 1
    carry = 0
    while i >= 0 and j >= 0:
        s = int(a0[i]) + int(a1[j]) + carry
        ret.append(str(s % 10))
        carry = 1 if s >= 10 else 0
        i -= 1
        j -= 1

    while i >= 0:
        s = int(a0[i]) + carry
        ret.append(str(s % 10))
        carry = 1 if s >= 10 else 0
        i -= 1

    while j >= 0:
        s = int(a1[j]) + carry
        ret.append(str(s % 10))
        carry = 1 if s >= 10 else 0
        j -= 1

    if carry:
        ret.append('1')
    return ''.join(ret[::-1])


def _multiply(s0, c):
    ret = []

    multiplier = int(c)
    if not multiplier:
        return '0'

    carry = 0
    i = len(s0) - 1
    while i >= 0:
        s = int(s0[i]) * multiplier + carry
        ret.append(str(s % 10))
        carry = s // 10
        i -= 1

    if carry:
        ret.append(str(carry))
    return ''.join(ret[::-1])


def multiply(s0, s1):
    if s0 == '0' or s1 == '0':
        return '0'
    j = len(s1)

    ret = '0'
    product = '0'
    while j > 0:
        j -= 1
        if s1[j] == '0':
            continue

        product = _multiply(s0, s1[j])
        product = product + '0' * (len(s1) - j - 1)
        ret = add(ret, product)
    return ret


class Solution(object):
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        return multiply(num1, num2)
