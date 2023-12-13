FROM alpine
ADD order-api-api /order-api-api
ENTRYPOINT [ "/order-api-api" ]
