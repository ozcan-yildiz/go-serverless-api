# Build Stage
FROM golang:1.24-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
# Lambda, Linux/AMD64 mimarisi bekler (Amazon Linux tabanlıdır)
RUN GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

# Final Stage (Lambda Runtime Environment)
FROM public.ecr.aws/lambda/provided:al2
COPY --from=builder /app/bootstrap ${LAMBDA_RUNTIME_DIR}/bootstrap
CMD [ "bootstrap" ]