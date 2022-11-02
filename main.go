package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Printf("echo is %v", e)
}
