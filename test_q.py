import unittest

from q import Q, quantize

class TestQuantize(unittest.TestCase):
    def testFullRange(self):
        steps = range(-5, 6, 1)
        steps = list(steps)

        fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

        expected = [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5]
        actual = Q(steps, fs)
        self.assertEqual(expected, actual)

    def testSingleValue(self):
        steps = range(-5, 6, 1)
        steps = list(steps)


        expected = -5
        actual = quantize(steps, -1)
        self.assertEqual(expected, actual)

if __name__ == '__main__':
    unittest.main()