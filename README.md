📚 Biblioteca Digital en Go
📌 Descripción del proyecto
Este proyecto es una aplicación web tipo backend desarrollada en el lenguaje Go (Golang) que simula un sistema de gestión de biblioteca digital.
El sistema permite administrar usuarios, libros y préstamos, aplicando conceptos de Programación Orientada a Objetos (POO) mediante estructuras (structs), además de implementar concurrencia con goroutines y mutex para el control de acceso simultáneo a los recursos.
________________________________________
🎯 Objetivo
Desarrollar una aplicación web en Go que permita:
•	Gestionar usuarios 
•	Registrar y listar libros 
•	Realizar préstamos de libros 
•	Simular concurrencia en solicitudes simultáneas 
•	Aplicar conceptos de POO y manejo de datos en memoria 
________________________________________
⚙️ Tecnologías utilizadas
•	Lenguaje: Go (Golang) 
•	Concurrencia: goroutines, sync.Mutex, sync.WaitGroup 
•	Entrada de datos: bufio / fmt 
•	Estructuras de datos: structs, slices, maps 
•	Control de versiones: Git / GitHub 
________________________________________
🧠 Conceptos aplicados
✔ Programación Orientada a Objetos (POO)
Se implementa mediante structs como:
•	Usuario 
•	Libro 
•	Prestamo 
✔ Encapsulación
Uso de estructuras para agrupar datos relacionados.
✔ Concurrencia
Se implementa mediante:
•	goroutines (ejecución simultánea) 
•	mutex (control de acceso a recursos compartidos) 
•	wait groups (sincronización) 
________________________________________
🏗️ Estructura del proyecto
📁 biblioteca-digital
 └── main.go
 └── README.md
________________________________________
▶️ Cómo ejecutar el proyecto
1. Clonar el repositorio
git clone https://github.com/tu-usuario/biblioteca-digital.git
2. Entrar al proyecto
cd biblioteca-digital
3. Ejecutar la aplicación
go run main.go
________________________________________
📋 Funcionalidades
👤 Usuarios
•	Registrar usuarios 
📚 Libros
•	Registrar libros 
•	Listar libros 
•	Ver disponibilidad 
🔄 Préstamos
•	Realizar préstamo de libros 
•	Validar disponibilidad 
•	Evitar doble préstamo 
⚡ Concurrencia
•	Simulación de dos usuarios solicitando el mismo libro al mismo tiempo 
•	Control de acceso con mutex 
________________________________________
🧪 Pruebas realizadas
✔ Pruebas unitarias (manuales)
•	Registro de usuario 
•	Registro de libro 
•	Validación de préstamo 
✔ Pruebas de integración
•	Flujo completo: usuario → libro → préstamo 
✔ Prueba de concurrencia
•	Dos usuarios intentan pedir el mismo libro simultáneamente 
•	Resultado: solo uno obtiene el préstamo 
________________________________________
📸 Evidencias
(Aquí debes agregar capturas de pantalla de tu ejecución)
________________________________________
⚠️ Observaciones
•	Los datos se almacenan en memoria (no usa base de datos). 
•	El sistema es una simulación educativa de backend. 
•	Diseñado para demostrar conceptos de POO y concurrencia en Go. 
________________________________________
👨‍💻 Autor
•	Proyecto académico – Unidad 4 
•	Lenguaje: Go 
•	Tema: Aplicación web con POO y concurrencia 
________________________________________
🚀 Conclusión
Este proyecto permite comprender la implementación de estructuras de datos, manejo de concurrencia y principios de programación orientada a objetos en Go, aplicados a un sistema funcional de gestión de biblioteca.
lioteca-Digital1
