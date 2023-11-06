// this one will not work because it is using regular queue, to make it work we have to sort the queue and place the elements which is adding in the queue while calling "reserve" method, so basically we have to modify for reserve method which could be bit costly if we are using brute force to sort it and place again in the queue, 

// to make it work , we will be using priority queue here instead of regular queue.

// class SeatManager {
//     Queue<Integer> queue = new LinkedList<>();

//     public SeatManager(int n) {
//         for (int i = 1; i <= n; i++) {
//             queue.offer(i);
//         }
//     }
    
//     public int reserve() {
//         return queue.poll();
//     }
    
//     public void unreserve(int seatNumber) {
//          queue.add(seatNumber);
//     }
// }

import java.util.PriorityQueue;

class SeatManager {
   PriorityQueue<Integer> queue = new PriorityQueue<>();
    public SeatManager(int n) {
        for (int i = 1; i <= n; i++) {
            queue.offer(i);
        }
    }
    
    public int reserve() {
        return queue.poll();
    }
    
    public void unreserve(int seatNumber) {
         queue.add(seatNumber);
    }
}

// /**
//  * Your SeatManager object will be instantiated and called as such:
//  * SeatManager obj = new SeatManager(n);
//  * int param_1 = obj.reserve();
//  * obj.unreserve(seatNumber);
//  */