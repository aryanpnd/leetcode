class Solution {
    public boolean isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if(sx==fx && sy == fy ) {
            if(t==0)return true;
            if(t<=1) return false;
            return true;
        };
        int dx = Math.abs(fx-sx);
        int dy = Math.abs(fy-sy);
        boolean canReach=false;
        if(Math.max(dx,dy)<=t){
            canReach =true;
        }
        return canReach;
    }
}