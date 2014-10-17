#-*- coding:utf-8 -*-


import unittest


class TestCase(unittest.TestCase):

    SolutionClass = None
    test_method = None
    data = ()

    def test_data(self):
        solution = self.SolutionClass()
        for param, expected in self.data:
            method = getattr(solution, self.test_method)
            ret = method(*param)
            self.assertEqual(expected, ret,
                             "expectd<%s> != actual<%s> for data %s" % (expected, ret, param))
