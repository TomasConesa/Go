package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

const (
	myKey    ctxKey = "my-key"
	myKeyInt ctxKey = "my-key-int"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, myKey, "my-value")
	ctx = context.WithValue(ctx, myKeyInt, 10)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Cuando se ha excedido el tiempo
	/* 	select {
	   	case <-ctx2.Done():
	   		fmt.Println("we've exceeded the deadline")
	   		fmt.Println(ctx2)
	   	} */ // Si tuviera mas case, si me sirve implementar el select sino lo puedo realizar como hice abajo

	<-ctx2.Done() // Es un canal que se cierra cuando el contexto se cancela o vence el timeout.
	fmt.Println("we've exceeded the deadline")
	fmt.Println(ctx2.Err())

}

func viewContext(ctx context.Context) {
	// El valor devuelto es interface{}, lo casteamos si hace falta
	if val, ok := ctx.Value(myKey).(string); ok {
		fmt.Printf("my value is %s \n", val)
	}
	if val, ok := ctx.Value(myKeyInt).(int); ok {
		fmt.Printf("my other is %d \n", val)
	}
}
