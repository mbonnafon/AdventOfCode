package year2023

import utils.Puzzle

object Day11 extends Puzzle {

  def findCoordonates(
      matrix: List[List[String]]
  ): Map[Int, (Int, Int)] = {
    val coordinates = for {
      (line, i) <- matrix.zipWithIndex
      (element, j) <- line.zipWithIndex
      if element == "#"
    } yield (i, j)

    coordinates.groupBy(idx => coordinates.indexOf(idx)).map {
      case (key, value) => (key + 1) -> value.head
    }
  }

  def manhattan(
      factor: Int,
      lineIndexes: List[Int],
      columnIndexes: List[Int],
      source: (Int, Int),
      dest: (Int, Int)
  ): Int = {
    val lineDelta =
      lineIndexes.filter(cdx =>
        (cdx > source._1 && cdx < dest._1) || (cdx < source._1 && cdx > dest._1)
      )
    val columnDelta =
      columnIndexes.filter(cdx =>
        (cdx > source._2 && cdx < dest._2) || (cdx < source._2 && cdx > dest._2)
      )

    val deltaX =
      Math.abs(source._1 - dest._1) + (lineDelta.size * (factor - 1))
    val deltaY =
      Math.abs(source._2 - dest._2) + (columnDelta.size * (factor - 1))

    deltaX + deltaY
  }

  def part1(input: List[String]): Long = {
    val matrix: List[List[String]] = input.map(_.split("").toList)
    val lineIndexes = matrix.zipWithIndex
      .collect {
        case (line, index) if line.distinct.length == 1 => index
      }
    val columnIndexes = matrix.transpose.zipWithIndex
      .collect {
        case (column, index) if column.distinct.length == 1 => index
      }

    val idxCoord = findCoordonates(matrix)

    val dists = for {
      (id1, cdx1) <- idxCoord.toList
      (id2, cdx2) <- idxCoord.toList
      if id1 < id2 // Avoid calculatin the dist twice for a same pair
    } yield (manhattan(2, lineIndexes, columnIndexes, cdx1, cdx2))
    dists.sum
  }

  def part2(input: List[String]): Long = {
    val matrix: List[List[String]] = input.map(_.split("").toList)
    val lineIndexes = matrix.zipWithIndex
      .collect {
        case (line, index) if line.distinct.length == 1 => index
      }
    val columnIndexes = matrix.transpose.zipWithIndex
      .collect {
        case (column, index) if column.distinct.length == 1 => index
      }

    val idxCoord = findCoordonates(matrix)

    val dists = for {
      (id1, cdx1) <- idxCoord.toList
      (id2, cdx2) <- idxCoord.toList
      if id1 < id2 // Avoid calculatin the dist twice for a same pair
    } yield (manhattan(1000000, lineIndexes, columnIndexes, cdx1, cdx2))

    dists.map(_.toLong).sum
  }

}
