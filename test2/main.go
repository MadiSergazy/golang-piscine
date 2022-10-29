package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		return
	} else if len(os.Args) > 2 {
		fmt.Println("ERROR: too many arguments")
		return
	}
	if !checkStdFile() { // защищаем standard.txt от изменени
		fmt.Println("ERROR: file was changed")
		return
	}
	data2, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("ERROR: open file")
		return
	}

	var res1 []string
	var s string
	counter := 0
	arttext := bufio.NewScanner(data2) // сканируем файл
	m2 := map[int][]string{}           // в standard.txt 94 символа по  ASCII 	от 32 до 126
	i := 32                            // начинаем с 32 чтобы совпадало по аски от 32 до 126
	for arttext.Scan() {
		s = arttext.Text() // сканируем по строчно standard.txt
		if s == "" {
			continue
		}
		res1 = append(res1, s) // добавляем каждую строку в массив
		counter++              // до тех пор пока счетчик станет 8
		if counter == 8 {      // восемь потому что каждый символ в standard.txt состоит из 8 строк
			m2[i] = append(m2[i], res1...) // собрали один символ в res1  и
			res1 = []string{}              //  добавляем собранный символ в m2
			s = ""                         // обнуляем все переменные для того чтобы собрать след символ
			counter = 0
			i++
		}
	}

	arg := os.Args[1]
	reg := regexp.MustCompile(`\\n`)
	arg = reg.ReplaceAllString(arg, "\n")
	arg1 := strings.Split(arg, "\n")                  // если аргументы подаются в несколько строк то
	m2[10] = []string{"", "", "", "", "", "", "", ""} // создаем новый массив и закидываем в новый массив
	for _, ch := range arg {                          // проверяем аргумент который подается
		if _, ok := m2[int(ch)]; !ok {
			fmt.Print("ERROR use only ascii symbol\n") // если в аргументе есть символы которые нету в нашей
			return                                     // map (карте) то выдаем ошибку
		}
	}

	if len(arg1) > 1 {
		if checkNewline(arg1) { //  эта функция нужна в случае когда
			arg2 := arg1[1:]                 // в качестве аргумента подается только "\n"
			for i := 0; i < len(arg2); i++ { // в качестве аргумента может подаватся несколько "\n\n\n"
				fmt.Println()
			}
		} else {
			for j := 0; j < len(arg1); j++ {
				if arg1[j] == "" {
					fmt.Println()
				} else {
					for i := 0; i < 8; i++ {
						for _, v := range arg1[j] {
							// fmt.Print(m[int(v)][i])
							fmt.Print(m2[int(v)][i])
						}
						fmt.Println()
					}
				}
			}
		}
	} else {
		if arg == "" { // тут подразумевается "\n" когда функ-я Split видит: "\n"
			fmt.Println() // создается [""] пустой массив
		} else {
			for i := 0; i < 8; i++ {
				for _, v := range arg {
					fmt.Print(m2[int(v)][i])
				}
				fmt.Println()
			}
		}
	}
}

func checkNewline(mass []string) bool { // проверяем все элемты на то пустые ли они.
	flag := true
	for i := 0; i < len(mass); i++ { // когда подаем (\n\n)  mass cодержит: ["", "", ""]
		if mass[i] != "" { // три пустых элемента и в ответ возвращает true
			flag = false
		}
	}
	return flag
}

func checkStdFile() bool {
	hasher := sha256.New()
	s, err := ioutil.ReadFile("standard.txt")
	hasher.Write(s)
	if err != nil {
		fmt.Println("Error CheckStdFile")
	}
	l := hasher.Sum(nil)
	hash_std := []byte{195, 236, 117, 132, 251, 126, 207, 189, 115, 158, 107, 63, 111, 99, 253, 31, 229, 87, 210, 174, 62, 36, 248, 112, 115, 13, 156, 248, 178, 85, 158, 148}
	return string(hash_std) == string(l)
}

//////////////////////////////////////
// func main() {
// 	for i := 1; i <= 3; i++ {
// 		file, err := os.Create(strconv.Itoa(i) + "text.txt")
// 		if err != nil {
// 			fmt.Println("error")
// 		}
// 		defer file.Close()
// 	}
// 	os.Rename("2text.txt", "4text.txt")
// }

//////////////////////////////////////
// func main() {
// 	data := []byte("Hello World!")
// 	errWrite := ioutil.WriteFile("test.txt", data, 0600)
// 	if errWrite != nil {
// 		fmt.Printf("%s", errWrite)
// 	}
// 	file, err := ioutil.ReadFile("test.txt")
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 	}
// 	fmt.Println(string(file))
// 	os.Remove("test.txt")

// 	filesFromDir, err := ioutil.ReadDir(".")
// 	if err != nil {
// 		fmt.Println("Error")
// 	}
// 	for _, file := range filesFromDir {
// 		// Проходим по всем найденным файлам и печатаем их имя и размер
// 		fmt.Printf("name: %s, size: %d\n", file.Name(), file.Size())
// 	}

// 	b := bytes.NewReader([]byte("Данные"))
// 	a, err := ioutil.ReadAll(b)
// 	if err != nil {
// 		fmt.Println("Error")
// 	}
// 	fmt.Println(string(a))

// 	file1, _ := os.Create("text.txt")
// 	file2, _ := os.Create("text.txt")
// 	info1, _ := file1.Stat() // функция Stat возвращает информацию о файле и ошибку
// 	info2, _ := file2.Stat()
// 	fmt.Println(os.SameFile(info1, info2))
// }

/////////////////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	text := "Test words"
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Printf("Ошибка создания файла [%s]", err.Error())
// 		return
// 	}
// 	stringLen, err := file.WriteString(text)
// 	if err != ni	data2, err := os.Open("standard.txt")
// if err != nil {
// 	fmt.Println("ERROR: open file")
// 	return
// }
// 	// fmt.Println(stringLen)
// 	openFile, err := os.Open(file.Name())
// 	if err != nil {
// 		fmt.Printf("Ошибка чтения из файла [%s]", err.Error())
// 		return
// 	}
// 	data := make([]byte, stringLen)

// 	for {
// 		_, err := openFile.Read(data)
// 		if err != io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Printf("Ошибка чтения из файла [%s]", err.Error())
// 			return
// 		}
// 	}
// 	fmt.Println(string(data))
// }
