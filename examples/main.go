package main

import (
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/k0kubun/pp"
)

func main() {
	c1, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Fatal(err)
	}
	id, err := c1.Put([]byte("KEY_01"), 1, 0, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1000 * time.Millisecond)
	pp.Println("id", id)

	c2, _ := beanstalk.Dial("tcp", "127.0.0.1:11300")
	f := func(c *beanstalk.Conn) {
		id, body, err := c.Reserve(2 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		pp.Println("consume id:", id, "body: ", string(body))
	}
	for i := 0; i < 2; i++ {
		f(c2)
	}

	// err = c2.Delete(id)
	// pp.Println(err)

	// stats, err := c2.Stats()
	// pp.Println(err)
	// pp.Println(stats)

	statsJob, err := c2.StatsJob(id)
	pp.Println(err)
	pp.Println(statsJob)
}
