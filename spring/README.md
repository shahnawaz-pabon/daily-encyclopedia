<div align="center">
  <a href="https://spring.io/">
    <img alt="spring" src="../logos/spring.png" height="96"/>
  </a>
</div>

# Table of Contents

- [Dockerize spring boot application with redis and mssql](#dockerize-spring-boot-application-with-redis-and-mssql)
- [Setting Up Swagger Documentation](#setting-up-swagger-documentation)
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

## Setting Up Swagger Documentation

### Include the `springdoc` dependency in your `build.gradle` file:

```groovy
implementation 'org.springdoc:springdoc-openapi-starter-webmvc-ui:2.2.0'
```
You can find this package on [Maven Repository](https://mvnrepository.com/artifact/org.springdoc/springdoc-openapi-starter-webmvc-ui).

### Create a Swagger Configuration File

Create a configuration file for Swagger, for example, SwaggerConfig.java, and [add the following content](https://gist.github.com/shahnawaz-pabon/7fa8ac82947c6bd1d54fcd33eecf4d2a):

### Access Swagger UI

You can access the Swagger UI by navigating to the following URL in your web browser, replacing your_port with your actual port number:

[http://localhost:your_port/swagger-ui/index.html](http://localhost:your_port/swagger-ui/index.html)

<br>

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
