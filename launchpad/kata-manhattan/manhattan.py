class Point:
    def __init__(self, x: int, y: int):
        self.__x = x
        self.__y = y

    def x_distance_to_origin(self) -> int:
        return self.__x

    def y_distance_to_origin(self) -> int:
        return self.__y


def distance(p1: Point, p2: Point) -> int:
    return abs(p1.x_distance_to_origin() - p2.x_distance_to_origin()) + abs(p1.y_distance_to_origin() - p2.y_distance_to_origin())
