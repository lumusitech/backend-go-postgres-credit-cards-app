package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "el pass que tenga"
	dbname   = "nombre de la base de datos"
)

//Si no se necesita password sacar ese dato de todos los lados donde se lo menciona
//En windows 10, podemos cambiar el archivo /c/Archivos de programas/Postgresql/11/data/pg_hba.conf
//Al final del archivo se ve esto:
//# TYPE  DATABASE        USER            ADDRESS                 METHOD
//# IPv4 local connections:
//host    all             all             127.0.0.1/32            md5    ----> ahí poner trust
//# IPv6 local connections:
//host    all             all             ::1/128                 md5    ----> ahí poner trust
//Cambiar el method a trust (en lugar de md5), solo en esas líneas. Listo,
//ahora ya no será necesario introducir password para conectarse

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
