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

  test "count number of occurances in a list" do
    assert Day1.count_occurances([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1) == 1
    assert Day1.count_occurances([1, 2, 2, 4, 5, 6, 7, 8, 9, 10], 2) == 2
  end

  test "calculate test similarty scores" do
    assert Day1.calc_similarity_score([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]) == 31
  end
end
