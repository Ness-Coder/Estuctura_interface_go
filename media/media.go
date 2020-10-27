package media

import (
	"fmt"
)

type Multimedia interface {
	Mostrar() string
}

type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (cw ContenidoWeb) Mostrar() {
	for _,i:= range cw.Multimedias{
		fmt.Println(i)
	}
}

type Imagen struct {
	Titulo string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() string {
	return i.Titulo + "\n" +  i.Formato + "\n" + i.Canales + "\n"
}

type Audio struct {
	Titulo string
	Formato string
	Duracion string
}

func (a *Audio) Mostrar() string {
	return a.Titulo + "\n" + a.Formato + "\n" + a.Duracion + "\n"
}

type Video struct {
	Titulo string
	Formato string
	Frames string
}

func (v Video) Mostrar() string {
	return v.Titulo + "\n" + v.Formato + "\n" + v.Frames + "\n"
}
