// https://leetcode.com/problems/number-of-segments-in-a-string/
class Solution {
public:
    int countSegments(string s) {
        string t;
        int cnt = 0;
        stringstream input(s);
        while (input >> t) {
            cnt++;
        }
        // return input.tellg();;
        return cnt;
    }
};
