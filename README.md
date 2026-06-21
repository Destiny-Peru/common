# common

Modulo compartido para concentrar piezas reutilizables entre servicios Go de Destiny Peru.

Incluye:

- helpers de configuracion y entorno
- errores genericos
- middlewares HTTP
- logger con `zap`
- envelope de respuesta HTTP
- health HTTP reutilizable

## Estructura

- `config`
- `errors`
- `http/health`
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
- health handler generico basado en `Checker`

## Versionado automatico

El repositorio queda preparado con GitHub Actions para CI y autoversionado.

- `CI`: ejecuta `go test ./...` en cada `push` a `main` y en cada `pull_request`
- `Release Please`: calcula la siguiente version, crea o actualiza el PR de release y publica el tag/release cuando ese PR se fusiona

### Convencion de commits

Usa Conventional Commits para que el versionado sea automatico:

- `fix: corrige bug en health handler` -> sube `patch` (`v1.0.1`)
- `feat: agrega middleware reutilizable` -> sube `minor` (`v1.1.0`)
- `feat!: cambia contrato de response` o `BREAKING CHANGE:` -> sube `major` (`v2.0.0`)

Si el commit no sigue esta convencion, Release Please no podra inferir correctamente el cambio de version.