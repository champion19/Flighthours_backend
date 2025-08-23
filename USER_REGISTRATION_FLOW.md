
# Flujo de Registro de Usuario

Este documento detalla el proceso paso a paso de cómo un nuevo usuario (empleado) se registra en el sistema.

## 1. Solicitud a la API

El flujo comienza con una solicitud `POST` al siguiente endpoint de la API:

```
POST /v1/employee
```

El cuerpo de la solicitud debe contener los datos del usuario en formato JSON, como nombre, correo electrónico, contraseña, etc.

## 2. Controlador (Handler)

La solicitud es recibida por la función `Save()` en `cmd/api/handler.go`. Esta función se encarga de:

- **Decodificar el JSON:** Extrae los datos del cuerpo de la solicitud y los mapea a una estructura `EmployeeRequest`.
- **Validar los Datos:** Llama al método `Validate()` de `EmployeeRequest` para asegurar que los datos recibidos son válidos (p. ej., que los campos obligatorios no estén vacíos).
- **Convertir al Dominio:** Convierte la estructura `EmployeeRequest` a un objeto del dominio `domain.Employee`.

## 3. Lógica de Negocio (Servicio)

El controlador invoca al método `Save()` del servicio de empleados, ubicado en `internal/domain/employee/service.go`. Aquí ocurren los siguientes pasos:

- **Verificación de Duplicados:** El servicio comprueba si ya existe un empleado con el mismo correo electrónico en la base de datos. Si es así, devuelve un error de "empleado duplicado".
- **Generación de ID y Hashing de Contraseña:**
    - Se genera un ID único para el nuevo empleado.
    - La contraseña del empleado se hashea utilizando `bcrypt` por razones de seguridad.
- **Llamada al Repositorio:** El servicio invoca al método `Save()` del repositorio para persistir los datos del nuevo empleado.

## 4. Persistencia de Datos (Repositorio)

El método `Save()` en `internal/platform/employee/repository.go` es responsable de interactuar con la base de datos:

- **Preparar la Consulta:** Se prepara una sentencia SQL `INSERT` para añadir el nuevo empleado a la tabla `employee`.
- **Ejecutar la Consulta:** Se ejecuta la consulta con los datos del empleado.
- **Manejo de Errores:** Se manejan posibles errores de la base de datos, como violaciones de claves únicas (que podrían indicar un empleado duplicado si la comprobación del servicio fallara).

## 5. Verificación por Correo Electrónico

Una vez que el empleado ha sido guardado exitosamente en la base de datos, el servicio `Save()` procede a:

- **Generar un Token JWT:** Se crea un token de validación (JSON Web Token) para el nuevo usuario.
- **Enviar Correo de Verificación:** Se envía un correo electrónico al usuario que contiene un enlace de verificación con el token JWT generado. Esto se hace a través del notificador `Resend`.

## 6. Respuesta al Cliente

Finalmente, el controlador `Save()` envía una respuesta `HTTP 201 Created` al cliente, indicando que el usuario ha sido creado exitosamente y que se ha enviado un correo de verificación.

Este flujo asegura que los datos del usuario sean validados, que no haya duplicados y que la contraseña se almacene de forma segura, además de implementar un mecanismo de verificación de correo electrónico.
