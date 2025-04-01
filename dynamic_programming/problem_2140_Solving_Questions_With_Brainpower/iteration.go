func mostPoints(questions [][]int) int64 {
    n := len(questions)
    maxPoints := make([]int64, n)
    
    maxPoints[n-1] = int64(questions[n-1][0])

    for i := n-2; i >= 0; i-- {
        points, brainpower := int64(questions[i][0]), questions[i][1]
        if i + brainpower + 1 > n-1 {
            maxPoints[i] = max(maxPoints[i+1], points)
        } else {
            maxPoints[i] = max(maxPoints[i+1], points + maxPoints[i+brainpower+1])
        }
        // fmt.Println(maxPoints)
    }
    return maxPoints[0]
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
