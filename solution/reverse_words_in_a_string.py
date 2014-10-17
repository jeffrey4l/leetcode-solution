#!/usr/bin/python
#-*- coding:utf-8 -*-

class Solution:
    # @param s, a string
    # @return a string
    def reverseWords(self, s):
        return " ".join(s.split()[::-1])
