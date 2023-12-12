<div align="center">
  <a href="https://fastapi.tiangolo.com/">
    <img height="96" width="96" alt="fastapi" src="../logos/fastapi.svg"/>
  </a>
  <h1>FastAPI</h1>
</div>

# Table of Contents

- [Dockerize FastAPI application with newrelic](#dockerize-fastapi-application-with-newrelic)

## Dockerize FastAPI application with newrelic

**`Dockerfile`**

```Dockerfile
# set base image (host OS)
FROM python:3.8-buster

# set work directory
WORKDIR /app

# set env variables
## Prevents Python from writing pyc files to disc (equivalent to python -B)
ENV PYTHONDONTWRITEBYTECODE 1
## PYTHONUNBUFFERED: Prevents Python from buffering stdout and stderr (equivalent to python -u)
ENV PYTHONUNBUFFERED 1

RUN pip install --no-cache-dir newrelic

# copy the dependencies file to the working directory
COPY requirements.txt .

# install dependencies
RUN pip install -r requirements.txt

ENV NEW_RELIC_LOG=stderr \
    NEW_RELIC_LOG_LEVEL=info \
    NEW_RELIC_ENABLED=true \
    NEW_RELIC_DISTRIBUTED_TRACING_ENABLED=true \
    NEW_RELIC_LICENSE_KEY=<license_key> \
    NEW_RELIC_APP_NAME="app-name"

# copy project
COPY . .

CMD ["newrelic-admin", "run-program", "gunicorn", "app.main:app", "--workers 2", "--worker-class", "uvicorn.workers.UvicornWorker", "--bind", "0.0.0.0"]
```

<br>

**`docker-compose.yml`**

```yml
# docker-compose.yml
version: "3"
services:
  web:
    build: .
    # command: gunicorn app.main:app --workers 2 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0
    volumes:
      - .:/app
    ports:
      - 8001:8000
```


```python
from fastapi import FastAPI, APIRouter
from fastapi.exceptions import HTTPException
from pydantic import BaseModel
from smtplib import SMTP_SSL
from email.mime.text import MIMEText

router = APIRouter()


OWN_EMAIL = ""  # this should be the email which the email will be sent from
OWN_EMAIL_PASSWORD = ""  # this should be app password


class EmailBody(BaseModel):
    to: str
    subject: str
    message: str


@router.post("/send-email")
async def send_email(body: EmailBody):
    try:
        msg = MIMEText(body.message, "html")
        msg['Subject'] = body.subject
        msg['From'] = f'Denolyrics <{OWN_EMAIL}>'
        msg['To'] = body.to

        port = 465  # For SSL

        # Connect to the email server
        server = SMTP_SSL("smtp.gmail.com", port)
        server.login(OWN_EMAIL, OWN_EMAIL_PASSWORD)

        # Send the email
        server.send_message(msg)
        server.quit()
        return {"message": "Email sent successfully"}

    except Exception as e:
        raise HTTPException(status_code=500, detail=e)
```