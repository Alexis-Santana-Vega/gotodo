package todo

import (
	"errors"
	"fmt"
	"time"
)

// Errores personalizados para las operaciones de Store
var (
	// Se retorna cuando intenta operar sobre un ID inexistente
	ErrNotFound = errors.New("Task not found")
	// Se retorna al intentar crear una tarea sin texto en el título
	ErrEmptyTitle = errors.New("Title cannot be empty")
)

// Store gestiona el ciclo de vida y almacenamiento de las tareas en memoria
type Store struct {
	tasks  []Task // Colección interna de tareas
	nextId int    // Contador para asignar identificadores únicos secuenciales
}

// New inicializa y retorna un puntero a un nuevo Store con el contador inicializado en 1
func New() *Store {
	return &Store{nextId: 1}
}

// Add valida, crea y añade una nueva tarea a Store
// Retonar la copia de Task creada o ErrEmptyTitle
func (s *Store) Add(title string, p Priority) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	t := Task{
		Id:        s.nextId,
		Title:     title,
		Priority:  p,
		CreatedAd: time.Now(),
	}
	s.tasks = append(s.tasks, t)
	s.nextId++
	return t, nil
}

// Complete busca una tarea por su Id y cambia su estado a completado
// Retorna un error envuelto (wrapped) si el Id proporcionado no existe
func (s *Store) Complete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks[i].Complete() // Invoca método Complete de Task
			return nil
		}
	}
	return fmt.Errorf("Complete: %w (id=%d)", ErrNotFound, id)
}

// Delete elimina físicamente una tarea del slice interno utilizando su ID
// Preserva el orden relativo de los elementos restantes. Retorna error si el Id no existe
func (s *Store) Delete(id int) error {
	// i representa el índice y t representa el struct
	for i, t := range s.tasks {
		if t.Id == id {
			// Elimina el elemento i rebanando y uniendo el slice
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Delete: %w (id=%d)", ErrNotFound, id)
}

// GetAll genera y retorna un clon superficial (shallow copy) del slice de tareas
// Esto evita que modificaciones externas al slice afecten el estado interno del Store
func (s *Store) GetAll() []Task {
	return append([]Task{}, s.tasks...)
}
