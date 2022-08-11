/*
- Problem: https://leetcode.com/problems/car-fleet/
*/

func carFleet(target int, position []int, speed []int) int {
    
    var psPair [][]int = [][]int{}
    var currentFleetTop []int = []int{-1,0}
    var totalFleet int = 0
    
    for i, p := range position {
        psPair = append(psPair, []int{p, speed[i]})
    }
    
    sort.Slice(psPair, func(i, j int) bool {
        return psPair[i][0] < psPair[j][0]
    })
    
    currentFleetTop = psPair[len(psPair) - 1]
    totalFleet++
    
    for i := len(psPair) - 2; i >= 0; i-- {
        if (float64((target - psPair[i][0])) / float64(psPair[i][1])) <= (float64((target - currentFleetTop[0])) / float64(currentFleetTop[1])) {
            continue
        } else {
            currentFleetTop = psPair[i]
            totalFleet++
        }
    }
    
    return totalFleet
    
}
