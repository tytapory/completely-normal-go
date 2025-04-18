// I saw a TikTok where "Hello, world" was bruteforced character by character.
// I wanted to optimize the process.
// This is my implementation of a simple brute-force string type.

package simpleprintablestring

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MySimpleString func(ch byte, i int) (bool, bool)

func (str MySimpleString) Print() {
	size := str.size()
	result := make([]byte, size)
	var wg sync.WaitGroup
	for i := range size {
		wg.Add(1)
		go str.bruteforceChar(i, &wg, result)
	}
	isBruteforceCompleted := make(chan struct{})
	printLoopDone := printLoop(isBruteforceCompleted, result)
	wg.Wait()
	close(isBruteforceCompleted)
	<-printLoopDone
	printFromBeginningOfLine(string(result))
}

func New(defaultString string) MySimpleString {
	str := []byte(defaultString)
	return func(ch byte, i int) (bool, bool) {
		if i < 0 || i >= len(str) {
			return false, false
		}
		return ch == str[i], true
	}
}

func (str MySimpleString) bruteforceChar(i int, wg *sync.WaitGroup, resultStr []byte) {
	chars := getRandomizedSliceOfPrintableChars()
	for _, ch := range chars {
		time.Sleep(time.Millisecond * 15)
		resultStr[i] = ch
		if ok, _ := str(ch, i); ok {
			break
		}
	}
	wg.Done()
}

func printLoop(stop chan struct{}, resultStr []byte) chan struct{} {
	done := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Millisecond * 50)
			select {
			case <-stop:
				done <- struct{}{}
				return
			default:
				printFromBeginningOfLine(string(resultStr))
			}
		}
	}()
	return done
}

func (str MySimpleString) size() int {
	size := 0
	for _, ok := str(0, 0); ok; _, ok = str(0, size) {
		size++
	}
	return size
}

func getRandomizedSliceOfPrintableChars() []byte {
	arr := make([]byte, 96)
	for i := range arr {
		arr[i] = byte(i + 32)
	}
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func printFromBeginningOfLine(str string) {
	fmt.Print("\r\033[2K")
	fmt.Print(str)
}
