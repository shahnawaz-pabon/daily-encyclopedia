<div align="center">
  <a href="https://www.djangoproject.com/">
    <img alt="django" src="../logos/django.png"/>
  </a>
  <h1>Django</h1>
</div>

# Table of Contents

- [Run a django project via virtual environment](#run-a-django-project-via-virtual-environment)
- [SQLite things](#sqlite-things)
- [Add cronjob in django](#add-cronjob-in-django)
- [Show active cronjobs](#show-active-cronjobs)
- [Remove cronjob in django](#remove-cronjob-in-django)

## Run a django project via virtual environment

Suppose project name is `django-test` having the `requirements.txt` file.

```sh
$ cd django-test
$ python3.8 -m venv venv
$ source venv/bin/activate
$ pip install --upgrade pip
$ pip install -r requirements.txt
```

Make sure you've run the migrations

```bash
$ python manage.py migrate
```

and finally run the project:

```sh
$ python manage.py runserver
```

You'll be able to see starting development server at **[http://127.0.0.1:8000/](http://127.0.0.1:8000/)**

<br>

## SQLite things

By default, `django` configuration uses SQLite. After creating django project you will be able to see `db.sqlite3` file in the project directory. To see tables of that db, you can run the following commands from the terminal

```sh
$ cd project_folder
$ sqlite3 db.sqlite3
$ .tables
```

Don't forget to have [sqlite](https://www.sqlite.org/download.html) in your OS if you want to run sqlite from your terminal.

To see sqlite's db tables via the client, you can download [DB Browser](https://sqlitebrowser.org/) and open **`db.sqlite3`** file.

<br>

## Add cronjob in django

Run this command to add all the defined CRONJOBS to crontab(\*nix cron utility). Make sure to run this command every time CRONJOBS is changed in any way.

```sh
python manage.py crontab add
```

<br>

## Show active cronjobs

To get all the active CRONJOBS, run the following command

```sh
python manage.py crontab show
```

<br>

## Remove cronjob in django

To remove all the defined CRONJOBS from crontab, run the following command

```sh
python manage.py crontab remove
```

**IMPORTANT: Make sure you've created `scheduled_job.log` this file into the home directory**
