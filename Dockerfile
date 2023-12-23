# Start van een officieel Golang-image met de specifieke versie.
FROM golang:1.21.5 as builder

# Stel de werkdirectory in de container in.
WORKDIR /app

# Kopieer de Go Modules-manifesten en download de afhankelijkheden.
# Dit wordt gedaan vóór het kopiëren van de broncode om de cache te benutten.
COPY go.mod go.sum ./
RUN go mod download

# Kopieer de broncode naar de container.
COPY . .

# Bouw de Go-app als een statisch gelinkt bestand.
# Dit is belangrijk voor een kleinere, meer veilige en efficiënte container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scheduling-service .

# Start een nieuwe fase vanaf een kleinere image om de grootte te minimaliseren.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Kopieer het binaire bestand van de builder-fase naar deze nieuwe image.
COPY --from=builder /app/scheduling-service .

# Exposeer de poort die door je applicatie wordt gebruikt.
EXPOSE 50051

# Voer de binaire uitvoer uit.
CMD ["./scheduling-service"]
