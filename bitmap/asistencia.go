package bitmap

import (
	"github.com/untref-ayp2/data-structures/set"
)

// Asistencias es una estructura que representa la asistencia de los alumnos a las clases
type Asistencias struct {
	// Implementar
}

// NewAsistencias crea un nuevo registro de asistencias
func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias {
	// Implementar
	return nil
}

// RegistrarAsistencia registra la asistencia de un alumno a una clase
// Si el alumno o la clase no existen, no hace nada
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {
	// Implementar
}

// Asistio devuelve true si el alumno asistió a la clase indicada
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	// Implementar
	return false
}

// ListarClase devuelve un set con los alumnos que asistieron a la clase indicada
func (a *Asistencias) ListarClase(clase uint8) set.Set[uint] {
	// Implementar
	return nil
}

// ListarAlumno devuelve un set con las clases a las que asistió el alumno indicado
func (a *Asistencias) ListarAlumno(alumno uint) set.Set[uint8] {
	// Implementar
	return nil
}

// ListarAsistencias devuelve un set con las clases a las que asistieron todos los alumnos
func (a *Asistencias) ListarAsistencias() set.Set[uint8] {
	// Implementar
	return nil
}

// ListarAsistenciaPerfecta devuelve un set con los alumnos que asistieron a todas las clases
func (a *Asistencias) ListarAsistenciaPerfecta() set.Set[uint] {
	// Implementar
	return nil
}
