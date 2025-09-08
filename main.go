package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Func struct {
	Name     string
	Function func()
}

type Film struct {
	Name   string
	Genres []string
	Mood   []string
}

type Music struct {
	Name   string
	Genres []string
	Mood   []string
}

type Merch struct {
	Name        string
	Type        string
	Price       float32
	Description string
}

//Тестовые данные сгенерированы ИИ. Кожаному слишком лень придумывать

var Films = []Film{
	{"Интерстеллар", []string{"фантастика", "драма"}, []string{"вдохновляющий", "задумчивый"}},
	{"Назад в будущее", []string{"комедия", "фантастика"}, []string{"весёлый", "ностальгия"}},
	{"Брат", []string{"драма", "боевик"}, []string{"серьёзный", "грустный"}},
}

var Musics = []Music{
	{"Imagine Dragons - Believer", []string{"рок", "альтернатива"}, []string{"энергичный", "мотивация"}},
	{"Queen - Bohemian Rhapsody", []string{"рок", "классика"}, []string{"ностальгия", "драматичный"}},
	{"Баста - Сансара", []string{"рэп", "русский"}, []string{"задумчивый", "грустный"}},
}

var Anecdotes = []string{
	"— Почему программисты путают Хэллоуин и Рождество? — Потому что OCT 31 == DEC 25.",
	"Вчера был баг, сегодня фича.",
	"Программист — это машина для преобразования кофе в код.",
}

var Merchs = []Merch{
	{"Футболка КепРоБот", "одежда", 799.99, "Стильная футболка с логотипом КепРоБот."},
	{"Кружка КепРоБот", "посуда", 399.50, "Керамическая кружка для настоящих фанатов."},
	{"Брелок КепРоБот", "аксессуар", 199.00, "Металлический брелок с символикой."},
}

var Functions = map[string]Func{
	"1": {"Найти фильм по жанру", GetFilmByGenre},
	"2": {"Найти фильм по настроению", GetFilmByMood},
	"3": {"Найти музыку по жанру", GetMusicByGenre},
	"4": {"Найти музыку по настроению", GetMusicByMood},
	"5": {"Рассказать анекдот", GetAnecdote},
	"6": {"Показать мерч", GetMerch},
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Добро пожаловать в КепРоБот!!!\n")
	var Answer string
	for {
		fmt.Println("Что вы хотите сделать?")
		for i := 1; i <= len(Functions); i++ {
			fmt.Printf("%v. %v\n", i, Functions[fmt.Sprintf("%v", i)].Name)
		}

		fmt.Scan(&Answer)
		fmt.Print("\n\033[H\033[2J")
		Functions[Answer].Function()
	}
}

func GetAnecdote() {
	idx := rand.Intn(len(Anecdotes))
	fmt.Printf("Анекдот: %v\n\n", Anecdotes[idx])
}

func GetFilmByGenre() {
	fmt.Printf("Введите жанр: ")
	var genre string
	fmt.Scan(&genre)

	var TryFilms []Film
	for _, film := range Films {
		for _, filmGenre := range film.Genres {
			if filmGenre == genre {
				TryFilms = append(TryFilms, film)
				break
			}
		}
	}
	if len(TryFilms) == 0 {
		fmt.Println("Таких фильмов не найдено\n")
		return
	}
	fmt.Println("Вот все фильмы с таким жанром:")
	for i, film := range TryFilms {
		fmt.Printf("%v. %v (жанр: %v)\n", i+1, film.Name, strings.Join(film.Genres, ", "))
	}
	fmt.Println("")
}

func GetFilmByMood() {
	fmt.Printf("Введите Настроение: ")
	var mood string
	fmt.Scan(&mood)

	var TryFilms []Film
	for _, film := range Films {
		for _, filmGenre := range film.Mood {
			if filmGenre == mood {
				TryFilms = append(TryFilms, film)
				break
			}
		}
	}
	if len(TryFilms) == 0 {
		fmt.Println("Таких фильмов не найдено")
		return
	}
	fmt.Println("Вот все фильмы с таким настроением:")
	for i, film := range TryFilms {
		fmt.Printf("%v. %v (настроение: %v)\n", i+1, film.Name, strings.Join(film.Mood, ", "))
	}
	fmt.Println("")
}

func GetMusicByGenre() {
	fmt.Printf("Введите жанр: ")
	var genre string
	fmt.Scan(&genre)

	var TryFilms []Music
	for _, film := range Musics {
		for _, filmGenre := range film.Genres {
			if filmGenre == genre {
				TryFilms = append(TryFilms, film)
				break
			}
		}
	}
	if len(TryFilms) == 0 {
		fmt.Println("Таких песен не найдено")
		return
	}
	fmt.Println("Вот все песни с таким жанром:")
	for i, film := range TryFilms {
		fmt.Printf("%v. %v (жанр: %v)\n", i+1, film.Name, strings.Join(film.Genres, ", "))
	}
	fmt.Println("")
}

func GetMusicByMood() {
	fmt.Printf("Введите Настроение: ")
	var mood string
	fmt.Scan(&mood)

	var TryFilms []Music
	for _, film := range Musics {
		for _, filmGenre := range film.Mood {
			if filmGenre == mood {
				TryFilms = append(TryFilms, film)
				break
			}
		}
	}
	if len(TryFilms) == 0 {
		fmt.Println("Таких песен не найдено")
		return
	}
	fmt.Println("Вот все песни с таким настроением:")
	for i, film := range TryFilms {
		fmt.Printf("%v. %v (настроение: %v)\n", i+1, film.Name, strings.Join(film.Mood, ", "))
	}
	fmt.Println("")
}

func GetMerch() {
	for i, merch := range Merchs {
		fmt.Printf("%v. %v (%v) - %.2f руб.\n", i+1, merch.Name, merch.Type, merch.Price)
		fmt.Printf("   Описание: %s\n", merch.Description)
	}
	fmt.Println("")
	fmt.Println("Что вы хотите сделать?\n1. Купить товар\n2. Вернуться назад")
	var action string
	fmt.Scan(&action)
	switch action {
	case "1":
		fmt.Print("Выберите товар для покупки: ")
		var choice int
		fmt.Scan(&choice)
		if choice < 1 || choice > len(Merchs) {
			fmt.Println("Такого мерча нет!")
			return
		}
		Select := Merchs[choice-1]
		fmt.Printf("\n%s (%s) - %.2f руб.\n", Select.Name, Select.Type, Select.Price)
		fmt.Print("Чтобы купить введите \"1\": ")
		var confirm string
		fmt.Scan(&confirm)

		if confirm == "1" {
			fmt.Printf("Вы купили %v.\n", Select.Name)
			fmt.Println("Спасибо за покупку!")
		} else {
			fmt.Println("Покупка отменена.")
		}
		fmt.Println()
	case "2":
		fmt.Println("Возвращаемся в главное меню...")
		return
	default:
		fmt.Println("Такого варианта ответа нет!")
		fmt.Println()
	}
}
