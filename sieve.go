// A concurrent prime sieve

package main

import "encoding/binary"
import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

type Packet struct {
	buffer []byte
}

func (p *Packet) Int16(i int) int16 {
	return int16(binary.BigEndian.Uint16(p.buffer[i : i+2]))
}
// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 9; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	p := &Packet{buffer: []byte{0x01, 0x02, 0x04}}
	fmt.Println("\nBytes: ", p.Int16(1))
}

