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
            $arr[] = $b[$j++];
        } else {
            $arr[] = $a[$i++];
        }
    }

    while (isset($a[$i])) {
        $arr[] = $a[$i++];
    }

    while (isset($b[$j])) {
        $arr[] = $b[$j++];
    }
    return $arr;
}

$a = [2,4,8,10,11,17,18];
$b = [1,3,12,20,25];

// var_dump(szz($a, $b));
// var_dump(kuaishu($sortarr));
// new MergeSort($sortarr);

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
            $this->mSort($arr, $left, $cet);
            $this->mSort($arr, $cet + 1, $right);
            $this->mergeArray($arr, $left, $cet, $right);
        }
    }

    public function mergeArray(&$arr, $left, $cet, $right)
    {
        echo '| ' . $left . ' - ' . $cet . ' - ' . $right . ' - ' . implode(',', $arr) . PHP_EOL;
        $a_i = $left;
        $b_i = $cet + 1;
        $temp = [];

        while ($a_i <= $cet && $b_i <= $right)
        {
            if ($arr[$a_i] > $arr[$b_i]) {
                $temp[] = $arr[$b_i++];
            } else {
                $temp[] = $arr[$a_i++];
            }
        }

        while ($a_i <= $cet) 
        {
            $temp[] = $arr[$a_i++];
        }

        while ($b_i <= $right) 
        {
            $temp[] = $arr[$b_i++];
        }
var_dump($temp);
        for ($i=0; $i < count($temp); $i++) { 
            $arr[$left + $i] = $temp[$i];
        }
    }
}

// 二分查找 log2n 空间O(1)/Olog2n
function efcz($arr, $low, $top, $target)
{
    if ($low <= $top) {
        $mid = floor(($low + $top)/2);
        if ($target > $arr[$mid]) {
            return efcz($arr, $mid+1, $top, $target);
        } elseif ($target < $arr[$mid]) {
            return efcz($arr, $low, $mid-1, $target);
        } else {
            return $mid;
        }
    } else {
        return -1;
    }
}

// 非递归
function efcz2($arr, $low, $top, $target)
{
    while ($low <= $top)
    {
        $mid = floor(($low + $top)/2);
        if ($target > $arr[$mid]) {
            $low = $mid + 1;
        } elseif ($target < $arr[$mid]) {
            $top = $mid - 1;
        } else {
           return $mid;
        }
    }
    return -1;
}
// var_dump(sizeof($a), count($a));
var_dump(efcz2($a, 0, sizeof($a), 10));