from fibonacci import fibonacci


def test_should_return_one_at_first_call():
    value = next(fibonacci())
    assert value == 1


def test_should_return_one_at_second_call():
    fib = fibonacci()
    next(fib)
    value = next(fib)
    assert value == 1

def test_should_return_two_at_third_call():
    fib = fibonacci()
    next(fib)
    next(fib)
    value = next(fib)
    assert value == 2

def test_should_return_three_at_fouth_call():
    fib = fibonacci()
    next(fib)
    next(fib)
    next(fib)
    value = next(fib)
    assert value == 3
