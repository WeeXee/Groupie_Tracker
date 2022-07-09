package function

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetArtists() error {
	resp, err := http.Get(URL + "/artists")
	if err != nil {
		return errors.New("Error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

func GetDates() error {
	resp, err := http.Get(URL + "/dates")
	if err != nil {
		return errors.New("Error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &Dates)
	return nil
}

func GetLocations() error {
	resp, err := http.Get(URL + "/locations")
	if err != nil {
		return errors.New("Error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &Locations)
	return nil
}

func GetRelations() error {
	resp, err := http.Get(URL + "/relation")
	if err != nil {
		return errors.New("Error by get")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &Relations)
	return nil
}

func GetData() error {
	if len(ArtistsFull) != 0 {
		return nil
	}
	err1 := GetArtists()
	err2 := GetLocations()
	err3 := GetDates()
	err4 := GetRelations()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return errors.New("Error by get data artists, locations, dates")
	}
	for i := range Artists {
		var tmpl MyArtistFull
		tmpl.ID = i + 1
		tmpl.Image = Artists[i].Image
		tmpl.Name = Artists[i].Name
		tmpl.Members = Artists[i].Members
		tmpl.CreationDate = Artists[i].CreationDate
		tmpl.FirstAlbum = Artists[i].FirstAlbum
		tmpl.Locations = Locations.Index[i].Locations
		tmpl.ConcertDates = Dates.Index[i].Dates
		tmpl.DatesLocations = Relations.Index[i].DatesLocations
		ArtistsFull = append(ArtistsFull, tmpl)
	}
	return nil
}

func GetArtistByID(id int) (MyArtist, error) {
	for _, artist := range Artists {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtist{}, errors.New("Not found")
}

func GetDateByID(id int) (MyDate, error) {
	for _, date := range Dates.Index {
		if date.ID == id {
			return date, nil
		}
	}
	return MyDate{}, errors.New("Not found")
}

func GetLocationByID(id int) (MyLocation, error) {
	for _, location := range Locations.Index {
		if location.ID == id {
			return location, nil
		}
	}
	return MyLocation{}, errors.New("Not found")
}

func GetRelationByID(id int) (MyRelation, error) {
	for _, relation := range Relations.Index {
		if relation.ID == id {
			return relation, nil
		}
	}
	return MyRelation{}, errors.New("Not found")
}

func GetFullDataById(id int) (MyArtistFull, error) {
	for _, artist := range ArtistsFull {
		if artist.ID == id {
			return artist, nil
		}
	}
	return MyArtistFull{}, errors.New("Not found")
}
