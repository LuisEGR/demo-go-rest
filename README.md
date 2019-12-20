# Ejemplo de microservicio en Go (golang)


## Librerías utilizadas: 

- Echo: https://echo.labstack.com/
- SQLX: http://jmoiron.github.io/sqlx/

---

### Compilar:

Este comando descargará las dependencias necesarias y generará el archivo ejecutable.

```
go build
```

### Ejecutar:

```
./app
```

---

## Docker

Para ejecutar este servicio en un contenedor de docker, solo hay que construir la imagen:

```
docker build -t myService .
```

y luego para ejecutarlo:

```
docker run -p 8080:8080 myService
```

---

## Database

Este microservicio se conecta a una base de datos de postgres, los valores de la conexión hay que modificarlos en el archivo `db/db.go`

El script para generar la tabla que se usa es este:

```sql
CREATE TABLE public.usuario (
	id numeric NULL,
	name varchar NULL,
	"time" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO public.usuario (id, "name", "time") VALUES(0, 'user0', CURRENT_TIMESTAMP);
INSERT INTO public.usuario (id, "name", "time") VALUES(1, 'user1', CURRENT_TIMESTAMP);
INSERT INTO public.usuario (id, "name", "time") VALUES(2, 'user2', CURRENT_TIMESTAMP);
```

---

## Peticiones soportadas en este microservicio:

### **GET /users**

    Obtiene todos los usuarios de la base de datos
    y los retorna como un arreglo en JSON

Response:
```json
[
    {
        "id": 0,
        "nombre": "user0",
        "time": "2019-12-20T12:06:57.286898Z"
    },
    {
        "id": 1,
        "nombre": "user1",
        "time": "2019-12-20T12:06:57.286898Z"
    },
    {
        "id": 2,
        "nombre": "user2",
        "time": "2019-12-20T12:06:57.286898Z"
    }
]
```

---

### **POST /user**

    Agrega un usuario nuevo a la base de datos

Request: 
```json
{
    "id": 3,
    "nombre": "user3"
}
```

Response: 
```json
{
    "id": 3,
    "nombre": "user3",
    "time": ""
}
```

---

### **DELETE /user/3**

    Elimina un usuario de la base de datos

Response `status 200`


---

## TO-DO

- [x] Múltiples métodos (GET, POST, DELETE)
- [x] Múltiples rutas (/users, /user)
- [x] Request body en JSON
- [x] Response body en JSON
- [x] Path Params (/user/:id)
- [x] Conexión con Postgresql
- [x] Logs
- [x] Dockerfile
- [x] Dependencias con `go modules`
- [ ] Variables de entorno (db, port)
- [ ] Makefile
- [ ] Unit Testing