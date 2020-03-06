// https://leetcode.com/problems/arranging-coins/submissions/
// arranging coins
class Solution {
public:
    int arrangeCoins(int n) {
        if (n > 0 && n <= 2) {
            return 1;
        }
        int loop = n / 2 + 2;
        for (int i = 1; i <= loop; ++i) {  
            if (n >= i) {
                n = n - i;
            } else {
                return i - 1;
            }
        }
        return 0;
    }
};
