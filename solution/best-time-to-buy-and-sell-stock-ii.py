#!/usr/bin/env python
#-*- coding:utf-8 -*-

# problem: https://oj.leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

import unittest

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

class SolutionTestCase(unittest.TestCase):

    def test_pass(self):
        data = (
            ([], 0),
            ([1, 2], 1),
            ([1, 2, 4], 3),
            ([1, 2, 4, 1], 3),
        )

        for prices, profit in data:
            calculated = Solution().maxProfit(prices)
            self.assertEqual(profit, calculated, "%s != %s for input: % s" % (profit, calculated, prices))
        

if __name__ == '__main__':
    unittest.main()
