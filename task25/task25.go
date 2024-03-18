package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func Sleep(n time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	defer cancel()
	<-ctx.Done()
}

func main() {
	fmt.Println("Start")
	Sleep(3)
	fmt.Println("Finish")
}
