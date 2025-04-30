package service

import (
	"fmt"
	"github.com/croatiangrn/packet_calculator/src/domain/repositories"
	"math"
	"slices"
)

type Calculator struct {
	repo repositories.PackageRepository
}

func NewCalculator(repo repositories.PackageRepository) *Calculator {
	return &Calculator{
		repo: repo,
	}
}

func (c *Calculator) CalculatePacks(order int, packages []int) (map[int]int, error) {
	if order <= 0 {
		return nil, fmt.Errorf("order must be positive")
	}
	if len(packages) == 0 {
		return nil, fmt.Errorf("no pack sizes provided")
	}

	// Create a sorted copy and remove duplicates
	sortedPacks := make([]int, len(packages))
	copy(sortedPacks, packages)
	slices.Sort(sortedPacks)
	sortedPacks = slices.Compact(sortedPacks)

	minPack := sortedPacks[0]

	// Early exit if order is smaller than the smallest pack
	if order < minPack {
		return map[int]int{minPack: 1}, nil
	}

	// Try greedy approach first to get an upper bound
	greedyCount := c.greedySolution(order, sortedPacks)
	upperBound := order
	if greedyCount > 0 {
		// Get tighter bound
		upperBound = min(order+minPack, order*2)
	}

	// Initialize DP table
	dp := make([]dpEntry, upperBound+1)
	for i := range dp {
		dp[i].count = math.MaxInt32
	}
	dp[0].count = 0

	// Fill DP table with early termination
	for i := 1; i <= upperBound; i++ {
		for _, s := range sortedPacks {
			if s > i {
				continue
			}
			if dp[i-s].count+1 < dp[i].count {
				dp[i].count = dp[i-s].count + 1
				dp[i].lastPack = s
			}
		}
		// Early exit if we've found the exact solution
		if i == order && dp[i].count != math.MaxInt32 {
			upperBound = i // No need to search further
			break
		}
	}

	// Find the best solution in [order, upperBound]
	bestQty := -1
	bestCount := math.MaxInt32
	for i := order; i <= upperBound; i++ {
		if dp[i].count < bestCount {
			bestCount = dp[i].count
			bestQty = i
		}
	}

	if bestQty == -1 {
		return nil, fmt.Errorf("no solution found")
	}

	// Backtrack to find packs used
	result := make(map[int]int)
	for current := bestQty; current > 0; current -= dp[current].lastPack {
		result[dp[current].lastPack]++
	}

	return result, nil
}

func (c *Calculator) greedySolution(order int, packs []int) int {
	count := 0
	remaining := order
	for i := len(packs) - 1; i >= 0; i-- {
		if packs[i] <= remaining {
			num := remaining / packs[i]
			count += num
			remaining -= num * packs[i]
			if remaining == 0 {
				break
			}
		}
	}
	if remaining > 0 {
		// Need one more smaller pack
		count++
	}
	return count
}

func (c *Calculator) GetPacksSizes() ([]int, error) {
	return c.repo.GetPacks()
}
