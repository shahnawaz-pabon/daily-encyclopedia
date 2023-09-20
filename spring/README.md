<div align="center">
  <a href="https://spring.io/">
    <img alt="spring" src="../logos/spring.png" height="96"/>
  </a>
</div>

# Table of Contents

- [Dockerize spring boot application with redis and mssql](#dockerize-spring-boot-application-with-redis-and-mssql)
- [Setting Up Swagger Documentation](#setting-up-swagger-documentation)
- [Role-Based API Access in Spring Boot 3](#role-based-api-access-in-spring-boot-3)
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

<br>

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

## Role-Based API Access in Spring Boot 3

This guide demonstrates how to configure role-based API access in a Spring Boot application using OAuth2 and JWT tokens. It includes the steps to include dependencies, create a security configuration, and verify JWT token roles.

### Step 1: Include Dependencies

First, include the following dependency in your `build.gradle` file:

```groovy
// Add OAuth2 Resource Server dependency
implementation 'org.springframework.boot:spring-boot-starter-oauth2-resource-server:3.1.3'
```

You can find this package on [Maven Repository](https://mvnrepository.com/artifact/org.springframework.boot/spring-boot-starter-oauth2-resource-server).

### Step 2: Create Security Configuration

Create a configuration file named SecurityConfig.java and [add the following code](https://gist.github.com/shahnawaz-pabon/162a99f43f4d8b47fa9437e758b003df).

### Step 3: Verify JWT Token Roles

In this code, all APIs with the pattern `"/api/v1/**"` require a JWT token with the `admin` role. You can check whether a JWT token has specific roles by inspecting the roles key in the token payload. You can decode and verify JWT tokens using tools like [jwt.io](https://jwt.io/).

That's it! You've successfully configured role-based API access in your Spring Boot application.

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
