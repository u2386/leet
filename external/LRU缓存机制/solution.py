# coding: utf-8


class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.m = {}
        self.q = []
        self.cap = capacity

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.m:
            return -1

        val = self.m[key]
        del self.m[key]
        self.q.remove(key)
        self.put(key, val)
        return val

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        if key in self.m:
            self.q.remove(key)

        self.q.append(key)
        self.m[key] = value

        if len(self.q) > self.cap:
            expire_key = self.q.pop(0)
            del self.m[expire_key]


def main():
    pass


if __name__ == '__main__':
    main()
