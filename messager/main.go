package main

import (
	"lib/models"
)

func main() {
	t := models.NewPayment()
	println(t.CaptureAmount)
}
