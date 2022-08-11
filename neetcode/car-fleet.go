/*
- Problem: https://leetcode.com/problems/car-fleet/
*/

func carFleet(target int, position []int, speed []int) int {
    
    var psPair [][]int = [][]int{}
    var currentFleetTop []int = []int{-1,0}
    var totalFleet int = 0
    
    // Creating an Array with pair of Postion and Speed
    for i, p := range position {
        psPair = append(psPair, []int{p, speed[i]})
    }
    
    // Sorting the (Postion, Speed) Pair Array
    sort.Slice(psPair, func(i, j int) bool {
        return psPair[i][0] < psPair[j][0]
    })
    
    // Initializing the current top of the fleet with the car in the nearest position to the target
    currentFleetTop = psPair[len(psPair) - 1]
    totalFleet++  
  
    for i := len(psPair) - 2; i >= 0; i-- {
        // Checking if the current indexed car will collide with the top car of the current fleet by comparing the time each one will take to reach the destination. 
        // If collides, it will join the fleet, Otherwise it will create a new fleet.
        if (float64((target - psPair[i][0])) / float64(psPair[i][1])) <= (float64((target - currentFleetTop[0])) / float64(currentFleetTop[1])) {
            continue
        } else {
            currentFleetTop = psPair[i]
            totalFleet++
        }
    }
    
    return totalFleet
    
}
