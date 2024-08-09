package main

import (
  "os/exec"
  "github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err != nil {
    panic(err)
  }

  cmd := exec.Command(
    "tern",
    "migrate",
    "--migrations",
    "internal/store/pgstore/migrations",
    "--config",
    "internal/store/pgstore/migrations/tern.conf",
  )

  if error := cmd.Run(); error != nil {
    panic(error)
  }
}
