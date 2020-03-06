// https://leetcode.com/problems/convert-a-number-to-hexadecimal/
class Solution {
public:
    string toHex(int num) {
        if (num == 0) {
            return "0";
        }
        unsigned uNum = num;
        string res;
        // unsignedでは負の値の時は補数になるので、16のmodを取れば良い
        char chars[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
        while (uNum > 0) {
            unsigned digit = uNum % 16;
            res = chars[digit] + res;
            uNum /= 16;
        }
        return res;
    }
};
