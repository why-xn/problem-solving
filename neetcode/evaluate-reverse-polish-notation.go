/*
- Problem: https://leetcode.com/problems/evaluate-reverse-polish-notation/
*/

func evalRPN(tokens []string) int {
    
    var stack []string = []string{}
    var n1, n2, tmp, res int = 0, 0, 0, 0
    
    var arithmaticExpression map[string]*byte = map[string]*byte {
        "+": nil,
        "-": nil,
        "/": nil,
        "*": nil,
    }
    
    for _, t := range tokens {
        if _, ok := arithmaticExpression[t]; ok {
            // check if the string is an arithmatic expression. If so, then do arithmatic operation on the top two items in the stack
            n1, _ = strconv.Atoi(stack[len(stack) - 1])
            n2, _ = strconv.Atoi(stack[len(stack) - 2])

            if t == "+" {
                tmp = n2 + n1
            } else if t == "-" {
                tmp = n2 - n1
            } else if t == "/" {
                tmp = n2 / n1
            } else if t == "*" {
                tmp = n2 * n1
            }
            
            // Pop the top 2 items from the stack and push the arithmatic operation result
            stack = stack[:len(stack) - 2]
            stack = append(stack, strconv.Itoa(tmp))
            
        } else {
            // if the string is not an arithmatic expression, just push the item in the stack
            stack = append(stack, t)
        }
    }

    // after all the operations, only one item should remain in the stack.
    res, _ = strconv.Atoi(stack[0])
    return res
    
}
