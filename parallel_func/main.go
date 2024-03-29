package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(4)
}

const (
	MAX_TIME_EXECUTION   = 500  // Максимальное время выполнения создаваемых функций
	MAX_ERROR_COUNT      = 6    // Максимальное допустимое количество ошибок после которого останавливаем если еще что-то осталось
	NUMBER_FUNCS         = 1000 //количество функций
	MAX_PACKET_EXECUTION = 5    // Максимальное количество одновременно выполняемых заданий
)

func main() {
	var executer Executer

	pf := getFuncs(NUMBER_FUNCS)
	err := executer.startTasks(pf[:], MAX_ERROR_COUNT, MAX_PACKET_EXECUTION, true)
	if err != nil {
		fmt.Printf("Возникла ошибка при выполнении: %v", err)
	}
}

// Создаем массив из функций с различным временем выполнения.
func getFuncs(n int) []Task {
	var result []Task
	for i := 0; i < n; i++ {
		z := i
		timeExec := time.Millisecond * time.Duration(rand.Intn(MAX_TIME_EXECUTION))
		f := func() error {
			time.Sleep(timeExec)
			var makeError bool
			if rand.Intn(100) > 90 { // 10 процентов функций из созданных возвращают ошибку
				makeError = true
			}
			fmt.Printf("Message from func #%v, time exection %v, error = %v\n", z, timeExec, makeError)
			if makeError {
				return errors.New("error in function")
			} else {
				return nil
			}

		}
		result = append(result, f)
	}

	return result
}
