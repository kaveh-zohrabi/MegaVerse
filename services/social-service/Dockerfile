FROM eclipse-temurin:21-jdk-alpine AS builder
WORKDIR /app
COPY pom.xml .
COPY src ./src
RUN apk --no-cache add maven && mvn clean package -DskipTests

FROM eclipse-temurin:21-jre-alpine
WORKDIR /root/
COPY --from=builder /app/target/*.jar app.jar
EXPOSE 8083
CMD ["java", "-jar", "app.jar"]
