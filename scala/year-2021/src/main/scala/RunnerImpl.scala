import adventofcode.year2021._
import utils.{Day, Puzzle, Runner}

object RunnerImpl extends Runner {
  override def puzzleMap: Map[Day, Puzzle] = Map(
    Day(1) -> Day01,
    Day(2) -> Day02,
    Day(3) -> Day03,
    Day(4) -> Day04
  )

}
