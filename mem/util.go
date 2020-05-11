package mem

func ConvertLocalToOutScope(local Address) Address {
	base := local - localstart
	return base + scopestart
}
