<div align="center">
  <a href="https://postgis.net/">
    <img alt="postgis" src="../logos/postgis.png"/>
  </a>
  <h1>PostGIS</h1>
</div>

# Table of Contents

- [Setup POSTGIS in Docker](#setup-postgis-in-docker)

## Setup POSTGIS in Docker

```sh
docker run -d -p 5432:5432 --name db kartoza/postgis:13.0
```

**Postgis docker’s Credentials:**

```sh
Username: docker
Password: docker
Database: gis
```
