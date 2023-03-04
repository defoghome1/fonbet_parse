package scrap

import (
	"awesomeProject1/bd"
	"awesomeProject1/secret"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)




type All struct{
	//Result string `json:"result"`
	//Request string `json:"request"`
	Place string `json:"place"`
	//Lang string `json:"lang"`
	Events []event `json:"events"`
}

type event struct {
	Id int `json:"id"`
	//Number int `json:"number"`

	SportName string `json:"skName"`
	Markets []markets `json:"markets"`
	LigaName string `json:"competitionCaption"`
	CompetitionId int `json:"competitionId"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	//AllFactorsCount int `json:"allFactorsCount"`
}

type markets struct {
	MarketID string `json:"marketId"`
	Ident string `json:"ident"`
	Rows []rows `json:"rows"`
}

type rows struct {
	Cells []cells `json:"cells"`
}

type cells struct {
	Caption string `json:"caption"`
	Param string `json:"paramText"`
}

func GetJson (url string,target interface{}) error{
	res,err := http.Get(url)
	if err!=nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}


type Base struct {
	Id int
	LigaName string
	Team1 string
	Team2 string
	Total int
	Time time.Time
}

func FonbetParse(k float64){
	now := time.Now()
	URL := secret.URLFonbet
	var m All

	err := GetJson(URL,&m)
	if err!=nil {
		log.Fatal(err)
	} else {
		for _,i:= range m.Events {
			if i.SportName == "Баскетбол"  {
				for _, j := range i.Markets {
					if j.Ident == "Totals"  {
						for _,a := range j.Rows{
							if bd.CheckId(i.Id) == false && a.Cells[0].Caption == "Тотал матча %P" {
								total, err := strconv.ParseFloat(a.Cells[0].Param,64)
								if err!=nil {
									log.Fatal(err)
								}
								bd.WriteToBase(i.Id, i.LigaName, i.Team1, i.Team2, total, now)
								fmt.Printf("В базу данных добавился матч %v - %v с тоталом %v\n",i.Team1,i.Team2,total)
								//fmt.Printf("Название Лиги: %v\nКоманда №1: %v\nКоманда №2: %v\nТотал: %v\nКогда записался тотал: %v\n\n",i.LigaName,i.Team1,i.Team2,total,now.Format("02-01-2006 15:04:05"))
							} else if a.Cells[0].Caption == "Тотал матча %P"{
								total, err := strconv.ParseFloat(a.Cells[0].Param,64)
								if err!=nil {
									log.Fatal(err)
								}
								bd.CheckTotal(total,i.Id,k,i.CompetitionId)
							}
						}
					}
				}
			}
		}
	}


}