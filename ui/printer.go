package ui

import (
	"fmt"
	"strings"

	"github.com/Alexis-Santana-Vega/gotodo/todo"
)

// PrintTasks imprime una colección de tareas estructurada en formato de tabla.
// Si la lista está vacía, muestra un mensaje indicando que no hay tareas
// El formato de la tabla linea dinámicamente las columnas de Id, estado, título y prioridad.
func PrintTasks(tasks []todo.Task) {
	if len(tasks) == 0 {
		fmt.Println("  (sin tareas)")
		return
	}
	// Imprime el encabezado de la tabla con espaciados fijos
	fmt.Printf("\n  %-4s %-3s %-32s %s\n", "ID", "Est", "Título", "Prioridad")
	// Dibuja una línea separadora visual de 52 caracteres de longitud
	fmt.Println("  " + strings.Repeat("─", 52))
	// Itera y formatea cada una de las tareas dentro del slice
	for _, t := range tasks {
		status := "○" // Símbolo para tareas pendientes
		if t.Done {
			status = "✓" // Símbolo para tareas completadas
		}
		// Imprime la fila alineando el texto a la izquierda con anchos fijos
		fmt.Printf("  %-4d %-3s %-32s %s\n",
			t.Id, status, t.Title, t.Priority)
	}
}

// PrintHelp muestra un menú gráfico en la terminal detallando la sintaxis
// de todos los comandos de consola disponibles en la aplicación gotodo.
func PrintHelp() {
	fmt.Println(`
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
╚═════════════════════════════════════════════════╝`)
}
