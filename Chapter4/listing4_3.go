func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied") // #A
	}
	return strings.Join(parts, " "), nil
}
