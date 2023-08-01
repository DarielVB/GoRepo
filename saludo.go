package saludo

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
    // Si no se proporcionó ningún nombre, devuelve un error con un mensaje.
    if name == "" {
        return "", errors.New("nombre vacío")
    }

    // Si se recibió un nombre, devuelve un valor que incluye el nombre
    // en un mensaje de saludo.
    message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido al banco de Bogota!", name)
    return message, nil
}