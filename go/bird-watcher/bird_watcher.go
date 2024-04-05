package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekNumber := 0
	targetWeekCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
		// modulo to have i go from 0 to 6 (7 week days) and back
		if i%7 == 0 {
			weekNumber++
		}
		if week == weekNumber {
			targetWeekCount += birdsPerDay[i]
		}
	}
	return targetWeekCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
