package utils

trait Puzzle {

  def solve(resource: String): (Long, Long) = {
    val input = Reader(resource).read()
    (part1(input), part2(input))
  }

  def part1(input: List[String]): Long

  def part2(input: List[String]): Long

}
