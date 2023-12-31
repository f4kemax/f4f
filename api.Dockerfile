# Используем базовый образ Go
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Устанавливаем зависимости и собираем приложение
RUN go mod download
RUN go build -o myapp .

# Определяем команду, которая будет запущена при запуске контейнера
CMD ["./myapp"]

EXPOSE 8080
