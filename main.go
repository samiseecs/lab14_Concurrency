package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func labTask1() {
	fmt.Println("Lab Task 1: Introduction to Multithreading")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Printf("Number: %d\n", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Printf("Square: %d\n", i*i)
		}
	}()

	wg.Wait()
	fmt.Println("Lab Task 1 completed.")
}

func labTask2() {
	fmt.Println("Lab Task 2: Thread Synchronization")
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)

	increment := func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	go increment()
	go increment()
	go increment()

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter)
	fmt.Println("Lab Task 2 completed.")
}

func labTask3() {
	fmt.Println("Lab Task 3: Concurrent Data Structures")
	var wg sync.WaitGroup
	sharedMap := sync.Map{}

	writeToMap := func(key, value int) {
		sharedMap.Store(key, value)
	}

	readFromMap := func(key int) {
		if value, ok := sharedMap.Load(key); ok {
			fmt.Printf("Key: %d, Value: %d\n", key, value)
		} else {
			fmt.Printf("Key %d not found\n", key)
		}
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			writeToMap(i, i*i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			readFromMap(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			writeToMap(i, i*i)
		}
	}()

	wg.Wait()
	fmt.Println("Concurrent data structure operations completed.")
	fmt.Println("Lab Task 3 completed.")
}

func labTask4() {
	fmt.Println("Lab Task 4: Simulation of Bank Transaction System")
	type BankAccount struct {
		balance int
		mu      sync.Mutex
	}

	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	client := func(id int) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			amount := rand.Intn(500)
			if rand.Intn(2) == 0 {
				account.mu.Lock()
				account.balance += amount
				account.mu.Unlock()
				fmt.Printf("Client %d Deposited: %d, Balance: %d\n", id, amount, account.balance)
			} else {
				account.mu.Lock()
				if account.balance >= amount {
					account.balance -= amount
					fmt.Printf("Client %d Withdrew: %d, Balance: %d\n", id, amount, account.balance)
				} else {
					fmt.Printf("Client %d Withdrawal of %d failed. Insufficient balance.\n", id, amount)
				}
				account.mu.Unlock()
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	wg.Add(3)

	go client(1)
	go client(2)
	go client(3)

	wg.Wait()
	fmt.Printf("Final Account Balance: %d\n", account.balance)
	fmt.Println("Lab Task 4 completed.")
}

func main() {
	//labTask1()
	//labTask2()
	//labTask3()
	labTask4()
}
