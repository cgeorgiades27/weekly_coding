

class Solution:
    table = {}

    def build(self, words):
        for i in words:
            if i[0] not in self.table:
                self.table[i[0]] = []
            self.table[i[0]].append(i)
        for i in self.table:
            self.table[i].sort()
        # print(self.table)

    def autocomplete(self, word):
        # Fill this in.
        print(self.table[word[0]])
        return self.table[word[0]]


s = Solution()
s.build(['dog', 'dark', 'cat', 'door', 'dodge'])
print(s.autocomplete('do'))
# ['dog', 'door', 'dodge']
