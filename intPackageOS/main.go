package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // Finaliza el programa
	}
	fmt.Println(file)
	v, _ := file.Stat()
	fmt.Println(v.Name(), v.Size())

	data := make([]byte, 1024)
	//fmt.Println(data)

	c, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data[:c]), c)                 // Con string podemos leerlo, sino se muestra la información en bytes
	fmt.Printf("read: %d bytes: %q \n", c, data[:c]) // Otra manera de leerlo

	// Obtener el valor de una variable de entorno
	fmt.Println(os.Getenv("USERNAME"))
	os.Setenv("MI_ENV", "my value")
	fmt.Println(os.Getenv("MI_ENV"))

	dir, _ := os.Getwd()
	fmt.Println(dir)

	// Para saber si una variable de entorno existe
	os.Setenv("ENV_EXISTS", "")
	env, ok := os.LookupEnv("ENV_EXISTS")
	fmt.Println(env, ok)

	os.Setenv("DB_USERNAME", "tomas")
	os.Setenv("DB_PASSWORD", "tomas")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "8080")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT:$DB_NAME")
	fmt.Println(dbURL)
}
