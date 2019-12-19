package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/allegro/bigcache"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "admin"
	password = "admin123"
	dbname   = "admin"
)

// Ejemplo de uso de caché para una petición get que no cambiará
var duracionCache = 10 * time.Minute
var cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(duracionCache))

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	// log.Fatal("Bye.")

	entry, _ := cache.Get("allDataDummy")
	// fmt.Println("ENTRY:" + string(entry))

	// La primer petición tarda alrededor de ~50ms
	// con caché las peticiones subsecuentes tardan ~2ms
	if entry != nil {
		print("From cache!")
		w.WriteHeader(http.StatusOK)
		w.Write(entry)
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Un arreglo para almacenar todos los elementos
	ds := []DatoDummy{}

	// Una estructura para almacenar cada elemento
	d := DatoDummy{}

	// Query SQL
	sqlStatement := "SELECT id, number, texto FROM datos_dummy"
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&d.ID, &d.Number, &d.Texto); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		// Agregamos el elemento al arreglo
		ds = append(ds, d)
		// fmt.Println(d.ID, d.Texto, d.Number)
	default:
		panic(err)
	}

	defer db.Close()

	// Convertimos de Struct a JSON(bytes)
	var jsonData []byte
	jsonData, err = json.Marshal(ds)
	if err != nil {
		// log.Println(err)
	}
	fmt.Println(string(jsonData))

	// Guardamos el JSON en el caché
	cache.Set("allDataDummy", jsonData)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
