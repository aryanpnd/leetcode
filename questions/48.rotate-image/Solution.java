class Solution {
    public void rotate(int[][] m) {
        int l = m.length;

        for (int i = 0; i < l; i++) {
            int a = 0;
            int b = l - 1;
            while (a <= b) {
                int t = m[a][i];
                m[a][i] = m[b][i];
                m[b][i] = t;
                a++;
                b--;
            }
        }

        for (int j = 0; j < l; j++) {
            for (int j2 = 0; j2 < l; j2++) {
                int t = m[j][j2];
                m[j][j2] = m[j2][j];
                m[j2][j] = t;
            }
        }
    }
}