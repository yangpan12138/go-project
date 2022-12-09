package animal

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

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
