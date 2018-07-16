FROM alpine
RUN apk update && apk upgrade
# Change timezone
RUN apk add --no-cache tzdata
ENV TZ Europe/Moscow
# Add binary
RUN mkdir /app
WORKDIR /app
COPY radio-t-bot .
EXPOSE 80
ENTRYPOINT [ "/app/radio-t-bot" ]