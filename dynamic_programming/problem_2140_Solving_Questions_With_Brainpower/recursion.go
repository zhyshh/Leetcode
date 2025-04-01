func max(a int64, b int64) int64 {
    if a>b {
        return a
    }
    return b
}

func dfs(questions [][]int, dp []int64, i int) int64 {
    n := len(dp)
    if i>=n {
        return 0
    }
    if dp[i] != -1 {
        return dp[i]
    }

    pts := int64(questions[i][0])
    next := questions[i][1]
    take := pts + dfs(questions, dp, i+next+1)
    skip := dfs(questions, dp, i+1)
    
    dp[i] = max(take,skip)
    return dp[i]
}

func mostPoints(questions [][]int) int64 {
    n := len(questions)
    dp := make([]int64, n)
    for i := range dp {
        dp[i] = -1
    } 

    return dfs(questions, dp, 0)
}


