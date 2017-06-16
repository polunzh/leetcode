class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        v = 0
        if x >= 0:
            v = int(str(x)[::-1])
        else:
            v = -int(str(-x)[::-1])

        if v < -2147483648 or v > 2147483647:
            return 0
        else:
            return v
