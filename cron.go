package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run dao.CleanAllTag...")
		//dao.CleanAllTag()
	})
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run dao.CleanAllArticle...")
		//dao.CleanAllArticle()
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