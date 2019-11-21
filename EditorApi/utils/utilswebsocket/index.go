package utilswebsocket

import (
	"time"
)

// executer la fonction a un temps donné
func setInterval(someFunc func(), milliseconds int) {
	interval := time.Duration(milliseconds) * time.Second
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			go someFunc()
		}
	}
}
