<div align="center">
  <a href="https://devhints.io/bash">
    <img alt="bash" src="../logos/bash.png"/>
  </a>
  <h1>Bash</h1>
</div>

# Table of Contents

- [Backup a Docker MySQL database with CRON](#backup-a-docker-mysql-database-with-cron)

## Backup a Docker MySQL database with CRON

Edit the crontab by this command **`crontab -e`** and add this line:

```sh
* * * * * cd /path/to/script && bash ./backup.sh >> /path/to/log/backup.log
```

**`backup.sh`** file should be like this:

```sh
#!/bin/bash
BACKUP_DIR="/home/backups"
BACKUP_FILE_NAME="$(date +"%d-%m-%y-%H%M%S.sql.gz")"

docker exec <CONTAINER_NAME> bash -c 'exec mysqldump --databases "$MYSQL_DATABASE" -h<DOCKER_MYSQL_SERVICE_NAME> -u"$MYSQL_USER" -p"$MYSQL_PASSWORD"' > gzip > "$BACKUP_DIR"/"$BACKUP_FILE_NAME";
```
