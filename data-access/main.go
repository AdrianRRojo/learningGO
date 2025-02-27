package main

import (

    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)
var db *sql.DB

type Album struct {
    ID int64
    Title string
    Artist string
    Price float32

}



func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    albums,err := albumsByArtist("John Coltrane")
    
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)
    
    idAlb, err := albumByID(3)
    if err != nil{

        log.Fatal(err)
    }
    fmt.Printf("ID FOUND: %v\n", idAlb)

    newAlbum, err := addAlbum(
        Album{
            Title: "Man on the Moon", 
            Artist: "Kid Cudi",
            Price: 10.99,
        })

    if err != nil{
        log.Fatal(err)
    }

    fmt.Printf("New Album just dropped homie: %v\n", newAlbum)

}


func albumsByArtist(name string) ([]Album, error){

    var albums []Album

    rows,err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil{
        return nil, fmt.Errorf("albumsByArtist %q: %v",name,err)
    }

    defer rows.Close()

    for rows.Next() {
        var alb Album

        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err!= nil{
            return nil, fmt.Errorf("albumsByArtist %q: %v",name,err)
        }
       albums = append(albums,alb)

    }

    if err != nil{
        return nil, fmt.Errorf("albumsByArtist %q: %v",name,err)
    }
    return albums, nil
}

func albumByID(id int) (Album, error){
    var album Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    
    if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err!=nil{
        if err != nil{
            return album, fmt.Errorf("albumByID %q: %v", id, err)
        }
        return album, fmt.Errorf("albumByID %q: %v", id, err)
    }
    return album, nil

}

func addAlbum(alb Album) (int64, error){
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
    
    if err != nil{
        return 0, fmt.Errorf("Could not add album %v: ", err)
    }

    id, err := result.LastInsertId()


    if err != nil{
        return 0, fmt.Errorf("Could not add album %v: ", err)
    }
    
    return id, nil


}
