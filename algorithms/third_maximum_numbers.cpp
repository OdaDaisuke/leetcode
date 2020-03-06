// https://leetcode.com/problems/third-maximum-number/
class Solution {
public:
    int thirdMax(vector<int>& nums) {
        sort(nums.begin(), nums.end(), greater<>());
        nums.erase(unique(nums.begin(), nums.end()), nums.end());
        if (nums.size() < 3) {
            return nums[0];
        }
        return nums[2];
    }
};
