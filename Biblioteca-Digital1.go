package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ===============================
// CATEGORÍAS FIJAS (ARRAY)
// ===============================
// Array de tamaño fijo (no cambia en ejecución)
var Categorias = []string{
	"Programación",
	"Bases de Datos",
	"Redes",
	"Ciberseguridad",
	"Inteligencia Artificial",
	"Drama",
}

// ===============================
// MODELOS (POO EN GO)
// ===============================

// Usuario representa una persona que usa la biblioteca
type Usuario struct {
	ID     int
	Nombre string
	Correo string
}

// Libro representa un libro en el sistema
type Libro struct {
	ID         int
	Titulo     string
	Autor      string
	Categoria  string
	Disponible bool // true = disponible, false = prestado
}

// Prestamo representa la acción de préstamo de un libro
type Prestamo struct {
	ID        int
	UsuarioID int
	LibroID   int
}

// ===============================
// BASE DE DATOS EN MEMORIA
// ===============================
// Simulamos una base de datos usando estructuras en memoria

var (
	Usuarios  = make([]Usuario, 0)
	Libros    = make([]Libro, 0)
	Prestamos = make([]Prestamo, 0)

	// Map para acceso rápido a libros por ID
	Catalogo = make(map[int]*Libro)

	// Mutex para evitar problemas de concurrencia
	mu sync.Mutex

	// Mutex para control de IDs
	idLock sync.Mutex
)

// ===============================
// CONSTRUCTORES (CREACIÓN DE OBJETOS)
// ===============================

// Crea un nuevo usuario
func NuevoUsuario(nombre, correo string) Usuario {
	idLock.Lock()
	defer idLock.Unlock()

	return Usuario{
		ID:     len(Usuarios) + 1,
		Nombre: nombre,
		Correo: correo,
	}
}

// Crea un nuevo libro
func NuevoLibro(titulo, autor, categoria string) Libro {
	idLock.Lock()
	defer idLock.Unlock()

	return Libro{
		ID:         len(Libros) + 1,
		Titulo:     titulo,
		Autor:      autor,
		Categoria:  categoria,
		Disponible: true,
	}
}

// ===============================
// FUNCIONES PRINCIPALES (CRUD)
// ===============================

// Registra un nuevo usuario en memoria
func RegistrarUsuario(nombre, correo string) {
	u := NuevoUsuario(nombre, correo)
	Usuarios = append(Usuarios, u)
	fmt.Println("✔ Usuario registrado correctamente")
}

// Registra un nuevo libro en el sistema
func RegistrarLibro(titulo, autor, categoria string) error {
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}

	l := NuevoLibro(titulo, autor, categoria)

	// Guardamos en slice y en mapa (búsqueda rápida)
	Libros = append(Libros, l)
	Catalogo[l.ID] = &Libros[len(Libros)-1]

	fmt.Println("✔ Libro registrado correctamente")
	return nil
}

// Muestra todos los libros registrados
func ListarLibros() {
	fmt.Println("\n📚 LISTA DE LIBROS:")

	for _, l := range Libros {
		estado := "Disponible"
		if !l.Disponible {
			estado = "Prestado"
		}

		fmt.Printf("ID:%d | %s | %s | %s\n", l.ID, l.Titulo, l.Autor, estado)
	}
}

// ===============================
// CONCURRENCIA (PARTE IMPORTANTE)
// ===============================
// Esta función simula préstamos simultáneos

func PrestarLibro(usuarioID, libroID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Bloqueamos acceso para evitar conflictos (race condition)
	mu.Lock()
	defer mu.Unlock()

	libro, existe := Catalogo[libroID]

	// Validación: libro existe
	if !existe {
		fmt.Println("❌ Libro no encontrado")
		return
	}

	// Validación: disponibilidad
	if !libro.Disponible {
		fmt.Println("❌ El libro ya está prestado")
		return
	}

	// Cambiamos estado del libro
	libro.Disponible = false

	// Registramos préstamo
	Prestamos = append(Prestamos, Prestamo{
		ID:        len(Prestamos) + 1,
		UsuarioID: usuarioID,
		LibroID:   libroID,
	})

	fmt.Println("✔ Préstamo realizado por usuario:", usuarioID)
}

// ===============================
// REPORTES DEL SISTEMA
// ===============================

// Muestra estadísticas generales del sistema
func Reporte() {
	fmt.Println("\n===== REPORTE GENERAL =====")
	fmt.Println("Usuarios registrados:", len(Usuarios))
	fmt.Println("Libros registrados:", len(Libros))
	fmt.Println("Préstamos realizados:", len(Prestamos))
}

// ===============================
// UTILIDAD
// ===============================

// Pausa la ejecución para que el usuario vea la información
func pausar(scanner *bufio.Scanner) {
	fmt.Print("\nPresiona ENTER para continuar...")
	scanner.Scan()
}

// ===============================
// MAIN (MENÚ PRINCIPAL)
// ===============================

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Menú principal del sistema
		fmt.Println("\n========================")
		fmt.Println("   BIBLIOTECA DIGITAL")
		fmt.Println("========================")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Registrar libro")
		fmt.Println("3. Listar libros")
		fmt.Println("4. Préstamo normal")
		fmt.Println("5. DEMO CONCURRENCIA")
		fmt.Println("6. Reporte general")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione opción: ")

		scanner.Scan()
		opcion, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		switch opcion {

		case 1:
			// Registro de usuario
			fmt.Print("Nombre: ")
			scanner.Scan()
			nombre := scanner.Text()

			fmt.Print("Correo: ")
			scanner.Scan()
			correo := scanner.Text()

			RegistrarUsuario(nombre, correo)
			pausar(scanner)

		case 2:
			// Registro de libro
			fmt.Print("Título: ")
			scanner.Scan()
			titulo := scanner.Text()

			fmt.Print("Autor: ")
			scanner.Scan()
			autor := scanner.Text()

			fmt.Print("Categoría: ")
			scanner.Scan()
			categoria := scanner.Text()

			err := RegistrarLibro(titulo, autor, categoria)
			if err != nil {
				fmt.Println("Error:", err)
			}

			pausar(scanner)

		case 3:
			// Mostrar libros
			ListarLibros()
			pausar(scanner)

		case 4:
			// Préstamo normal (con goroutine)
			var wg sync.WaitGroup

			fmt.Print("ID Usuario: ")
			scanner.Scan()
			u, _ := strconv.Atoi(scanner.Text())

			fmt.Print("ID Libro: ")
			scanner.Scan()
			l, _ := strconv.Atoi(scanner.Text())

			wg.Add(1)
			go PrestarLibro(u, l, &wg)
			wg.Wait()

			pausar(scanner)

		case 5:
			// DEMOSTRACIÓN DE CONCURRENCIA
			var wg sync.WaitGroup

			fmt.Println("🔴 Simulación: 2 usuarios intentando el mismo libro")

			wg.Add(2)
			go PrestarLibro(1, 1, &wg)
			go PrestarLibro(2, 1, &wg)

			wg.Wait()

			pausar(scanner)

		case 6:
			// Reporte general
			Reporte()
			pausar(scanner)

		case 0:
			// Salida del sistema
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
