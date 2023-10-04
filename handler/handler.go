package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/yessikaplata/rps-game/rps"
	"github.com/yessikaplata/rps-game/services"
)

const (
	errorLoadTemplate = "An error occurred while loading the template."
	errorParseForm    = "An error occurred while processing the form."
	templatePath      = "template/"
	baseTemplate      = "base.html"
)

type Player struct {
	Name string
}

var player Player
var rpsGame = rps.NewRPSGame(services.RandomGenerator{})

func Index(w http.ResponseWriter, r *http.Request) {
	resetGame()
	renderTemplate(w, baseTemplate, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	resetGame()
	renderTemplate(w, baseTemplate, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, errorParseForm, http.StatusInternalServerError)
			return
		}
		player.Name = r.FormValue("name")
		renderTemplate(w, baseTemplate, "game.html", player)
	}
	resetGame()
	http.Redirect(w, r, "/new", http.StatusFound)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("option"))
	resultRound := rpsGame.PlayRound(playerChoice)
	jsonResult, error := json.MarshalIndent(resultRound, "", "    ")
	if error != nil {
		log.Println(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTemplate, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	page = fmt.Sprintf("%s%s", templatePath, page)
	base = fmt.Sprintf("%s%s", templatePath, baseTemplate)
	templateIndex := template.Must(template.ParseFiles(base, page))
	err := templateIndex.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, errorLoadTemplate, http.StatusInternalServerError)
		return
	}
}

func resetGame() {
	player.Name = ""
}
