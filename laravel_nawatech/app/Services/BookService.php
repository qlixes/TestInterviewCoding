<?php

namespace App\Services;

final class BookService
{
    protected $json;

    function source(string $raw)
    {
        $this->json = json_decode($raw);

        return $this;
    }

    function getStatus()
    {
        return $this->json->status;
    }

    function getMessage()
    {
        return $this->json->message;
    }

    function getData()
    {
        return $this->json->data;
    }

    function generate($name, $email, $booking_number, $book_date, $ahass_code, $ahass_name, $ahass_address = '', $ahass_contact = '', $ahass_distance = 0, $motorcycle_ut_code, $motorcyle)
    {
        $result = [];
        $result['name']                 = $name;
        $result['email']                = $email;
        $result['booking_number']       = $booking_number;
        $result['book_date']            = $book_date;
        $result['ahass_code']           = $ahass_code;
        $result['ahass_name']           = $ahass_name;
        $result['ahass_address']        = $ahass_address;
        $result['ahass_contact']        = $ahass_contact;
        $result['ahass_distance']       = $ahass_distance;
        $result['motorcycle_ut_code']   = $motorcycle_ut_code;
        $result['motorcycle']           = $motorcyle;

        return $result;
    }
}
