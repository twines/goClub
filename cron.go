package main

import (
	. "github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
