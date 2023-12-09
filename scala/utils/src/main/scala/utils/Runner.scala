package utils

case class Day(value: Int)

abstract class Runner extends App {

  val year = args.headOption.map(_.toInt).getOrElse(2021)
  val day = args.lift(1).map(_.toInt).getOrElse(1)
  def puzzleMap: Map[Day, Puzzle]

  puzzleMap.get(Day(day)) match {
    case None => println(s"Puzzle for Day $day (Year $year) is not yet solved!")
    case Some(puzzle) =>
      println(s"Solving puzzle for Day $day (Year $year)")

      val (result1, result2) =
        puzzle.solve(s"$year/Day${"%02d".format(day)}.txt")
      println(s"Part 1: $result1")
      println(s"Part 2: $result2")
  }
}
