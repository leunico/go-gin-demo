<?php

$sortarr = [
    23, 45, 66, 2, 16, 77, 21, 34, 24, 8, 49
];

// O(n^2), O(1)
function maopao(array $arr): array
{
    $count = count($arr);
    for ($i=1; $i < $count; $i++) { 
        for ($j=0; $j < $count-$i; $j++) { 
            if ($arr[$j] > $arr[$j+1]) {
                // $temp = $arr[$j];
                // $arr[$j] = $arr[$j+1];
                // $arr[$j+1] = $temp;
                list($arr[$j], $arr[$j+1]) = [$arr[$j+1], $arr[$j]];
            }
        }
    }
    return $arr;
}

// 同上
function xuanzhe(array $arr)
{
    $count = count($arr);
    for ($i=0; $i<$count-1; $i++) {
        $minIndex = $i;
        for ($j=$i+1; $j < $count; $j++) { 
            $minIndex = $arr[$minIndex] > $arr[$j] ? $j : $minIndex;
        }
        list($arr[$i], $arr[$minIndex]) = [$arr[$minIndex], $arr[$i]];
    }
    return $arr;
}

// 同上
function charu(array $arr)
{
    $count = count($arr);
    for ($i=1; $i<$count; $i++) { 
        for ($j=$i; $j>0; $j--) { 
            if ($arr[$j] < $arr[$j-1]) {
                list($arr[$j], $arr[$j-1]) = [$arr[$j-1], $arr[$j]];
            }
        }
    }
    return $arr;
}

function kuaishu(array $arr)
{
    if (count($arr) < 2) {
        return $arr;
    }
    $left = $right = [];
    $mid = $arr[1];
    for ($i=0; $i < count($arr); $i++) {
        if ($arr[$i] > $mid) {
            $right[] = $arr[$i];
        } elseif ($arr[$i] < $mid) {
            $left[] = $arr[$i];
        }
    }
    $left = kuaishu($left);
    $right = kuaishu($right);
    return array_merge($left, [$mid], $right);
}

function fdgkuaishu(array $arr)
{
    $sarr[0] = ['left' => 0, 'right' => count($arr) - 1];
    $i = 0;
    $n = 1;
    while ($i < $n) {
        $left = $sarr[$i]['left'];
        $right = $sarr[$i]['right'];

    }

}

// 双指针
function szz($a, $b)
{
    $arr = [];
    $i = $j = $k = 0;
    while (isset($a[$i]) && isset($b[$j])) {
        if ($a[$i] > $b[$j]) {
            $arr[$k++] = $b[$j++];
        } else {
            $arr[$k++] = $a[$i++];
        }
    }

    while (isset($a[$i])) {
        $arr[$k++] = $a[$i++];
    }

    while (isset($b[$j])) {
        $arr[$k++] = $b[$j++];
    }
    return $arr;
}

$a = [2,4,8,10,11,17,18];
$b = [1,3,12,20,25];

var_dump(szz($a, $b));
// var_dump(kuaishu($sortarr));

class MergeSort
{
    public function __construct(array $arr)
    {
        $this->mSort($arr, 0, count($arr) - 1);
        var_dump($arr);
    }

    public function mSort(&$arr, $left, $right)
    {
        if ($left < $right) {
            $cet = floor(($left+$right)/2);
        }
    }
}