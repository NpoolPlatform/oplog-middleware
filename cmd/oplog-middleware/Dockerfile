FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY oplog-middleware /app

ENTRYPOINT ["/app/oplog-middleware"]

EXPOSE 50780 50781 50782 50783

CMD ["run"]
