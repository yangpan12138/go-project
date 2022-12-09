package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	dog, err := NewAnimal(0)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	now := time.Now()
	dog.Eat("肉肉", 2)
	dog.Play("公园", now.Format("2006-01-02 15:04:05"), []string{"小花", "大白", "狗蛋"})
	dog.Sleep(now.Add(1 * time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Println("-------------------------------------------分割---------------------------------")
	cat, err := NewAnimal(1)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	cat.Eat("小鱼", 3)
	cat.Play("公园", now.Format("2006-01-02 15:04:05"), []string{"二哈", "大白", "狗蛋"})
	cat.Sleep(now.Add(1 * time.Hour).Format("2006-01-02 15:04:05"))
}

type Animal interface {
	Eat(food string, num int) string
	Sleep(tm string) string
	Play(place string, tm string, friends []string) string
}

func NewAnimal(body int) (Animal, error) {
	var animal Animal
	var err error
	switch body {
	case 0:
		if animal, err = NewDog("二哈", "哈士奇", 3); err != nil {
			return nil, err
		}
	case 1:
		if animal, err = NewCat("小花", "橘猫", 2); err != nil {
			return nil, err
		}
	}

	return animal, nil
}

type Dog struct {
	Name string
	Type string
	Age  int
}

func NewDog(name, Type string, age int) (*Dog, error) {
	dog := &Dog{
		Name: name,
		Type: Type,
		Age:  age,
	}
	return dog, nil
}

func (c *Dog) Eat(food string, num int) string {
	str := fmt.Sprintf("<%s>在吃 '%s',一共吃了 %d 块", c.Name, food, num)
	log.Info().Msg(str)
	return str
}
func (c *Dog) Sleep(tm string) string {
	str := fmt.Sprintf("<%s>要开始睡觉了,现在是 %s", c.Name, tm)
	log.Info().Msg(str)
	return str
}
func (c *Dog) Play(place string, tm string, friends []string) string {
	str := fmt.Sprintf("现在时间是：%s, <%s>正在和它的好朋友在一起玩,好朋友是：", tm, c.Name)
	for _, f := range friends {
		str += fmt.Sprintf("%s、", f)
	}
	log.Info().Msg(str)
	return str
}

type Cat struct {
	Name string
	Type string
	Age  int
}

func NewCat(name, Type string, age int) (*Cat, error) {
	cat := &Cat{
		Name: name,
		Type: Type,
		Age:  age,
	}
	return cat, nil
}

func (c *Cat) Eat(food string, num int) string {
	str := fmt.Sprintf("<%s>在吃 '%s',一共吃了 %d 条", c.Name, food, num)
	log.Info().Msg(str)
	return str
}
func (c *Cat) Sleep(tm string) string {
	str := fmt.Sprintf("<%s>要开始睡觉了,现在是 %s", c.Name, tm)
	log.Info().Msg(str)
	return str
}
func (c *Cat) Play(place string, tm string, friends []string) string {
	str := fmt.Sprintf("现在时间是：%s, <%s>正在和它的好朋友在一起玩,好朋友是：", tm, c.Name)
	for _, f := range friends {
		str += fmt.Sprintf("%s、", f)
	}
	log.Info().Msg(str)
	return str
}
