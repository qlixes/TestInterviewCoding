<?php

use App\Http\Controllers\UserController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::middleware("auth:sanctum")->group(function () {

    Route::controller(UserController::class)->group(function () {
        Route::get("/users/{id}", "view");
        Route::post("/register", "register");
        Route::post("/login", "login");
        Route::post("/logout", "logout");
    });

    Route::controller(PostController::class)->group(function () {
        Route::get("/posts", "index");
        Route::get("/posts/{id}", "view");
        Route::post("/posts", "store");
        Route::patch("/posts/{id}", "update");
        Route::delete("/posts/{id}", "delete");
    });
});
