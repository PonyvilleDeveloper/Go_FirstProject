package logger

import (
	"fmt"
	"os"
	"time"
)

func Log(pack string, mes string, vals ...any) {
	log := fmt.Sprintf("%v [%s]: %s", time.Now(), pack, mes)
	for v := range vals {
		log += fmt.Sprintf(" %v", v)
	}
	log += "\n"
	filename := fmt.Sprintf("logs/log %d.%d.%d", time.Now().Day(), time.Now().Month(), time.Now().Year())
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(log)
	fmt.Print(log)
}
