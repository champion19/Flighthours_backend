# Flujo de Verificación de Correo Electrónico

Este documento describe el proceso de verificación de correo electrónico en la aplicación Flighthours, que se inicia cuando un usuario hace clic en un enlace de verificación enviado a su bandeja de entrada.

## Diagrama del Flujo

```mermaid
sequenceDiagram
    participant User
    participant Frontend
    participant API_Gateway
    participant AuthService

    User->>+Frontend: 1. Clic en el enlace de verificación del correo
    Frontend->>+API_Gateway: 2. GET /v1/employees/verify-email?token=<jwt>
    API_Gateway->>+AuthService: 3. Procesa la solicitud y llama a VerifyEmail(token)
    AuthService->>+AuthService: 4. Valida el token JWT
    alt Token Válido
        AuthService->>+Database: 5. Obtiene el empleado por ID
        Database-->>-AuthService: 6. Devuelve los datos del empleado
        alt Correo no verificado
            AuthService->>+Database: 7. Actualiza email_confirmed a true
            Database-->>-AuthService: 8. Confirmación de actualización
            AuthService-->>-API_Gateway: 9. Verificación exitosa
            API_Gateway-->>-Frontend: 10. Devuelve respuesta de éxito (HTTP 200)
            Frontend-->>-User: 11. Muestra mensaje de éxito
        else Correo ya verificado
            AuthService-->>-API_Gateway: 12. Error: Correo ya verificado
            API_Gateway-->>-Frontend: 13. Devuelve error (HTTP 409)
            Frontend-->>-User: 14. Muestra mensaje de error
        end
    else Token Inválido
        AuthService-->>-API_Gateway: 15. Error: Token inválido
        API_Gateway-->>-Frontend: 16. Devuelve error (HTTP 400)
        Frontend-->>-User: 17. Muestra mensaje de error
    end
```

## Descripción Detallada del Proceso

1.  **Inicio del Flujo (Clic del Usuario):**
    *   El usuario recibe un correo electrónico de verificación después de registrarse.
    *   Al hacer clic en el enlace del correo, se abre una URL en el navegador que apunta al frontend de la aplicación, con un token JWT adjunto como parámetro de consulta (`?token=...`).

2.  **Solicitud del Frontend al Backend:**
    *   El frontend recibe la solicitud y extrae el token de la URL.
    *   Realiza una solicitud `GET` al endpoint de la API: `/v1/employees/verify-email`.

3.  **Procesamiento en el Backend (API Handler):**
    *   El manejador (`handler`) en `cmd/api/handler.go` recibe la solicitud.
    *   Extrae el token del parámetro de consulta.
    *   Invoca al método `service.VerifyEmail(token)` para iniciar la lógica de verificación.

4.  **Lógica de Verificación en el Servicio:**
    *   El método `VerifyEmail` en `internal/domain/employee/service.go` es responsable de la lógica de negocio.
    *   **Validación del Token:** Se utiliza `tokenGenerator.ValidateJWT(token)` para verificar la autenticidad y validez del token. Si el token es inválido o ha expirado, se devuelve un error `ErrInvalidToken`.
    *   **Extracción del ID de Empleado:** Si el token es válido, se extrae el `employeeID`.

5.  **Interacción con la Base de Datos:**
    *   **Búsqueda del Empleado:** Se utiliza el `employeeID` para buscar al empleado en la base de datos. Si no se encuentra, se devuelve un error `ErrEmployeeCannotGet`.
    *   **Verificación del Estado:** Se comprueba si el campo `email_confirmed` del empleado ya es `true`. Si es así, se devuelve un error `ErrEmailAlreadyVerified` para evitar procesar la misma verificación varias veces.
    *   **Actualización del Estado:** Si el correo no ha sido verificado, se actualiza el campo `email_confirmed` a `true`.

6.  **Respuesta de la API:**
    *   **Éxito:** Si la verificación es exitosa, la API devuelve una respuesta `HTTP 200 OK` con un mensaje de confirmación.
    *   **Errores:**
        *   `HTTP 400 Bad Request`: Si el token es inválido (`ErrInvalidToken`).
        *   `HTTP 404 Not Found`: Si el empleado no se encuentra (`ErrEmployeeCannotGet`).
        *   `HTTP 409 Conflict`: Si el correo ya ha sido verificado (`ErrEmailAlreadyVerified`).

7.  **Finalización del Flujo en el Frontend:**
    *   El frontend recibe la respuesta de la API.
    *   Muestra un mensaje apropiado al usuario, ya sea de éxito o de error, completando así el flujo de verificación.
