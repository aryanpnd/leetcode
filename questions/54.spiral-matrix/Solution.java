import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> spiralOrder(int[][] m) {
        
        int r = 0;
        int R = m.length;
        int c = 0;
        int C = m[0].length;
        List<Integer> a = new ArrayList<Integer>();
        
        while(R>r && C>c){

            for(int i = c; i<C;i++){
                a.add(m[r][i]);
            }
            r++;

            for(int i = r;i<R;i++){
                a.add(m[i][C-1]);
            }
            C--;

            if(r<R){
            for(int i = C-1;i>=c;i--){
                a.add(m[R-1][i]);
            }
            R--;}

            if(c<C){
            for(int i = R-1;i>=r;i--){
                a.add(m[i][c]);
            }
            c++;}
        }
    return a;
    }
}