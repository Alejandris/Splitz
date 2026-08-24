# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar archivos de definición de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente restante
COPY . .

# Compilar el binario optimizado para producción
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o splitz-api ./cmd/api/main.go

# Production stage
FROM alpine:3.20

# Instalar certificados CA y curl para el healthcheck / diagnóstico
RUN apk --no-cache add curl ca-certificates

WORKDIR /app

# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder /app/splitz-api .

# Exponer el puerto por defecto
EXPOSE 8080

# Variable de entorno de puerto por defecto
ENV PORT=8080

# Comando de ejecución
ENTRYPOINT ["./splitz-api"]

