# 🧩 Order API - REST Mock & Coding Exercises

API REST mock para pruebas + colección de ejercicios de LeetCode y HackerRank.

## 🏗️ Estructura del Proyecto

```
order-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point de la aplicación
├── internal/                    # Código privado del proyecto
│   ├── handlers/               # HTTP handlers
│   │   └── order_handler.go
│   ├── models/                 # Modelos de datos
│   │   └── order.go
│   ├── middleware/             # Middlewares HTTP
│   │   └── logging.go
│   └── repository/             # Capa de datos (mock)
│       └── order_repository.go
├── pkg/                        # Código público reutilizable
│   └── exercises/              # Ejercicios de código
│       ├── leetcode/
│       │   ├── arrays.go       # TwoSum, MinimumAbsDifference
│       │   ├── strings.go      # FirstUniqueChar
│       │   └── design.go       # ATM Machine
│       └── hackerrank/
│           └── problems.go
├── examples/                   # Ejemplos de uso
│   └── exercises_demo.go
├── utils/                      # Utilidades generales
│   ├── demo.go                # Demo de concurrencia
│   ├── exercises.go           # (Legacy - migrado a pkg/)
│   ├── helpers.go             # IsPalindrome
│   └── palindrome.go          # (Legacy - migrado a helpers.go)
├── go.mod
├── go.sum
├── main.go                    # (Legacy - usar cmd/server/main.go)
└── README.md
```

## 🚀 Inicio Rápido

### Ejecutar el servidor

```bash
go run cmd/server/main.go
```

El servidor iniciará en `http://localhost:8080`

### Ejecutar solo los ejercicios

```go
import "example.com/order-api/pkg/exercises/leetcode"

result := leetcode.TwoSum([]int{2, 7, 11, 15}, 9)
fmt.Println(result) // [0, 1]
```

## 📚 API Endpoints

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET    | `/order` | Lista todas las órdenes |
| GET    | `/order/{id}` | Obtiene una orden por ID |
| POST   | `/order` | Crea una nueva orden |
| PUT    | `/order/{id}` | Actualiza una orden |
| DELETE | `/order/{id}` | Elimina una orden |

### Ejemplo de uso

```bash
# Listar todas las órdenes
curl http://localhost:8080/order

# Obtener una orden específica
curl http://localhost:8080/order/1111

# Crear una nueva orden
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"customer":"Carlos","amount":3500}'

# Actualizar una orden
curl -X PUT http://localhost:8080/order/1111 \
  -H "Content-Type: application/json" \
  -d '{"customer":"Ana García","amount":1500}'

# Eliminar una orden
curl -X DELETE http://localhost:8080/order/1111
```

## 💡 Ejercicios Implementados

### LeetCode

- **Listado de ejercicios**: Revisar [pkg/exercises/README.md](pkg/exercises/leetcode/README.md) para más detalles.

### HackerRank

- Listo para agregar problemas en `pkg/exercises/hackerrank/`

## 🛠️ Tecnologías

- **Go 1.21+**
- **gorilla/mux** - Router HTTP
- **google/uuid** - Generación de UUIDs

## 📖 Filosofía del Proyecto

Este proyecto combina dos objetivos:

1. **API Mock**: Servidor REST simple para pruebas y demos
2. **Ejercicios de Código**: Colección organizada de soluciones a problemas de coding

La estructura sigue las convenciones de Go:
- `internal/` - Código privado del API
- `pkg/` - Código reutilizable (ejercicios)
- `cmd/` - Aplicaciones ejecutables

## 🎯 Próximos Pasos

- [ ] Agregar tests unitarios
- [ ] Implementar más ejercicios de HackerRank
- [ ] Agregar documentación Swagger
- [ ] Implementar autenticación JWT
- [ ] Agregar persistencia con base de datos

## ✍ Autor

**Greg Perez**  
📫 [gregperezm@gmail.com](mailto:gregperezm@gmail.com)  
🌍 Santiago, Chile
