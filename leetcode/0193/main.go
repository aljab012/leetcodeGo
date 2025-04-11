package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(dp); i++ {
		if dp[i] {
			for _, word := range wordDict {
				if i+len(word) <= len(s) && word == s[i:i+len(word)] {
					dp[i+len(word)] = true
				}
			}
		}
	}
	return dp[len(s)]
}
