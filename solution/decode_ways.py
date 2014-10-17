#-*- coding:utf-8 -*-

# Problem: https://oj.leetcode.com/problems/decode-ways/

class Solution:
    # @param s, a string
    # @return an integer
    cache = {}
    def numDecodings(self, s):
        if not s or s == '0':
            return 0
        if len(s) == 1:
            return 1
        elif len(s) == 2:
            return 1 if int(s) > 26 or int(s) < 11 else 2

        if s in self.cache:
            return self.cache[s]
        else:
            head = int(s[:2])
            if head > 26:
                ret = self.numDecodings(s[1:])
            elif head > 10:
                ret = self.numDecodings(s[1:]) + self.numDecodings(s[2:])
            elif head >= 0:
                ret = self.numDecodings(s[2:])
            self.cache[s] = ret
            return ret
