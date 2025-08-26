FROM alpine:latest

# Set permission dan working directory
WORKDIR /app
COPY build/linux/app_v1.2.0 .

RUN chmod +x app_v1.2.0

# Jalankan aplikasi
CMD ["./app_v1.2.0"]
