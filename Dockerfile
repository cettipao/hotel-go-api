FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o /app/mvc-go .

# Exponer el puerto 8000
EXPOSE 8000

# Ejecutar la aplicación
CMD ["/app/mvc-go"]

