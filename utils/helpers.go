package utils

// Must function is used to parse if there is an error being thrown or not.
func Must(err error) {
    if err != nil {
        panic(err)
    }
}
