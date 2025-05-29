/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const arrOdd = []
    const arrEven = []
    for (let index = 0; index < s.length; index++) {
        if(index%2===0){
            arrOdd.push(s[index])
        }else{
            arrEven.push(s[index])
        }
    }
    console.log(arrOdd)
    console.log(arrEven)
    // if(arr.length===s.length) return true;
    // else return false
};

console.log(isValid("()"))