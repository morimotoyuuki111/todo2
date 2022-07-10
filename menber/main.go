package main

import (
    "gin_test/server"
)

func main() {
    r := server.GetRouter()
    r.Run(":8080")
}
