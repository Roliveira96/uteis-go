package cnpj

import (
	"fmt"
	"regexp"
	"strconv"
)

// ValidateCNPJs validates a list of CNPJs
func ValidateCNPJs(cnpjs []string) map[string]bool {
	validityMap := make(map[string]bool)
	for _, cnpj := range cnpjs {
		validityMap[AddCNpjMask(cnpj)] = ValidateCNPJ(cnpj)
	}
	return validityMap
}

// RemoveSpecialCharacters removes all non-digits from a string.
func RemoveSpecialCharacters(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}

// ValidateCNPJ validates a CNPJ number, it can be formatted or unformatted
func ValidateCNPJ(cnpj string) bool {
	cnpj = RemoveSpecialCharacters(cnpj)
	if len(cnpj) != 14 || isNumberSequence(cnpj) {
		return false
	}

	numbers := StringToDigits(cnpj)
	digit1, digit2 := CalcCheckDigit(numbers)

	return digit1 == strconv.Itoa(numbers[12]) && digit2 == strconv.Itoa(numbers[13])
}

// isNumberSequence checks if the CNPJ number is from a numbers sequence, and returns true if it is.
func isNumberSequence(cnpj string) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			return false
		}
	}
	return true
}

// StringToDigits converts a string to a slice of digits as integers
func StringToDigits(cnpj string) []int {
	var digits []int
	for _, char := range cnpj {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		}
	}
	return digits
}

// CalcCheckDigit calculates the verification check digits (digit1 and digit2)
func CalcCheckDigit(numbers []int) (string, string) {
	firstFactors := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	secondFactors := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	digit1 := CalcDigit(numbers[0:12], firstFactors)
	digit2 := CalcDigit(numbers[0:13], secondFactors)

	return digit1, digit2
}

// CalcDigit calculates a verification digit.
func CalcDigit(numbers []int, factors []int) string {
	var total int
	for i, digit := range numbers {
		total += digit * factors[i]
	}
	remainder := total % 11
	if remainder < 2 {
		return "0"
	}
	return strconv.Itoa(11 - remainder)
}

// AddCNpjMask adds a mask to a CNPJ string (xx.xxx.xxx/xxxx-xx)
func AddCNpjMask(cnpj string) string {
	if HasCnpjMask(cnpj) {
		return cnpj
	}

	if len(cnpj) != 14 {
		return cnpj
	}

	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[0:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:14])
}

// HasCnpjMask checks if a string already has the CNPJ mask applied.
func HasCnpjMask(cnpj string) bool {
	maskedPattern := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	return maskedPattern.MatchString(cnpj)
}
