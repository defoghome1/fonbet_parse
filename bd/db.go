package bd

import (
	"awesomeProject1/secret"
	"awesomeProject1/send-message"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math"
	"time"
)

type Match struct {
	id    int
	liga  string
	Team1 string
	Team2 string
	Total float64
	time  time.Time

}

func WriteToBase(Id int, Liga string, Team1 string, Team2 string, total float64, w_time time.Time) {
	connStr := secret.Connstr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_ , err = db.Exec("insert into matches (m_no, liga, team1, team2, total, w_time) values ($1, $2, $3, $4, $5, $6)",
		Id, Liga, Team1,Team2,total,w_time)
	if err != nil{
		panic(err)
	}
}

func ReadFromBase(){
	connStr := secret.Connstr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from matches")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	matches := []Match{}

	for rows.Next() {
		m:=Match{}
		err := rows.Scan(&m.id, &m.liga, &m.Team1, &m.Team2, &m.Total,&m.time)
		if err != nil{
			fmt.Println(err)
			continue
		}
		matches = append (matches, m)
	}
	for _, m := range matches{
		fmt.Println(m.id, m.liga, m.Team1, m.Team2, m.Total,m.time)
	}
}

func CheckId (Id int) bool {
	connStr := secret.Connstr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row , err := db.Query("select * from matches where m_no = $1",Id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return row.Next()

}

func CheckTotal (Total float64,Id int, K float64, CID int){
	connStr := secret.Connstr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m:=Match{}

	row , err := db.Query("select * from matches where m_no = $1",Id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&m.id, &m.liga, &m.Team1, &m.Team2, &m.Total, &m.time)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}
	if (math.Abs(Total - m.Total) >= K){
		fmt.Printf("В матче %v - %v тотал изменился %v -> %v",m.Team1,m.Team2,m.Total,Total)
		send_message.SendMessage(m.Team1,m.Team2,m.Total,Total,CID,Id)
		_, err := db.Exec("update matches set total = $1 where m_no = $2",Total,Id)
		if err != nil{
			log.Fatal(err)
		} else {
			fmt.Printf("\tСсылка на матч : https://www.fon.bet/live/basketball/%v/%v\n",CID,Id)
		}
	}
}