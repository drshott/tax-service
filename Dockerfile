# Stage 1: Build React Frontend
FROM node:18 AS frontend
WORKDIR /app/frontend
COPY tax-frontend/ ./
RUN npm install && npm run build

# Stage 2: Build Go Backend
FROM golang:1.23 AS backend
WORKDIR /app
COPY taxCalculator/ ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o taxCalculator main.go

# Stage 3: Create Final Image
FROM alpine:latest
WORKDIR /root/
COPY --from=backend /app/taxCalculator ./
COPY --from=frontend /app/frontend/build  ./frontend

EXPOSE 8080

CMD ["./taxCalculator"]
