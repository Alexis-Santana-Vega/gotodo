# gotodo

Gestor de tareas interactivo en línea de comandos escrito en Go. Las tareas viven en memoria durante la sesión — sin base de datos ni archivos. Proyecto de práctica para aprender structs con métodos, slices como almacén de datos, paquetes propios, manejo de errores tipados y tests unitarios con estado.

---

## Características

- **CRUD completo**: agregar, completar y eliminar tareas
- **Prioridades**: alta, media y baja
- **Filtros**: listar solo pendientes, completados, por prioridad, o búsqueda por título
- **Menú interactivo**: REPL en consola, sin flags ni argumentos al arrancar

---

## Requisitos

- [Go 1.22+](https://go.dev/dl/)

```bash
go version
```

---

## Instalación

```bash
git clone https://github.com/Alexis-Santana-Vega/gotodo.git
cd gotodo
go run .
```

O compila el binario primero:

```bash
go build -o gotodo .
./gotodo
```

---

## Uso

Al ejecutar el programa aparece el menú de ayuda y el prompt `>` esperando comandos.

```
╔═════════════════════════════════════════════════╗
║              gotodo — gestor CLI                ║
║    tipo de prioridades: alta, media, baja       ║
╠═════════════════════════════════════════════════╣
║ add <priority> <title >                         ║
║ done <id>                                       ║
║ delete <id>                                     ║
║ priority <priority>                             ║
║ search <query>                                  ║
║ list · pending · complete · help · quit         ║
╚═════════════════════════════════════════════════╝
```

### Agregar tareas

```
> add alta Estudiar goroutines en Go
Tarea creada: [○] #1 Estudiar goroutines en Go       [alta]

> add media Leer documentación de net/http
Tarea creada: [○] #2 Leer documentación de net/http  [media]

> add baja Configurar el editor
Tarea creada: [○] #3 Configurar el editor            [baja]
```

### Listar tareas

```
> list

  ID   Est Título                            Prioridad
  ────────────────────────────────────────────────────
  1    ○   Estudiar goroutines en Go         alta
  2    ○   Leer documentación de net/http    media
  3    ○   Configurar el editor              baja
```

### Completar y eliminar

```
> done 1
Tarea #1 marcada como completada ✓

> delete 3
Tarea #3 eliminada.
```

Alias disponibles: `del` = `delete`, `ls` = `list`, `q` = `quit`.

---

## Estructura del proyecto

```
gotodo/
├── go.mod               # definición del módulo
├── main.go              # punto de entrada — conecta Store con UI (5 líneas)
├── Makefile             # atajos de build, test y ejecución
├── todo/
│   ├── task.go          # struct Task, tipo Priority, métodos y fmt.Stringer
│   ├── store.go         # Store: CRUD en memoria con []Task
│   ├── filter.go        # filtros puros: ByStatus, ByPriority, Search
│   └── store_test.go    # tests unitarios del Store y los filtros
└── ui/
    ├── menu.go          # REPL: loop de lectura con bufio.Scanner
    └── printer.go       # tabla de salida y mensaje de ayuda
```

El proyecto se divide en tres capas con responsabilidades claras:

- **`todo/`** — dominio puro. No sabe nada de consola ni de presentación.
- **`ui/`** — presentación. Solo lee input e imprime; no conoce los detalles del Store.
- **`main.go`** — conecta ambas capas. No contiene lógica propia.

---

## Tests

```bash
go test ./...              # todos los tests
go test -v ./todo/         # con detalle por caso
```

Los tests cubren:

- Crear tareas con título válido e inválido (vacío)
- Completar tareas por ID existente e inexistente
- Eliminar tareas y verificar que el slice se actualiza
- Filtros `ByStatus` y `ByPriority` sobre el resultado de `GetAll`

Cada test crea su propio `Store` con `todo.New()` para garantizar aislamiento total entre casos.

---

## Makefile

```bash
make build    # compila el binario ./bin/gotodo
make run      # ejecuta con go run .
make test     # corre todos los tests con -v
make clean    # elimina el binario
```
