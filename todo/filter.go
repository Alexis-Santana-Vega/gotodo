package todo

import "strings"

// ByStatus filtra la lista de tareas y retorna únicamente aquellas que coinciden con el status solicitado
// Retorna un slice vacío (no nulo) si ninguna tarea cumple con el criterio
func ByStatus(tasks []Task, done bool) []Task {
	result := []Task{}
	for _, t := range tasks {
		if t.Done == done {
			result = append(result, t)
		}
	}
	return result
}

// ByPriority filtra la lista de tareas y retorna únicamente aquellas que coinciden con la prioridad solicitada
// Returna un slice vacío (no nulo) si ninguna tarea cumple con el criterio
func ByPriority(tasks []Task, p Priority) []Task {
	result := []Task{}
	for _, t := range tasks {
		if t.Priority == p {
			result = append(result, t)
		}
	}
	return result
}

// Search busca y retorna las tareas cuyo título contenga la subcadena provista en query
// La búsqueda se realiza sin distinción entre mayúsculas y minúsculas (case-insentive)
func Search(tasks []Task, query string) []Task {
	result := []Task{}
	for _, t := range tasks {
		// Normaliza ambos strings a minúsculas para asegurar una búsqueda flexible
		if strings.Contains(
			strings.ToLower(t.Title),
			strings.ToLower(query),
		) {
			result = append(result, t)
		}
	}
	return result
}
