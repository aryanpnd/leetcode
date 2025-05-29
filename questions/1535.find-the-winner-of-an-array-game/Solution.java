class Solution {
    public int getWinner(int[] arr, int k) {
        if(k==1) return Math.max(arr[0],arr[1]);
        if(k>=arr.length) return Arrays.stream(arr).max().getAsInt();

        int w = arr[0];
        int wC = 0;

        for(int i=1;i<arr.length;i++){
            if(w>arr[i]) wC++;
            else{
                w = arr[i];
                wC=1;
            }

            if(wC==k) return w;
        }
        return w;
    }
}