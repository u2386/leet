class Solution:
    def validWordAbbreviation(self, word: str, abbr: str) -> bool:

        sl = []
        i = 0
        while i < len(abbr):
            c = abbr[i]
            if not c.isdigit():
                sl.append(c)
                i += 1
                continue

            j = i+1
            while j < len(abbr):
                d = abbr[j]
                if not d.isdigit():
                    break
                j += 1
            sl.append(abbr[i:j])
            i = j

        for e in sl:
            if not e.isdigit():
                continue

            if e == '0':
                return False
            if e != str(int(e)):
                return False

        i = 0
        for e in sl:
            if i >= len(word):
                return False

            if not e.isdigit():
                if not word[i] == e:
                    return False
                i += 1
                continue

            i += int(e)
            if i > len(word):
                return False
        return i == len(word)



s = Solution()
print(s.validWordAbbreviation("hi", "1i"))
