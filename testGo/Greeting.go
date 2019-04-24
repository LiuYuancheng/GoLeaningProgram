package testGo

func translate(s string ) string{
	switch s{
	case "en-US":
		return "Hello"
	case "fr-Fr":
		return "Bonjour"
	case "it-IT":
		return "Ciao"
	default:
		return "Hello"
	}
}

func Greeting0(name, locale string) string{
	salutation := translate(locale)
	return (salutation+name)
}
func Greeting( s string) string{
	return ("Hello " +s )
}
