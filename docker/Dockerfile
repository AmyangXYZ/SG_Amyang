# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sg-amyang
FROM alpine

ADD templates/ /templates
ADD sg-amyang /

CMD [ "/sg-amyang" ]