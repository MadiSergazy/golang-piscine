package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if n < 2 {
		return
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			print(i, "*")
			n = n / i
			i = 1
		}
	}
	print(n)
	fmt.Println()
}

/*
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func main() {
	c := Circle{Radius: 10}
	PrintArea(c)

	s := Square{Height: 10, Width: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%v is the smallest\n", l)
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
*/
////////////////////////////////////////////////////////////////////////

// type User struct {
// 	Name string
// 	Id   int64
// }

// func main() {
// 	users := []User{

// 		{
// 			Id:   1,
// 			Name: "A",
// 		},
// 		{
// 			Id:   2,
// 			Name: "B",
// 		},
// 		{
// 			Id:   3,
// 			Name: "C",
// 		},
// 		{
// 			Id:   1,
// 			Name: "A",
// 		},
// 	}

// 	userMap := make(map[int64]User, len(users))
// 	for _, user := range users {
// 		if _, ok := userMap[user.Id]; !ok {
// 			userMap[user.Id] = user
// 		}
// 	}
// 	fmt.Println(findmap(1, userMap))
// }

// func findmap(id int64, usermap map[int64]User) *User {
// 	if user, ok := usermap[id]; ok {
// 		return &user
// 	}
// 	return nil
// }

//////////////////////////////////////////////////////////////////////////////////////
// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	arrays()
// }

// func arrays() {

// 	people := [2]Person{
// 		{
// 			Age:  30,
// 			Name: "Kate",
// 		},
// 		{
// 			Age:  20,
// 			Name: "Bil",
// 		},
// 	}
// 	fmt.Printf("type: %T val: %v\n", people, people)
// }

///////////////////////////////////////////////////////////////////////////////////////////////////
// type Builder interface {
// 	Build()
// }

// type WoodBuilder struct {
// 	Person
// }

// func (wb WoodBuilder) Build() {
// 	fmt.Println("work wirh wood")
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// type BrickBuilder struct {
// 	Person
// }

// func (bb BrickBuilder) Build() {
// 	fmt.Println("work with brick")
// }type Builder interface {
// 	Build()
// }

// type WoodBuilder struct {
// 	Person
// }

// func (wb WoodBuilder) Build() {
// 	fmt.Println("work wirh wood")
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// type BrickBuilder struct {
// 	Person
// }

// func (bb BrickBuilder) Build() {
// 	fmt.Println("work with brick")
// }

// type Building struct {
// 	Builder
// 	Name string
// }

// func main() {
// 	uscace()
// }

// func uscace() {
// 	woodBuilding := Building{
// 		Builder: WoodBuilder{Person{
// 			Name: "Kate",
// 			Age:  40,
// 		}},
// 		Name: "Wood house",
// 	}
// 	woodBuilding.Build()
// 	brickBuilder := Building{
// 		Builder: BrickBuilder{
// 			Person{
// 				Name: "Jane",
// 				Age:  30,
// 			},
// 		},
// 		Name: "Brick house",
// 	}
// 	brickBuilder.Build()
// }

// type Building struct {
// 	Builder
// 	Name string
// }

// func main() {
// 	uscace()
// }

// func uscace() {
// 	woodBuilding := Building{
// 		Builder: WoodBuilder{Person{
// 			Name: "Kate",
// 			Age:  40,
// 		}},
// 		Name: "Wood house",
// 	}
// 	woodBuilding.Build()
// 	brickBuilder := Building{
// 		Builder: BrickBuilder{
// 			Person{
// 				Name: "Jane",
// 				Age:  30,
// 			},
// 		},
// 		Name: "Brick house",
// 	}
// 	brickBuilder.Build()
// }
///////////////////////////////////////////////////////////////

// type Runner interface { // интерфейс Runner у него есть метод-Run который должен вернуть стринг
// 	Run() string
// }

// type Swimmer interface {
// 	Swim() string
// }

// type Flyer interface {
// 	Fly() string
// }

// type Ducker interface { // в интеурфейс даккер встроили другие интерфейсы
// 	Runner
// 	Swimmer
// 	Flyer
// }

// type Human struct {
// 	Name string
// }

// func (h Human) Run() string {
// 	return "Человек бежит"
// }

// func main() {
// 	interfaceValues()
// 	typeas()
// }

// func interfaceValues() {
// 	var runner Runner
// 	fmt.Printf("Type: %T Value : %v\n", runner, runner)

// 	if runner == nil {
// 		fmt.Println("Runner is nil")
// 	}
// 	asdf := &Human{Name: "Dias"}
// 	fmt.Printf("Type %T Value: %#v\n", asdf, asdf)
// 	asdlfj(asdf)
// }

// func asdlfj(dsaf Runner) {
// 	fmt.Println(dsaf.Run())
// }

// func typeas() {
// 	var runner Runner
// 	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
// }Person64) {
// 	global--
// 	fmt.Println("saving...")
// 	return uint64(data.counter)
// }

// func main() {
// 	if err, data := create("", 0); err != nil {
// 		defer func() {
// 			data.save()
// 			fmt.Println("saved")
// 			fmt.Println("global:", global)
// 		}()
// 		panic(err)
// 	}
// }

// func create(info string, count int) (error, data) {
// 	if len(info) == 0 {
// 		return errors.New("len is 0"), data{info: "", counter: 0, add: 0}
// 	}
// 	data := data{
// 		info:    info,
// 		counter: count,
// 		add:     1,
// 	}
// 	global++
// 	return nil, data
// }

// func main() {
// 	groupCity := map[int][]string{
// 		10:   {"xromtay", "shieli"},         // города с населением 10-99 тыс. человек
// 		100:  {"Aktobe", "Aktau", "Atyrau"}, // Gorodа с населением 100-999 тыс. человек
// 		1000: {"Almaty"},                    // Gorodа с населением 1000 тыс. человек и более
// 	}
// 	cityPopulation := map[string]int{
// 		"xromtay": 10,
// 		"Aktobe":  100,
// 		"Almaty":  1000,
// 		"Aktau":   100,
// 		"Atyrau":  100,
// 	}

// 	for key := range cityPopulation {
// 		var exist bool
// 		for _, city := range groupCity[100] {
// 			if city == key {
// 				exist = true
// 				break
// 			}
// 		}
// 		if !exist {
// 			delete(cityPopulation, key)
// 		}
// 	}

// 	fmt.Println(cityPopulation)
// }
