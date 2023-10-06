package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
)

const ADDRESS = ":8080"

type Days = []struct {
	Name   string
	Dishes []string
}

type Cards = map[string]struct {
	Name string
	Img  string
	Desc []string
}

type PageTemplate struct {
	Current int
	Days    []string
	Cards   Cards
}

func MakeTemplate(dayindex int) ([]byte, error) {
	jsonbytes, err := os.ReadFile("menu-days.json")
	if err != nil {
		return nil, err
	}
	days := Days{}
	if err = json.Unmarshal(jsonbytes, &days); err != nil {
		return nil, err
	}

	if dayindex < 0 || dayindex >= len(days) {
		return nil, errors.New("invalid day index")
	}

	jsonbytes, err = os.ReadFile("menu-cards.json")
	if err != nil {
		return nil, err
	}
	cards := make(Cards)
	if err = json.Unmarshal(jsonbytes, &cards); err != nil {
		return nil, err
	}

	pagetemp := PageTemplate{
		Days:  []string{},
		Cards: make(Cards),
	}
	for _, day := range days {
		pagetemp.Days = append(pagetemp.Days, day.Name)
	}
	for _, dishkey := range days[dayindex].Dishes {
		if card, exists := cards[dishkey]; exists {
			pagetemp.Cards[dishkey] = card
		}
	}

	t, err := template.New("index.html").ParseFiles("wwwroot/index.html")
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err = t.Execute(&buffer, pagetemp); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)

		if r.URL.Path != "/" {
			http.ServeFile(w, r, "wwwroot/"+r.URL.Path)
			return
		}

		dayindex := 0
		switch r.URL.Query().Get("day") {
		case "kedd":
			dayindex = 1
		case "szerda":
			dayindex = 2
		case "csutortok":
			dayindex = 3
		case "pentek":
			dayindex = 4
		case "szombat":
			dayindex = 5
		case "vasarnap":
			dayindex = 6
		}

		bytes, err := MakeTemplate(dayindex)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}

		if _, err := w.Write(bytes); err != nil {
			log.Println(err)
		}
	})

	log.Println("Started server on", ADDRESS)
	if err := http.ListenAndServe(ADDRESS, nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
