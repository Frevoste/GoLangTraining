package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Stałe zmienne
const (
	firma     = "Sii"
	data      = "15.04.2023"
	szkolenie = "Golang"

	// mon = iota  Automatycznie numeruje zmienne
	// tue
	// wed
)

// Kanały
func saySync(c chan int, s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		println(s)
	}

	c <- 1 // Zapis do kanału że już raz się skończylo
}

func Kanały() {
	//make(chan czego)
	c := make(chan int)
	go saySync(c, "Siema", 3000)
	go saySync(c, "Cześć", 1000)

	<-c
	fmt.Println("50%")
	<-c
	//Oczekuje dwóch znaków do kanału że się wykonało
	//Jezeli coroutine coś wyśle do kanału to odczeka aż ktoś to odczyta
}

// GoRoutines trzeba ustawić ilość dostepnych procesorow zeby działała
func Say(s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		println(s)
	}
}

func Rutyny() {
	go Say("Hello", 1000)
	go Say("Hola", 1000)
	time.Sleep(time.Second * 10)
}

//Interfejsy

type Greeter interface {
	Hello() string
}

func Greet(greeter Greeter) {
	fmt.Println(greeter.Hello())
}

type EnglishGreeter struct{}

func (eg EnglishGreeter) Hello() string {
	return "Hello World"
}

type SpanishGreeter struct{}

func (sg SpanishGreeter) Hello() string {
	return "Hola Mundo!"
}

// Dobra praktyka - Definiować interfejsy blisko ich użycia
func Interfejsy() {
	// eg := EnglishGreeter{}
	sg := SpanishGreeter{}
	Greet(sg)
}

// Dane z dużej litery są eksportowalne, dane z małej litery są prywatne, prywatne działają wobec jednej paczki, paczka to folder
// Struktury
type User struct {
	Name  string  `json:"name,omitempty"` // `json ... to są tagi, możemy dodawać ich ile potrzebujemy, i funkcje takie jak json.Marshal wyszukują te tagi by wiedzieć co zrobić
	Age   float64 `json:"age,omitempty"`
	phone string  `json:"phone,omitempty"`
}

// Zagnieżdzanie funkcji na strukturze
func (u *User) SetPhone(phone string) {
	u.phone = phone
}
func (u User) GetPhone() string {
	return u.phone
}

func struktury() {
	u := User{
		phone: "1231",
		Name:  "Andrzej",
		Age:   30,
	}
	u2 := User{}

	fmt.Printf("%+v\n", u)
	//Serializacja
	bytes, _ := json.Marshal(u)
	fmt.Printf("Marshal %+v \n", string(bytes))
	_ = json.Unmarshal(bytes, &u2)

	fmt.Printf("Unmarshaled: %+v \n", u2)
	u.SetPhone("XD")
	fmt.Println(u.GetPhone())

}

// Pointers * oznacza wskaźnik na komórkę, czyli adres
func mut(x *int) {
	*x = *x + 1
}

func pointers() {

	// a:= new(int)
	//*a = 1
	a := 1
	mut(&a) //wyciągnięcie adresu z a to &
}

// Obsługa błędów
func div(a, b float64) (float64, error) {
	if b == 0.0 {
		return math.Inf(1), errors.New("zero divison")
	}
	return a / b, nil
}

func błąd() {
	res, err := div(4, 1)
	if err != nil {
		log.Fatalf("błąd działania %v", err)
	}
	fmt.Println(res)
}

// Przyjmuje dwa argumenty a i b i zwraca dwa argumenty int
func reszta(a, b int) (int, int) {
	return a % b, (a - (a % b)) / b
}
func add(a, b int) int {
	return a + b
}

func inc(x int) {
	x++
}

// Funckja w funkcji
func up(s string) string {
	return strings.ToUpper(s)
}

func low(s string) string {
	return strings.ToLower(s)
}

func mapText(s string, f func(string) string) {
	fmt.Println(f(s))
}

func main() {
	//Zmienne
	var a int64 = 0
	var b int32 = 1
	var c string = "0"
	var d int = rand.Intn(10)

	//Casting
	a = int64(b)
	c = strconv.Itoa(d)

	fmt.Println("A value:", a)
	fmt.Println("Your random number is:", d)
	fmt.Println("Your random number converted to string is:", c)

	//Typy złożone
	str := "HelloSii"

	fmt.Println("Wartość byte pola 0 ze stringa to:", str[0]) //Wyciąga wartość Byte ze stringa
	fmt.Println("Pierwsza litera stringa to:", str[0:1])      //SLice ze stringa to dalej string

	//Tablice i listy
	l := []int{0, 1, 2, 3, 4} //Deklaracja listy
	t := [3]int{0, 1, 2}      //Deklaracja tablicy lub [...]int{0,1,2,3}
	l2 := []int{4, 3, 2, 1}
	makeSlice := make([]int, 10) //Tworzy slice o długości 10 i wypełnia go 0

	fmt.Println("Wycinek tablicy to:", t[0:2])
	fmt.Println("Pole 0 z listy to: ", l[0])
	fmt.Println("Rezultat makeSLice to:", makeSlice)

	//Wielowymiarowe tablice [ile][jakdlugich]int np 2 tablice 3 elementowe [2][3]
	l2d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Twoja dwuwymiarowa tablica to: ", l2d)
	//T.append todo
	l = append(l, l2...)

	//Mapy map[klucz]wartość{}
	mapa := map[string]string{
		"test": "true",
	}
	mapa["new"] = "true" //Dodanie kolejnego klucza
	fmt.Println("Twoja mapa to:", mapa)

	//Dwuwymiarowa mapa
	mapa2d := map[string]map[string]string{
		"test": {
			"case": "true",
		},
	}
	mapa2d["test"]["new"] = "false"                      // Dodanie klucza do istniejącej mapy
	mapa2d["test1"] = map[string]string{"case1": "true"} // Dodanie klucza mapy i klucza do tej mapy
	fmt.Println("Twoja mapa dwuwymiarowa to:", mapa2d)   // Wyciągnięcie danego case to mapa2d["test"]["case"]

	//If domyślnie bez nawiasów
	flag := true

	if flag {
		fmt.Println("If się sprawdził")
	} else if !flag {
		fmt.Println("If się nie sprawdził")
	} else {
		fmt.Println("Boo")
	}

	if r := 6 % 2; r == 0 {
		fmt.Println("Parzyste")
	} else {
		fmt.Println("Nieparzyste")
	}

	//Switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Niewiadomo")
	}

	//Zrzucenie typu do zmiennej
	// intMy := 3
	// typ := intMy.(type)

	//Pętle
	for i := 0; i < 10; i++ {
		fmt.Println("x")
	}
	//While
	warunek := true
	for warunek {
		fmt.Println("While")
		warunek = false
	}

	//Foreach
	r := []int{5, 4, 3, 2, 1}

	for i, v := range r {
		fmt.Printf("idx: %d, v: %d \n", i, v)
	}

	//For na mapie

	mapaFor := map[string]string{
		"a": "x",
		"b": "c",
	}

	for k, v := range mapaFor {
		fmt.Printf("k: %s , v: %s\n", k, v)
	}

	//Dwuwymiarowy For
	for2d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//label:
	for _, r := range for2d {
		for _, c := range r {
			fmt.Printf("%d, ", c)
			//break label
		}
		fmt.Println()
	}
	//Funkcje
	fmt.Println(add(2, 2))
	//Mapy są mutalne to znaczy że ich zmiana w funkcji jest zmianą globalną, większość zmiennych mutalna nie jest więć zmieniamy je tylko lokalnie

	//Zwracanie wielu zmiennych z funkcji
	resz, cal := reszta(10, 3)
	fmt.Println(resz, cal)

	mapText("XDerere", low)
	mapText("XDerere", up)
	błąd()
	pointers()
	struktury()
	Interfejsy()
	Rutyny()
	Kanały()
}
