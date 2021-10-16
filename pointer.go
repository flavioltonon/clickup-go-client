package clickup

func dereferenceString(p *string) string {
	if p == nil {
		return ""
	}

	return *p
}

func dereferenceInt(p *int) int {
	if p == nil {
		return 0
	}

	return *p
}

func dereferenceBool(p *bool) bool {
	if p == nil {
		return false
	}

	return *p
}
