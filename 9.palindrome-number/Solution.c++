class Solution {
public:
    bool isPalindrome(int x) {
        string str = to_string(x);
        int strSize = str.size();
        for(int i = 0; i<strSize;i++)
        {
            if(str[i]!=str[strSize-1-i]) return false;
        }
        return true;
    }
};
// Time complexcity will O(n) for the above code...


// --------------------------------------------------------------
// Not using because Time complexcity will be O(logn), We divided the input by 10 for every iteration, so the time complexity is O(\log_{10}(n))

// class Solution {
// public:
//     bool isPalindrome(int x) {
//         if(x < 0 || (x != 0 && x%10 == 0)) return false;
        
//         int reverted = 0;
//         while(x > reverted){
//             reverted = reverted*10 + x%10;
//             x/=10;
//         }
        
//         return x==reverted || x==reverted/10;
//     }
// };