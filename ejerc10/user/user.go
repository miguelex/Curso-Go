package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombe string, alta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombe
	this.FechaAlta = alta
	this.Status = status
}
