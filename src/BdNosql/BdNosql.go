package BdNosql

import (
	"UI"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	bolt "github.com/coreos/bbolt"
)

var db *bolt.DB

type Cliente struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjeta struct {
	Nrotarjeta   string
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra float64
	Estado       string
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compra struct {
	Nrooperacion int
	Nrotarjeta   string
	Nrocomercio  int
	Fecha        time.Time
	Monto        float64
	Pagado       bool
}

func CrearBdNosql(nombre string) *bolt.DB {
	bd, error := bolt.Open(nombre, 0600, nil)
	if error != nil {
		log.Fatal(error)
	}
	return bd
}

func Escritura(bd *bolt.DB, nombreBucket string, llave []byte, valor []byte) error {
	//se abre una transacción de Escritura.
	transcc, error := bd.Begin(true) //me devuelve una transaccion.
	if error != nil {
		return error
	}
	defer transcc.Rollback() //metodo que controla consistencia y cierra transaccion.

	//creo el backet con el nombre que me pasaron
	backet, _ := transcc.CreateBucketIfNotExists([]byte(nombreBucket))

	error = backet.Put(llave, valor) //se coloca el valor asociado.
	if error != nil {
		return error
	}
	//se cierra la transacción de Escritura.
	if error := transcc.Commit(); error != nil { //se guardan los datos
		return error
	}
	return nil
}

//funcion que lee un solo valor debe recibir como key algo como una pk.
func LecturaUnica(bd *bolt.DB, nombreBucket string, llave []byte) ([]byte, error) {

	var buffer []byte //slice de bytes para valor de retorno

	// abre una transacción de LecturaUnica con la func View que espera como param
	// otra func y controla esa transaccion.(rollback, commit).
	error := bd.View(func(transcc *bolt.Tx) error {

		//se pasa el bck a slice, es decir, busca el bck del nombre ej alumno.
		backet := transcc.Bucket([]byte(nombreBucket))

		// a ese bck le pido el valor ej legajo y la guardo para devolverlo.
		buffer = backet.Get(llave)

		return nil
	})
	return buffer, error
}

func PoblarDatos(bd *bolt.DB) {
	//los clientes para agregar a la backet
	julia := Cliente{4781, "Julia", "Montiel", "Paraguay 778", "4657-1866"}
	data, error := json.Marshal(julia)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "cliente", []byte(strconv.Itoa(julia.Nrocliente)), data)

	juan := Cliente{4975, "Juan", "Melgarejo", "Chacabuco 465", "4738-4823"}
	data, error = json.Marshal(juan)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "cliente", []byte(strconv.Itoa(juan.Nrocliente)), data)

	romina := Cliente{4807, "Romina", "Guerreiro", "Lacroze 1533", "1534888006"}
	data, error = json.Marshal(romina)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "cliente", []byte(strconv.Itoa(romina.Nrocliente)), data)

	//las tarjetas para agregar
	tarjeta1 := Tarjeta{"5175847480130436", julia.Nrocliente, "201612", "202612", "0489", 27342, "vigente"}
	data, error = json.Marshal(tarjeta1)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "tarjeta", []byte(tarjeta1.Nrotarjeta), data)

	tarjeta2 := Tarjeta{"5239202881623321", juan.Nrocliente, "201210", "202211", "6310", 20000, "vigente"}
	data, error = json.Marshal(tarjeta2)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "tarjeta", []byte(tarjeta2.Nrotarjeta), data)

	tarjeta3 := Tarjeta{"4754982137770169", romina.Nrocliente, "201204", "202205", "0778", 10750, "vigente"}
	data, error = json.Marshal(tarjeta3)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "tarjeta", []byte(tarjeta3.Nrotarjeta), data)

	//los comercios para insertar
	calzados := Comercio{59335, "Vallejo Calzados", "Callao 6643", "b6740fda", "4638-9813"}
	data, error = json.Marshal(calzados)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "comercio", []byte(strconv.Itoa(calzados.Nrocomercio)), data)

	moda := Comercio{14910, "Full Moda", "Alem 3062", "b1669fda", "9351-6043"}
	data, error = json.Marshal(moda)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "comercio", []byte(strconv.Itoa(moda.Nrocomercio)), data)

	tech := Comercio{26782, "Mundo Tech", "Indios 4289", "b1862fda", "4642-8374"}
	data, error = json.Marshal(tech)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "comercio", []byte(strconv.Itoa(tech.Nrocomercio)), data)

	//Listado de Compras
	fecha := time.Now()
	comprauno := Compra{45042, tarjeta1.Nrotarjeta, tech.Nrocomercio, fecha, 18000, true}
	data, error = json.Marshal(comprauno)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "compra", []byte(strconv.Itoa(comprauno.Nrooperacion)), data)

	fecha = time.Now()
	comprados := Compra{50467, tarjeta2.Nrotarjeta, calzados.Nrocomercio, fecha, 8500, false}
	data, error = json.Marshal(comprados)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "compra", []byte(strconv.Itoa(comprados.Nrooperacion)), data)

	fecha = time.Now()
	compratres := Compra{60987, tarjeta3.Nrotarjeta, moda.Nrocomercio, fecha, 9870, true}
	data, error = json.Marshal(compratres)
	if error != nil {
		log.Fatal(error)
	}
	Escritura(bd, "compra", []byte(strconv.Itoa(compratres.Nrooperacion)), data)
}

//Mostrar Datos en Pantalla
func ImprimirDatos(bd *bolt.DB) {

	fmt.Printf("Los Clientes que han sido Ingresados son:\n\n")
	valor, _ := LecturaUnica(bd, "cliente", []byte(strconv.Itoa(4781)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "cliente", []byte(strconv.Itoa(4975)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "cliente", []byte(strconv.Itoa(4807)))
	fmt.Printf("%s\n\n", valor)

	fmt.Printf("Las Tarjetas Ingresadas son:\n\n")
	valor, _ = LecturaUnica(bd, "tarjeta", []byte("5175847480130436"))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "tarjeta", []byte("5239202881623321"))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "tarjeta", []byte("4754982137770169"))
	fmt.Printf("%s\n\n", valor)

	fmt.Printf("Los Comercios Ingresados son:\n\n")
	valor, _ = LecturaUnica(bd, "comercio", []byte(strconv.Itoa(59335)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "comercio", []byte(strconv.Itoa(14910)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "comercio", []byte(strconv.Itoa(26782)))
	fmt.Printf("%s\n\n", valor)

	fmt.Printf("Las Compras Registradas son:\n\n")
	valor, _ = LecturaUnica(bd, "compra", []byte(strconv.Itoa(45042)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "compra", []byte(strconv.Itoa(50467)))
	fmt.Printf("%s\n", valor)
	valor, _ = LecturaUnica(bd, "compra", []byte(strconv.Itoa(60987)))
	fmt.Printf("%s\n\n", valor)
}

func OpcionNoSQL() {
	opcion := noSQLUI()

	switch opcion {
	case 1:
		db = CrearBdNosql("Base")
		UI.Esperar(2)
		OpcionNoSQL()
	case 2:
		PoblarDatos(db)
		UI.Esperar(2)
		OpcionNoSQL()
	case 3:
		ImprimirDatos(db)
		var q string
		fmt.Print("Presiones enter para terminar...")
		fmt.Scanf("%d", &q)
		UI.Esperar(2)
		OpcionNoSQL()
	case 4:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		OpcionNoSQL()
	}
}

func noSQLUI() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Gestionar base de datos no relacional")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Crear la base de datos
	2 --> Escribir información en ella
	3 --> Leer la información de la base de datos
	---------------------------------------
	4 --> Salir
	`)
	fmt.Print("\n	Opción: ")
	fmt.Scanf("%d\n", &eleccion)
	return eleccion
}

func msjAuxiliar(msj string) {
	UI.Clean()
	UI.BloqueDeTexto("*", msj)
	//UI.Esperar(2)
}

// Muestra mensaje de salida
func salir() {
	//msjAuxiliar("Gracias por utilizar la aplicación")
	UI.Clean()
	return
}
