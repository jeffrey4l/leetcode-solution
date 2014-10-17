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
            self.assertEqual(expected, method(*param))
