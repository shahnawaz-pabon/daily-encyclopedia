<div align="center">
  <a href="https://devhints.io/bash">
    <img alt="bash" src="../logos/bash.png"/>
  </a>
  <h1>Bash</h1>
</div>

# Table of Contents

- [Backup a Docker MySQL database with CRON](#backup-a-docker-mysql-database-with-cron)
- [Taking data from the file and writing STDOUT to a new file](#taking-data-from-the-file-and-writing-stdout-to-a-new-file)
- [Pipe in the bash script](#pipe-in-the-bash-script)
- [Using nohup to put the task in the background](#using-nohup-to-put-the-task-in-the-background)
- [Find the PID](#find-the-pid)

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

<br>

## Taking data from the file and writing STDOUT to a new file

Let's say `sports.txt` has following lines

```sh
football
basketball
swimming
```

```sh
cat sports.txt 1> new_sports.txt
```

If you run `cat new_sports.txt` you'll see the above lines.

<br>

## Pipe in the bash script

Consider a file `animals.txt`

```sh
magpie, bird
emu, bird
kangaroo, marsupial
wallaby, marsupial
shark, fish
```

We would like to count animals in each group with including pipes in our Bash script. Here is the command:

```sh
cat animals.txt | cut -d " " -f 2 | sort | uniq -c
```

## Using nohup to put the task in the background

```sh
nohup my_command > my.log 2>&1 &
```

## Find the PID

```sh
ps -ef | grep "command name"
```
