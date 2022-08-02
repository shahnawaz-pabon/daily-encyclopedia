<div align="center">
  <a href="https://www.djangoproject.com/">
    <img alt="django" src="../logos/django.png"/>
  </a>
  <h1>Django</h1>
</div>

# Table of Contents

- [Run a django project via virtual environment](#run-a-django-project-via-virtual-environment)

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
