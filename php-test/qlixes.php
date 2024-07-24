<?php

namespace Qlixes;

class Qlixes
{
    function switch_number(int $x, int $y): array
    {
        $a = $x;
        $b = $y;

        if ($x < $y) {
            $a = $y;
            $b = $x;
        }

        $a = $b - $a;
        $b = $b - $a + $a;
        $a = $b - $a;

        return [$a, $b];
    }
}
