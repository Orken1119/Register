# Используем базовый образ Golang
FROM golang:1.23

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы проекта в контейнер
COPY . /app

# Загружаем зависимости
RUN go mod download

# Компилируем приложение
RUN go build -tags netgo -ldflags '-s -w' -o app

# Открываем нужный порт (например, 8080)
EXPOSE 8080

# Запускаем приложение
CMD ["./app"]