FROM alpine:latest

# Create a non-root user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

# Set environment variable for app environment
ARG APP_ENV
ENV APP_ENV=$APP_ENV

WORKDIR /app

COPY twauter /app/twauter
COPY appconfig.json /app/appconfig.json

RUN chmod 755 /app/twauter && chmod 644 /app/appconfig.json && \
    chown -R nonroot:nonroot /app && chmod -R 775 /app

EXPOSE 8080

USER nonroot:nonroot

CMD ["./twauter"]
