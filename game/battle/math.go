package battle

const (
	I_10_000 = 10000
	F_10_000 = 10000.0
)

// NumberFormatPermyriadInt converts int value to permyriad format (万份比)
func NumberFormatPermyriadInt(value int) float64 {
	return float64(value*I_10_000) / F_10_000
}

// NumberFormatPermyriadFloat converts float value to permyriad format
func NumberFormatPermyriadFloat(value float64) float64 {
	return value / F_10_000
}

// NumberFormatThousands converts value to thousands format
func NumberFormatThousands(value float64) float64 {
	return value * 0.0001
}

// NumberFormatPermyriadWithPercentage converts int value with percentage to permyriad format
func NumberFormatPermyriadWithPercentage(value int, percentage int) float64 {
	return float64(value*(I_10_000+percentage)) / F_10_000
}