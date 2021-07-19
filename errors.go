package main

const (
	// InvalidStartingNumberError is an error that is thrown when the input to the fibonacci function is invalid
	InvalidStartingNumberError = RepoNameHereError("invalid starting number specified")
)

// RepoNameHereError is a custom error type
type RepoNameHereError string

// Error is a function that allows RepoNameHereError to implement the Error interface
func (e RepoNameHereError) Error() string {
	return string(e)
}
