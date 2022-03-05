package must

type SetOfRules map[string][]Rule

type Rule interface {
	ErrorMessage(inputField string, inputValue string) string
	Valid(inputField string, inputValue string) bool
}
