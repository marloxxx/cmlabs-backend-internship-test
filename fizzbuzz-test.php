<?php

function fizzbuzzWithConstraint($N) {
    for ($i = 1; $i <= $N; $i++) {
        if ($i % 3 === 0 && $i % 5 === 0) {
            echo "FizzBuzz\n";
        } elseif ($i % 3 === 0) {
            echo "Fizz\n";
        } elseif ($i % 5 === 0) {
            echo "Buzz\n";
        } else {
            echo $i . "\n";
        }
    }
}

// Input dari pengguna
$N = (int) readline("Masukkan nilai N: ");
fizzbuzzWithConstraint($N);

?>
