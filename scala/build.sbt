lazy val commonSettings = Seq(
  scalaVersion := "2.13.12"
)

lazy val utils = project
  .settings(commonSettings)

lazy val `year-2021` = project
  .dependsOn(utils)
  .settings(commonSettings)

lazy val `year-2023` = project
  .dependsOn(utils)
  .settings(commonSettings)

lazy val `adventOfCode` = project
  .in(file("."))
  .aggregate(
    `year-2021`,
    `year-2023`,
    utils
  )
