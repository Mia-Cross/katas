from typing import Iterator


def fibonacci() -> Iterator[int]:
    n_minus_one = 1
    n_minus_two = 0
    while True:
        yield n_minus_one
        n = n_minus_one + n_minus_two
        n_minus_two = n_minus_one
        n_minus_one = n
