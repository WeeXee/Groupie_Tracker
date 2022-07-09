package main

import (
	"errors"
	"fmt"
	"groupie/function"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var HTML function.MyArtist

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", Home)
	server.HandleFunc("/artist", Redirect)
	server.HandleFunc("/404", Redirect)
	//server.HandleFunc("/info", Info)
	fmt.Println("server listening on http://localhost:8000")
	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	server.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.ListenAndServe(":8000", server)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(w, nil)
	err := function.GetData()
	if err != nil {
		errors.New("Error")
	}
	search := r.FormValue("artist")
	if search == "" && len(function.ArtistsFull) != 0 {
		Search(search)
	}
	return
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("artist")
	searchResults := Search(input)
	if searchResults != nil {
		tmpl := template.Must(template.ParseFiles("html/artist.html"))
		_ = tmpl.Execute(w, searchResults[0])
	} else {
		tmpl := template.Must(template.ParseFiles("html/404.html"))
		_ = tmpl.Execute(w, nil)
	}
}

func Search(search string) []function.MyArtistFull {
	if search == "" {
		return function.ArtistsFull
	}
	art, err := ConverterStructToString()
	if err != nil {
		errors.New("Error by converter")
	}
	var search_artist []function.MyArtistFull

	for i, artist := range art {
		lower_band := strings.ToLower(artist)
		for i_name, l_name := range []byte(lower_band) {
			lower_search := strings.ToLower(search)
			if lower_search[0] == l_name {
				lenght_name := 0
				indx := i_name
				for _, l := range []byte(lower_search) {
					if l == lower_band[indx] {
						if indx+1 == len(lower_band) {
							break
						}
						indx++
						lenght_name++
					} else {
						break
					}
				}
				if len(search) == lenght_name {
					band, _ := function.GetFullDataById(i + 1)
					search_artist = append(search_artist, band)
					break
				}
			}
		}

	}
	return search_artist
}

func ConverterStructToString() ([]string, error) {
	var data []string
	for i := 1; i <= len(function.Artists); i++ {
		artist, err1 := function.GetArtistByID(i)
		locations, err2 := function.GetLocationByID(i)
		date, err3 := function.GetDateByID(i)
		if err1 != nil || err2 != nil || err3 != nil {
			return data, errors.New("Error by converter")
		}

		str := artist.Name + " "
		for _, member := range artist.Members {
			str += member + " "
		}
		str += strconv.Itoa(artist.CreationDate) + " "
		str += artist.FirstAlbum + " "
		for _, location := range locations.Locations {
			str += location + " "
		}
		for _, d := range date.Dates {
			str += d + " "
		}
		data = append(data, str)
	}
	return data, nil
}
