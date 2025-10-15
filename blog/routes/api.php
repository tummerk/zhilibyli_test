<?php

use App\Http\Controllers\PostController;


Route::get('posts/count', [PostController::class, 'count']);
Route::apiResource('posts', PostController::class);
