package main

import (
	"fmt"
	"./media"
)

func main(){
	//s1 := media.Imagen{Titulo: "Miami", Formato: "JPG", Canales: "RGB"}

	//fmt.Println(s1.Mostrar())

	cw := media.ContenidoWeb{}

	var opcion int 
for{	
	fmt.Println("\t Menu de medios \n\t 1) Capturar imagen \n\t 2) Capturar audio \n\t 3) Capturar video \n\t 4) Mostrar \n\t 5) Salir \n\t Selecciona una opción:")
	fmt.Scan(&opcion)
	
	if opcion == 1 {
		fmt.Println("\n Capturando Imagen")
		image := media.Imagen{}
		fmt.Println("\n Ingrese el titulo de la imagen: ")
		fmt.Scan(&image.Titulo)
		fmt.Println("\n Ingrese el formato de la imagen: ")
		fmt.Scan(&image.Formato)
		fmt.Println("\n Ingrese el canal de la imagen: ")
		fmt.Scan(&image.Canales)
		cw.Multimedias = append(cw.Multimedias,&image)
	}else if opcion == 2 {
		fmt.Println("\n Capturando Audio")
		var song media.Audio
		fmt.Println("\n Ingrese el titulo del audio: ")
		fmt.Scan(&song.Titulo)
		fmt.Println("\n Ingrese el formato del audio: ")
		fmt.Scan(&song.Formato)
		fmt.Println("\n Ingrese la duracion del audio: ")
		fmt.Scan(&song.Duracion)
		cw.Multimedias = append(cw.Multimedias,&song)
	}else if opcion == 3 {
		fmt.Println("\n Capturando Video")
		var movie media.Video
		fmt.Println("\n Ingrese el titulo del video: ")
		fmt.Scan(&movie.Titulo)
		fmt.Println("\n Ingrese el formato del video: ")
		fmt.Scan(&movie.Formato)
		fmt.Println("\n Ingrese el frame del video: ")
		fmt.Scan(&movie.Frames)
		cw.Multimedias = append(cw.Multimedias,&movie)
	}else if opcion == 4 {
		fmt.Println("\n  Mostrando Datos")
		cw.Mostrar()	
	}else if opcion == 5{
		break
	}
}
}