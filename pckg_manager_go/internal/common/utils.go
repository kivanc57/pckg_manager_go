package common

import (
  "log"
  "os/exec"
)

func ExecPowerShellCommand(command string) (string, error) {
  cmd := exec.Command("powershell", "-Command", command)
  output, err := cmd.CombinedOutput()
  if err != nil {
    log.Printf("PowerShell command error: %s\nOutput: %s", err, string(output))
  }
  return string(output), err
}
