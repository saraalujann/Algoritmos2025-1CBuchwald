# TP1: Conversión de notación infija a posfija

## Contenido
- **Previo al enunciado**
- **Introducción**
- **Conversor desde notación infija**
  - Formato de entrada
  - Asociatividad y precedencia
- **Criterios de aprobación**

## Previo al enunciado
Como se ha indicado en clase, esperamos para la elaboración de este trabajo práctico que ya tengan conocimiento pleno del uso de Go, lo cual incluye todo lo visto en clase y lo explicado en los videos sobre el lenguaje. Si no se ha visto alguno de los videos, es necesario que primero lo revisen porque este enunciado asume que esto es sabido.

También, recomendamos volver a revisar el video sobre *cómo armar los módulos en Go*, en particular para los TPs.

## Introducción
Nuestra forma usual de escribir operaciones matemáticas (operador operando operador, ejemplo 3 + 2) se llama notación infija. Otro tipo de notación puede ser la notación polaca inversa (también llamada notación posfija, en inglés reverse Polish notation). Esta notación se utiliza, por ejemplo, en calculadoras financieras, especialmente de hace algunas décadas.

Mostramos algunos ejemplos de pasaje de notación infija a posfija:

```
3 + 2         → 3 2 +
20 / -3       → 20 -3 /
20^10         → 20 10 ^
20 * 10^5 - 2 → 20 10 5 ^ * 2 -
```

## Conversor desde notación infija
Se desea implementar un conversor de notación infija a notación posfija. El programa debe leer por entrada estándar operaciones en notación infija (una por línea), e imprimirá por salida estándar la representación en postfijo de la misma operación. Ejemplo:
```sh
$ cat arith.txt
3 + 5
5 - 3
8 / 2 + 1
9 - 2 * 4
(9-2) * 4
5 + 4 ^ 3 ^ 2

$ ./infix < arith.txt
3 5 +
5 3 -
8 2 / 1 +
9 2 4 * -
9 2 - 4 *
5 4 3 2 ^ ^ +
```

Como referencia bibliográfica, la conversión se puede realizar mediante el algoritmo *shunting yard* (ver página de Wikipedia en castellano o en inglés).

## Formato de entrada
Cada línea de la entrada consistirá de una secuencia de tokens; cada uno de ellos podrá ser:

- uno de los cinco operadores aritméticos `+` `-` `*` `/` `^`
- un paréntesis de apertura, `(`; o de cierre, `)`
- un número entero, no negativo y en base decimal

Se debe aceptar espacios en blanco en cualquier punto de la expresión, excepto entre los dígitos de un valor numérico.

**Se garantiza, por último, que todas las expresiones de entrada estarán bien formadas**: paréntesis balanceados, cantidad de operadores correctos, etc.

## Asociatividad y precedencia
Todos los operadores asocian por la izquierda, excepto la exponenciación, `^`, que asocia por la derecha.

Por otra parte, ^ es el operador de **mayor precedencia**, seguido de `*` y `/` (ambos al mismo nivel); `+` y `-` son, ambos, los operadores de **menor precedencia**.

## Criterios de aprobación
El código entregado debe ser claro y legible y ajustarse a las especificaciones de la consigna. Debe compilar sin advertencias y correr sin errores.

La entrega incluye, obligatoriamente, todos los archivos involucrados en la realización del TP (es decir, el módulo del trabajo en sí, que debe llamarse `infix`), así como el módulo `tdas` en caso de haber utilizado al menos alguno de los tipos de datos implementados anteriormente.

La entrega se realiza únicamente en forma digital a través del sistema de entregas, con todos los archivos mencionados en un único archivo ZIP.
