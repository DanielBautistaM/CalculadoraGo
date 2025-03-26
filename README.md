# Calculadora Criptográfica

Una aplicación de consola interactiva desarrollada en Go que implementa diversos algoritmos criptográficos, técnicas de codificación y operaciones matemáticas modulares.

## Descripción

Esta calculadora criptográfica es una herramienta educativa diseñada para estudiantes y profesionales interesados en ciberseguridad. La aplicación permite explorar desde conceptos básicos de matemáticas modulares hasta algoritmos de criptografía moderna, todo a través de una interfaz de consola intuitiva que muestra los procesos paso a paso.

## Características

La aplicación está organizada en cinco módulos principales:

### 1. Operaciones Matemáticas Modulares

- **Cálculo de módulo**: Implementación de la operación `a mod n = b`
- **Inverso aditivo**: Cálculo del valor que sumado al original resulta en cero bajo un módulo específico
- **Inverso XOR**: Demostración de las propiedades del operador XOR
- **Máximo Común Divisor**: Algoritmo de Euclides con análisis de coprimalidad
- **Inverso multiplicativo (método tradicional)**: Búsqueda exhaustiva con explicación detallada
- **Inverso multiplicativo (AEE)**: Implementación del Algoritmo Extendido de Euclides

### 2. Criptografía Clásica

- **Cifrado Módulo 27**: Adaptación para el alfabeto español (incluye Ñ)
- **Cifrado César**: Implementación clásica con tabla de sustitución
- **Cifrado Vernam**: Operación XOR bit a bit con visualización del proceso
- **Cifrado ATBASH**: Sustitución por inversión del alfabeto
- **Transposición Columnar**: Reorganización de caracteres con clave alfabética
- **Cifrado Afín**: Implementación de la función `(ax + b) mod m`
- **Cifrado de Sustitución**: Sistema monoalfabético con múltiples opciones

### 3. Criptografía Moderna

- **Diffie-Hellman**: Simulación de intercambio seguro de claves
- **RSA**: Generación de claves, cifrado y descifrado
- **Exponenciación Rápida**: Algoritmo optimizado para cálculos con grandes exponentes

### 4. Algoritmos Hash

- **MD5**: Generación y análisis de digest de 128 bits
- **SHA-1**: Implementación con información sobre sus propiedades de seguridad
- **SHA-512**: Hash avanzado con visualización detallada

### 5. Codificación

- **Binario**: Conversión bidireccional texto-binario
- **Hexadecimal**: Codificación y decodificación con explicaciones
- **Base64**: Transformación para transmisión segura de datos binarios

## Requisitos

- Go 1.16 o superior

## Instalación

1. Clone el repositorio:
   ```
   git clone https://github.com/DanielBautistaM/CalculadoraGo.git
   cd CalculadoraGo
   ```

2. Compile el programa:
   ```
   go build -o CalculadoraGo
   ```

3. Ejecute la aplicación:
   ```
   ./CalculadoraGo
   go run .
   ```

## Uso

La calculadora presenta un menú interactivo navegable mediante las teclas de flecha (↑↓) y Enter para seleccionar opciones. Cada función incluye instrucciones claras que guían al usuario durante todo el proceso.

## Compilación multiplataforma

Para generar ejecutables para diferentes sistemas operativos:

```bash
# Para Windows
GOOS=windows GOARCH=amd64 go build -o calculadora_criptografica.exe

# Para Linux
GOOS=linux GOARCH=amd64 go build -o calculadora_criptografica_linux

# Para macOS
GOOS=darwin GOARCH=amd64 go build -o calculadora_criptografica_mac
```

## Estructura del proyecto

```
calculadora_criptografica/
├── main.go                  # Punto de entrada y manejo de menús
├── go.mod                   # Definición del módulo
├── go.sum                   # Checksums de dependencias
├── funciones/
│   ├── caso1.go             # Operaciones matemáticas modulares
│   ├── caso2.go             # Criptografía clásica
│   ├── caso3.go             # Criptografía moderna
│   ├── caso4.go             # Algoritmos hash
│   ├── caso5.go             # Codificación
│   └── helpers.go           # Funciones auxiliares para I/O
└── README.md                # Documentación
```
