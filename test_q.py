import unittest

from q import Q, quantize

class TestBalancedDomain(unittest.TestCase):
    def testFullRange(self):
        steps = range(-5, 6, 1)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

    def testSingleValue(self):
        steps = range(-5, 6, 1)
        steps = list(steps)

        expected = -5
        actual = quantize(steps, -1)
        self.assertEqual(expected, actual)

class TestUnbalancedDomain(unittest.TestCase):
    def testFullRange(self):
        steps = range(-10, 21, 3)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-10, -7, -4, -1, 2, 5, 8, 11, 14, 17, 20]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

    def testSingleValue(self):
        steps = range(-10, 21, 3)
        steps = list(steps)

        expected = -10
        actual = quantize(steps, -1)
        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
