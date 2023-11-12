package consumo

import (
	"UI"
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

////////////////////////////////////////////////////////////////////////

// Resetear : Borra todos los consumos realizados
func Resetear() {
	//Se procede a la conexión de la base de datos
	fmt.Print("\n\tIngrese el nombre de la base de datos: ")
	var dbname string
	fmt.Scanf("%s", &dbname)

	//Se crea el string con los datos para acceder a la base de datos
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Abrimos la base de datos
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		fmt.Println("\n	", err)
	} else { //Si no hay error en la conexión se procede
		_, err = db.Exec(`DELETE FROM consumo;`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tSe han eliminado todos los consumos realizados")
		}
	}
}

////////////////////////////////////////////////////////////////////////

// Nuevo : Realiza un consumo nuevo a partir de datos ofrecidos
func Nuevo() {

	//Se intenta la conexión con postgresql y la bbdd
	fmt.Print("\n\tIngrese el nombre de la base de datos: ")
	var dbname string
	fmt.Scanf("%s", &dbname)

	if dbname == "" {
		dbname = "sinNombre"
	}

	//Se crea el string con los datos para acceder a la base de datos
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Abrimos la base de datos
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		fmt.Println("\n	", err)
	} else {
		err := db.Ping() //Para evitar errores, se prueba la conexión
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			tarjeta := elegirTarjeta(db)
			if tarjeta != "" {
				codigo := ingresarCodigo()
				if codigo != 0 {
					comercio := elegirComercio(db)
					if comercio != 0 {
						monto := ingresarMonto()
						if monto != 0 {
							UI.Clean()
							UI.BloqueDeTexto("*", "Consumo: ")
							fmt.Println("\n\tNro de tarjeta elegido: ", tarjeta)
							fmt.Println("\n\tCódigo de la tarjeta elegida: ", codigo)
							fmt.Println("\n\tComercio elegido: ", comercio)
							fmt.Println("\n\tMonto de la compra: $", monto)

							msj := insertarConsumo(db, tarjeta, codigo, comercio, monto)
							fmt.Println("\n\t", msj)
						} else {
							fmt.Println("\n\tError al ingresar el monto!")
						}
					} else {
						fmt.Println("\n\tError al seleccionar el comercio!")
					}
				} else {
					fmt.Println("\n\tError al ingresar el código!")
				}
			} else {
				fmt.Println("\n\tError al seleccionar la tarjeta!")
			}

		}
	}

	//Usado para salir cuando el usuario presione enter
	fmt.Print("\n\n\tPresione enter para continuar...")
	var q int
	fmt.Scanf("%d\n", &q)

}

///////////////////////////////////////////////////////////

func elegirTarjeta(db *sql.DB) string {
	UI.Clean()
	type Tarjeta struct {
		nro    string
		codigo int
		limite float32
		estado string
	}
	rows, err := db.Query(`SELECT nrotarjeta, codseguridad, limitecompra, estado FROM tarjeta`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var contador int
	contador = 1
	var lista []Tarjeta
	UI.BloqueDeTexto("*", "Seleccionar la tarjeta para realizar el consumo")
	for rows.Next() {
		tarjeta := Tarjeta{}
		err = rows.Scan(&tarjeta.nro, &tarjeta.codigo, &tarjeta.limite, &tarjeta.estado)
		if err != nil {
			panic(err)
		}
		lista = append(lista, tarjeta)
		fmt.Printf("%d --> Nro: %v | Código: %v | Límite: %v | Estado: %v\n",
			contador, lista[contador-1].nro, lista[contador-1].codigo,
			lista[contador-1].limite, lista[contador-1].estado)
		contador++
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("\n\t", err)
	}
	//Seleccionar la tarjeta para realizar el consumo
	var nrotarjetaElegida int
	fmt.Print("\n\tOpción: ")
	fmt.Scanf("%d\n", &nrotarjetaElegida)
	if nrotarjetaElegida <= 0 || nrotarjetaElegida > contador { //Si no es uno de los números de la lista
		return ""
	}

	return lista[nrotarjetaElegida-1].nro
}

////////////////////////////////////////////////////////////////////////////

func ingresarCodigo() int {
	UI.Clean()
	fmt.Print("\n\tIngrese el código de seguridad de su tarjeta: ")
	var codigo int
	_, err := fmt.Scanf("%v", &codigo)
	//Comprobación mínima de código
	if err != nil {
		fmt.Println("\n\tDato ingresado inválido, se esperan un número")
		return 0
	}
	//Se valida la cantidad de dígitos
	contador := 1
	num := codigo
	for num/10 > 0 {
		num = num / 10
		contador++
	}
	if contador != 4 {
		fmt.Println("\n\tDato ingresado inválido, se espera un número de 4 dígitos")
		return 0
	}

	return codigo
}

//////////////////////////////////////////////////////////////////////////////////////////

func elegirComercio(db *sql.DB) int {
	UI.Clean()
	type Comercio struct {
		nrocomercio  int
		nombre       string
		domicilio    string
		codigopostal string
	}
	rows, err := db.Query(`SELECT nrocomercio, nombre, domicilio, codigopostal FROM comercio`)
	if err != nil {
		fmt.Println("\n\t", err)
		return 0
	}
	defer rows.Close()
	contador := 1
	var lista []Comercio
	UI.BloqueDeTexto("*", "Seleccionar el comercio donde realizará el consumo")
	for rows.Next() {
		comercio := Comercio{}
		err = rows.Scan(&comercio.nrocomercio, &comercio.nombre, &comercio.domicilio, &comercio.codigopostal)
		if err != nil {
			fmt.Println("\n\t", err)
			return 0
		}
		lista = append(lista, comercio)
		fmt.Printf("%d --> Nro: %v | Nombre: %v | Domicilio: %v | Código postal: %v\n",
			contador, lista[contador-1].nrocomercio, lista[contador-1].nombre,
			lista[contador-1].domicilio, lista[contador-1].codigopostal)
		contador++
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("\n\t", err)
		return 0
	}
	//Seleccionar el comercio para realizar el consumo
	var nrocomercioElegido int
	fmt.Print("\n\tOpción: ")
	fmt.Scanf("%d\n", &nrocomercioElegido)

	if nrocomercioElegido <= 0 || nrocomercioElegido > contador { //Si no es uno de los números de la lista
		return 0
	}

	return lista[nrocomercioElegido-1].nrocomercio
}

//////////////////////////////////////////////////////////////////////////

func ingresarMonto() float32 {
	UI.Clean()
	fmt.Print("\n\tIngrese el monto: $")
	var monto float32
	_, err := fmt.Scanf("%f", &monto)
	if err != nil {
		fmt.Println("\n\tDato ingresado incorrecto, se admiten solo montos enteros o decimales")
		return 0
	}
	return monto
}

//////////////////////////////////////////////////////////////////////////

func insertarConsumo(db *sql.DB, tarjeta string, codigo, comercio int, monto float32) string {
	var msj string

	//Vemos el nro de la última compra
	var nrocompra int
	rows, err := db.Query("SELECT MAX(nrooperacion) FROM compra")
	if rows.Next() {
		rows.Scan(&nrocompra)
		fmt.Println("\n\tNro. inicial de compra: ", nrocompra)
	}
	if err != nil {
		fmt.Println("\n\t", err)
		msj = ""
	}

	//Se inserta el consumo que será procesado por un procedimiento almacenado disparado por un trigger
	sql := `INSERT INTO consumo (nrotarjeta, codseguridad, nrocomercio, monto) VALUES($1, $2, $3, $4)`
	_, err2 := db.Exec(sql, tarjeta, codigo, comercio, monto)
	if err2 != nil {
		fmt.Println("\n\t", err2)
		msj = ""
	}

	//Vemos el nro de la última compra nuevamente
	var nrocompraPost int
	rows3, err3 := db.Query("SELECT MAX(nrooperacion) FROM compra")
	if rows3.Next() {
		rows3.Scan(&nrocompraPost)
		fmt.Println("\n\tNro. final de compra: ", nrocompraPost)
	}
	if err3 != nil {
		fmt.Println("\n\t", err3)
		msj = ""
	}

	//Vemos si se insertó una nueva compra o fue rechazada
	if nrocompra == nrocompraPost { //si son iguales es que no se insertó en venta, sino en rechazo

		//vemos el último rechazo para informar el detalle
		nrorechazo := 0
		sqlrechazo := `SELECT MAX(nrorechazo) FROM rechazo`
		rowRechazo := db.QueryRow(sqlrechazo)
		errRechazo := rowRechazo.Scan(&nrorechazo)
		if errRechazo != nil || nrorechazo == 0 {
			fmt.Println("\n\t", errRechazo)
		} else {
			var detalle string
			sqldetalle := `SELECT motivo FROM rechazo WHERE nrorechazo=$1`
			rowDetalle := db.QueryRow(sqldetalle, nrorechazo)
			errDetalle := rowDetalle.Scan(&detalle)
			if errDetalle != nil {
				fmt.Println("\n\t", errDetalle)
			} else {
				fmt.Println("\n\tLa compra fue rechazada con el motivo: ", detalle)
			}
		}

	} else {
		msj = "La compra fue aprobada!!! "
	}

	return msj
}
