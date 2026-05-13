package todo

import (
	"testing"
)

// TestAdd verifica que la creación de tareas válidas asigne correctamente
// los identificadores secuenciales, el estado inicial y actualice el Store
func TestAdd(t *testing.T) {
	s := New()                                      // Se crea un nuevo Store
	task, err := s.Add("Aprender Go", PriorityHigh) // Se llama la función para agregar una tarea

	if err != nil {
		t.Fatalf("Add inesperado: %v", err)
	}
	if task.Id != 1 {
		t.Errorf("ID esperado 1, got %d", task.Id)
	}
	if task.Done {
		t.Error("tarea nueva no debe estar completada")
	}
	if len(s.GetAll()) != 1 {
		t.Error("el store debe tener 1 tarea")
	}
}

// TestAddEmptyTitle comprueba que el sistema rechace la creación de tareas
// que no tengan un título descriptivo provisto
func TestAddEmptyTitle(t *testing.T) {
	s := New()
	_, err := s.Add("", PriorityLow)
	if err == nil {
		t.Error("se esperaba error con título vacío")
	}
}

// TestComplete valida el flujo de marcar una tarea como resulta a través de su Id,
// asegurando también que falle correctamente ante identificadores inexistentes
func TestComplete(t *testing.T) {
	s := New()
	s.Add("Tarea 1", PriorityMedium)

	if err := s.Complete(1); err != nil {
		t.Fatalf("Complete(1) error: %v", err)
	}
	if !s.GetAll()[0].Done {
		t.Error("la tarea debería estar marcada como hecha")
	}
	if err := s.Complete(99); err == nil {
		t.Error("Complete(99) debería retornar error")
	}
}

// TestDelete confirma que la eliminación física de un elemento reduzco el tamaño
// del almacén y maneje adecuadamente los errores de desbordamiento o Ids inválidos
func TestDelete(t *testing.T) {
	s := New()
	s.Add("A", PriorityLow)
	s.Add("B", PriorityLow)

	if err := s.Delete(1); err != nil {
		t.Fatalf("Delete(1) error: %v", err)
	}
	if len(s.GetAll()) != 1 {
		t.Error("el store debería tener 1 tarea tras eliminar")
	}
	if err := s.Delete(99); err == nil {
		t.Error("Delete(99) debería retornar error")
	}
}

// TestSearch asegura el correcto comportamiento del filtro de búsqueda por query
func TestSearch(t *testing.T) {
	s := New()
	s.Add("Red", PriorityLow)
	s.Add("Blue", PriorityHigh)
	s.Add("Yellow", PriorityMedium)
	s.Add("Green", PriorityLow)
	if result := Search(s.tasks, "l"); len(result) != 2 {
		t.Error("Search(l) debería retornar 2 registros")
	}
}

// TestFilter asegura el correcto comportamiento de las funciones utilitarias
// de segmentación por propiedades físicas de la tarea (estado y criticidad)
func TestFilter(t *testing.T) {
	s := New()
	s.Add("Urgente", PriorityHigh)
	s.Add("Normal", PriorityMedium)
	s.Complete(1)

	done := ByStatus(s.GetAll(), true)
	if len(done) != 1 {
		t.Errorf("esperaba 1 completada, got %d", len(done))
	}

	high, _ := ByPriority(s.GetAll(), PriorityHigh)
	if len(high) != 1 {
		t.Errorf("esperaba 1 alta prioridad, got %d", len(high))
	}
}
