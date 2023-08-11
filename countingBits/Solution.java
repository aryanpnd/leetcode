// approach 1
// class Solution {
//     public int[] countBits(int n) {
//         int[] a = new int[n+1];
//         for (int i = 0; i <= n; i++) {
//             String temp = Integer.toBinaryString(i);
//             int count = 0;
//             for (int j = 0; j < temp.length(); j++) {
//                 if (temp.charAt(j) == '1') {
//                     count++;
//                 }
//             }
//             a[i] = count;
//         }
//         return a;
//     }
// }

// approach 2
class Solution {
    public int[] countBits(int n) {
        
        int[] arr = new int[n+1];
        
        for(int j =  0; j<=n; j++){
            int count = 0;
            int i = j;
            while( i != 0){
                i = i & (i-1);
                count++;
            }
            arr[j] = count;
        }
        return arr;
    }
}