<?php

namespace App\Services;

final class AhassService
{
    protected $json;

    function source(string $raw)
    {
        $this->json = json_decode($raw);

        return $this;
    }

    function filter($code)
    {
        $collection = collect($this->json->data);

        return $collection->where('code', $code)->first();
    }
}
