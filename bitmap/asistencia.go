package bitmap

import (
	"github.com/untref-ayp2/data-structures/set"
)

// Asistencias es una estructura que representa la asistencia de los alumnos a las clases
type Asistencias struct {
	clases      uint8
	alumnos     uint
	asistencias []*BitMap
}

// NewAsistencias crea un nuevo registro de asistencias
func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias {

	asistencias := make([]*BitMap, cantAlumnos)
	for i := range asistencias {
		asistencias[i] = NewBitMap()
	}

	return &Asistencias{alumnos: cantAlumnos, clases: cantClases, asistencias: asistencias}
}

// RegistrarAsistencia registra la asistencia de un alumno a una clase
// Si el alumno o la clase no existen, no hace nada
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {
	if clase > a.clases || alumno > a.alumnos {
		return
	}
	a.asistencias[alumno].On(clase)
}

// Asistio devuelve true si el alumno asistió a la clase indicada
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	if clase > a.clases || alumno > a.alumnos {
		return false
	}
	isOn, _ := a.asistencias[alumno].IsOn(clase)
	return isOn
}

// ListarClase devuelve un set con los alumnos que asistieron a la clase indicada
func (a *Asistencias) ListarClase(clase uint8) set.Set[uint] {
	set := set.NewListSet[uint]()
	for i := range a.asistencias {
		isOn, _ := a.asistencias[i].IsOn(clase)
		if isOn {
			set.Add(uint(i))
		}
	}
	return set
}

// ListarAlumno devuelve un set con las clases a las que asistió el alumno indicado
func (a *Asistencias) ListarAlumno(alumno uint) set.Set[uint8] {
	set := set.NewListSet[uint8]()
	for i := 0; i < int(a.clases); i++ {
		isOn, _ := a.asistencias[alumno].IsOn(uint8(i))
		if isOn {
			set.Add(uint8(i))
		}
	}
	return set
}

// ListarAsistencias devuelve un set con las clases a las que asistieron todos los alumnos
func (a *Asistencias) ListarAsistencias() set.Set[uint8] {
	set := set.NewListSet[uint8]()
	var alumno int
	for clase := 0; clase < int(a.clases); clase++ {
		for alumno = 0; alumno < int(a.alumnos); alumno++ {
			isOff, _ := a.asistencias[alumno].IsOff(uint8(clase))
			if isOff {
				break
			}
		}
		if alumno == int(a.alumnos) {
			set.Add(uint8(clase))
		}
	}
	return set
}

// ListarAsistenciaPerfecta devuelve un set con los alumnos que asistieron a todas las clases
func (a *Asistencias) ListarAsistenciaPerfecta() set.Set[uint] {
	set := set.NewListSet[uint]()
	var clase int
	for alumno := 0; alumno < int(a.alumnos); alumno++ {
		for clase = 0; clase < int(a.clases); clase++ {
			isOff, _ := a.asistencias[alumno].IsOff(uint8(clase))
			if isOff {
				break
			}
		}
		if clase == int(a.clases) {
			set.Add(uint(alumno))
		}
	}
	return set
}
