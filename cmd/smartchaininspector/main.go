// cmd/smartchaininspector/main.go
package main

import (
"flag"
"log"
"os"

"smartchaininspector/internal/smartchaininspector"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := smartchaininspector.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
