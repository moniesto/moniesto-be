package core

import (
	"math"
	"time"

	"github.com/moniesto/moniesto-be/model"
)

// CalculateApproxScore calculates the approximately score of the post
func CalculateApproxScore(duration time.Duration, startPrice float64, endPrice float64) float64 {
	return CalculatePredictionScore(startPrice, endPrice, duration)
}

// CalculateScore calculates the exact score of the post (runs after post time run out)
func CalculateScore(post model.CreatePostResponse) float64 {
	var minValueInPostDuration, maxValueInPostDuration float64

	if time.Now().After(post.Duration) { // Is post duration exceeded
		_ = minValueInPostDuration
		_ = maxValueInPostDuration
	}

	return 0.0
}

// Create post responseda duration data type ? hour day minutes?
func CalculateChangeInPercentage(start float64, end float64) float64 {
	return ((end - start) / start) * 100
}
func CalculateTimeBonus(duration time.Duration) float64 {
	durationInDays := float64(duration / (24 * time.Hour))
	if durationInDays <= 14 {
		return (float64(3)/float64(338))*math.Pow((durationInDays-14), 2) + 1
	}
	if durationInDays < 45 {
		return (float64(9)/float64(9610))*math.Pow((durationInDays-45), 2) + 0.1
	}
	return 0.1
}
func CalculateBaseScore(change float64) float64 {
	if change < 27.667 {
		return 0.5*math.Log2(change) + math.Sqrt(change)
	}
	return 0.01 * math.Pow(change, 2)
}
func CalculatePredictionScore(startPrice, endPrice float64, duration time.Duration) float64 {
	change := CalculateChangeInPercentage(startPrice, endPrice)
	baseScore := CalculateBaseScore(change)
	timeBonus := CalculateTimeBonus(duration)
	return baseScore * timeBonus
}
