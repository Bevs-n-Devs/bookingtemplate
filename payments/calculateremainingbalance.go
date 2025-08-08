package payments

// calulate remaining balance of a service
//
// returns remaining amount as int
func CalculateRemainingBalance(deposit, cost int) int {
	result := deposit - cost
	return result
}
