FROM alpine:3.12.1

CMD mkdir /app
COPY bin/svc /app/svc
RUN ls -la /app/
ENTRYPOINT ["/app/svc","-config-file","/app/config.json"]