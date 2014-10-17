#!/usr/bin/python
#-*- coding:utf-8 -*-

class Solution:
    # @param s, a string
    # @return an integer
    def lengthOfLastWord(self, s):
        return len(s.strip().rsplit(' ')[-1])
