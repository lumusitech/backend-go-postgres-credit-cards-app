package main

//declaracion de importaciones

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//creacion de struct para realizar una consulta.

type Cliente struct {
	legajo           int
	nombre, apellido string
}

//conectarse a una Base de Datos.

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

//creacion de tablas e insertar valores.

func main() {
	createDatabase()

	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=Trajeta-UNGS sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//creacion de tablas
	_, err = db.Exec(`create table Cliente (nrocliente int, nombre text, apellido text, domicilio text, telefono char (12));
                    create table tarjeta (nrotarjeta int, validadesde char (6), validahasta char (6), codseguridad char (4), limitecompra decimal (8,2), estado char (10));
                    create table Comercio (nrocomercio int, nombre text, apellido text, domicilio text, telefono char (12));
                    create table Compra (nrooperacion int, nrotarjeta char (16), nrocomercio int, fecha timestamp, monto decimal (7,2), pagado boolean);
                    create table Rechazo (nrorechazo int, nrotarjeta char (16), nrocomercio int, fecha timestamp, monto decimal (7,2), motivo text );
                    create table Cierre (año int, mes int, terminacion int, fecahainicio date, fechacierre date, fechavto date);
                    create table Cabecera (nroresumen int, nombre text, apellido text, domicilio text, nrotarjeta char(16), desde date, hasta date, vence date, total float);
                    create table Detalle (nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto float);
                    create table Alerta (nroalerta int, nrotarjeta char (16), fecha timestamp, nrorechazo int, codalerta int, descripcion text);
					create table Consumo (nrotarjeta char (16), codseguridad char (4), nrocomercio int, monto float);
						--creacion de PKs--
					alter table Cliente add constraint Cliente_pk primary key (nrocliente);
					alter table tarjeta add constraint tarjeta_pk primary key (nrotarjeta);
					alter table Comercio add constraint Comercio_pk primary key (nrocomercio);
					alter table Compra add constraint Compra_pk primary key (nrooperacion);
					alter table Rechazo add constraint Rechazo_pk primary key (nrorechazo);
					alter table Cierre add constraint Cierre_pk primary key (año, mes, terminacion);
					alter table Cabecera add constraint Cabecera_pk primary key (nroresumen);
					alter table Detalle add constraint Detalle_pk primary key (nroresumen, nrolinea);
					alter table Alerta add constraint Alerta_pk primary key (nroalerta);
					alter table Consumo add constraint Consumo_pk primary key (nrotarjeta);
							--creacion de FKs
					alter table tarjeta add constraint tarjeta_nrocliente_fk foreing key (nrocliente) references Cliente (nrocliente);
					alter table Compra add constraint Compra_nrotarjeta_fk foreing key (nrotarjeta) references tarjeta (nrotarjeta);
					alter table Compra add constraint Compra_nrocomercio_fk foreing key (nrocomercio) references Comercio (nrocomercio);
					alter table Rechazo add constraint Rechazo_nrotarjeta_fk foreing key (nrotarjeta) references tarjeta (nrotarjeta);
					alter table Rechazo add constraint Rechazo_nrocomercio_fk foreing key (nrocomercio) references Comercio (nrocomercio);
					alter table Cabecera add constraint Cabecera_nrotarjeta_fk foreing key (nrotarjeta) references tarjeta (nrotarjeta);
					alter table Detalle add constraint Detalle_nroresumen_fk foreing key (nroresumen) references Cabecera (nroresumen);
					alter table Alerta add constraint Alerta_nrotarjeta_fk foreing key (nrotarjeta) references tarjeta (nrotarjeta);
					alter table Alerta add constraint Alerta_nrorechazo_fk foreing key (nrorechazo) references Rechazo (nrorechazo);
					alter table Consumo add constraint Consumo_nrotarjeta_fk foreing key (nrotarjeta) references tarjeta (nrotarjeta);
					alter table Consumo add constraint Consumo_nrocomercio_fk foreing key (nrocomercio) references Comercio (nrocomercio);`)
	if err != nil {
		log.Fatal(err)
	}
	//insert de clientes
	_, err = db.Exec(`insert into Cliente values (2681532370671820, 'Paola', 'Melgar', 'Calle Martín Peschell Nº 1150', 4666-8146);
					insert into Cliente values (4975165789802240, 'Juan', 'Melgarejo', 'Chacabuco 465', '4738-4823');
					insert into Cliente values (9536965128195590, 'Pablo', 'Carrasco', 'Suipacha  5850', '4729-2370');
					insert into Cliente values (7833870731668110, 'Esteban', 'Janiot', '25 de Mayo 637', '4754-4840');
					insert into Cliente values (8356725894325500, 'Martin', 'Jaimez', 'Emancipador 10344', '4768-5290');
					insert into Cliente values (8310410213988620, 'Mariana', 'Herrera', 'Parage Fernandez 5115', '4720-8525');
					insert into Cliente values (3646599595246680, 'Carla', 'Guillermaz', 'Santiago del Estero N°2067', '4757-7529');
					insert into Cliente values (4807436291637920, 'Romina', 'Guerriero', 'Lacroze 1533', '1534888006');
					insert into Cliente values (9118448214526070, 'Miguel', 'Salas', 'Charlone Nº 9170', '4844-4253');
					insert into Cliente values (3582633399586100, 'Claudio', 'Rizzo', 'Las delicias 1007', '4722-4214');
					insert into Cliente values (1978979267369580, 'Adrian', 'Peralta', 'Profesor Aguer 6207', '4769-3639');
					insert into Cliente values (1572475819595620, 'Silvia', 'Palavecino', 'Benito Perez Galdos 9720', '1568663634');
					insert into Cliente values (5657290738364560, 'Alexis', 'Ovejero', 'Gobernador Castro 107', '4764-1612');
					insert into Cliente values (4448692263743990, 'Micaela', 'Otero', 'Manuel Medina 255', '1567323708');
					insert into Cliente values (7210481742111880, 'Jorge', 'Nievas', 'Pasaje de Octubre 532', '4767-0638');
					insert into Cliente values (2254417777146620, 'Cesar', 'Cores', 'Rio Negro N°8879', '4849-1817');
					insert into Cliente values (6435561126795810, 'Sergio', 'Chazarreta', 'A Alcorta 660', '5290-4027');
					insert into Cliente values (9112843311144330, 'Mariano', 'Cerrudo', 'San Martin Nº 8901', '4720-1788');
					insert into Cliente values (7833870731668110, 'Pablo', 'Cernadas', 'Asamblea 1290', '4844-1507');
					insert into Cliente values (4781954512188850, 'Julia', 'Montiel', 'Paraguay 778', '4657-1866');
								--insert de Comercios
					insert into Comercio values (23553, 'Cetrogar', 'Montiel', 'b1770fda', '4660-5670');
					insert into Comercio values (15491, 'Gasmarket', 'Cuba 658', 'b2740fda', '3881-4501');
					insert into Comercio values (80176, 'Pintureria Silva', 'Chacabuco 1924', 'b1605fda', '6541-2245');
					insert into Comercio values (42525 , 'Chikle Kids', 'San Martin 6239 ', 'b6430fda', '3833-4235');
					insert into Comercio values (80395, 'Colt Jeans', 'Rivadavia 910', 'b2752fda', '5647-8967');
					insert into Comercio values (47446, 'Sommier Center', 'Pacheco 3278', 'b8504fda', '5865-5612');
					insert into Comercio values (59335, 'Vallejo Calzados', 'Callao 6643', 'b1712fda', '4638-9813');
					insert into Comercio values (32112, 'Sastreria Valentino', 'Bolivar 571', 'b6740fda', '6445-0738');
					insert into Comercio values (71556, 'Carabus Viajes', 'Sarmiento 2981', 'b7130fda', '8431-2971');
					insert into Comercio values (82496, 'Optica del Carmen', 'Santa Rosa 1563', 'b1684fda', '5467-8734');
					insert into Comercio values (14910, 'Full Moda', 'Alem 3062', 'b1669fda', '9351-6043');
					insert into Comercio values (58700, 'Super Show Deportes', '9 de Julio 3823', 'b1824fda', '3541-4293');
					insert into Comercio values (18159, 'El Rey del Colchon', 'Lima 5838', 'b1615fda', '7865-6101');
					insert into Comercio values (26782, 'Mundo Tech', 'Indios 4289', 'b1862fda', '4642-8374');
					insert into Comercio values (48456, 'Sanitarios Saltanovich', 'Alvear 1529', 'b6034fda', '4523-8769');
					insert into Comercio values (32700, 'America Muebles', 'Av Colon 3852', 'b6000fda', '4671-5571');
					insert into Comercio values (99000, 'Masterphone', 'Boucherville 1050', 'b2805fda', '4783-9078');
					insert into Comercio values (82025, 'Tecno Vision', 'Lavalle 2341', 'b1657fda', '4240-4013');
					insert into Comercio values (41735, 'Bebelandia', 'Dorrego 5643', 'b1854fda', '8976-4567');
					insert into Comercio values (53960, 'Libreria Lerma', 'Uruguay 1024', 'b6700fda', '4954-4614');`)

	if err != nil {
		log.Fatal(err)
	}

	//Insert de tarjetas ----------------->  NO ESTÁN LO NROS DE CLIENTES A LOS QUE PERTENECE CADA TARJETA!

	_, err = db.Exec(`insert into tarjeta values('5239202881623321', 201210, '202211', 0631, 20000, 'vigente');
					insert into tarjeta values('4262319016061870', 200910, '201911', 0372, 50000, 'vigente');
					insert into tarjeta values('5477699847594811', 201201, '201802', 0456, 15000,'anulada');
 					insert into tarjeta values('5266053569153247', 201312, '201912', 0332, 13000, 'vigente');
					insert into tarjeta values('4147000443655564', 201505, '201910', 0554, 25000,'suspendida');
					insert into tarjeta values('4637226106139300', 201203, '201802', 0432, 12500, 'anulada'); 
					insert into tarjeta values('4306431000959569', 201805, '202307', 0554, 28350, 'suspendida');
					insert into tarjeta values('4754982137770169', 201204, '202205', 0778, 10750, 'vigente');
					insert into tarjeta values('3714573296598442', 201507, '202507', 0741, 27245, 'vigente');
					insert into tarjeta values('5356431514233846', 200801, '201801', 0542, 10000, 'anulada'); 
					insert into tarjeta values('4384218611107331', 201703, '202703', 0003, 23270, 'vigente');
					insert into tarjeta values('3748158326942841', 201504, '202504', 0456, 47750, 'vigente');
					insert into tarjeta values('5238726338571998', 201901, '202402', 0753, 35000, 'suspendida');
					insert into tarjeta values('5149464739074432', 201106, '202106', 0168, 18440, 'vigente');
					insert into tarjeta values('3728436353847153', 201608, '202207', 0468, 15780, 'vigente');
					insert into tarjeta values('3452416342729618', 201209, '202209', 0015, 26000, 'suspendida');
					insert into tarjeta values('5292603592286865', 201902, '202901', 0187, 27870, 'vigente');
					insert into tarjeta values('4841267341470649', 201712, '201901', 0418, 58750, 'anulada');
					insert into tarjeta values('5344070958087352', 201005, '201705', 0373, 29457, 'anulada');
					insert into tarjeta values('5101810313449976', 201201, '202202', 0439, 87545, 'vigente');
					insert into tarjeta values('3731941594695921', 200508, '201507', 0887, 50000, 'anulada');
					insert into tarjeta values('5175847480130436', 201612, '202612', 0489, 27342, 'vigente');`)

	if err != nil {
		fmt.Println(`Error al insertar los datos`)
	}

	//realizar consultas.

	rows, err := db.Query(`select * from Cliente`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var a Cliente
	for rows.Next() {
		if err := rows.Scan(&a.legajo, &a.nombre, &a.apellido); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v\n", a.legajo, a.nombre, a.apellido)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
