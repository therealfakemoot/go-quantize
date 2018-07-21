import unittest

from q import Q

class TestSimpleDomain(unittest.TestCase):
    def testFullRange(self):
        steps = range(-5, 6, 1)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

class TestBalancedDomains(unittest.TestCase):
    def testOddStepCount(self):
        steps = range(-10, 21, 3)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-10, -7, -4, -1, 2, 5, 8, 11, 14, 17, 20]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

    def testEvenStepCount(self):
        steps = range(-10, 21, 3)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-10, -7, -4, -1, 2, 5, 8, 11, 14, 17, 20]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))


class TestUnbalancedDomains(unittest.TestCase):
    def testPositiveHeavyDomain(self):
        steps = list(range(-10, 21, 4))

        fs = [-1, -.75, -.5, -.25, .25, .5, .75, 1]

        expected = [-10, -6, -2, 2, 6, 10, 14, 18]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

    def testNegativeHeavyDomain(self):
        steps = list(range(-20, 11, 4))

        fs = [-1, -.75, -.5, -.25, .25, .5, .75, 1]

        expected = [-20, -16, -12, -8, -4, 0, 4, 8]
        actual = Q(steps, fs)
        self.assertEqual(expected, list(actual))

if __name__ == '__main__':
    unittest.main()
