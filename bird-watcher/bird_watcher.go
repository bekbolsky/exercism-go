package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalNumber := 0
	for i := range birdsPerDay {
		totalNumber += birdsPerDay[i]
	}
	return totalNumber
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInGivenWeek := birdsPerDay[(week-1)*7 : week*7]
	return TotalBirdCount(birdsInGivenWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// for i := 0; i < len(birdsPerDay); i += 2 {
	// 	birdsPerDay[i]++
	// }
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
