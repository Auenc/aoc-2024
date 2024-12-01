defmodule Day1Test do
  use ExUnit.Case
  doctest Day1

  test "calculates distance between two numbers" do
    assert Day1.calc_distance(1, 3) == 2
    assert Day1.calc_distance(2, 3) == 1
    assert Day1.calc_distance(3, 3) == 0
    assert Day1.calc_distance(3, 4) == 1
    assert Day1.calc_distance(3, 5) == 2
    assert Day1.calc_distance(4, 9) == 5
  end

  test "calculates the total distance between a list of pairs" do
    assert Day1.total_distance([{1, 3}, {2, 3}, {3, 3}, {3, 4}, {3, 5}, {4, 9}]) == 11
  end
end
