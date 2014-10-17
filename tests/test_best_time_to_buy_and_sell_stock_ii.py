#-*- coding:utf-8 -*-

from tests import base
from solution import best_time_to_buy_and_sell_stock_ii

class SolutionTestCase(base.TestCase):
       SolutionClass = best_time_to_buy_and_sell_stock_ii.Solution 
       test_method = 'maxProfit'
       data = (
           (([],), 0),
           (([1, 2],), 1),
           (([1, 2, 4],), 3),
           (([1, 2, 4, 1],), 3),
       )
