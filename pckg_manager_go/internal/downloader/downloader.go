package downloader

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "path/filepath"
  "strings"

  "your-project/internal/common"
)

func downloadFile(url, outputPath string) error {
  resp, err := http.Get(url)
  if err != nil {
    return fmt.Errorf("HTTP GET error: %v", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
  }

  os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)

  file, err := os.Create(outputPath)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.Copy(file, resp.Body)
  if err != nil {
    return err
  }

  log.Printf("Download Completed: %s\n", outputPath)
  return nil
}

func ProcessDownloads(inputFilePath string, links []string) error {
  file, err := os.Open(inputFilePath)
  if err != nil {
    return fmt.Errorf("file open error: %v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    pkgParts := strings.SplitN(line, ":", 2)
    if len(pkgParts) != 2 {
      continue
    }
    pckgName, pckgVersion := pkgParts[0], pkgParts[1]

    for _, link := range links {
      outputFolder := filepath.Join("output", pckgName, pckgVersion)
      outputFile := fmt.Sprintf("%s.%s.nupkg", pckgName, pckgVersion)
      outputPath := filepath.Join(outputFolder, outputFile)

      fullURL := fmt.Sprintf("%s&path=%s", link, outputFile)
      if err := downloadFile(fullURL, outputPath); err != nil {
        log.Println("Download error:", err)
        continue
      }

      powerShellCommands := []string{
        fmt.Sprintf(`Rename-Item -Path "%s" -NewName "%s.zip"`, outputPath, strings.TrimSuffix(outputPath, ".nupkg")),
        fmt.Sprintf(`Expand-Archive -Path "%s.zip" -DestinationPath "%s"`, strings.TrimSuffix(outputPath, ".nupkg"), outputFolder),
        fmt.Sprintf(`Get-ChildItem -Path "%s" -Recurse -Include *.exe,*.msi | ForEach-Object { Write-Host "Installer found: $($_.FullName)" }`, outputFolder),
        fmt.Sprintf(`Get-ChildItem "%s" -Recurse | Where { $_.Extension -notin ".ps1",".nuspec" } | Remove-Item -Force`, outputFolder),
        fmt.Sprintf(`Get-ChildItem "%s" -Recurse -Directory | Where { ($_.GetFiles().Count -eq 0 -and $_.GetDirectories().Count -eq 0) } | Remove-Item -Force`, outputFolder),
      }

      for _, cmd := range powerShellCommands {
        common.ExecPowerShellCommand(cmd)
      }
    }
  }

  return scanner.Err()
}
