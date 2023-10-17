FROM golang:1.19

ARG SERVICE
ARG PORT

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD CartService ./CartService
ADD CheckoutService ./CheckoutService
ADD PaymentService ./PaymentService
ADD PaymentService ./PaymentService
ADD ProductCatalogService ./ProductCatalogService
ADD ShippingService ./ShippingService

RUN CGO_ENABLED=0 GOOS=linux  go build -o main ./$SERVICE/cmd/*.go

EXPOSE $PORT

CMD ["./main"]