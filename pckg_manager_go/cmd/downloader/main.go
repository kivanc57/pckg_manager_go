package main

import (
  "log"
  "pckg_manager_go/internal/common"
  "pckg_manager_go/internal/packagemanager"
)

func main() {
  common.LoadEnv()

  cfg := packagemanager.Config{
    Repositories: common.ParseEnvArray("REPOSITORIES"),
    ExcludedPackages: common.ParseEnvArray("EXCLUDED_PACKAGES"),
    Extensions: common.ParseEnvArray("EXTENSIONS"),
    DataDir: "data",
    OutputFile: "output/output.txt",
  }

  err := packagemanager.ProcessPackages(cfg)
  if err != nil {
    log.Fatal(err)
  }
}
