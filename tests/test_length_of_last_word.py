from tests import base
from solution import length_of_last_word


class SolutionTestCase(base.TestCase):
    SolutionClass = length_of_last_word.Solution
    test_method = 'lengthOfLastWord'
    data = (
        (("Hello World",), 5),
        ((" Hello World ",), 5),
    )
