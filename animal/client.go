package animal

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
