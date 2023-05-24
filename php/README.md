<div align="center">
  <a href="https://www.php.net/">
    <img alt="ubuntu" src="../logos/php.png"/>
  </a>
  <h1>PHP and Laravel</h1>
</div>

# Table of Contents

- [Dockerize Laravel application with newrelic](#dockerize-laravel-application-with-newrelic)
- [Laravel Commands](#laravel-commands)
- [Chunk arrays in blade](#chunk-arrays-in-blade)
- [Get specific column](#get-specific-column)
- [Upload multiple files to the third party API](#upload-multiple-files-to-the-third-party-api)

## Dockerize Laravel application with newrelic

**`Dockerfile`**

```Dockerfile
FROM ubuntu:bionic

RUN apt-get update && apt-get -y upgrade && apt-get install -y software-properties-common tzdata
RUN echo "Asia/Karachi" > /etc/timezone && rm -f /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
RUN add-apt-repository -y ppa:ondrej/php
RUN apt-get update
RUN apt install -y wget php8.0 php8.0-fpm php8.0-common php8.0-mysql php8.0-xml php8.0-dev php8.0-xmlrpc php8.0-curl php8.0-mbstring php8.0-zip
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php composer-setup.php
RUN mv composer.phar /usr/local/bin/composer
RUN apt-get install -y php8.0-mongodb unzip
RUN apt-get install -y php8.0-gd
#RUN apt-get install -y mysql-client
WORKDIR /app
COPY . .



## New relic installation
#RUN wget -r -l1 -nd -A"linux.tar.gz" https://download.newrelic.com/php_agent/release/
RUN wget https://download.newrelic.com/php_agent/release/newrelic-php5-9.18.1.303-linux.tar.gz

# Unzip it and enter it
RUN gzip -dc newrelic*.tar.gz | tar xf -
WORKDIR '/app/newrelic-php5-9.18.1.303-linux'
RUN ls -l

# Delete any previous New Relic extension
RUN rm -f /usr/lib/php/20200930/newrelic.so

# Copy the downloaded extension to your php's extension directory
# If is ZTS is compiled in, the file name should have "-zts" appended. See examples below.
RUN cp ./agent/x64/newrelic-20200930.so /usr/lib/php/20200930/newrelic.so

# Install New Relic's daemon
RUN cp ./daemon/newrelic-daemon.x64 /usr/bin/newrelic-daemon

WORKDIR /app

# Copy your newrelic.ini custom file to your php conf.d dir. Make sure your php reads .ini files from this directory.
RUN cp /app/newrelic.ini /etc/php/8.0/fpm/conf.d/20-newrelic.ini
RUN cp /app/newrelic.ini /etc/php/8.0/cli/conf.d/20-newrelic.ini
```

<br>

**`init.sh`**

```sh
#!/bin/bash
#mysql -uroot -pexample --host mysql -e "CREATE DATABASE IF NOT EXISTS maya_local;"
composer install --ignore-platform-reqs
#php artisan sentry:publish --dsn=http://cbe2201f8f374c9785b10697886ce902@3.0.117.15:9000/3
php artisan optimize:clear
php artisan cache:clear
php artisan config:clear
php artisan key:generate
php artisan passport:keys
#php artisan migrate
#php artisan db:seed

#New relic
/usr/bin/newrelic-daemon start

# Server application
php artisan serve --host 0.0.0.0
```

<br>

**`docker-compose.yml`**

```yml
version: "3"

services:
  client:
    build: ./
    command: tail -f /dev/null
    ports:
      - 5500:8000
    #    depends_on:
    #      - mysql
    volumes:
      - ./:/app
    entrypoint: ./init.sh

  #  mysql:
  #    image: mariadb
  #    environment:
  #      MYSQL_ROOT_PASSWORD: example
  #      MYSQL_USER: root
  #    volumes:
  #      - ./mysqldbdir:/var/lib/mysql
  #    ports:
  #      - 3307:3306

  redis:
    image: redis:latest
    ports:
      - 6379:6379

networks:
  default:
    external:
      name: custom-bot
```

## Laravel Commands

```sh
php artisan make:factory MedicineOrderDetailsFactory --model=MedicineOrderDetails
php artisan make:factory MedicineOrderFactory --model=MedicineOrder
php artisan make:test MedicineOrderController
php artisan make:model MedicineOrder -m
php artisan make:model MedicineOrderDetails -m
php artisan make:model Pharmacy -m
php artisan make:model Models/Pharmacy -m
php artisan config:clear && php artisan config:cache
# create migration to specific table
php artisan make:migration add_votes_to_users_table --table=users
# run specific migration
php artisan migrate --path=/database/migrations/2021_06_01_123739_your_file.php
php artisan make:test Http/Controllers/App/V5/MedicineOrderController
php artisan make:job JobName
php artisan make:command CommandName
# add custom middleware
php artisan make:middleware Role
# rollback specific migration
php artisan migrate:rollback --path=/database/migrations/2022_02_03_144724_your_file.php
# rollback 1 step
php artisan migrate:rollback --step=1
```

## Chunk arrays in blade

Suppose you're return an array from the controller, you want to show maximum 3 items in a row in your blade view. Then you can do something like this:

```php
@foreach(array_chunk($your_array_got_from_the_controller, 3) as $chunk)
    <div class="row">
        @foreach($chunk as $item)

            <div class="col-md-4">
                <h1>{{ $item }}</h1>
            </div>

        @endforeach
    </div>
@endforeach
```

## Get specific column

Get specific column in laravel eloquent:

```php
// User is the name of a model
User::where('id', $id)->first()->field_name
```

## Upload multiple files to the third party API

```php

use GuzzleHttp\Client;
use GuzzleHttp\Psr7;

public function __construct()
{
    $this->client = new Client([
        'base_uri' => 'https://base_url'
    ]);
}

public function processFiles(){
    $process_urls = $orderData['file_urls'][0];
    $prescription_urls = preg_replace("^[\"\[\] ]^", "", $process_urls);
    $prescription_urls = explode(',', $prescription_urls); // for multiple urls

    $multipartArrays = [];
    $fileNameArrays = [];

    foreach ($prescription_urls as $url){
        // Use basename() function to return the base name of file
        $file_name = 'files/' . basename($url);
//            dd($file_name);
        file_put_contents($file_name, file_get_contents($url)); // save file in public folder
        $multipartObj = [
            'name' => 'prescription_files',
            'contents' => fopen(public_path($file_name), 'rb')
        ];
        array_push($multipartArrays, $multipartObj);
//            dd(File::exists(public_path($file_name))); // check whether file exists
//            dd(fopen(public_path($file_name), 'rb'));//verify that file exists
        array_push($fileNameArrays, $file_name);

    }

    [$statusCode, $response] = $this->uploadMultipleFiles($multipartArrays);

    return $response
}

public function uploadMultipleFiles($multipartArrays){
  $apiRequest = $this->client->request('POST', 'upload/files',
    [
        'headers' => [
            'Authorization' => $this->getAuthorizationToken()
        ],
        'multipart' => $multipartArrays,
        'verify' => false
    ]);
return [json_decode($apiRequest->getStatusCode()), json_decode($apiRequest->getBody())];
}
```
