package validate

import (
	"fmt"
	"regexp"
	"strconv"
)

// ValidateCPFs validates a list of CPFs
func ValidateCPFs(cpfs []string) map[string]bool {
	validityMap := make(map[string]bool)
	for _, cpf := range cpfs {
		cpf = AddCPFMask(cpf)
		validityMap[cpf] = ValidateCPF(cpf)
	}
	return validityMap
}

// ValidateCPF validates a CPF number, it can be formatted or unformatted.
func ValidateCPF(cpf string) bool {
	cpf = RemoveSpecialCharacters(cpf)
	if len(cpf) != 11 || isNumericSequence(cpf) {
		return false
	}

	numbers := StringToDigits(cpf)
	digit1, digit2 := CalculateCPFCheckDigit(numbers)

	return digit1 == strconv.Itoa(numbers[9]) && digit2 == strconv.Itoa(numbers[10])
}

// isNumericSequence checks if the CPF number is a numeric sequence and returns true if it is.
func isNumericSequence(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// CalculateCPFCheckDigit calculates the verification check digits (digit1 and digit2) for CPF.
func CalculateCPFCheckDigit(numbers []int) (string, string) {
	firstFactors := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	secondFactors := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	digit1 := CalculateCheckDigit(numbers[0:9], firstFactors)
	digit2 := CalculateCheckDigit(numbers[0:10], secondFactors)

	return digit1, digit2
}

// CalculateCheckDigit calculates a verification digit.
func CalculateCheckDigit(numbers []int, factors []int) string {
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

// AddCPFMask adds a mask to a CPF string (xxx.xxx.xxx-xx).
func AddCPFMask(cpf string) string {
	if HasCPFMask(cpf) {
		return cpf
	}

	if len(cpf) != 11 {
		return cpf
	}

	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

// HasCPFMask checks if a string already has the CPF mask applied.
func HasCPFMask(cpf string) bool {
	maskPattern := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return maskPattern.MatchString(cpf)
}
