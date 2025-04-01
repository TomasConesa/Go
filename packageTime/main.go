package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now() // Fecha actual
	p(now)

	then := time.Date(2000, 4, 1, 15, 40, 28, 891165, time.UTC) // UTC fecha global
	p(then)

	then = then.Add(24 * time.Hour) // Agrego un día
	p(then)

	// Funciones a las que puedo acceder
	p(then.Year())
	p(then.Day())
	p(then.Month())
	p(then.Hour())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	// Validaciones o comparaciones con las fechas
	p("Then es antes que now:", then.Before(now))
	p("Then es despues que now:", then.After(now))
	p("Then es igual que now:", then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours() / 24)         // Días
	p((diff.Hours() / 24) / 365) // Años

}
