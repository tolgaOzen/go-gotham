package helpers

import (
	"fmt"
	"regexp"
)

func ClearNonAlphanumericalCharacters(val string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(val, ""), nil
}

func ByteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
