defmodule Day1 do
  def start(_type, _args) do
    {lefts, rights} = read_file("../../data/day1/data.txt")
    sortedLefts = Enum.sort(lefts)
    sortedRights = Enum.sort(rights)
    combined = Enum.zip(sortedLefts, sortedRights)
    IO.inspect(total_distance(combined), label: "Total distance")

    Supervisor.start_link([], strategy: :one_for_one)
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
end
