// class Solution {
//     public int[] restoreArray(int[][] adjacentPairs) {
//          HashSet<Integer> h = new HashSet<Integer>(); 
//          for(int i =0;i<adjacentPairs.length;i++){
//              for(int j = 0;j<adjacentPairs[i].length;j++){
//                  h.add(adjacentPairs[i][j]);
//              }
//          }
//          int[] arr = new int[h.size()];
//         int index = 0;
//         for (int num : h) {
//             arr[index] = num;
//             index++;
//         }
//          return arr;
//     }
// }

class Solution {
    public int[] restoreArray(int[][] adjacentPairs) {
        Map<Integer, List<Integer>> graph = new HashMap<>();

        for (int[] pair : adjacentPairs) {
            graph.computeIfAbsent(pair[0], k -> new ArrayList<>()).add(pair[1]);
            graph.computeIfAbsent(pair[1], k -> new ArrayList<>()).add(pair[0]);
        }

        List<Integer> result = new ArrayList<>();

        for (Map.Entry<Integer, List<Integer>> entry : graph.entrySet()) {
            if (entry.getValue().size() == 1) {
                result.add(entry.getKey());
                result.add(entry.getValue().get(0));
                break;
            }
        }

        while (result.size() < adjacentPairs.length + 1) {
            int last = result.get(result.size() - 1);
            int prev = result.get(result.size() - 2);
            List<Integer> candidates = graph.get(last);
            int nextElement = candidates.get(0) != prev ? candidates.get(0) : candidates.get(1);
            result.add(nextElement);
        }

        return result.stream().mapToInt(Integer::intValue).toArray();        
    }
}