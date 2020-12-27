package bank

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock() 
	defer mu.RUnlock()
	b := balance
	return b
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

// この関数は必ずMutexをロックしてから使うこと
func deposit(amount int) { balance = balance + amount }
