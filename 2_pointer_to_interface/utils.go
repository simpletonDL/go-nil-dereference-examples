package main

type ValidationError struct {
	Reason string
}

func (v *ValidationError) Error() string {
	return v.Reason
}
