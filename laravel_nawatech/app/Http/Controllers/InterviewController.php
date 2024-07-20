<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;

class InterviewController extends Controller
{
    function index(Request $request)
    {
        $file_book = Storage::disk('assets')->get('18HKxBeM.json');
        $file_ahass = Storage::disk('assets')->get('2fBXAWqh.json');

        $book = $this->book->source($file_book);
        $workshop = $this->ahass->source($file_ahass);

        $response = [];
        $response['status'] = $book->getStatus();
        $response['mesage'] = $book->getMessage();

        $result = [];

        foreach ($book->getData() as $data) {

            $data->workshop = $workshop->filter($data->booking->workshop->code);

            $result[] = $this->book->generate(
                $data->name,
                $data->email,
                $data->booking->booking_number,
                $data->booking->book_date,
                ($data->workshop->code) ?? '',
                ($data->workshop->name) ?? '',
                ($data->workshop->address) ?? '',
                ($data->workshop->phone_number) ?? '',
                ($data->workshop->distance) ?? 0,
                $data->booking->motorcycle->ut_code,
                $data->booking->motorcycle->name
            );
        }

        $final = (collect($result))->sortBy('ahass_distance');


        $response['data'] = $final->values()->all();
        return response()->json($response);
    }
}
