// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {

// 	links := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://stackoverflow.com",
// 		"http://golang.org",
// 		"http://amazon.com",
// 	}

// 	c := make(chan string)

// 	for _, v := range links {
// 		go checkLink(v, c)

// 	}

// 	for l := range c {

// 		go func(link string) {

// 			time.Sleep(5 * time.Second)
// 			checkLink(link, c)

// 		}(l)

// 	}

// }

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- link
// 		return
// 	}

// 	fmt.Println(link, "is up!")
// 	c <- link

// }

package main

import (
	"fmt"
	"time"
)

func count(thing string, ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- thing                        // mengirim nilai ke channel
		time.Sleep(time.Millisecond * 500) // delay
	}
	close(ch) // menutup channel setelah selesai mengirim
}

func main() {
	ch := make(chan string)
	go count("sheep", ch) // menjalankan count sebagai Go Routine
	for msg := range ch { // menerima nilai dari channel
		fmt.Println(msg)
	}
}
