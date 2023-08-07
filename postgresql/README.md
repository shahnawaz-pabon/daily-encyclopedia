<div align="center">
  <a href="https://www.postgresql.org/">
    <img alt="postgresql" src="../logos/postgis.png"/>
  </a>
  <h1>PostgreSQL</h1>
</div>

# Table of Contents

- [List all databases](#list-all-databases)
- [List all table names of the current database](#list-all-table-names-of-the-current-database)

## List all databases

```sql
SELECT datname FROM pg_database WHERE datistemplate = false;
```

<br>

## List all table names of the current database

```sql
SELECT table_name
  FROM information_schema.tables
 WHERE table_schema='public'
   AND table_type='BASE TABLE';
```

<br>

```yml
version: "3.8"
services:
  db:
    image: postgres:14-alpine
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
volumes:
  postgres-data:
```
