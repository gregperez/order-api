# 🧩 Orders API & Go Practice

Este repositorio contiene una API REST sencilla para la gestión de órdenes, junto con ejercicios de lógica en Go que uso como práctica personal y preparación para desafíos técnicos.

## 📦 Contenido

- **API CRUD de órdenes**:
  - Endpoints REST con `gorilla/mux`
  - Almacenamiento en memoria (`map[string]Order`)
  - Validaciones básicas y uso de UUIDs
- **Ejercicios prácticos**:
  - Verificador de palíndromos (`utils/palindrome.go`)
  - Manejo de argumentos por consola (`utils/demo.go`)

---

## 🚀 Ejecutar localmente

### Requisitos
- Go 1.18 o superior

### Clonar y correr
```bash
git clone https://github.com/tu-usuario/orders-api-go.git
cd orders-api-go
go run main.go
```

La API estará disponible en: [http://localhost:8080](http://localhost:8080)

---

## 📚 Endpoints disponibles

| Método | Endpoint          | Descripción                    |
|--------|-------------------|--------------------------------|
| GET    | `/orders`         | Obtener todas las órdenes      |
| GET    | `/order/{id}`     | Obtener una orden por ID       |
| POST   | `/order`          | Crear una nueva orden          |
| PUT    | `/order/{id}`     | Actualizar una orden existente |
| DELETE | `/order/{id}`     | Eliminar una orden por ID      |

---

## 🧠 Ejercicios incluidos

### ✅ Verificador de Palíndromos
Ubicado en `utils/palindrome.go`, valida si una palabra es palíndroma comparando extremos del string.

### ✅ Demo con argumentos de consola
En `utils/demo.go`, se simula una lectura/escritura simple usando `os.Args`, ideal para testing rápido de inputs por CLI.

---

## 📁 Estructura del proyecto

```
ORDERS-API/
├── main.go
├── go.mod
├── go.sum
├── utils/
│   ├── demo.go
│   ├── palindrome.go
```

---

## 📚 Dependencias

- [`gorilla/mux`](https://github.com/gorilla/mux)
- [`google/uuid`](https://github.com/google/uuid)

---

## ✍ Autor

**Greg Perez**  
📫 [gregperezm@email.com](mailto:gregperezm@email.com)  
🌍 Santiago, Chile

---
