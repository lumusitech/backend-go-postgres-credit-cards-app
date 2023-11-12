package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type alumno struct {
	legajo           int
	nombre, apellido string
}

func createDatabase() {

	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create database guarani`)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createDatabase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=guarani sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`create table alumno (legajo int, nombre text, apellido text);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`insert into alumno values (1, 'Cristina', 'Kirchner');
					  insert into alumno values (2, 'Juan Domingo', 'Perón');`)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`select * from alumno;`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var a alumno
	for rows.Next() {
		if err := rows.Scan(&a.legajo, &a.nombre, &a.apellido); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v\n", a.legajo, a.nombre, a.apellido)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Prueba de GO")
}

// ResumenConGo : Inserta un nuevo resumen
// func ResumenConGo(cliente, desde, hasta int) {

// 	//Se procede a la conexión de la base de datos
// 	fmt.Print("\n\tIngrese el nombre de la base de datos: ")
// 	var dbname string
// 	fmt.Scanf("%s", &dbname)

// 	// 	//Se crea el string con los datos para acceder a la base de datos
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	//Probar si funciona esto: SELECT 1 WHERE nrotarjeta_recibido IN(SELECT nrotarjeta FROM tarjeta WHERE estado = 'vigente')

// 	// 	//Abrimos la base de datos
// 	db, err := sql.Open("postgres", psqlInfo)
// 	defer db.Close()
// 	if err != nil {
// 		fmt.Println("\n	", err)
// 	} else {

// 		// datos para completar cabecera
// 		var nombre string
// 		var apellido string
// 		var domicilio string
// 		var nrotarjeta string
// 		var vence string
// 		var total string

// 		sqlStatement1 := `SELECT nombre FROM cliente WHERE nrocliente=$1`
// 		sqlStatement2 := `SELECT apellido FROM cliente WHERE nrocliente=$1`
// 		sqlStatement3 := `SELECT domicilio FROM cliente WHERE nrocliente=$1`
// 		sqlStatement4 := `SELECT nrotarjeta FROM tarjeta WHERE nrocliente=$1`
// 		sqlStatement5 := `SELECT validahasta FROM tarjeta WHERE nrocliente=$1`
// 		sqlStatement6 := `SELECT SUM(monto) FROM compra WHERE nrotarjeta IN(
// 						 	 SELECT nrotarjeta FROM tarjeta WHERE nrocliente=$1
// 						  )
// 						  AND fecha >= to_date($2, 'YYYYMMDD')
// 						  AND fecha <= to_date($3, 'YYYYMMDD') `

// 		row1 := db.QueryRow(sqlStatement1, cliente)
// 		row2 := db.QueryRow(sqlStatement2, cliente)
// 		row3 := db.QueryRow(sqlStatement3, cliente)
// 		row4 := db.QueryRow(sqlStatement4, cliente)
// 		row5 := db.QueryRow(sqlStatement5, cliente)
// 		row6 := db.QueryRow(sqlStatement6, cliente, desde, hasta)

// 		err1 := row1.Scan(&nombre)
// 		err2 := row2.Scan(&apellido)
// 		err3 := row3.Scan(&domicilio)
// 		err4 := row4.Scan(&nrotarjeta)
// 		err5 := row5.Scan(&vence)
// 		err6 := row6.Scan(&total)

// 		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
// 			if err == sql.ErrNoRows {
// 				fmt.Println("\n\tNo se encontraron resultados")
// 			} else {
// 				fmt.Println("\n\t	", err)
// 			}
// 		} else {
// 			fmt.Println("\t\nNombre: " + nombre)
// 			fmt.Println("\t\nApellido: " + apellido)
// 			fmt.Println("\t\nDomicilio: " + domicilio)
// 			fmt.Println("\t\nTarjeta Nro: " + nrotarjeta)
// 			fmt.Println("\t\nVencimiento: " + vence)
// 			fmt.Println("\t\nSaldo a pagar: " + total)

// 			sqlStatement7 := `
// 			INSERT INTO cabecera(nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence, total)
// 			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// 			_, err = db.Exec(sqlStatement7, nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence, total) //pensar si vence no es hasta + 10 dias
// 			if err != nil {
// 				fmt.Println("\n\t	", err)
// 			} else {
// 				// datos para completar cada detalle
// 				var nroresumen string
// 				var cantidadDeLineas int
// 				var nrolinea string
// 				var fecha string
// 				var nombrecomercio string
// 				var monto string

// 				sqlStatement8 := `SELECT nroresumen FROM cabecera WHERE nrotarjeta=$1 AND desde=$2 AND hasta=$3`
// 				sqlStatement9 := `SELECT COUNT(*) FROM compra WHERE nrotarjeta=$1 AND desde=$2 AND hasta=$3 AND pagado=$4`

// 				row8 := db.QueryRow(sqlStatement8, nrotarjeta, desde, hasta)
// 				row9 := db.QueryRow(sqlStatement9, nrotarjeta, desde, hasta, false)

// 				err8 := row8.Scan(&nroresumen)
// 				err9 := row9.Scan(&cantidadDeLineas)

// 				if err8 != nil || err9 != nil {
// 					if err == sql.ErrNoRows {
// 						fmt.Println("\n\tNo se encontraron datos para la tarjeta y período seleccionado")
// 					} else {
// 						fmt.Println("\n\t	", err)
// 					}
// 				} else {
// 					//bucle para agregar cada línea en el resumen

// 				}
// 			}

// 		}
// 	}
// }
