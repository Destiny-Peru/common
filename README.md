# common

Modulo compartido para concentrar piezas reutilizables entre servicios Go de Destiny Peru.

Incluye:

- helpers de configuracion y entorno
- errores genericos
- middlewares HTTP
- logger con `zap`
- envelope de respuesta HTTP

## Estructura

- `config`
- `errors`
- `http/middleware`
- `logger`
- `response`

## Enfoque

Este modulo no contiene reglas de negocio ni configuracion especifica de un servicio.
La idea es que cada proyecto defina su propia estructura de configuracion y reutilice aqui:

- tipos base (`App`, `CORS`, `Database`, `HTTPServer`)
- helpers para `.env`
- logger
- middlewares HTTP
- response envelope
