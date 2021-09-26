package main

import (
	"html/template"
	"net/http"
	"os"
)


func main() {

	// Make HTTP GET request
	http.HandleFunc("/", homePage)
	http.HandleFunc("/achievement", achievementPage)
	http.HandleFunc("/experience", experiencePage)
	http.HandleFunc("/project", projectPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	portNumber, ok := os.LookupEnv("PORT")
	if !ok {
		portNumber = "8080"
	} else {
	}
	http.ListenAndServe(":"+portNumber, nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"static/html/MyPage.html",
		"static/html/base_template.html",
	}

	template, _ := template.ParseFiles(files[0], files[1])
	template.Execute(writer, nil)
}

func achievementPage(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"static/html/AchievementPage.html",
		"static/html/base_template.html",
	}

	template, _ := template.ParseFiles(files[0], files[1])
	template.Execute(writer, nil)
}

func experiencePage(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"static/html/ExperiencePage.html",
		"static/html/base_template.html",
	}

	template, _ := template.ParseFiles(files[0], files[1])
	template.Execute(writer, nil)
}

func projectPage(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"static/html/ProjectPage.html",
		"static/html/base_template.html",
	}

	template, _ := template.ParseFiles(files[0], files[1])
	template.Execute(writer, nil)
}
