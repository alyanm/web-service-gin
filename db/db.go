package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB is the database handle
var DB *sql.DB

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetAlbums() ([]Album, error) {
	rows, err := DB.Query("SELECT * FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []Album{}
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}

	return albums, nil
}

func GetAlbumByID(id string) (Album, error) {
	row := DB.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	var alb Album
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		return Album{}, err
	}

	return alb, nil
}

func AddAlbum(alb Album) (int64, error) {
	res, err := DB.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}	

func UpdateAlbumByID(id string, alb Album) (int64, error) {
	res, err := DB.Exec("UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteAlbumByID(id string) (int64, error) {
	res, err := DB.Exec("DELETE FROM albums WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}