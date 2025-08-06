package payments

func Service1RemainingBalance30Mins() int {
	result := servcie1Deposit30Mins - service1FullAmount30Mins
	return result
}

func Service1RemainingBalance45Mins() int {
	result := servcie1Deposit45Mins - service1FullAmount45Mins
	return result
}

func Service1RemainingBalance60Mins() int {
	result := servcie1Deposit60Mins - service1FullAmount60Mins
	return result
}
