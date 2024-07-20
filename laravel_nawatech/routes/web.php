<?php

use App\Http\Controllers\InterviewController;
use Illuminate\Support\Facades\Route;

Route::get("/", [InterviewController::class, 'index']);
