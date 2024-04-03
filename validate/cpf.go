package validate

import (
	"fmt"
	"regexp"
	"strconv"
)

// ValidateCPFs valida uma lista de CPFs
func ValidateCPFs(cpfs []string) map[string]bool {
	validityMap := make(map[string]bool)
	for _, cpf := range cpfs {
		cpf = AddCPFMascara(cpf)
		validityMap[cpf] = ValidateCPF(cpf)
	}
	return validityMap
}

// RemoveCaracteresEspeciais remove todos os caracteres não numéricos de uma string.
func RemoveCaracteresEspeciais(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}

// ValidateCPF valida um número de CPF, pode ser formatado ou não formatado.
func ValidateCPF(cpf string) bool {
	cpf = RemoveCaracteresEspeciais(cpf)
	if len(cpf) != 11 || isSequenciaNumerica(cpf) {
		return false
	}

	numeros := StringParaDigitos(cpf)
	digito1, digito2 := CalcularDigitoVerificadorCPF(numeros)

	return digito1 == strconv.Itoa(numeros[9]) && digito2 == strconv.Itoa(numeros[10])
}

// isSequenciaNumerica verifica se o número de CPF é uma sequência numérica e retorna true se for.
func isSequenciaNumerica(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// StringParaDigitos converte uma string em uma fatia de dígitos como inteiros.
func StringParaDigitos(cpf string) []int {
	var digitos []int
	for _, char := range cpf {
		if digito, err := strconv.Atoi(string(char)); err == nil {
			digitos = append(digitos, digito)
		}
	}
	return digitos
}

// CalcularDigitoVerificadorCPF calcula os dígitos de verificação do CPF (digito1 e digito2)
func CalcularDigitoVerificadorCPF(numeros []int) (string, string) {
	primeirosFatores := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	segundosFatores := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	digito1 := CalcularDigito(numeros[0:9], primeirosFatores)
	digito2 := CalcularDigito(numeros[0:10], segundosFatores)

	return digito1, digito2
}

// CalcularDigito calcula um dígito de verificação.
func CalcularDigito(numeros []int, fatores []int) string {
	var total int
	for i, digito := range numeros {
		total += digito * fatores[i]
	}
	resto := total % 11
	if resto < 2 {
		return "0"
	}
	return strconv.Itoa(11 - resto)
}

// AddCPFMascara adiciona uma máscara a uma string de CPF (xxx.xxx.xxx-xx)
func AddCPFMascara(cpf string) string {
	if HasCPFMascara(cpf) {
		return cpf
	}

	if len(cpf) != 11 {
		return cpf
	}

	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

// HasCPFMascara verifica se uma string já tem a máscara de CPF aplicada.
func HasCPFMascara(cpf string) bool {
	mascaraPadrao := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return mascaraPadrao.MatchString(cpf)
}
