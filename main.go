package main

import (

	// OWM
	openweathermap "github.com/briandowns/openweathermap"

	// Fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// Others
	"fmt"
	"strconv"
	"time"

	"io/ioutil"

	"math/rand"

)

func main() {

	a := app.New()
	w := a.NewWindow("Погода")

	icon, _ := fyne.LoadResourceFromPath("icon.png") // set icon

	w.Resize(fyne.NewSize(220, 200)) // set sizes for window
	w.SetFixedSize(true) // set unresizable window
	w.SetIcon(icon) // set icon for window panel






	// Entries
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Напишите город")

	// Recommends
	recLabel := widget.NewLabel("Рекомендаций пока нет")

	// const text
	tempLabel := widget.NewLabel("Температура:           ")
	tempMaxLabel := widget.NewLabel("Макс. Температура:")
	tempMinLabel := widget.NewLabel("Мин. температура: ")

	windSpeedLabel := widget.NewLabel("Скорость ветра:       ")
	humidityLabel := widget.NewLabel("Влажность:                ")

	// variables of weather
	tempLabelInt := widget.NewLabel("")
	tempMaxLabelInt := widget.NewLabel("")
	tempMinLabelInt := widget.NewLabel("")

	windSpeedLabelInt := widget.NewLabel("")
	humidityLabelInt := widget.NewLabel("")








	// buttons
	searchButton := widget.NewButton("Узнать погоду", func() { // function

		w, err := openweathermap.NewCurrent("C", "EN", "YOUR_API")
		if err != nil {
			fmt.Println(err)
		}

		w.CurrentByName(entry.Text)

		// Demonstration inf

		if w.Main.Temp == 0.0 && // if uncorrectable city
			w.Main.TempMax == 0 &&
			w.Main.TempMin == 0 &&
			w.Wind.Speed == 0 &&
			w.Main.Humidity == 0 {
			
			// set error message in entry
			entry.SetText("Ваш город введен некорректно") 
		
			time.Sleep(time.Second)
			entry.SetText("")
		
		} else {

			recommends := make(chan string)

			go recommend(w, recommends)

			go tempLabelInt.SetText(strconv.FormatFloat(w.Main.Temp, 'f', -1, 64) + "°C")
			go tempMaxLabelInt.SetText(strconv.FormatFloat(w.Main.TempMax, 'f', -1, 64) + "°C")
			go tempMinLabelInt.SetText(strconv.FormatFloat(w.Main.TempMin, 'f', -1, 64) + "°C")
			go windSpeedLabelInt.SetText(strconv.FormatFloat(w.Wind.Speed, 'f', -1, 64) + "м/с")
			go humidityLabelInt.SetText(fmt.Sprint(w.Main.Humidity) + "%")

			go recLabel.SetText(<-recommends)

			time.Sleep(time.Nanosecond / 1000000000000000000)

		}

	})

	clearEntryButton := widget.NewButton("  Очистить  ", func() {

		go entry.SetText("")

		go recLabel.SetText("Рекомендаций пока нет")

		go tempLabelInt.SetText("")
		go tempMaxLabelInt.SetText("")
		go tempMinLabelInt.SetText("")

		go windSpeedLabelInt.SetText("")
		go humidityLabelInt.SetText("")

		time.Sleep(time.Nanosecond / 1000000000000000000 / 1000000000000000000 / 1000000000000000000 / 1000000000000000000 / 1000000000000000000)

	})





	// Menu buttons
	saveCityButton := widget.NewButton("Сохр. город", func() {

		if entry.Text == "" ||
			entry.Text == "Ваш город введен некорректно" ||
			entry.Text == "Введите город в поле" ||
			entry.Text == "Вы не сохранили город" ||
			entry.Text == "Успех!" {

				entry.SetText("Введите город в поле")
				
				time.Sleep(time.Second)
				entry.SetText("")

		} else {

			err := ioutil.WriteFile("saved_city.txt", []byte(entry.Text), 0644)
			if err != nil {

				fmt.Println("Ошибка при сохранении файла:", err)

			}

			entry.SetText("Успех!")

		}

	})

	weatherSavedCityButton := widget.NewButton("Исп. сохр. город", func() {

		city, err := ioutil.ReadFile("saved_city.txt")
		if err != nil {
			
			entry.SetText("Вы не сохранили город")
			time.Sleep(time.Second)
			entry.SetText("")

		} else {

			go entry.SetText(string(city))

			// ButtonSearch
			w, err := openweathermap.NewCurrent("C", "EN", "649896f57f241e916a80292f65b1f6f6")
			if err != nil {
				fmt.Println(err)
			}

			w.CurrentByName(string(city))

			// Demonstration inf

			recommends := make(chan string)

			go recommend(w, recommends)

			go tempLabelInt.SetText(strconv.FormatFloat(w.Main.Temp, 'f', -1, 64) + "°C")
			go tempMaxLabelInt.SetText(strconv.FormatFloat(w.Main.TempMax, 'f', -1, 64) + "°C")
			go tempMinLabelInt.SetText(strconv.FormatFloat(w.Main.TempMin, 'f', -1, 64) + "°C")
			go windSpeedLabelInt.SetText(strconv.FormatFloat(w.Wind.Speed, 'f', -1, 64) + "м/с")
			go humidityLabelInt.SetText(fmt.Sprint(w.Main.Humidity) + "%")
			
			go recLabel.SetText(<-recommends)

			time.Sleep(time.Nanosecond / 1000000000000000000 / 1000000000000000000 / 1000000000000000000 / 1000000000000000000 / 1000000000000000000)

		}

	})

	// set content
	w.SetContent(

		container.NewVBox(

			entry,
			
			container.NewHBox(

				searchButton,
				clearEntryButton,

			),

			recLabel,

			container.NewHBox(
				tempLabel,
				tempLabelInt,
			),

			container.NewHBox(
				tempMaxLabel,
				tempMaxLabelInt,
			),

			container.NewHBox(
				tempMinLabel,
				tempMinLabelInt,
			),

			container.NewHBox(
				windSpeedLabel,
				windSpeedLabelInt,
			),

			container.NewHBox(
				humidityLabel,
				humidityLabelInt,
			),

			container.NewHBox(
				saveCityButton,
				weatherSavedCityButton,
			),
			
		),
	)
	
	w.Show()
	w.SetMaster()

	a.Run()
}


func recommend(w *openweathermap.CurrentWeatherData, ch chan string) {

	var temp1LvlRec = [5]string {

		"Утепляйся и не болей",
        "Оденься потеплее",
        "Утеплись получше",
        "Зима, утеплись",
        "Надевай зимнюю куртку",

	}

    var temp2LvlRec = [5]string {

		"Советую тепло одеться",
        "Хорошо будет одеться тепло",
        "Оденься тепло и будет ок",
        "Тепло оденься, если выйдешь",
        "Потеплее оденьтесь, сэр",

	}

    var temp3LvlRec = [5]string {

		"Надень куртку",
        "Оденься по-осеннему",
        "Одень куртку и тогда выходи",
        "Советую надеть куртку",
        "Наденьте что-нибудь осеннее",

	}

    var temp4LvlRec = [5]string {

		"Одень шорты и футболку",
        "Наслаждайся жарой на пляже",
		"Надевай плавки, иди купаться",
        "Оденься в футболку и шорты.",
        "Оденься по-летнему",

	}

    var temp5LvlRec = [5]string {

		"ИДИ В ТЕНЬ ИЛИ СГОРИШЬ!",
        "ДАЖЕ ГОЛЫМ ТЕБЕ КОНЕЦ!",
        "НЕ ВЫХОДИ НА УЛИЦУ!",
        "КОНДИЦИОНЕР - СПАСЕНИЕ!",
        "ВРУБАЙ КОНДИЦИОНЕР!",

	}

	if w.Main.Temp <= -5 {

		ch <- temp1LvlRec[rand.Intn(4)]

	} else if w.Main.Temp > -5 && w.Main.Temp <= 7 { // temp in (-5; 10]
		
		ch <- temp2LvlRec[rand.Intn(4)]
		
	} else if w.Main.Temp > 7 && w.Main.Temp < 18 { // temp in (10; 18)

		ch <- temp3LvlRec[rand.Intn(4)]
		
	} else if w.Main.Temp >= 18 && w.Main.Temp <= 33 { 	// temp in [18; 33]
		
		ch <- temp4LvlRec[rand.Intn(4)]
		
	} else if w.Main.Temp > 33 { 	// temp in [18; 33]

		ch <- temp5LvlRec[rand.Intn(4)]
		
	} else {

		ch <- "Странная погодка, правда"

	}

}
