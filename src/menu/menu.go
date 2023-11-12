package menu

import (
	"BdNosql"
	"fmt"

	"../UI"
	"../bd"
	"../consumo"
	"../storedprocedures"
	"../tablas"
)

var (
	opcion int
)

////////// PRESENTACIÓN //////////

// Presentacion : Muestra Info del Proyecto
func Presentacion() {
	UI.Clean()
	UI.BloqueDeTexto("*", "Trabajo práctico", "Base de datos 1 - UNGS")
	UI.Esperar(2)
	UI.Clean()
	UI.BloqueDeTexto("*", "Administración para tarjetas de crédito")
	UI.Esperar(2)
	UI.Clean()
}

////////// MANEJO DE OPCIONES //////////

// Inicio : Inicia el menú de opciones
func Inicio() {
	opcion = principal()
	switch opcion {
	case 1:
		opcionBBDD()
	case 2:
		opcionTablas()
	case 3:
		opcionDatos()
	case 4:
		opcionFunciones()
	case 5:
		opcionConsumo()
	case 6:
		opcionResumen()
	case 7:
		BdNosql.OpcionNoSQL()
	case 8:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		Inicio()
	}
}

func opcionBBDD() {
	opcion = bbdd()
	switch opcion {
	case 1:
		bd.Crear()
		UI.Esperar(2)
		opcionBBDD()
	case 2:
		bd.Eliminar()
		UI.Esperar(2)
		opcionBBDD()
	case 3:
		bd.Renombrar()
		UI.Esperar(2)
		opcionBBDD()
	case 4:
		Inicio()
	case 5:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		opcionBBDD()
	}
}

func opcionTablas() {
	opcion = tablasUI()
	switch opcion {
	case 1:
		UI.Clean()
		msjAuxiliar("Creación de una tabla: ")
		tablas.Crear()
		UI.Esperar(4)
		opcionTablas()
	case 2:
		UI.Clean()
		msjAuxiliar("Eliminación de una tabla: ")
		tablas.Eliminar()
		UI.Esperar(2)
		opcionTablas()
	case 3:
		UI.Clean()
		msjAuxiliar("Cambiar nombre de una tabla: ")
		tablas.Renombrar()
		UI.Esperar(2)
		opcionTablas()
	case 4:
		UI.Clean()
		msjAuxiliar("Insertar PK a una tabla: ")
		tablas.PK()
		UI.Esperar(2)
		opcionTablas()
	case 5:
		UI.Clean()
		msjAuxiliar("Insertar FK a una tabla: ")
		tablas.FK()
		UI.Esperar(10)
		opcionTablas()
	case 6:
		msjAuxiliar("Creación de todas las tablas del TP: ")
		tablas.CrearTodas()
		UI.Esperar(2)
		opcionTablas()
	case 7:
		msjAuxiliar("Inserción de todas las Primary keys a las tablas del TP: ")
		tablas.AgregarTodasLasPks()
		UI.Esperar(2)
		opcionTablas()

	case 8:
		msjAuxiliar("Inserción de todas las Foreign keys a las tablas del TP: ")
		tablas.AgregarTodasLasFks()
		UI.Esperar(2)
		opcionTablas()
	case 9:
		msjAuxiliar("Eliminación de todas las Primary keys y Foreign keys de las tablas del TP: ")
		tablas.BorrarPKsFKs()
		UI.Esperar(2)
		opcionTablas()
	case 10:
		Inicio()
	case 11:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		opcionTablas()
	}
}

func opcionDatos() {
	opcion = datosUI()
	switch opcion {
	case 1:
		tablas.PoblarTodas()
		UI.Esperar(2)
		opcionDatos()
	case 2:
		tablas.ResetearTodas()
		UI.Esperar(2)
		opcionDatos()
	case 3:
		Inicio()
	case 4:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		opcionDatos()
	}
}

func opcionFunciones() {
	msjAuxiliar("Inserción de stored procedures y Triggers")
	storedprocedures.AutorizacionCompra()
	storedprocedures.GenerarResumen()
	storedprocedures.Alertas()
	UI.Esperar(2)
	Inicio()
}

func opcionConsumo() {
	opcion = consumoUI()
	switch opcion {
	case 1:
		UI.Clean()
		msjAuxiliar("Consumo nuevo: ")
		consumo.Nuevo()
		opcionConsumo()
	case 2:
		consumo.Resetear()
		UI.Esperar(2)
		opcionConsumo()
	case 3:
		Inicio()
	case 4:
		salir()
	default:
		msjAuxiliar("Debe ingresar una opción correcta")
		opcionConsumo()
	}
}

func opcionResumen() {
	UI.Clean()
	msjAuxiliar("Generación de resumen:")
	storedprocedures.Resumen()
	UI.Esperar(2)
	Inicio()
}

////////// INTERFAZ DE USUARIO //////////

// Muestra el menú principal de la aplicación
func principal() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Menú principal")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Gestionar base de datos
	2 --> Gestionar tablas
	3 --> Gestionar datos
	4 --> Insertar funciones
	5 --> Gestionar consumos
	6 --> Generar resumen
	7 --> Gestionar NoSQL
	-----------------------------
	8 --> Salir
	`)
	fmt.Print("\n	Opción: ")
	fmt.Scanf("%d\n", &eleccion)
	return eleccion
}

// Muestra el submenú para gestionar la BBDD
func bbdd() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Gestionar base de datos")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Crear una base de datos
	2 --> Eliminar una base de datos
	3 --> Renombrar una base de datos
	----------------------------------
	4 --> Volver al menú principal
	5 --> Salir
	`)
	fmt.Print("\n	Opción: ")
	fmt.Scanf("%d\n", &eleccion)
	return eleccion
}

// Muestra el submenú para gestionar tablas
func tablasUI() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Gestionar tablas")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Crear una tabla
	2 --> Eliminar una tabla
	3 --> Renombrar una tabla
	-----------------------------------------
	4 --> Agregar PK a tabla
	5 --> Agregar FK a tabla
	-----------------------------------------
	6 --> Generar tablas del TP
	7 --> Insertar todas las PKs
	8 --> Insertar todas las FKs
	9 --> Eliminar todas las PKs y Fks
	-----------------------------------------
	10 --> Volver al menú principal
	11 --> Salir
	`)
	fmt.Print("\n	Opción: ")
	fmt.Scanf("%d\n", &eleccion)
	return eleccion
}

func datosUI() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Gestionar consumos")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Insertar todos los datos a las tablas del TP
	2 --> Resetear todos los datos a las tablas del TP
	--------------------------------------------------
	3 --> Volver al menú principal
	4 --> Salir
	`)
	fmt.Print("\n	Opción: ")
	fmt.Scanf("%d\n", &eleccion)
	return eleccion
}

func consumoUI() int {
	var eleccion int
	UI.Clean()
	UI.BloqueDeTexto("*", "Gestionar consumos")
	fmt.Print(`
	Seleccione la opción numérica deseada:

	1 --> Realizar un consumo
	2 --> Resetear los consumos realizados
	---------------------------------------
	3 --> Volver al menú principal
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
