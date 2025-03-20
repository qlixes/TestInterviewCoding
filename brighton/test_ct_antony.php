<?php

$numbers = 77777;

echo $numbers;
echo "<br/>";

echo number_format($numbers);
echo "<br/>";

$nama = "Antony";

$name = "{$numbers}CrossTechnoDeveloper";

echo $name;
echo "<br/>";

echo str_replace($name, "/^[oO]$/", "A");
echo "<br/>";

$data = [];

$data[1] = "Surabaya";
$data[2] = "Jakarta";
$data[3] = "Semarang";
$data[4] = "Makassar";

echo json_encode($data);
echo "<br/>";

$data[3] = "Samarinda";
$data[5] = "Aceh";

echo json_encode($data);
echo "<br/>";

$buah = [];

$buah = [
    "fruit" => "Orange",
    "animal" => "Dog",
    "bird" => "Eagle",
    "food" => "Rice",
];

echo json_encode($buah);
echo "<br/>";

$buah['animal'] = "Cat";

echo json_encode($buah);
echo "<br/>";

echo json_encode(array_merge($data, $buah));
echo "<br/>";

function gambarPattern(int $row, int $length)
{
    $chars = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];

    $lengths = floor($length / 2);

    $low = array_slice($chars, 0, $lengths);
    $high = array_slice($chars, 26 - $lengths, 25);

    $result = [];

    for ($j = 0; $j < $length; $j++) {
        $start = ($j % 2 == 0) ? $low : $high;

        for ($i = 0; $i <= floor($row / $length); $i++) {
            if($j % 2 == 0)
            {
                if($i % 2 == 1) {
                    $start = array_merge($start, $low);
                } else {
                    $start = array_merge($start, $high);
                }
            } else {
                if($i % 2 == 1) {
                    $start = array_merge($start, $high);
                } else {
                    $start = array_merge($start, $low);
                }
            }
        }

        $result[$j] = $start;
    }

    return $result;
}

echo json_encode(gambarPattern(12, 6));
echo "<br/>";

class Vehicle
{
    private string $name;
    private int $kapasitas_mesin;
    private int $roda;

    function __construct(string $name, int $kapasitas_mesin, int $roda)
    {
        $this->name = strtoupper($name);
        $this->kapasitas_mesin = $kapasitas_mesin;
        $this->roda = $roda;
    }

    function show()
    {
        return "vehichle with name {$this->name} dengan kapasitas_mesin {$this->kapasitas_mesin} dan roda {$this->roda}";
    }
}

$init = new Vehicle("Mobile", 2, 4);

echo $init->show();
echo "<br/>";

class Car extends Vehicle
{
    private string $name;
    private int $kapasitas_mesin;
    private int $roda;
    private string $merk;
    private int $tahun_pembuatan;
    private string $nomor_rangka;

    function __construct(string $name, int $kapasitas_mesin, int $roda, string $merk, int $tahun_pembuatan, string $nomor_rangka)
    {
        $this->name = strtoupper($name);
        $this->kapasitas_mesin = $kapasitas_mesin;
        $this->roda = $roda;
        $this->merk = strtoupper($merk);
        $this->tahun_pembuatan = $tahun_pembuatan;
        $this->nomor_rangka = strtoupper($nomor_rangka);
    }

    function show()
    {
        return "vehichle with name {$this->name} dengan kapasitas_mesin {$this->kapasitas_mesin} dan roda {$this->roda} dan merk {$this->merk} dan produksi tahun {$this->tahun_pembuatan} dengan nomor rangka {$this->nomor_rangka}";
    }
}

$mobil1 = new Car("Mobil", 1, 4, "Honda", 2002, "RK123");
echo $mobil1->show();
echo "<br/>";

$mobil2 = new Car("Mobi2", 1, 4, "Toyota", 2002, "QRIS0101");
echo $mobil2->show();
echo "<br/>";
