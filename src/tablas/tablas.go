package tablas

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

// CrearTodas : crea todas las tablas del TP
func CrearTodas() {

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
	} else {
		_, err = db.Exec(`create table cliente (nrocliente int,	nombre text, apellido text,	domicilio text,	telefono char(12));
		create table tarjeta (nrotarjeta char(16), nrocliente int, validadesde char(6), validahasta char(6), codseguridad char(4), limitecompra decimal(7,2), estado char(10));
		create table comercio (nrocomercio int, nombre text, domicilio text, codigopostal char(8), telefono char(12));
		create table compra (nrooperacion SERIAL NOT NULL, nrotarjeta char(16),	nrocomercio int, fecha timestamp, monto decimal(7,2), pagado boolean);
		create table rechazo (nrorechazo SERIAL NOT NULL, nrotarjeta char(16), nrocomercio int,	fecha timestamp, monto decimal(7,2), motivo text);
		create table cierre (año int, mes int, terminacion int, fechainicio date, fechacierre date,	fechavto date);
		create table cabecera (nroresumen SERIAL NOT NULL, nombre text,	apellido text, domicilio text, nrotarjeta char(16),	desde date,	hasta date, vence date,	total decimal(7,2));
		create table detalle (nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto decimal(7,2));
		create table alerta (nroalerta SERIAL NOT NULL, nrotarjeta char(16), fecha timestamp, nrorechazo int, codalerta int, descripcion text);
		create table consumo (nroconsumo SERIAL NOT NULL, nrotarjeta char(16), codseguridad char(4), nrocomercio int, monto decimal(7,2))`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tTodas las tablas fueron creadas exitosamente!")
		}
	}
}

////////////////////////////////////////////////////////////////////////

// AgregarTodasLasPks : Agrega todas las primary keys del TP
func AgregarTodasLasPks() {
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
	} else {
		_, err = db.Exec(`alter table cliente add constraint cliente_pk primary key (nrocliente);
		alter table tarjeta add constraint tarjeta_pk primary key (nrotarjeta);
		alter table comercio add constraint comercio_pk primary key (nrocomercio);
		alter table compra add constraint compra_pk primary key (nrooperacion);
		alter table rechazo add constraint rechazo_pk primary key (nrorechazo);
		alter table cierre add constraint cierre_pk primary key (año, mes, terminacion);
		alter table cabecera add constraint cabecera_pk primary key (nroresumen);
		alter table detalle add constraint detalle_pk primary key (nroresumen, nrolinea);
		alter table alerta add constraint alerta_pk primary key (nroalerta);
		alter table consumo add constraint consumo_pk primary key (nroconsumo);`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tTodas las Primary keys fueron agregadas exitosamente!")
		}
	}
}

////////////////////////////////////////////////////////////////////////

// AgregarTodasLasFks : Agrega todas las Foreign keys del TP
func AgregarTodasLasFks() {
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
	} else {
		_, err = db.Exec(`alter table tarjeta add constraint tarjeta_nrocliente_fk foreign key (nrocliente) references cliente (nrocliente);
		alter table compra add constraint compra_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
		alter table compra add constraint compra_nrocomercio_fk foreign key (nrocomercio) references comercio (nrocomercio);
		alter table rechazo add constraint rechazo_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
		alter table rechazo add constraint rechazo_nrocomercio_fk foreign key (nrocomercio) references comercio (nrocomercio);
		alter table cabecera add constraint cabecera_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
		alter table detalle add constraint detalle_nroresumen_fk foreign key (nroresumen) references cabecera (nroresumen);
		alter table alerta add constraint alerta_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
		alter table alerta add constraint alerta_nrorechazo_fk foreign key (nrorechazo) references rechazo (nrorechazo);
		`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tTodas las Foreign keys fueron agregadas exitosamente!")
		}
	}
}

////////////////////////////////////////////////////////////////////////

// BorrarPKsFKs : Borra todas las Primary y foreign keys del TP
func BorrarPKsFKs() {
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
	} else {
		_, err = db.Exec(`ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_nrocliente_fk;
		ALTER TABLE compra DROP CONSTRAINT compra_nrotarjeta_fk;
		ALTER TABLE compra DROP CONSTRAINT compra_nrocomercio_fk;
		ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrotarjeta_fk;
		ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrocomercio_fk;
		ALTER TABLE cabecera DROP CONSTRAINT cabecera_nrotarjeta_fk;
		ALTER TABLE detalle DROP CONSTRAINT detalle_nroresumen_fk;
		ALTER TABLE alerta DROP CONSTRAINT alerta_nrotarjeta_fk;
		ALTER TABLE alerta DROP CONSTRAINT alerta_nrorechazo_fk;`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			_, err = db.Exec(`ALTER TABLE cliente DROP CONSTRAINT cliente_pk;
			ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_pk;
			ALTER TABLE comercio DROP CONSTRAINT comercio_pk;
			ALTER TABLE compra DROP CONSTRAINT compra_pk;
			ALTER TABLE rechazo DROP CONSTRAINT rechazo_pk;
			ALTER TABLE cierre DROP CONSTRAINT cierre_pk;
			ALTER TABLE cabecera DROP CONSTRAINT cabecera_pk;
			ALTER TABLE detalle DROP CONSTRAINT detalle_pk;
			ALTER TABLE alerta DROP CONSTRAINT alerta_pk;
			ALTER TABLE consumo DROP CONSTRAINT consumo_pk;`)
			if err != nil {
				fmt.Println("\n	", err)
			} else {
				fmt.Println("\n\tTodas las Primary keys y Foreign keys fueron eliminadas exitosamente!")
			}
		}

	}
}

////////////////////////////////////////////////////////////////////////

// PoblarTodas : inserta todos los datos a las tablas del TP
func PoblarTodas() {
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
	} else {
		_, err = db.Exec(`insert into cliente values (2681, 'Paola', 'Melgar', 'Calle Martín Peschell Nº 1150', '4738-8146');
			insert into cliente values (4975, 'Juan', 'Melgarejo', 'Chacabuco 465', '4738-4823');
			insert into cliente values (9536, 'Pablo', 'Carrasco', 'Suipacha  5850', '4729-2370');
			insert into cliente values (7833, 'Esteban', 'Janiot', '25 de Mayo 637', '4754-4840');
			insert into cliente values (8356, 'Martin', 'Jaimez', 'Emancipador 10344', '4768-5290');
			insert into cliente values (8310, 'Mariana', 'Herrera', 'Parage Fernandez 5115', '4720-8525');
			insert into cliente values (3646, 'Carla', 'Guillermaz', 'Santiago del Estero N°2067', '4757-7529');
			insert into cliente values (4807, 'Romina', 'Guerriero', 'Lacroze 1533', '1534888006');
			insert into cliente values (9118, 'Miguel', 'Salas', 'Charlone Nº 9170', '4844-4253');
			insert into cliente values (3582, 'Claudio', 'Rizzo', 'Las delicias 1007', '4722-4214');
			insert into cliente values (1978, 'Adrian', 'Peralta', 'Profesor Aguer 6207', '4769-3639');
			insert into cliente values (1572, 'Silvia', 'Palavecino', 'Benito Perez Galdos 9720', '1568663634');
			insert into cliente values (5657, 'Alexis', 'Ovejero', 'Gobernador Castro 107', '4764-1612');
			insert into cliente values (4448, 'Micaela', 'Otero', 'Manuel Medina 255', '1567323708');
			insert into cliente values (7210, 'Jorge', 'Nievas', 'Pasaje de Octubre 532', '4767-0638');
			insert into cliente values (2254, 'Cesar', 'Cores', 'Rio Negro N°8879', '4849-1817');
			insert into cliente values (6435, 'Sergio', 'Chazarreta', 'A Alcorta 660', '5290-4027');
			insert into cliente values (9112, 'Mariano', 'Cerrudo', 'San Martin Nº 8901', '4720-1788');
			insert into cliente values (7837, 'Pablo', 'Cernadas', 'Asamblea 1290', '4844-1507');
			insert into cliente values (4781, 'Julia', 'Montiel', 'Paraguay 778', '4657-1866');
						--insert de Comercios
			insert into comercio values (23553, 'Cetrogar', 'Montiel', 'b1770fda', '4660-5670');
			insert into comercio values (15491, 'Gasmarket', 'Cuba 658', 'b1770fda', '3881-4501');
			insert into comercio values (80176, 'Pintureria Silva', 'Chacabuco 1924', 'b1605fda', '6541-2245');
			insert into comercio values (42525 , 'Chikle Kids', 'San Martin 6239 ', 'b6430fda', '3833-4235');
			insert into comercio values (80395, 'Colt Jeans', 'Rivadavia 910', 'b2752fda', '5647-8967');
			insert into comercio values (47446, 'Sommier Center', 'Pacheco 3278', 'b8504fda', '5865-5612');
			insert into comercio values (59335, 'Vallejo Calzados', 'Callao 6643', 'b1712fda', '4638-9813');
			insert into comercio values (32112, 'Sastreria Valentino', 'Bolivar 571', 'b6740fda', '6445-0738');
			insert into comercio values (71556, 'Carabus Viajes', 'Sarmiento 2981', 'b7130fda', '8431-2971');
			insert into comercio values (82496, 'Optica del Carmen', 'Santa Rosa 1563', 'b1684fda', '5467-8734');
			insert into comercio values (14910, 'Full Moda', 'Alem 3062', 'b1669fda', '9351-6043');
			insert into comercio values (58700, 'Super Show Deportes', '9 de Julio 3823', 'b1824fda', '3541-4293');
			insert into comercio values (18159, 'El Rey del Colchon', 'Lima 5838', 'b1615fda', '7865-6101');
			insert into comercio values (26782, 'Mundo Tech', 'Indios 4289', 'b1862fda', '4642-8374');
			insert into comercio values (48456, 'Sanitarios Saltanovich', 'Alvear 1529', 'b6034fda', '4523-8769');
			insert into comercio values (32700, 'America Muebles', 'Av Colon 3852', 'b6000fda', '4671-5571');
			insert into comercio values (99000, 'Masterphone', 'Boucherville 1050', 'b2805fda', '4783-9078');
			insert into comercio values (82025, 'Tecno Vision', 'Lavalle 2341', 'b1657fda', '4240-4013');
			insert into comercio values (41735, 'Bebelandia', 'Dorrego 5643', 'b6700fda', '8976-4567');
			insert into comercio values (53960, 'Libreria Lerma', 'Uruguay 1024', 'b6700fda', '4954-4614');
						--insert de Tarjetas
			insert into tarjeta values('5239202881623321', 4975, '201210', '202211', 6310, 20000, 'vigente');
			insert into tarjeta values('4262319016061870', 4975, '200910', '201911', 1372, 50000, 'vigente');
			insert into tarjeta values('5477699847594811', 9536, '201201', '201802', 1456, 15000,'vigente');
			insert into tarjeta values('5266053569153247', 7833, '201312', '201912', 1332, 13000, 'vigente');
			insert into tarjeta values('4147000443655564', 8356, '201505', '201910', 1554, 25000,'suspendida');
			insert into tarjeta values('4637226106139300', 8310, '201203', '201802', 1432, 12500, 'anulada'); 
			insert into tarjeta values('4306431000959569', 3646, '201805', '202307', 1554, 28350, 'suspendida');
			insert into tarjeta values('4754982137770169', 4807, '201204', '202205', 1778, 10750, 'vigente');
			insert into tarjeta values('3714573296598442', 9118, '201507', '202507', 1741, 27245, 'vigente');
			insert into tarjeta values('5356431514233846', 3582, '200801', '201801', 1542, 10000, 'anulada'); 
			insert into tarjeta values('4384218611107331', 1978, '201703', '202703', 1003, 23270, 'vigente');
			insert into tarjeta values('3748158326942841', 1572, '201504', '202504', 1456, 47750, 'vigente');
			insert into tarjeta values('5238726338571998', 5657, '201901', '202402', 1753, 35000, 'suspendida');
			insert into tarjeta values('5149464739074432', 4448, '201106', '202106', 1168, 18440, 'vigente');
			insert into tarjeta values('3728436353847153', 7210, '201608', '202207', 1468, 15780, 'vigente');
			insert into tarjeta values('3452416342729618', 2254, '201209', '202209', 1015, 26000, 'suspendida');
			insert into tarjeta values('5292603592286865', 6435, '201902', '202901', 1187, 27870, 'vigente');
			insert into tarjeta values('4841267341470649', 9112, '201712', '201901', 1418, 58750, 'anulada');
			insert into tarjeta values('5344070958087352', 7837, '201005', '201705', 1373, 29457, 'anulada');
			insert into tarjeta values('5101810313449976', 7837, '201201', '202202', 1439, 87545, 'vigente');
			insert into tarjeta values('3731941594695921', 4781, '200508', '201507', 1887, 50000, 'anulada');
			insert into tarjeta values('5175847480130436', 4781, '201612', '202612', 1489, 27342, 'vigente');
						--insert de cierres
			insert into cierre values (2019, 1, 0, '2019-1-7', '2019-1-10', '2019-1-17');
			insert into cierre values (2019, 2, 0, '2019-2-7', '2019-2-10', '2019-2-17');
			insert into cierre values (2019, 3, 0, '2019-3-7', '2019-3-10', '2019-3-17');
			insert into cierre values (2019, 4, 0, '2019-4-7', '2019-4-10', '2019-4-17');
			insert into cierre values (2019, 5, 0, '2019-5-7', '2019-5-10', '2019-5-17');
			insert into cierre values (2019, 6, 0, '2019-6-7', '2019-6-10', '2019-6-17');
			insert into cierre values (2019, 7, 0, '2019-7-7', '2019-7-10', '2019-7-17');
			insert into cierre values (2019, 8, 0, '2019-8-7', '2019-8-10', '2019-8-17');
			insert into cierre values (2019, 9, 0, '2019-9-7', '2019-9-10', '2019-9-17');
			insert into cierre values (2019, 10, 0, '2019-10-7', '2019-10-10', '2019-10-17');
			insert into cierre values (2019, 11, 0, '2019-11-7', '2019-11-10', '2019-11-17');
			insert into cierre values (2019, 12, 0, '2019-12-7', '2019-12-10', '2019-12-17');
			insert into cierre values (2019, 1, 1, '2019-1-8', '2019-1-11', '2019-1-18');
			insert into cierre values (2019, 2, 1, '2019-2-8', '2019-2-11', '2019-2-18');
			insert into cierre values (2019, 3, 1, '2019-3-8', '2019-3-11', '2019-3-18');
			insert into cierre values (2019, 4, 1, '2019-4-8', '2019-4-11', '2019-4-18');
			insert into cierre values (2019, 5, 1, '2019-5-8', '2019-5-11', '2019-5-18');
			insert into cierre values (2019, 6, 1, '2019-6-8', '2019-6-11', '2019-6-18');
			insert into cierre values (2019, 7, 1, '2019-7-8', '2019-7-11', '2019-7-18');
			insert into cierre values (2019, 8, 1, '2019-8-8', '2019-8-11', '2019-8-18');
			insert into cierre values (2019, 9, 1, '2019-9-8', '2019-9-11', '2019-9-18');
			insert into cierre values (2019, 10, 1, '2019-10-8', '2019-10-11', '2019-10-18');
			insert into cierre values (2019, 11, 1, '2019-11-8', '2019-11-11', '2019-11-18');
			insert into cierre values (2019, 12, 1, '2019-12-8', '2019-12-11', '2019-12-18');
			insert into cierre values (2019, 1, 2, '2019-1-9', '2019-1-12', '2019-1-19');
			insert into cierre values (2019, 2, 2, '2019-2-9', '2019-2-12', '2019-2-19');
			insert into cierre values (2019, 3, 2, '2019-3-9', '2019-3-12', '2019-3-19');
			insert into cierre values (2019, 4, 2, '2019-4-9', '2019-4-12', '2019-4-19');
			insert into cierre values (2019, 5, 2, '2019-5-9', '2019-5-12', '2019-5-19');
			insert into cierre values (2019, 6, 2, '2019-6-9', '2019-6-12', '2019-6-19');
			insert into cierre values (2019, 7, 2, '2019-7-9', '2019-7-12', '2019-7-19');
			insert into cierre values (2019, 8, 2, '2019-8-9', '2019-8-12', '2019-8-19');
			insert into cierre values (2019, 9, 2, '2019-9-9', '2019-9-12', '2019-9-19');
			insert into cierre values (2019, 10, 2, '2019-10-9', '2019-10-12', '2019-10-19');
			insert into cierre values (2019, 11, 2, '2019-11-9', '2019-11-12', '2019-11-19');
			insert into cierre values (2019, 12, 2, '2019-12-9', '2019-12-12', '2019-12-19');
			insert into cierre values (2019, 1, 3, '2019-1-10', '2019-1-13', '2019-1-20');
			insert into cierre values (2019, 2, 3, '2019-2-10', '2019-2-13', '2019-2-20');
			insert into cierre values (2019, 3, 3, '2019-3-10', '2019-3-13', '2019-3-20');
			insert into cierre values (2019, 4, 3, '2019-4-10', '2019-4-13', '2019-4-20');
			insert into cierre values (2019, 5, 3, '2019-5-10', '2019-5-13', '2019-5-20');
			insert into cierre values (2019, 6, 3, '2019-6-10', '2019-6-13', '2019-6-20');
			insert into cierre values (2019, 7, 3, '2019-7-10', '2019-7-13', '2019-7-20');
			insert into cierre values (2019, 8, 3, '2019-8-10', '2019-8-13', '2019-8-20');
			insert into cierre values (2019, 9, 3, '2019-9-10', '2019-9-13', '2019-9-20');
			insert into cierre values (2019, 10, 3, '2019-10-10', '2019-10-13', '2019-10-20');
			insert into cierre values (2019, 11, 3, '2019-11-10', '2019-11-13', '2019-11-20');
			insert into cierre values (2019, 12, 3, '2019-12-10', '2019-12-13', '2019-12-20');
			insert into cierre values (2019, 1, 4, '2019-1-11', '2019-1-14', '2019-1-21');
			insert into cierre values (2019, 2, 4, '2019-2-11', '2019-2-14', '2019-2-21');
			insert into cierre values (2019, 3, 4, '2019-3-11', '2019-3-14', '2019-3-21');
			insert into cierre values (2019, 4, 4, '2019-4-11', '2019-4-14', '2019-4-21');
			insert into cierre values (2019, 5, 4, '2019-5-11', '2019-5-14', '2019-5-21');
			insert into cierre values (2019, 6, 4, '2019-6-11', '2019-6-14', '2019-6-21');
			insert into cierre values (2019, 7, 4, '2019-7-11', '2019-7-14', '2019-7-21');
			insert into cierre values (2019, 8, 4, '2019-8-11', '2019-8-14', '2019-8-21');
			insert into cierre values (2019, 9, 4, '2019-9-11', '2019-9-14', '2019-9-21');
			insert into cierre values (2019, 10, 4, '2019-10-11', '2019-10-14', '2019-10-21');
			insert into cierre values (2019, 11, 4, '2019-11-11', '2019-11-14', '2019-11-21');
			insert into cierre values (2019, 12, 4, '2019-12-11', '2019-12-14', '2019-12-21');
			insert into cierre values (2019, 1, 5, '2019-1-12', '2019-1-15', '2019-1-22');
			insert into cierre values (2019, 2, 5, '2019-2-12', '2019-2-15', '2019-2-22');
			insert into cierre values (2019, 3, 5, '2019-3-12', '2019-3-15', '2019-3-22');
			insert into cierre values (2019, 4, 5, '2019-4-12', '2019-4-15', '2019-4-22');
			insert into cierre values (2019, 5, 5, '2019-5-12', '2019-5-15', '2019-5-22');
			insert into cierre values (2019, 6, 5, '2019-6-12', '2019-6-15', '2019-6-22');
			insert into cierre values (2019, 7, 5, '2019-7-12', '2019-7-15', '2019-7-22');
			insert into cierre values (2019, 8, 5, '2019-8-12', '2019-8-15', '2019-8-22');
			insert into cierre values (2019, 9, 5, '2019-9-12', '2019-9-15', '2019-9-22');
			insert into cierre values (2019, 10, 5, '2019-10-12', '2019-10-15', '2019-10-22');
			insert into cierre values (2019, 11, 5, '2019-11-12', '2019-11-15', '2019-11-22');
			insert into cierre values (2019, 12, 5, '2019-12-12', '2019-12-15', '2019-12-22');
			insert into cierre values (2019, 1, 6, '2019-1-13', '2019-1-16', '2019-1-23');
			insert into cierre values (2019, 2, 6, '2019-2-13', '2019-2-16', '2019-2-23');
			insert into cierre values (2019, 3, 6, '2019-3-13', '2019-3-16', '2019-3-23');
			insert into cierre values (2019, 4, 6, '2019-4-13', '2019-4-16', '2019-4-23');
			insert into cierre values (2019, 5, 6, '2019-5-13', '2019-5-16', '2019-5-23');
			insert into cierre values (2019, 6, 6, '2019-6-13', '2019-6-16', '2019-6-23');
			insert into cierre values (2019, 7, 6, '2019-7-13', '2019-7-16', '2019-7-23');
			insert into cierre values (2019, 8, 6, '2019-8-13', '2019-8-16', '2019-8-23');
			insert into cierre values (2019, 9, 6, '2019-9-13', '2019-9-16', '2019-9-23');
			insert into cierre values (2019, 10, 6, '2019-10-13', '2019-10-16', '2019-10-23');
			insert into cierre values (2019, 11, 6, '2019-11-13', '2019-11-16', '2019-11-23');
			insert into cierre values (2019, 12, 6, '2019-12-13', '2019-12-16', '2019-12-23');
			insert into cierre values (2019, 1, 7, '2019-1-14', '2019-1-17', '2019-1-24');
			insert into cierre values (2019, 2, 7, '2019-2-14', '2019-2-17', '2019-2-24');
			insert into cierre values (2019, 3, 7, '2019-3-14', '2019-3-17', '2019-3-24');
			insert into cierre values (2019, 4, 7, '2019-4-14', '2019-4-17', '2019-4-24');
			insert into cierre values (2019, 5, 7, '2019-5-14', '2019-5-17', '2019-5-24');
			insert into cierre values (2019, 6, 7, '2019-6-14', '2019-6-17', '2019-6-24');
			insert into cierre values (2019, 7, 7, '2019-7-14', '2019-7-17', '2019-7-24');
			insert into cierre values (2019, 8, 7, '2019-8-14', '2019-8-17', '2019-8-24');
			insert into cierre values (2019, 9, 7, '2019-9-14', '2019-9-17', '2019-9-24');
			insert into cierre values (2019, 10, 7, '2019-10-14', '2019-10-17', '2019-10-24');
			insert into cierre values (2019, 11, 7, '2019-11-14', '2019-11-17', '2019-11-24');
			insert into cierre values (2019, 12, 7, '2019-12-14', '2019-12-17', '2019-12-24');
			insert into cierre values (2019, 1, 8, '2019-1-15', '2019-1-18', '2019-1-25');
			insert into cierre values (2019, 2, 8, '2019-2-15', '2019-2-18', '2019-2-25');
			insert into cierre values (2019, 3, 8, '2019-3-15', '2019-3-18', '2019-3-25');
			insert into cierre values (2019, 4, 8, '2019-4-15', '2019-4-18', '2019-4-25');
			insert into cierre values (2019, 5, 8, '2019-5-15', '2019-5-18', '2019-5-25');
			insert into cierre values (2019, 6, 8, '2019-6-15', '2019-6-18', '2019-6-25');
			insert into cierre values (2019, 7, 8, '2019-7-15', '2019-7-18', '2019-7-25');
			insert into cierre values (2019, 8, 8, '2019-8-15', '2019-8-18', '2019-8-25');
			insert into cierre values (2019, 9, 8, '2019-9-15', '2019-9-18', '2019-9-25');
			insert into cierre values (2019, 10, 8, '2019-10-15', '2019-10-18', '2019-10-25');
			insert into cierre values (2019, 11, 8, '2019-11-15', '2019-11-18', '2019-11-25');
			insert into cierre values (2019, 12, 8, '2019-12-15', '2019-12-18', '2019-12-25');
			insert into cierre values (2019, 1, 9, '2019-1-16', '2019-1-19', '2019-1-26');
			insert into cierre values (2019, 2, 9, '2019-2-16', '2019-2-19', '2019-2-26');
			insert into cierre values (2019, 3, 9, '2019-3-16', '2019-3-19', '2019-3-26');
			insert into cierre values (2019, 4, 9, '2019-4-16', '2019-4-19', '2019-4-26');
			insert into cierre values (2019, 5, 9, '2019-5-16', '2019-5-19', '2019-5-26');
			insert into cierre values (2019, 6, 9, '2019-6-16', '2019-6-19', '2019-6-26');
			insert into cierre values (2019, 7, 9, '2019-7-16', '2019-7-19', '2019-7-26');
			insert into cierre values (2019, 8, 9, '2019-8-16', '2019-8-19', '2019-8-26');
			insert into cierre values (2019, 9, 9, '2019-9-16', '2019-9-19', '2019-9-26');
			insert into cierre values (2019, 10, 9, '2019-10-16', '2019-10-19', '2019-10-26');
			insert into cierre values (2019, 11, 9, '2019-11-16', '2019-11-19', '2019-11-26');
			insert into cierre values (2019, 12, 9, '2019-12-16', '2019-12-19', '2019-12-26');`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tTodos los datos fueron agregados exitosamente!")
		}
	}
}

// ResetearTodas : Borra todos los valores de las tablas
func ResetearTodas() {
	UI.BloqueDeTexto("*", "Falta agregar la lógica de la función")
}

// Crear : crea una nueva tabla
func Crear() {

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
	} else { //Si no hay error en la conexión se procede a solicitar los datos de la tabla

		fmt.Print("\n\tIngrese el nombre de la tabla: ")
		var tabla string
		fmt.Scanf("%s", &tabla)

		if tabla == "" {
			fmt.Println("\n\tEl nombre está vacío, se le asignará --> sin_nombre")
			tabla = "sin_nombre"
		}

		//Ingreso de atributos y tipos de la tabla
		var parametros []string
		var tipos []string
		continuar := true

		for continuar {
			fmt.Print("\n\tIngrese el nombre del atributo ó q para terminar: ")
			var parametro string
			fmt.Scanf("%s", &parametro)
			if parametro != "q" && parametro != "" {
				parametros = append(parametros, parametro)
				fmt.Print("\n\tIngrese el tipo del atributo ó q para terminar: ")
				var tipo string
				fmt.Scanf("%s", &tipo)
				if tipo != "q" && tipo != "" {
					tipos = append(tipos, tipo)
				} else {
					continuar = false
				}
			} else {
				continuar = false
			}
		}

		//Se comprueba si es una tabla válida (nombre de aributo --> tipo)
		if len(parametros) != 0 && len(tipos) != 0 && len(parametros) == len(tipos) {
			fmt.Println("\n\tSe intentará crear la tabla con el nombre: " + tabla)
			var consulta string
			for i := 0; i < len(parametros); i++ {
				if i < len(parametros)-1 {
					consulta += parametros[i] + " " + tipos[i] + ", "
				} else {
					consulta += parametros[i] + " " + tipos[i]
				}
			}

			_, err = db.Exec("CREATE TABLE " + tabla + "(" + consulta + ");")
			if err != nil {
				fmt.Println("\n	", err)
			} else {
				fmt.Println("\n\tLa tabla " + tabla + " fue creada exitosamente!")
			}
		} else {
			fmt.Println("\n\tLa tabla " + tabla + " no es válida. No se puede crear")
		}
	}
}

////////////////////////////////////////////////////////////////////////

// Eliminar : Elimina la tabla indicada si existe
func Eliminar() {
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
	} else { //Si no hay error en la conexión se procede a solicitar el nombre de la tabla a borrar
		fmt.Print("\n\tIngrese el nombre de la tabla a borrar: ")
		var tabla string
		fmt.Scanf("%s", &tabla)

		//Se procede a eliminar
		_, err = db.Exec("DROP TABLE " + tabla)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tSe eliminó la tabla " + tabla)
		}
	}
}

////////////////////////////////////////////////////////////////////////

// Renombrar : Renombra la tabla indicada si existe
func Renombrar() {
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
	} else { //Si no hay error en la conexión se procede a solicitar el nombre de la tabla a renombrar

		//Se solicita el nombre de la tabla a renombrar
		fmt.Print("\n\tIngrese el nombre de la tabla que desea renombrar: ")
		var tablaOld string
		fmt.Scanf("%s", &tablaOld)

		//Se solicita el nuevo nombre
		fmt.Print("\n\tIngrese el nuevo nombre para la tabla: ")
		var tablaNew string
		fmt.Scanf("%s", &tablaNew)

		//Se procede a renombrar
		_, err = db.Exec("ALTER DATABASE " + tablaOld + " RENAME TO " + tablaNew)
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tSe realizó el cambio " + tablaOld + " --> " + tablaNew)
		}
	}
}

////////////////////////////////////////////////////////////////////////

// PK : Agrega la PK indicada si la tabla existe
func PK() {
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
	} else { //Si no hay error en la conexión se procede a solicitar el nombre de la tabla

		//Se solicita el nombre de la tabla para agregar la PK
		fmt.Print("\n\tIngrese el nombre de la tabla: ")
		var tabla string
		fmt.Scanf("%s", &tabla)

		//Se solicitan los atributos de la tabla para agregarles la PK
		fmt.Print("\n\tIngrese los nombres de los atributos que tendrán la primary key: ")
		var atributo string
		fmt.Scanf("%s", &atributo)

		//Se procede a insertar la PK
		_, err = db.Exec("ALTER TABLE " + tabla + " ADD CONSTRAINT " + tabla + "_pk PRIMARY KEY (" + atributo + ");")
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tSe insertó la primary key al atributo " + atributo + " de la tabla " + tabla)
		}
	}
}

////////////////////////////////////////////////////////////////////////

// FK : Agrega las FKs indicadas si la tabla existe
func FK() {
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
	} else { //Si no hay error en la conexión se procede a solicitar el nombre de la tabla

		//Se solicita el nombre de la tabla para agregar la FK
		fmt.Print("\n\tIngrese el nombre de la tabla: ")
		var tabla string
		fmt.Scanf("%s", &tabla)

		//Se solicitan los atributos de la tabla para agregarles la FK
		fmt.Print("\n\tIngrese los nombres de los atributos referenciados que tendrán la foreign key: ")
		var atributoReferenciado string
		fmt.Scanf("%s", &atributoReferenciado)
		fmt.Print("\n\tIngrese el nombre de la tabla a la que pertenece el atributo referenciado: ")
		var tablaReferenciada string
		fmt.Scanf("%s", &tablaReferenciada)

		//Se procede a insertar la FK
		//alter table Tarjeta add constraint Tarjeta_nrocliente_fk foreing key (nrocliente) references Cliente (nrocliente);
		_, err = db.Exec("ALTER TABLE " + tabla + " ADD CONSTRAINT " + tabla + "_" + atributoReferenciado + "_fk FOREIGN KEY (" + atributoReferenciado + ") REFERENCES " + tablaReferenciada + " (" + atributoReferenciado + ");")
		if err != nil {
			fmt.Println("\n	", err)
		} else {
			fmt.Println("\n\tSe insertó como foreign key al atributo referenciado " + atributoReferenciado + " de la tabla " + tablaReferenciada)
		}
	}
}
