package saludo

import (
	"errors"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var tests = []struct {
		testName	   string
		name           string
		expectedResult string
		expectedError  error
	}{
		{
			testName: 		"Test1",
			name:           "Juan",
			expectedResult: "¡Hola, Juan! ¡Bienvenido al banco de Bogota!",
			expectedError:  nil,
		},
		{
			testName: 		"Test2",
			name:           "",
			expectedResult: "",
			expectedError:  errors.New("nombre vacío"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Ejecutando test: ", tt.testName)
			message, err := Hello(tt.name)

			if err != nil {
				if tt.expectedError == nil {
					t.Errorf("Test '%s': Esperaba que no ocurriera un error, pero obtuve: %v", tt.name, err)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("Test '%s': El mensaje de error esperado no coincide. Esperaba '%v', pero obtuve '%v'", tt.name, tt.expectedError.Error(), err.Error())
				}
			} else {
				if message != tt.expectedResult {
					t.Errorf("Test '%s': El mensaje esperado no coincide. Esperaba '%v', pero obtuve '%v'", tt.name, tt.expectedResult, message)
				}
			}
		})
	}
}