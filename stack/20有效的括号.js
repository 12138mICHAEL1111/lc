/**
 * @param {string} s
 * @return {boolean}
 */

 var isValid = function(s) {
    const map = new Map([
        [']','['],
        ['}','{'],
        [')','(']
    ])
    
    if(s.length%2 ===1){
        return false
    }
    
    const stack = []
    for(let char of s){
        //当前为右括号 比较当前符号是否和stack末尾匹配
        if(map.has(char)){
            if(!stack.length || stack[stack.length-1] !== map.get(char)){
                return false
            }
            stack.pop()
        }
        //当前左括号 放入stack
        else{
            stack.push(char)
        }
    }

    return !stack.length
};
isValid('()')