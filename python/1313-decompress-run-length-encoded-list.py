class Solution:
    def decompressRLElist(self, nums):
        res = []
        it = iter(nums)
        for x in it:
          t = next(it)
          for _ in range(0, x):
            res.append(t)
        
        return res

c = Solution()
print(c.decompressRLElist([1,1,2,3]))
print(c.decompressRLElist([1, 2, 3, 4]))