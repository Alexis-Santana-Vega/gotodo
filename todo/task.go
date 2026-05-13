package todo

import (
	"fmt"
	"time"
)

// Priority define los niveles de urgencia de una tarea
// Priority (un string) brinda seguridad de tipos. Si usamos una función
// que espera Priority, no podremos pasarle un string cualquiera por accidente
type Priority string

// Constantes de tipo Priority que representan los niveles de prioridad
const (
	PriorityLow    Priority = "baja"
	PriorityMedium Priority = "media"
	PriorityHigh   Priority = "alta"
)

// Task representa una tarea individual en el sistema
type Task struct {
	Id        int
	Title     string
	Done      bool
	Priority  Priority
	CreatedAd time.Time
}

// String implementa la interfaz fmt.Stringer para el tipo Task
// Devuelve una representación en texto formateada de la tarea,
// mostrando su estado (completado o pendiente), id, título y prioridad.
func (t Task) String() string {
	// 1. Determina el símbolo visual según el estado de la tarea
	status := "○" // Pendiente
	if t.Done {
		status = "✓" // Completado
	}
	// 2. Retorna la cadena formateada
	return fmt.Sprintf("[%s] #%d %-30s [%s]", status, t.Id, t.Title, t.Priority)
}

// IsValid verifica que la tarea contenga datos mínimos obligatorios y válidos
// Retorna true si el título no está vacío y la prioridad corresponde a uno de
// los valores definidos (PriorityLow, PriorityMedium o PriorityHigh).
func (t Task) IsValid() bool {
	// Valida que el título tenga contenido y la prioridad sea una de las permitidas
	return t.Title != "" && (t.Priority == PriorityLow ||
		t.Priority == PriorityMedium || t.Priority == PriorityHigh)
}

// Complete marca la tarea como completada
// Es un receptor de puntero: modifica el original
func (t *Task) Complete() {
	t.Done = true
}
