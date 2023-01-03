FROM alpine:latest
RUN mkdir /app
COPY logApp /app
CMD [ "/app/logApp" ]
