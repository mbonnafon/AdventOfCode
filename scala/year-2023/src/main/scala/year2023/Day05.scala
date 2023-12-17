package year2023

import utils.Puzzle

object Day05 extends Puzzle {
  def part1(input: List[String]): Int = {
    val seeds = input.head.split(":").last.split(" ").toList.drop(1)
    val mapping: Map[String, List[String]] = input
      .drop(1)
      .drop(1)
      .foldLeft((Map.empty[String, List[String]], "")) {
        case ((accMap, currentKey), line: String) if (line.endsWith("map:")) =>
          (accMap, line)
        case ((accMap, currentKey), line: String) if (line.nonEmpty) =>
          val newAcc = accMap.updatedWith(currentKey)(list =>
            list match {
              case Some(value) => Some(value.appended(line))
              case None        => Some(List(line))
            }
          )
          (newAcc, currentKey)
        case ((accMap, currentKey), line) =>
          (accMap, "")
      }
      ._1

    seeds
      .map { seed =>
        val calculator: (String, List[String]) => Option[String] = {
          (target0, maps) =>
            val target = target0.toLong
            val destination = maps(0).toLong
            val source = maps(1).toLong
            val offset = maps(2).toLong - 1
            if (source to (source + offset) contains target) {
              Some((target - source + destination).toString)
            } else {
              None
            }
        }
        def process(target: String, list: List[String]) = {
          list
            .map(line => calculator(target, line.split(" ").toList))
            .filter(_.isDefined)
            .flatten
            .headOption
            .getOrElse(target)
        }

        val soil = process(seed, mapping.get("seed-to-soil map:").get)
        val fertilizer =
          process(soil, mapping.get("soil-to-fertilizer map:").get)
        val water =
          process(fertilizer, mapping.get("fertilizer-to-water map:").get)
        val light = process(water, mapping.get("water-to-light map:").get)
        val temperature =
          process(light, mapping.get("light-to-temperature map:").get)
        val humidity =
          process(temperature, mapping.get("temperature-to-humidity map:").get)
        val location =
          process(humidity, mapping.get("humidity-to-location map:").get)

        location
      }
      .min
      .toInt

  }

  def part2(input: List[String]): Int = {
    1
  }

}
