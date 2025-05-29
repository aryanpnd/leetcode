public class Solution {
    public boolean repeatedSubstringPattern(String s) {
        int a = s.length();
        String sFold = s.substring(1, a) + s.substring(0, a-1);
        return sFold.contains(s);
    }
}
