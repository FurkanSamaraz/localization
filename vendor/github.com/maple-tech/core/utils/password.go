package utils

//ValidatePassword takes a possible password and checks against validity rules, returns TRUE if ok
//Current check is length
func ValidatePassword(password string) bool {
	return len(password) >= 6
}
