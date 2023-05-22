<div align="center">
  <a href="https://spring.io/">
    <img alt="spring" src="../logos/spring.png" height="96"/>
  </a>
  <h1>Spring</h1>
</div>

# Table of Contents

- [Dockerize spring boot application with redis and mssql](#dockerize-spring-boot-application-with-redis-and-mssql)
- [Parse datetime with specific locale](#parse-datetime-with-specific-locale)

## Dockerize spring boot application with redis and mssql

**Dockerfile**

```Dockerfile
FROM openjdk:17-jdk-alpine AS build
WORKDIR /app
COPY build.gradle .
COPY gradlew .
COPY settings.gradle .
COPY gradle gradle
COPY src src
RUN chmod +x gradlew
RUN ./gradlew bootJar

FROM openjdk:17-jdk-alpine
WORKDIR /app
COPY --from=build /app/build/libs/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
```

**docker-compose.yml**

```yml
version: "3.8"
services:
  app:
    build: .
    ports:
      - "5000:5000"
    depends_on:
      - db
      - redis
  db:
    restart: always
    image: mcr.microsoft.com/mssql/server:2019-latest
    user: root
    environment:
      - ACCEPT_EULA=Y
      - SA_PASSWORD=shahnawaz@123
    ports:
      - "1433:1433"
    volumes:
      - ./data:/var/opt/mssql/data
  redis:
    image: redis:latest
    ports:
      - "6379:6379"
```

## Parse datetime with specific locale

```java
// Import
import java.util.Date;
import java.text.DateFormat;
import java.util.Locale;

Locale localeEn = new Locale("en", "US");
DateFormat dateFormatEn = DateFormat.getDateTimeInstance(DateFormat.DEFAULT, DateFormat.SHORT, localeEn);
String dateEn = dateFormatEn.format(new Date());

Locale localeBn = new Locale("bn", "BD");
DateFormat dateFormatBn = DateFormat.getDateTimeInstance(DateFormat.DEFAULT, DateFormat.SHORT, localeBn);
String dateBn = dateFormatBn.format(new Date());
```
