// class Solution {
//     public int eliminateMaximum(int[] dist, int[] speed) {
//         int mE = 0;
//         for(int i = 0;i<dist.length;i++){

//             for(int j=0;j<speed.length;j++){
//             dist[j]=dist[j]/speed[j];
//             }
//             mE++;

//             int killCount=0;
//             for(int j=0;j<speed.length;j++){
//                 if (killCount>1) return mE;
//                 if(dist[j]<1) killCount++;
//             }
//         }
//         return mE;
//     }
// }
public class Solution {
    public int eliminateMaximum(int[] dist, int[] speed) {
        double[] t = new double[dist.length];
        
        for (int i = 0; i < dist.length; i++) {
            t[i] = (double)dist[i] / speed[i];
        }
        
        Arrays.sort(t);
        
        for (int i = 0; i < dist.length; i++) {
            if (t[i] <= i) {
                return i;
            }
        }
        
        return dist.length;
    }
}