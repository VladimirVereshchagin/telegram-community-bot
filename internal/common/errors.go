package common

import "log"

// HandleError обрабатывает ошибки и выводит сообщение в лог
var HandleError = func(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
