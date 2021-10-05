package main

import (
	Sunucu "KDAP/sunucu"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

var a = app.New()

func sunucu() {

	var w = fyne.CurrentApp().NewWindow("Sunucu")
	durumlbl := widget.NewLabel("Durum: Boşta")
	durum := container.NewVScroll(durumlbl)
	durum.SetMinSize(fyne.NewSize(400, 0))

	port := widget.NewEntry()
	sifre := widget.NewPasswordEntry()

	//HTTP
	httpCheck := widget.NewCheck("Aç", func(value bool) {
		if value {
			sifre.Disable()
		} else {
			sifre.Enable()
		}
	})

	//ipList
	ipList := widget.NewLabel("")
	go ipList.SetText("Yerel Adresler:\n" + Sunucu.YerelIP() + "\nGenel Adres:\n\t" + Sunucu.GenelIP())

	ipListCont := container.NewHScroll(
		container.NewVScroll(
			ipList,
		),
	)
	ipListCont.SetMinSize(fyne.NewSize(0, 300))

	//konumlbl
	konumlbl := widget.NewLabel("Dosya Seçilmedi!")

	isKlasor := false
	fileName := ""
	//Pencereyi doldur
	w.SetContent(container.NewHBox(
		durum,
		container.NewVBox(

			widget.NewForm(
				widget.NewFormItem("Port", port),
				widget.NewFormItem("Şifre", sifre),
				widget.NewFormItem("HTTP", httpCheck),
			),
			container.NewGridWithColumns(
				//burası ve klasör seçme şeyi
				//geliştirilmesi lazım!
				2,
				widget.NewButton("Dosya Seç", func() {
					isKlasor = false
					file, err := dialog.File().Title("Dosya Seç").Filter("All Files", "*").Load()
					if err != nil {
						konumlbl.SetText("Dosya seçerken bir hata oluştu!")
						return
					}
					konumlbl.SetText(file)
					fileName = file
				}),
				widget.NewButton("Klasör Seç", func() {
					isKlasor = true
					file, err := dialog.Directory().Title("Klasör Seç").Browse()
					if err != nil {
						konumlbl.SetText("Klasör seçerken bir hata oluştu!")
						return
					}
					konumlbl.SetText(file)
					fileName = file
				}),
			),
			container.NewHScroll(
				konumlbl,
			),
			widget.NewButton("Başlat", func() {

				//burada sunucu başlayacak
				//eğer şifre boşsa hata ver
				if sifre.Text == "" && !httpCheck.Checked {
					durumlbl.SetText("Durum: Şifre girin!")
					return
				}
				//eğer port sayı değilse hata ver
				portNum, err := strconv.Atoi(port.Text)
				if err != nil {
					durumlbl.SetText("Durum: \"Port\" sayı olmak zorunda!")
					return
				}
				durumlbl.SetText("Durum: Sunucu Başlatılıyor.")

				Sunucu.Baslat(fileName, portNum, sifre.Text, httpCheck.Checked, isKlasor)
			}),
			ipListCont,
		),
	))
	w.Resize(fyne.NewSize(600, 500))
	w.SetFixedSize(true)
	w.Show()
}

func istemci() {
	var w = fyne.CurrentApp().NewWindow("İstemci")

	durum := container.NewVScroll(widget.NewLabel("Durum: Boşta"))
	durum.SetMinSize(fyne.NewSize(400, 0))

	menu := container.NewVBox(

		widget.NewForm(
			widget.NewFormItem("IP", widget.NewEntry()),
			widget.NewFormItem("Port", widget.NewEntry()),
			widget.NewFormItem("Şifre", widget.NewPasswordEntry()),
		),
		widget.NewButton("Klasör Seç", func() {
			//burada dosya falan seçme şeyi
		}),
		widget.NewLabel("Klasör Seçilmedi!"),
		widget.NewButton("Başlat", func() {
			//burada sunucu başlayacak
		}),
		widget.NewLabel("-----------------------------------"),
	)

	w.SetContent(container.NewHBox(
		durum,
		menu,
	))
	w.Resize(fyne.NewSize(600, 500))
	w.SetFixedSize(true)
	w.Show()
}

func anasayfa() {

	var w = a.NewWindow("KDAP")

	//karanlık tema
	a.Settings().SetTheme(theme.DarkTheme())

	w.SetContent(container.NewVBox(
		widget.NewLabel("Eray Mercan - 2021"),
		widget.NewButton("Sunucu", func() {
			sunucu()
		}),
		widget.NewButton("İstemci", func() {
			istemci()
		}),
	))
	w.SetOnClosed(func() {
		a.Quit()
		os.Exit(0)
	})
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}

func main() {
	anasayfa()
}
