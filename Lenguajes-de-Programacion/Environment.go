package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func argCommandLine(args []string) {
	fmt.Println("Argumentos de línea de comando:")
	fmt.Println("Número de argumentos:", len(args)-1)
	fmt.Println("Argumentos:", args[1:])
}

func variableAmbiente() {
	fmt.Println("Variables de Entorno:")
	fmt.Println("PATH: ", os.Getenv("PATH"))
}

func fileIO() {
	nombre := "C:\\Users\\solca\\Downloads\\GO\\prueba.txt"
	file, err := os.Create(nombre)
	if err != nil {
		fmt.Println("Error al crear archivo: ", err)
		return
	}

	file.WriteString("Mensaje de Prueba")

	fmt.Println("Texto en archivo")
	contenido, err := os.ReadFile(nombre)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return
	}
	fmt.Println(string(contenido))
}

func networkIO() {
	host := "intec.edu.do"
	port := "80"

	ip, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Error al encontrar IP: ", err)
		return
	}
	fmt.Printf("Direccion IP de '%s' es %s", host, ip)

	con, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error al establecer conexion: ", err)
		return
	}
	defer con.Close()

	fmt.Fprintf(con, "GET / HTTP/1.0\r\n\r\n")
	response := bufio.NewScanner(con)
	for response.Scan() {
		fmt.Println(response.Text())
	}

}
func main() {

	for {
		fmt.Println("1) Mostrar Argumentos de Linea de Comando")
		fmt.Println("2) Mostrar Variables de Ambiente")
		fmt.Println("3) File I/O")
		fmt.Println("4) Network I/O")
		fmt.Println("5) Salir")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			argCommandLine(os.Args)
		case 2:
			variableAmbiente()
		case 3:
			fileIO()
		case 4:
			networkIO()
		case 5:
			return

		}

	}
}
