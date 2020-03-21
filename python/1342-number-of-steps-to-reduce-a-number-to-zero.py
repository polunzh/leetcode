class Solution:
    def numberOfSteps (self, num):
      steps = 0
      while num > 0:
        steps += 1
        if num % 2 != 0:
          num -= 1
          continue

        if num == 0:
          break
        
        num = int(num / 2)
      
      return steps

print(Solution().numberOfSteps(14))