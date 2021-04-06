package main

import (
    "log"
    "sync"
    "time"
)


type Bank struct {
    balance int
    mux     sync.Mutex    // read write lock
//    mux     sync.RWMutex    // read write lock
}

func (b *Bank) Deposit(amount int) {
    b.mux.Lock()            // write lock
    time.Sleep(time.Second)
    b.balance += amount
    b.mux.Unlock()          // wirte unlock
}

func (b *Bank) Balance() (balance int) {
    b.mux.Lock()          // read lock
//    b.mux.RLock()          // read lock
    time.Sleep(time.Second)
    balance = b.balance
    b.mux.Unlock()        // read unlock
//    b.mux.RUnlock()        // read unlock
    return
}


func main() {
    var wg sync.WaitGroup
    b := &Bank{}

    n := 5
    wg.Add(n)
    for i := 1; i <= n; i++ {
        go func() {
            b.Deposit(1000)
            log.Printf("Write: deposit amonut: %v", 1000)
            wg.Done()
        }()
    }
    wg.Add(n)
    for i := 1; i <= n; i++ {
        go func() {
            log.Printf("Read: balance: %v", b.Balance())
            wg.Done()
        }()
    }

    wg.Wait()
}
