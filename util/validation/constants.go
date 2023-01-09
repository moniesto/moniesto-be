package validation

// TODO: fill it according to links in url
var InvalidUsernames = []string{
	"",
}

var ValidPasswordLength = 6
var UsernameRegex = `^[A-Za-z][A-Za-z0-9_]{0,29}$`
var EmailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
