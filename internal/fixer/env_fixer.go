package fixer

import (
        "fmt"
        "os"
        "strings"

        "github.com/devpreflight/devpreflight/internal/config"
        "github.com/fatih/color"
        "github.com/joho/godotenv"
)

func FixEnvParity(cfg *config.Config, autoConfirm bool) error {
        // Check if .env.example exists
        if _, err := os.Stat(".env.example"); os.IsNotExist(err) {
                return fmt.Errorf(".env.example not found")
        }
        
        // Load .env.example
        exampleEnv, err := godotenv.Read(".env.example")
        if err != nil {
                return fmt.Errorf("failed to parse .env.example: %w", err)
        }
        
        // Load existing .env (or create empty map)
        actualEnv := make(map[string]string)
        if _, err := os.Stat(".env"); err == nil {
                actualEnv, err = godotenv.Read(".env")
                if err != nil {
                        return fmt.Errorf("failed to parse .env: %w", err)
                }
        }
        
        // Find missing keys
        var missing []string
        for key := range exampleEnv {
                if _, exists := actualEnv[key]; !exists {
                        missing = append(missing, key)
                }
        }
        
        green := color.New(color.FgGreen).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        cyan := color.New(color.FgCyan).SprintFunc()
        
        if len(missing) == 0 {
                fmt.Printf("%s No missing environment variables\n", green("✓"))
                return nil
        }
        
        fmt.Printf("%s Found %d missing environment variables:\n", yellow("!"), len(missing))
        for _, key := range missing {
                fmt.Printf("  %s %s\n", cyan("•"), key)
        }
        
        if !autoConfirm {
                fmt.Print("\nAdd these keys to .env with __REPLACE_ME__ placeholders? (y/n): ")
                var response string
                fmt.Scanln(&response)
                if strings.ToLower(response) != "y" {
                        fmt.Println("Cancelled")
                        return nil
                }
        }
        
        // Add missing keys to actualEnv
        for _, key := range missing {
                actualEnv[key] = "__REPLACE_ME__"
        }
        
        // Write back to .env
        var lines []string
        for key, value := range actualEnv {
                lines = append(lines, fmt.Sprintf("%s=%s", key, value))
        }
        
        content := strings.Join(lines, "\n") + "\n"
        if err := os.WriteFile(".env", []byte(content), 0644); err != nil {
                return fmt.Errorf("failed to write .env: %w", err)
        }
        
        fmt.Printf("\n%s Added %d placeholders to .env\n", green("✓"), len(missing))
        fmt.Printf("\n%s Next steps:\n", cyan("→"))
        fmt.Println("  1. Open .env")
        fmt.Println("  2. Replace __REPLACE_ME__ with actual values")
        fmt.Println("  3. Run 'devpreflight check' to verify")
        
        return nil
}
