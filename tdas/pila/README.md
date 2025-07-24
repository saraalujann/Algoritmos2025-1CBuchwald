# Pila

El trabajo a realizar es el de una implementación de pila dinámica (es decir, una pila sobre un arreglo que pueda crecer o reducirse según la cantidad de elementos) que contenga datos genéricos.

En el adjunto en el sitio de descargas encontrarán el archivo `pila.go` que tienen que utilizar. En este archivo están definidas las primitivas que tendrán que implementar, con su correspondiente documentación. Todas las primitivas tienen que funcionar en tiempo constante.

**Importante:** La función debe devolver un elemento de tipo **Pila**, ya que la estructura no debe exportarse.

## Implementación

Hay que escribir el archivo `pila_dinamica.go`, con la implementación de la estructura de la pila y de cada una de las primitivas incluidas en el encabezado. Además de las primitivas, pueden tener funciones auxiliares, de uso interno, que no deben estar declaradas dentro de `pila.go`. En `pila.go` se encuentran únicamente las primitivas que el usuario de la pila tiene que conocer.

Además, deben incluir la primitiva para crear la pila:

```go
func CrearPilaDinamica[T any]() Pila[T] {
    // ...
}
```

Les sugerimos la siguiente implementación para la estructura interna de la pila:

```go
type pilaDinamica[T any] struct {
    datos    []T
    cantidad int // Cantidad de elementos almacenados.
}
```

## Pruebas

Además de `pila_dinamica.go`, tienen que entregar otro archivo `pila_test.go`, que contenga las pruebas unitarias para verificar que la pila funciona correctamente, y que al ejecutarlo puede verificarse que todo funciona bien. Se permite y recomienda usar funciones auxiliares en donde crean necesario. Es necesario separar los tests en diferentes funciones que prueben los escenarios. No es válido tener un único test que pruebe todo.

El archivo `pila_dinamica.go` debe ser parte del paquete `pila` (tal y como se encuentra en el archivo base adjunto), y el archivo de pruebas debe ser parte del paquete `pila_test`.

Las pruebas deberán verificar que:

1. **Crear una Pila vacía**:
   - Se debe poder crear una Pila vacía, y debe comportarse como tal.

2. **Apilar y Desapilar elementos**:
   - Se deben poder apilar elementos, y al desapilarlos, se debe mantener el invariante de pila (LIFO).
   - Probar con elementos diferentes y verificar que salgan en el orden esperado.

3. **Prueba de volumen**:
   - Se deben poder apilar muchos elementos (1,000, 10,000 elementos o el volumen que corresponda).
   - Hacer crecer la pila y desapilar hasta que esté vacía, verificando que siempre cumpla el invariante.
   - Comprobar que el tope de la pila sea el correcto en cada paso, y que el nuevo tope después de cada desapilar también sea el correcto.

4. **Condición de borde (Pila vacía)**:
   - Comprobar que una pila vacía se comporte como recién creada.

5. **Condición de borde (Desapilar en Pila vacía)**:
   - Las acciones para desapilar y ver el tope de una pila recién creada deben ser inválidas.

6. **Condición de borde (Pila vacía)**:
   - La acción para ver si una pila recién creada está vacía debe devolver verdadero.

7. **Condición de borde (Desapilar en Pila vacía después de operaciones)**:
   - Las acciones para desapilar y ver el tope de una pila que se ha apilado y desapilado hasta quedar vacía deben ser inválidas.

8. **Probar con diferentes tipos de datos**:
   - Probar apilar y desapilar diferentes tipos de datos: enteros, cadenas, etc.

Además de todos los casos no descriptos que ustedes crean **necesarios**.

La estructura de directorios del módulo/proyecto en el que se deberá trabajar es:

```
tdas
   └── pila
       ├── pila.go
       ├── pila_dinamica.go
       └── pila_test.go
   └── go.mod
```

Revisar el video de cómo realizar los módulos para las diferentes entregas.

##Forma de compilación.

Compilar todo el código:

```sh
go build pila
```

Ejecutar las pruebas:

```sh
go test pila_test.go
```

Adicionalmente, pueden utilizar el flag `-v` para que se muestre la salida de cada prueba, incluyendo logs (esto es lo que se hace en el corrector automático).

## Entrega

Al igual que en los casos anteriores, deberán entregar el código en formato digital subiendo el código a la página de entregas de la materia, con el código completo. Se debe entregar un archivo zip con la carpeta `pila` (que dentro debe tener `pila_dinamica.go` y `pila_test.go`) así como el archivo `go.mod` (tanto la carpeta como el archivo `go.mod` deben estar en el mismo nivel, sin `.`). No es necesario enviar el archivo `pila.go`. Incluso, este será ignorado para evitar cualquier modificación sobre las primitivas.

No olviden revisar las preguntas frecuentes de la pila.

## Bibliografía Recomendada

- Weiss, Mark Allen, “Data Structures and Algorithm Analysis”: 3.3. The Stack ADT.
- Cormen, Thomas H. “Introduction to Algorithms”: 10.1. Stacks and queues.
