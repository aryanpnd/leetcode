class Solution {
    public int reverse(int x) {
        long tempNum = 0;
        while(x!=0){
            int lastNum = x%10;
            tempNum = tempNum + lastNum;
            tempNum = tempNum * 10;
            x = x/10;
        }
        tempNum = tempNum/10;
        if(tempNum>Integer.MAX_VALUE || tempNum<Integer.MIN_VALUE){
            return 0;
        }

        if(x<0){
            return (int)(tempNum*-1);
        }else{
            return (int)(tempNum);
        }
    }
}