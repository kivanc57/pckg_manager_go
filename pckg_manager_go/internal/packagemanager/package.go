package packagemanager

import (
  "encoding/json"
  "fmt"
  "io/fs"
  "log"
  "os"
  "path"
  "strings"

  "your-project/internal/common"
)

type Package struct {
  Path  string              `json:"path"`
  Props map[string][]string `json:"props"`
}

type Config struct {
  Repositories     []string
  ExcludedPackages []string
  Extensions       []string
  DataDir          string
  OutputFile       string
}

func listFiles(srcDir, pattern string) ([]string, error) {
  root := os.DirFS(srcDir)
  files, err := fs.Glob(root, pattern)
  if err != nil {
    return nil, err
  }

  var paths []string
  for _, file := range files {
    paths = append(paths, path.Join(srcDir, file))
  }
  return paths, nil
}

func ProcessPackages(config Config) error {
  files, err := listFiles(config.DataDir, "*.json")
  if err != nil {
    return err
  }

  output, err := os.Create(config.OutputFile)
  if err != nil {
    return err
  }
  defer output.Close()

  for _, filePath := range files {
    data, err := os.ReadFile(filePath)
    if err != nil {
      return err
    }

    var packages []Package
    if err := json.Unmarshal(data, &packages); err != nil {
      return err
    }

    for _, pkg := range packages {
      skip := false
      for _, repo := range config.Repositories {
        if strings.HasPrefix(pkg.Path, repo) {
          for _, ext := range config.Extensions {
            if strings.HasSuffix(pkg.Path, ext) {
              for _, excl := range config.ExcludedPackages {
                if strings.Contains(pkg.Path, excl) {
                  skip = true
                  break
                }
              }
            }
          }
        }
      }
      if skip {
        continue
      }

      id, ver := pkg.Props["nuget.id"], pkg.Props["nuget.version"]
      if len(id) == 0 || len(ver) == 0 {
        continue
      }

      chocoCmd := fmt.Sprintf("choco find %s --exact --version %s", id[0], ver[0])
      out, _ := common.ExecPowerShellCommand(chocoCmd)

      if strings.Contains(out, "0 packages found") {
        output.WriteString(pkg.Path + "\n")
      }
    }
  }
  return nil
}
