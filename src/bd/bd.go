package bd

import (
	"database/sql"
	"fmt"

	// driver de postgresql
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
)

//Crear : Crea una BBDD
func Crear() {
	//Se crea el string con los datos para acceder a postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	// Entramos a postgres
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Se procede a la creación de la base de datos
	fmt.Print("\n\tIngrese el nombre de la base de datos que desea crear: ")
	var dbname string
	fmt.Scanf("%s", &dbname)

	_, err = db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		//Agregar logica para dar msj si la base de datos ya existe
		fmt.Println("\n	", err)
	} else {

		//Procedemos a hacerle ping para asegurarnos que todo salió bien

		//Se crea el string con los datos para acceder a la base de daos creada
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		//Abrimos la base de datos
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		//Le hacemos ping
		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n\tLa base de datos %s, se creó exitosamente!", dbname)
	}

}

// Eliminar : Elimina una base de datos
func Eliminar() {

	//Se crea el string con los datos para acceder a postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	// Entramos a postgres
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Se solicita el nombre de la base de datos a eliminar
	fmt.Print("\n\tIngrese el nombre de la base de datos que desea eliminar: ")
	var dbname string
	fmt.Scanf("%s", &dbname)

	//Se procede a eliminar
	_, err = db.Exec("DROP DATABASE " + dbname)
	if err != nil {
		// Agregar lógica para dar msj si la base de datos no existe
		fmt.Println("\n	", err)
	} else {
		fmt.Println("\n\tSe eliminó la base de datos " + dbname)
	}
}

// Renombrar : Renombra una base de datos
func Renombrar() {

	//Se crea el string con los datos para acceder a postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	// Entramos a postgres
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Se solicita el nombre de la base de datos a renombrar
	fmt.Print("\n\tIngrese el nombre de la base de datos que desea renombrar: ")
	var dbnameOld string
	fmt.Scanf("%s", &dbnameOld)

	//Se solicita el nuevo nombre
	fmt.Print("\n\tIngrese el nuevo nombre para la base de datos: ")
	var dbnameNew string
	fmt.Scanf("%s", &dbnameNew)

	//Se procede a renombrar
	_, err = db.Exec("ALTER DATABASE " + dbnameOld + " RENAME TO " + dbnameNew)
	if err != nil {
		// Agregar lógica para dar msj si la base de datos no existe
		fmt.Println("\n	", err)
	} else {
		fmt.Println("\n\tSe realizó el cambio " + dbnameOld + " --> " + dbnameNew)
	}
}

// Conectar : conecta a una bbd y devuelve el puntero a dicha base
func Conectar() *sql.DB {
	//Se procede a la conexión de la base de datos
	fmt.Print("\n\tIngrese el nombre de la base de datos: ")
	var dbname string
	fmt.Scanf("%s", &dbname)

	//Se crea el string con los datos para acceder a la base de datos
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Abrimos la base de datos
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("\n\t	", err)
	}
	return db

}
