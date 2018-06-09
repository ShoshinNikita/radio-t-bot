FROM alpine
# Add binary
RUN mkdir /app && mkdir /app/logs
WORKDIR /app
COPY app .
EXPOSE 80
ENTRYPOINT [ "/app/radio-t-bot" ]