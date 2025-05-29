class Solution {
    public double findMedianSortedArrays(int[] n1, int[] n2) {
        int nl  = n1.length;
        int n2l  = n2.length;
        int[] tA = new int[nl + n2l];
        int tl  = tA.length;

        int i = 0,j=0;
        while (i < nl)  {
            tA[i]=n1[i];
            i++;
        }
        while (j < n2l)  {
            tA[i++]=n2[j];
            j++;
        }
        Arrays.sort(tA);
        for (int j2 = 0; j2 < tl; j2++) {
            System.out.print(tA[j2]);
        }
        System.out.println();

        if(tl%2==0){
            int r = tA[tl/2];
            int l = tA[(tl/2)-1];
            return (double)(r+l)/2;
        }else{
            return (double)tA[(tl/2)];
        }
    }
}