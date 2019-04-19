package queue

import (
	"fmt"
	"github.com/dulumao/Guten-utils/conv"
	"time"
)

func test() {
	q := Default()

	q.Add("test_delay", func(name string) {
		time.AfterFunc(2*time.Second, func() {
			fmt.Println(name)
			q.Done()
		})
	})

	for i := 1; i <= 5; i++ {
		q.Add("test"+conv.String(i), func(name string) {
			time.Sleep(10 * time.Second)
			fmt.Println(name)
			q.Done()
		})
	}

	q.Add("test_sleep", func(name string) {
		time.Sleep(5 * time.Second)
		fmt.Println(name)
		q.Done()
	})


	q.Wait()
}
