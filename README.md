# gofib - Fast Doubling Fibonacci Arbitrary Precision Calculator

# with Docker container for all environments

![Go Fibonacci!](doc/fibonacci.png)

This is a high-digit Fibonacci calculator in GO lang.  Same difficulty as Python.  Better performance than C++.

Records:

| Index | Time | Notes | Digits |
| :---: | :---: | :---: | :---: |
| F(10M) | 3.306s | Go! | 2,089,898 |
| F(10M) | 50.041s | Python3 | 2,089,898 |
| F(20M) | 7.958s | Go! | 4,179,774 |
| F(50M) | 27.119s | Go! (Intel) | 10,449,403 |
| F(60M) | 27.646s | Go! (Apple) | 12,539,280 |
| F(80M) | 27.592s | Go! (Apple) | 16,719,032 |
| F(100M) | 39.590s | Go! (Apple M1 Max) | 20,898,786 |
| F(100M) | 37.880s | Go! (Apple M2 Max) | 20,898,786 |
| F(100M) | 23.828s | Go! (Apple M3 Max) | 20,898,786 |
| F(100M) | 25.265s | Go! (Apple M5 Max) | 20,898,786 |
| F(200M) | 1m56.757s | Go! (Apple M1 Max) | 41,797,550 |
| F(200M) | 1m49.757s | Go! (Apple M2 Max) | 41,797,550 |
| F(200M) | 1m9.664s | Go! (Apple M3 Max) | 41,797,550 |
| F(200M) | 1m14.216s | Go! (Apple M5 Max) | 41,797,550 |
| F(300M) | 3m25.590s | Go! (Apple M2 Max) | 62,696,314 |
| F(300M) | 2m9.151s | Go! (Apple M3 Max) | 62,696,314 |
| F(300M) | 2m19.422s | Go! (Apple M5 Max) | 62,696,314 |
