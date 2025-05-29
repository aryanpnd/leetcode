class Solution {
    public int[][] transpose(int[][] m) {
        int r = m.length;
        int c = m[0].length;
        int[][] a = new int[c][r];

        for(int i = 0;i<c;i++){
            for(int j = 0;j<r;j++){
            a[i][j]=m[j][i];
        }
        }
        return a;
    }
}