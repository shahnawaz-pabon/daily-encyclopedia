<div align="center">
  <a href="https://www.php.net/">
    <img alt="ubuntu" src="../logos/php.png"/>
  </a>
  <h1>PHP and Laravel</h1>
</div>

# Table of Contents

- [Laravel Commands](#laravel-commands)
- [Chunk arrays in blade](chunk-arrays-in-blade)

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
php artisan make:migration add_votes_to_users_table --table=users // create migration to specific table
php artisan migrate --path=/database/migrations/2021_06_01_123739_category_name.php // run specific migration
php artisan make:test Http/Controllers/App/V5/MedicineOrderController
php artisan make:job JobName
php artisan make:command CommandName
php artisan make:middleware Role // add custom middleware
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
