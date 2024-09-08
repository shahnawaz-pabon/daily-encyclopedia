# Table of Contents

- [MariaDB Docker Setup](#create-kinesis-stream-from-aws-cli)


## MariaDB Docker Setup

Create `docker-compose.yml`:

```yml
version: '3.8'

services:
  mariadb:
    image: mariadb:latest
    environment:
      MYSQL_ROOT_PASSWORD: my-secret-pw
      MYSQL_DATABASE: mydb
      MYSQL_USER: myuser
      MYSQL_PASSWORD: mypassword
    ports:
      - "3306:3306"
    volumes:
      - mariadb-data:/var/lib/mysql

volumes:
  mariadb-data:
```

Start the MariaDB Container:

```sh
docker compose up -d
```

Install MySQL Client:

```sh
# Debian/Ubuntu
sudo apt-get install mysql-client

# macOS (Homebrew)
brew install mysql
```

Test the Connection:

```sh
mysql -h 127.0.0.1 -P 3306 -u myuser -p
```

Stop and Remove Containers:

```sh
docker compose down
```

**Important: Problem with pip install mariadb - mariadb_config not found**

Solution:

```sh
sudo apt-get update -y
sudo apt-get install -y libmariadb-dev
```