# coding: utf-8


class AllOne(object):

    class FreqNode(object):

        def __init__(self, freq):
            self.freq = freq
            self.prev = None
            self.after = None
            self.keys = []

    class FreqList(object):

        def __init__(self):
            self.head = AllOne.FreqNode(None)
            self.rear = AllOne.FreqNode(None)
            self.head.after = self.rear
            self.rear.prev = self.head

        def get_max(self):
            return self.rear.prev.keys[-1] if self.rear.prev.keys else ''

        def get_min(self):
            return self.head.after.keys[-1] if self.head.after.keys else ''

        def add(self, node, prev, after):
            node.prev = prev
            node.after = after
            prev.after = node
            after.prev = node

        def remove(self, node):
            node.prev.after = node.after
            node.after.prev = node.prev

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.kv = {}
        self.order = [None] * 10000
        self.freq = AllOne.FreqList()

    def inc(self, key):
        """
        Inserts a new key <Key> with value 1. Or increments an existing key by 1.
        :type key: str
        :rtype: None
        """
        if key not in self.kv:
            self.kv[key] = 1
            freq = self.order[1]
            if not freq:
                freq = AllOne.FreqNode(1)
                self.freq.add(freq, self.freq.head, self.freq.head.after)
                self.order[1] = freq
            freq.keys.append(key)
            return

        v = self.kv[key]
        freq = self.order[v]
        self.kv[key] += 1
        after_freq = self.order[v+1]
        if not after_freq:
            after_freq = AllOne.FreqNode(v+1)
            self.freq.add(after_freq, freq, freq.after)
            self.order[v+1] = after_freq
        after_freq.keys.append(key)

        freq.keys.remove(key)
        if not freq.keys:
            self.freq.remove(freq)
            self.order[freq.freq] = None

    def dec(self, key):
        """
        Decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
        :type key: str
        :rtype: None
        """
        if key not in self.kv:
            return

        v = self.kv[key]
        if v == 1:
            del self.kv[key]
        else:
            self.kv[key] -= 1

        freq = self.order[v]
        if v > 1:
            prev_freq = self.order[v-1]
            if not prev_freq:
                prev_freq = AllOne.FreqNode(v-1)
                self.freq.add(prev_freq, freq.prev, freq)
                self.order[v-1] = prev_freq
            prev_freq.keys.append(key)
        freq.keys.remove(key)
        if not freq.keys:
            self.freq.remove(freq)
            self.order[freq.freq] = None

    def getMaxKey(self):
        """
        Returns one of the keys with maximal value.
        :rtype: str
        """
        return self.freq.get_max()

    def getMinKey(self):
        """
        Returns one of the keys with Minimal value.
        :rtype: str
        """
        return self.freq.get_min()


# Your AllOne object will be instantiated and called as such:
# obj = AllOne()
# obj.inc(key)
# obj.dec(key)
# param_3 = obj.getMaxKey()
# param_4 = obj.getMinKey()


def main():
    one = AllOne()
    one.inc('hello')
    one.inc('hello')
    one.inc('world')
    one.inc('world')
    one.inc('hello')
    one.dec('world')
    print(one.getMaxKey())
    print(one.getMinKey())

    one.inc('world')
    one.inc('world')
    one.inc('leet')
    print(one.getMaxKey())
    print(one.getMinKey())


if __name__ == '__main__':
    main()
