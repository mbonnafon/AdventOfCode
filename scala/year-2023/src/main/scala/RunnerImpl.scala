import utils.{Day, Puzzle, Runner}
import year2023._

object RunnerImpl extends Runner {
  override def puzzleMap: Map[Day, Puzzle] = Map(
    Day(1) -> Day01,
    Day(2) -> Day02
  )

}
