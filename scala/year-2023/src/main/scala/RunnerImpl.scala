import utils.{Day, Puzzle, Runner}
import year2023._

object RunnerImpl extends Runner {
  override def puzzleMap: Map[Day, Puzzle] = Map(
    Day(1) -> Day01,
    Day(2) -> Day02,
    Day(3) -> Day03,
    Day(4) -> Day04,
    Day(5) -> Day05,
    Day(6) -> Day06,
    Day(7) -> Day07,
    Day(8) -> Day08
  )

}
