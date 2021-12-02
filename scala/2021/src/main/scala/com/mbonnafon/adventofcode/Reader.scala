package com.mbonnafon.adventofcode

import scala.io.Source

case class Reader(filename: String) {
  def read[A](convert: String => A): List[A] = {
    Source.fromResource(filename).getLines().toList.map(convert)
  }
}
