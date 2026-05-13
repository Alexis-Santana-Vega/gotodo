package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Alexis-Santana-Vega/gotodo/todo"
)

// Run inicia el bucle iterativo de la interfaz en la terminal (REPL)
// Lee los comandos ingresados por el usuario a través de la entrada estándar (stdin),
// los procesa en base a argumentos y delega las operaciones hacia el store de tareas.
func Run(store *todo.Store) {
	scanner := bufio.NewScanner(os.Stdin)
	PrintHelp()
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		// Divide el texto de entrada ignorando espacios múltiples o saltos de línea
		parts := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(parts) == 0 {
			continue
		}
		cmd, args := parts[0], parts[1:]
		switch cmd {
		case "add":
			handleAdd(store, args)
		case "done":
			handleDone(store, args)
		case "delete", "del":
			handleDelete(store, args)
		case "list", "ls":
			PrintTasks(store.GetAll())
		case "pending":
			PrintTasks(todo.ByStatus(store.GetAll(), false))
		case "complete":
			PrintTasks(todo.ByStatus(store.GetAll(), true))
		case "priority":
			handlePriority(store, args)
		case "search":
			handleSearch(store, args)
		case "help":
			PrintHelp()
		case "quit", "q":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Printf("Comando desconocido: %s. Escribe 'help'.\n", cmd)
		}
	}
}

// handleAdd procesa el comando de creación de una tarea.
// Requiere la prioridad como primer parámetro y une los subsecuentes para el título.
func handleAdd(store *todo.Store, args []string) {
	if len(args) < 2 {
		fmt.Println("uso: add <priority> <title...>   prioridades: alta media baja")
		return
	}
	priority := todo.Priority(args[0])
	// Une el resto de los argumentos en una sola cadena para soportar títulos con espacios
	title := strings.Join(args[1:], " ")
	t, err := store.Add(title, priority)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Tarea creada: %s\n", t)
}

// handleDone procesa el comando para marcar una tarea existente como completada mediante su ID.
func handleDone(store *todo.Store, args []string) {
	id, err := parseId(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := store.Complete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Tarea #%d marcada como completada ✓\n", id)
}

// handleDelete procesa la remoción definitiva de una tarea utilizando su identificador.
func handleDelete(store *todo.Store, args []string) {
	id, err := parseId(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := store.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Tarea #%d eliminada.\n", id)
}

// handlePriority filtra por prioridad
func handlePriority(store *todo.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("uso: priority <priority>   prioridades: alta media baja")
		return
	}
	p := todo.Priority(args[0])
	result, err := todo.ByPriority(store.GetAll(), p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	PrintTasks(result)
}

func handleSearch(store *todo.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("uso: search <query>")
		return
	}
	query := strings.Join(args[0:], " ")
	PrintTasks(todo.Search(store.GetAll(), query))
}

// parseId valida los argumentos del subcomando y convierte el primer elemento
// de tipo texto a un identificador numérico de tipo entero.
func parseId(args []string) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("se requiere un Id numérico")
	}
	return strconv.Atoi(args[0])
}
