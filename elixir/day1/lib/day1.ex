defmodule Day1 do
  def start(_type, _args) do
    {lefts, rights} = read_file("../../data/day1/data.txt")
    puzzle_one(lefts, rights)
    puzzle_two(lefts, rights)

    Supervisor.start_link([], strategy: :one_for_one)
  end

  def puzzle_one(lefts, rights) do
    sortedLefts = Enum.sort(lefts)
    sortedRights = Enum.sort(rights)
    combined = Enum.zip(sortedLefts, sortedRights)
    IO.inspect(total_distance(combined), label: "Total distance")
  end

  def puzzle_two(lefts, rights) do
    IO.inspect(calc_similarity_score(lefts, rights), label: "Similarity score")
  end

  def read_file(fpath) do
    fpath
    |> File.stream!()
    |> Enum.reduce({[], []}, fn line, {ls, rs} ->
      [l, r] = String.split(line) |> Enum.map(&String.to_integer/1)
      {[l | ls], [r | rs]}
    end)
  end

  def calc_distance(left, right) do
    abs(left - right)
  end

  def total_distance(list) do
    Enum.reduce(list, 0, fn {l, r}, acc -> acc + calc_distance(l, r) end)
  end

  def count_occurances(list, value) do
    Enum.reduce(list, 0, fn x, acc -> if x == value, do: acc + 1, else: acc end)
  end

  def calc_similarity_score(left, right) do
    Enum.reduce(left, 0, fn x, acc -> acc + x * count_occurances(right, x) end)
  end
end
