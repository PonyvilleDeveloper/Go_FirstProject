package service

import (
	"fmt"
	"os"
	"time"
)

func log(str string, vals ...any) {
	mes := fmt.Sprintf(str, vals)
	file, err := os.OpenFile(fmt.Sprintf("logs/log %v.txt", time.Now()), os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		fmt.Sprintf("\t\t[SERVICE]: Encoding current json item error: %v\n", err)
	}
	file.WriteString(mes)
	fmt.Print(mes)
}
