package storedprocedures

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

// AutorizacionCompra : Verifica si una compra es posible y devuelve un booleano
func AutorizacionCompra() {
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
		_, err = db.Exec(`CREATE OR REPLACE FUNCTION compra_valida() RETURNS TRIGGER AS $$

		DECLARE

			existe int;
			codigo_valido int;
			suma_consumos_previos decimal(8,2);
			limite decimal(8,2);
			vencida int;
			suspendida int;
			
		BEGIN

			--Inicialización de variales:

			SELECT COUNT(nrotarjeta) INTO existe FROM tarjeta
				WHERE NEW.nrotarjeta = nrotarjeta AND estado = 'vigente';

			SELECT COUNT(codseguridad) INTO codigo_valido FROM tarjeta
				WHERE NEW.nrotarjeta = nrotarjeta AND NEW.codseguridad = codseguridad;

			SELECT SUM(monto) INTO suma_consumos_previos FROM compra
				WHERE pagado = FALSE AND nrotarjeta = NEW.nrotarjeta;


			SELECT limitecompra INTO limite FROM tarjeta
				WHERE NEW.nrotarjeta = nrotarjeta;

			SELECT COUNT(nrotarjeta) INTO vencida FROM tarjeta
				WHERE NEW.nrotarjeta = nrotarjeta AND ( current_date < to_date(tarjeta.validadesde, 'YYYYMM') OR current_date > to_date(tarjeta.validahasta, 'YYYYMM') );

			SELECT COUNT(nrotarjeta) INTO suspendida FROM tarjeta
			WHERE NEW.nrotarjeta = nrotarjeta AND estado = 'suspendida';

			--Validación de datos:

			IF existe THEN
				IF codigo_valido THEN
					IF vencida = 0 THEN
						IF suspendida = 0 THEN
							IF (NEW.monto > limite) OR (suma_consumos_previos + NEW.monto > limite) THEN
								INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo)
								VALUES(
									NEW.nrotarjeta,
									NEW.nrocomercio,
									CURRENT_TIMESTAMP,
									NEW.monto,
									'supera límite de tarjeta'   
								);	
									
							ELSE INSERT INTO compra(nrotarjeta, nrocomercio, fecha, monto, pagado)
								VALUES(
									NEW.nrotarjeta,
									NEW.nrocomercio,
									CURRENT_TIMESTAMP,
									NEW.monto,
									FALSE
								);
																					
							END IF;

						ELSE INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo)
							VALUES(
								NEW.nrotarjeta,
								NEW.nrocomercio,
								CURRENT_TIMESTAMP,
								NEW.monto,
								'la tarjeta se encuentra suspendida'
							);
							
						END IF;

					ELSE INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo)
						VALUES(
							NEW.nrotarjeta,
							NEW.nrocomercio,
							CURRENT_TIMESTAMP,
							NEW.monto,
							'plazo de vigencia expirado'
						);

					END IF;

				ELSE INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo)
					VALUES(
						NEW.nrotarjeta,
						NEW.nrocomercio,
						CURRENT_TIMESTAMP,
						NEW.monto,
						'código de seguridad inválido'
					);
								
				END IF;

			ELSE INSERT INTO rechazo(nrotarjeta, nrocomercio, fecha, monto, motivo)
				VALUES(
					NEW.nrotarjeta,
					NEW.nrocomercio,
					CURRENT_TIMESTAMP,
					NEW.monto,
					'tarjeta no válida ó no vigente'
				);

			END IF;
			RETURN NULL;
		END;
		$$ LANGUAGE plpgsql;`)

		if err != nil {
			fmt.Println("\n	", err)
		} else {
			_, err = db.Exec(`CREATE TRIGGER compra_valida_trg
				AFTER INSERT ON consumo
				FOR EACH ROW
				EXECUTE PROCEDURE compra_valida();`)

			if err != nil {
				fmt.Println("\n	", err)
			} else {
				fmt.Println("\n\tProcedimiento almacenado para la autorización de compra, agregado")
			}
		}
	}
}

// generación del resumen el trabajo práctico deberá contener la lógica que reciba
// como parámetros el número de cliente, y el periodo del año, y que guarde en las
// tablas que corresponda los datos del resumen con la siguiente información: nombre
// y apellido, dirección, número de tarjeta, periodo del resumen, fecha de vencimiento,
// todas las compras del periodo, y total a pagar.

// GenerarResumen : Crea un procedimiento almacenado para que, posteriormente, se generen resúmenes
func GenerarResumen() {
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

		//Se crea el procedimiento almacenado
		_, err = db.Exec(`
		CREATE OR REPLACE FUNCTION resumen(nrocliente_recibido int, desde char(8), hasta char(8)) RETURNS void AS $$

		DECLARE
		
			nrocabecera int;
			fecha_desde date;
			fecha_hasta date;
			fecha_vencimiento date;
			i record;
			j record;
			linea_actual int;
			fila_tarjeta record;
			nombre_comercio text;
			cliente_temporal record;
			suma_total decimal(7,2);
		
		BEGIN
		
			linea_actual := 1;
			
			fecha_desde := to_date(desde, 'YYYYMMDD');
			fecha_hasta := to_date(hasta, 'YYYYMMDD');
			fecha_vencimiento := to_date(hasta, 'YYYYMMDD') + 10;
		
			--Para cada tarjeta del cliente (puede tener más de una)
			FOR i IN SELECT * FROM tarjeta WHERE(tarjeta.nrocliente = nrocliente_recibido) loop
		
				suma_total := 0;
		
				--Se verifica la tarjeta. Si existe, se guarda la fila completa
				SELECT * INTO fila_tarjeta FROM tarjeta WHERE i.nrotarjeta = nrotarjeta;
				IF FOUND THEN
					
					--Se guarda al cliente con todos sus datos
					SELECT * INTO cliente_temporal FROM cliente WHERE(nrocliente = fila_tarjeta.nrocliente);
				   
					--Agregamos todos los datos obtenidos a la cabecera
					INSERT INTO cabecera(nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence, total) 
					VALUES (cliente_temporal.nombre, cliente_temporal.apellido, cliente_temporal.domicilio, fila_tarjeta.nrotarjeta, fecha_desde, fecha_hasta, fecha_vencimiento, suma_total); 
				   
					--Guardamos el número mayor de cabecera, después de la inserción
					SELECT MAX(nroresumen) INTO nrocabecera FROM cabecera;
				  
				   -- SELECT count(*) INTO cabecera_id  FROM cabecera;
						
					--Se revisan todas las compras de ese cliente con esa tarjeta
					FOR j IN SELECT * FROM compra WHERE (compra.nrotarjeta = fila_tarjeta.nrotarjeta AND compra.fecha <= fecha_hasta AND compra.fecha >= fecha_desde) loop
						
						--Se guardan el nombre del comercio
						SELECT nombre INTO nombre_comercio FROM comercio where(comercio.nrocomercio = j.nrocomercio);
						
						INSERT INTO detalle VALUES (nrocabecera, linea_actual, j.fecha, nombre_comercio, j.monto);
						
						suma_total = suma_total + j.monto;
		
						linea_actual := linea_actual + 1;
					
					end loop;
		
					--Se actualiza 
					UPDATE cabecera SET total = suma_total WHERE (nroresumen = nrocabecera);
				   
				ELSE
					 RAISE NOTICE 'Error en la solicitud, verifique los datos ingresados';
					
				END IF;
			end loop;    
		END;
		$$ LANGUAGE plpgsql;
		`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {

			fmt.Println("\n\tProcedimiento almacenado para la generación de resumen, agregado")

		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

// Resumen : Inserta el
func Resumen() {
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

		//Se obtienen los datos que recibirá el procedimiento almacenado
		var nrocliente int
		var desde string
		var hasta string

		nrocliente = elegirCliente(db)
		if nrocliente != 0 {
			UI.Clean()
			fmt.Print("\n\tIngrese la fecha de inicio en el formato (YYYYMMDD) Ejemplo --> 20190610: ")
			fmt.Scanf("%s", &desde)

			fmt.Print("\n\tIngrese la fecha de fin en el formato (YYYYMMDD) Ejemplo --> 20190710: ")
			fmt.Scanf("%s", &hasta)

			if desde > hasta {
				fmt.Println("\n\tLa fecha de inicio es mayor a la de fin!")
			} else {
				//Se prepara la consulta
				sqlStatement := `SELECT resumen($1, $2, $3)`
				_, err = db.Exec(sqlStatement, nrocliente, desde, hasta)
				if err != nil {
					fmt.Println("\n\t	", err)
				} else {
					fmt.Println("\n\tSe ha creado el resumen a partir del stored procedures --> resumen")

					//verResumen();
				}
			}
		} else {
			fmt.Println("\n\tError al elegir el cliente de la lista!")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////

func elegirCliente(db *sql.DB) int {
	UI.Clean()
	type Cliente struct {
		nrocliente int
		nombre     string
		apellido   string
	}
	rows, err := db.Query(`SELECT nrocliente, nombre, apellido FROM cliente`)
	if err != nil {
		fmt.Println("\n\t", err)
		return 0
	}
	defer rows.Close()
	contador := 1
	var lista []Cliente
	UI.BloqueDeTexto("*", "Seleccionar el Cliente: ")
	for rows.Next() {
		cliente := Cliente{}
		err = rows.Scan(&cliente.nrocliente, &cliente.nombre, &cliente.apellido)
		if err != nil {
			fmt.Println("\n\t", err)
			return 0
		}
		lista = append(lista, cliente)
		fmt.Printf("%d --> Nro: %v | Apellido y nombre: %v, %v\n", contador, lista[contador-1].nrocliente, lista[contador-1].nombre, lista[contador-1].apellido)
		contador++
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("\n\t", err)
		return 0
	}
	//Seleccionar el cliente
	var nroclienteElegido int
	fmt.Print("\n\tOpción: ")
	fmt.Scanf("%d\n", &nroclienteElegido)

	if nroclienteElegido <= 0 || nroclienteElegido > contador { //Si no es uno de los números de la lista
		return 0
	}

	return lista[nroclienteElegido-1].nrocliente
}

/////////////////////////////////////////////////////////////////////////////////////////

// Alertas : Crea el procedimiento almacenado y trigger para gestionar alertas
func Alertas() {
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

		//Se crea el procedimiento almacenado
		_, err = db.Exec(`
			CREATE OR REPLACE FUNCTION rechazo() RETURNS TRIGGER AS $$
			DECLARE 
			BEGIN
				
				INSERT INTO alerta (nrotarjeta, fecha, nrorechazo, codalerta, descripcion) 
				VALUES (NEW.nrotarjeta, NEW.fecha, NEW.nrorechazo, 2, NEW.motivo);
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			----------TRIGGER------------

			CREATE TRIGGER rechazo_trg
				AFTER INSERT ON rechazo
				FOR EACH ROW
					EXECUTE PROCEDURE rechazo();
					
			--------------------------------------------------------------------------------------------
			--------------------------------------------------------------------------------------------
			
			CREATE OR REPLACE FUNCTION compras_consecutivas_1minuto() RETURNS TRIGGER AS $$
			DECLARE

				compra_consecutiva_1minuto record;
				cp_comercio char(8);
				codigo_alerta int;

			BEGIN
				codigo_alerta := 1000; --Se elige el tiempo de 1 minuto en milisegundos como código de alerta

				SELECT codigopostal INTO cp_comercio FROM comercio WHERE nrocomercio = NEW.nrocomercio;

				SELECT * INTO compra_consecutiva_1minuto FROM compra
				WHERE 
					NEW.nrotarjeta = compra.nrotarjeta --misma tarjeta
					AND NEW.fecha > ( NOW() - 1 * INTERVAL '1 minute' ) --en menos de un minuto
					AND NEW.nrocomercio NOT IN(SELECT nrocomercio FROM compra WHERE nrotarjeta = NEW.nrotarjeta) --en diferentes comercios
					AND cp_comercio IN (SELECT codigopostal FROM comercio WHERE nrocomercio = compra.nrocomercio); --en el mismo cod postal

				IF FOUND THEN
					INSERT INTO alerta(nrotarjeta, fecha, codalerta, descripcion)
					VALUES(
						NEW.nrotarjeta,
						NEW.fecha,
						codigo_alerta,
						'Compras consecutivas en lapso menor a 1 minuto, en diferentes comercios, en el mismo código postal'
					);
				END IF;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			----------TRIGGER------------

			CREATE TRIGGER compras_consecutivas_1minuto_trg
				BEFORE INSERT ON compra
				FOR EACH ROW
					EXECUTE PROCEDURE compras_consecutivas_1minuto();
			
			--------------------------------------------------------------------------------------------
			--------------------------------------------------------------------------------------------
			
			CREATE OR REPLACE FUNCTION compras_consecutivas_5minutos() RETURNS TRIGGER AS $$
			DECLARE
				compra_consecutiva_5minutos record;
				cp_comercio char(8);
				codigo_alerta int;
			BEGIN
				codigo_alerta := 5000; --Se elige el tiempo de 5 minutos en milisegundos como código de alerta

				SELECT codigopostal INTO cp_comercio FROM comercio WHERE nrocomercio = NEW.nrocomercio;

				SELECT * INTO compra_consecutiva_5minutos FROM compra
				WHERE 
					NEW.nrotarjeta = compra.nrotarjeta --misma tarjeta
					AND NEW.fecha > ( NOW() - 5 * INTERVAL '1 minute' ) --en menos de 5 minutos
					AND NEW.nrocomercio NOT IN(SELECT nrocomercio FROM compra WHERE nrotarjeta = NEW.nrotarjeta) --en diferentes comercios
					AND cp_comercio NOT IN (SELECT codigopostal FROM comercio WHERE nrocomercio = compra.nrocomercio); --en el distintos cod postales

				IF FOUND THEN
					INSERT INTO alerta(nrotarjeta, fecha, codalerta, descripcion)
					VALUES(
						NEW.nrotarjeta,
						NEW.fecha,
						codigo_alerta,
						'Compras consecutivas en lapso menor a 5 minutos, en diferentes comercios, con distinto código postal'
					);
				END IF;
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;

			----------TRIGGER------------

			CREATE TRIGGER compras_consecutivas_5minutos_trg
				BEFORE INSERT ON compra
				FOR EACH ROW
					EXECUTE PROCEDURE compras_consecutivas_5minutos();

			--------------------------------------------------------------------------------------------
			--------------------------------------------------------------------------------------------
			
			CREATE OR REPLACE FUNCTION compras_rechazadas_limite() RETURNS TRIGGER AS $$
			DECLARE
				fila_rechazo record;
				codigo_alerta int;
				--fecha_nueva_anio date;
				--fecha_nueva_mes date;
				--fecha_nueva_dia date;
			
			BEGIN
				codigo_alerta := 2; --Se elige como código el número de rechazos por exceso de límite

                --Se extraen de la fecha de rechazo timestamp solo el año, mes y día
                --SELECT EXTRACT(YEAR FROM NEW.fecha) INTO fecha_nueva_anio;	
                --SELECT EXTRACT(MONTH FROM NEW.fecha) INTO fecha_nueva_mes;
                --SELECT EXTRACT(DAY FROM NEW.fecha) INTO fecha_nueva_dia;
                
				
                SELECT * INTO fila_rechazo FROM rechazo
				WHERE
					NEW.nrotarjeta = rechazo.nrotarjeta
                    --Comparamos con otros rechazos, para ver si se dio el mismo año, mes, día
					--AND fecha_nueva_anio = EXTRACT(YEAR FROM rechazo.fecha)
                    --AND fecha_nueva_mes = EXTRACT(MONTH FROM rechazo.fecha)
					--AND fecha_nueva_dia = EXTRACT(DAY FROM rechazo.fecha)
					AND cast(new.fecha as date) = cast(rechazo.fecha as date)
					AND NEW.motivo = rechazo.motivo
					AND NEW.motivo = 'supera límite de tarjeta';
				
				IF FOUND THEN
					UPDATE tarjeta SET estado='suspendida' WHERE nrotarjeta = NEW.nrotarjeta;

					INSERT INTO alerta(nrotarjeta, fecha, codalerta, descripcion)
					VALUES(
						NEW.nrotarjeta,
						NEW.fecha,
						codigo_alerta,
						'Suspensión de tarjeta por exceder dos veces el límite, el mismo día'
					);
				END IF;

				RETURN NEW;

			END;
			$$ LANGUAGE plpgsql;

			CREATE TRIGGER compras_rechazadas_limite_trg
			BEFORE INSERT ON rechazo
			FOR EACH ROW
				EXECUTE PROCEDURE compras_rechazadas_limite();`)
		if err != nil {
			fmt.Println("\n	", err)
		} else {

			fmt.Println("\n\tProcedimientos almacenados y triggers para la gestión de alertas, agregados")

		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
