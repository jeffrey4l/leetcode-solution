#-*- coding:utf-8 -*-

from tests import base
from solution import decode_ways


class SolutionTestCase(base.TestCase):
    SolutionClass = decode_ways.Solution
    test_method = 'numDecodings'
    data = (
        (('',), 0),
        (('12',), 2),
        (('123',), 3),
        (('0',), 0),
        (('4757562545844617494555774581341211511296816786586787755257741178599337186486723247528324612117156948',), 589824),
    )
