package main

import (
	"fmt"
	"go-interface/animal"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	dog, err := animal.NewAnimal(0)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	now := time.Now()
	dog.Eat("肉肉", 2)
	dog.Play("公园", now.Format("2006-01-02 15:04:05"), []string{"小花", "大白", "狗蛋"})
	dog.Sleep(now.Add(1 * time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Println("-------------------------------------------分割---------------------------------")
	cat, err := animal.NewAnimal(1)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	cat.Eat("小鱼", 3)
	cat.Play("公园", now.Format("2006-01-02 15:04:05"), []string{"二哈", "大白", "狗蛋"})
	cat.Sleep(now.Add(1 * time.Hour).Format("2006-01-02 15:04:05"))
}
