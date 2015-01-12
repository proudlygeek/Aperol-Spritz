package main

import (
    "github.com/proudlygeek/Aperol-Spritz/lib"
)

func main() {
    a := aperol.Spritz()

    a.Get("/", func() string {
        return "<h1>Hello Dude!\nGimme Booze!</h1><img style='width: 500px;' src='http://www.etnaristorante.hu/userfiles/images/6e09f1638613de13fc87753b12207885Aperol_Spritz_tumbler_vine.jpg'/>"
    })

    a.Run(":8080")
}
