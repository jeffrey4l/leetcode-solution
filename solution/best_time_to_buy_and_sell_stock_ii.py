#!/usr/bin/env python
#-*- coding:utf-8 -*-

# problem: https://oj.leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

class Solution:
    def maxProfit(self, prices):
        profit = 0
        
        def it():
            data = []
            for price in prices:
                if not data:
                    data = [price]
                elif price >= data[-1]:
                    data.append(price)
                elif price < data[-1]:
                    yield data
                    data = [price]
            if data:
                yield data

        for d in it():
            if len(d) > 1:
                profit += (d[-1] - d[0])

        return profit
