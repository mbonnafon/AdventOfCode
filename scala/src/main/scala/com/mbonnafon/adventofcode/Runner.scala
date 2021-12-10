package com.mbonnafon.adventofcode

object Runner extends App {
  val year = args.headOption.map(_.toInt).getOrElse(2021)
  val day = args.lift(1).map(_.toInt).getOrElse(1)

  def puzzleMap = Map(
    (2021, 1) -> year2021.Day01,
    (2021, 2) -> year2021.Day02,
    (2021, 3) -> year2021.Day03,
    (2021, 4) -> year2021.Day04
  )

  puzzleMap.get(year, day) match {
    case None => println(s"Puzzle for Day $day (Year $year) is not yet solved!")
    case Some(puzzle) =>
      println(s"Solving puzzle for Day $day (Year $year)")

      val (result1, result2) =
        puzzle.solve(s"$year/Day${"%02d".format(day)}.txt")
      println(s"Part 1: $result1")
      println(s"Part 2: $result2")
  }
}
