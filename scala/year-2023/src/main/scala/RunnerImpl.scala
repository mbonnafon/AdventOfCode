import utils.{Day, Puzzle, Runner}
import year2023.Day01

object RunnerImpl extends Runner {
  override def puzzleMap: Map[Day, Puzzle] = Map(
    Day(1) -> Day01
  )

}
