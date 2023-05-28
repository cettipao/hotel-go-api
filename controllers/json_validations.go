package controllers

import "strings"

// Función para verificar si un campo está vacío
func IsEmptyField(field string) bool {
	return len(strings.TrimSpace(field)) == 0
}
