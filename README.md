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
| F(50M) | 27.119s | Go! (Intel amd64) | 10,449,403 |
| F(60M) | 27.646s | Go! (Apple arm64) | 12,539,280 |
| F(80M) | 27.592s | Go! (Apple arm64) | 16,719,032 |
| F(100M) | 39.590s | Go! (Apple arm64) | 20,898,786 |
| F(200M) | 1m56.757s | Go! (Apple arm64) | 41,797,550 |


