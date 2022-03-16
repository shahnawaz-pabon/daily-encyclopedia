<div align="center">
  <a href="https://www.php.net/">
    <img alt="ubuntu" src="../logos/php.png"/>
  </a>
  <h1>PHP and Laravel</h1>
</div>

# Table of Contents

- [Laravel Commands](#laravel-commands)
- [Chunk arrays in blade](#chunk-arrays-in-blade)
- [Get specific column](#get-specific-column)
- [Upload multiple files to the third party API](#upload-multiple-files-to-the-third-party-api)

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
