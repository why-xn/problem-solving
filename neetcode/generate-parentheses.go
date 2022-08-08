/*
- https://leetcode.com/problems/generate-parentheses/
*/


func generateParenthesis(n int) []string {
    
    var combination string
    var res []string
    
    // Starting a combination with an open parenthesis
    pushParenthesis(combination, &res, "(", 0, 0, n)
    
    return res
}

func pushParenthesis(combination string, res *[]string, input string, openCount int, closeCount int, n int) {
    // Adding the parenthesis to the combination stack
    combination = combination + input
    if input == "(" {
        openCount++
    } else {
        closeCount++
    }
    
    if len(combination) == (n * 2) {
        // If the length of combination matches with expected length then add the combination to result array
        *res = append(*res, combination)
    } else {
        // Run recursion for with all possible next paranthesis
        if openCount >= closeCount && openCount < n {
            pushParenthesis(combination, res, "(", openCount, closeCount, n)
        }

        if closeCount < openCount && closeCount < n {
            pushParenthesis(combination, res, ")", openCount, closeCount, n)
        }
    }
}
