import java.util.Stack;

class Solution {
    public boolean isValid(String parens) {
      Stack<Character> s = new Stack<Character>();
      for(char c: parens.toCharArray()){
          if(c=='('||c=='{'||c=='['){
              s.push(c);
          }
          else if(c==')'){
              try{
                  char t = s.pop();
                  if(t!='('){
                      return false;
                  }
              }catch(Exception e){
                  return false;
              }
          }
          else if(c=='}'){
              try{
                  char t = s.pop();
                  if(t!='{'){
                      return false;
                  }
              }catch(Exception e){
                  return false;
              }
          }
          else if(c==']'){
              try{
                  char t = s.pop();
                  if(t!='['){
                      return false;
                  }
              }catch(Exception e){
                  return false;
              }
          }
      }
      if(s.empty()){
          return true;
      }else{
          return false;
      }
    }
}
