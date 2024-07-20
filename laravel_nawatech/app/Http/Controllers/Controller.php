<?php

namespace App\Http\Controllers;

use App\Services\AhassService;
use App\Services\BookService;

abstract class BaseController
{
    //
}

class Controller extends BaseController
{
    protected $book, $ahass;

    function __construct(BookService $book, AhassService $ahass)
    {
        $this->book = $book;
        $this->ahass = $ahass;
    }
}
