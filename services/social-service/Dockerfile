FROM openjdk:21-jdk-slim AS builder
WORKDIR /app
COPY . .
RUN ./gradlew bootJar --no-daemon

FROM eclipse-temurin:21-jre-alpine
WORKDIR /root/
COPY --from=builder /app/build/libs/*.jar app.jar
EXPOSE 8083
CMD ["java", "-jar", "app.jar"]
