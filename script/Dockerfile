# Go base image
FROM golang:1.16-alpine

# Çalışma dizinini ayarla
WORKDIR /app

# Kaynak kodu kopyala
COPY . .

# Bağımlılıkları indir (eğer varsa)
RUN go mod download

# Script'i derle
RUN go build -o script main.go

# Script'i çalıştır
CMD ["./script"]
