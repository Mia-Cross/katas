import manhattan


def test_dist_same_point_should_be_zero():
    p = manhattan.Point(9, 4)
    assert manhattan.distance(p, p) == 0


def test_dist_two_points_when_only_x_differs_should_be_two():
    p1 = manhattan.Point(2, 2)
    p2 = manhattan.Point(4, 2)
    assert manhattan.distance(p1, p2) == 2


def test_dist_two_points_when_only_x_differs_in_reverse_order_should_still_be_two():
    p1 = manhattan.Point(2, 2)
    p2 = manhattan.Point(4, 2)
    assert manhattan.distance(p2, p1) == 2


def test_dist_two_points_when_only_y_differs_should_be_five():
    p1 = manhattan.Point(4, 2)
    p2 = manhattan.Point(4, 7)
    assert manhattan.distance(p1, p2) == 5


def test_dist_two_points_when_only_y_differs_in_reverse_order_should_still_be_five():
    p1 = manhattan.Point(4, 2)
    p2 = manhattan.Point(4, 7)
    assert manhattan.distance(p2, p1) == 5


def test_dist_two_points_when_both_differ_should_be_four():
    p1 = manhattan.Point(4, 2)
    p2 = manhattan.Point(7, 1)
    assert manhattan.distance(p1, p2) == 4


def test_dist_two_points_when_both_differ_in_reverse_order_should_still_be_four():
    p1 = manhattan.Point(4, 2)
    p2 = manhattan.Point(7, 1)
    assert manhattan.distance(p2, p1) == 4

def test_dist_two_points_with_one_negative_value_should_be_seven():
    p1 = manhattan.Point(3, -2)
    p2 = manhattan.Point(7, 1)
    assert manhattan.distance(p1, p2) == 7


def test_dist_two_points_with_one_negative_value_in_reverse_order_should_still_be_seven():
    p1 = manhattan.Point(3, -2)
    p2 = manhattan.Point(7, 1)
    assert manhattan.distance(p2, p1) == 7


def test_dist_two_points_with_only_negative_value_should_be_three():
    p1 = manhattan.Point(-3, -3)
    p2 = manhattan.Point(-2, -1)
    assert manhattan.distance(p1, p2) == 3


def test_dist_two_points_with_only_negative_value_in_reverse_order_should_still_be_three():
    p1 = manhattan.Point(-3, -3)
    p2 = manhattan.Point(-2, -1)
    assert manhattan.distance(p2, p1) == 3


def test_coding_dojo():
    assert manhattan.distance(manhattan.Point(1, 1), manhattan.Point(1, 1)) == 0
    assert manhattan.distance(manhattan.Point(5, 4), manhattan.Point(3, 2)) == 4
    assert manhattan.distance(manhattan.Point(1, 1), manhattan.Point(0, 3)) == 3
