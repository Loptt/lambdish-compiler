package mem

func ConvertLocalToOutScope(local Address) Address {
	base := local - Localstart
	return base + Scopestart
}
