package bootstrap

import (
    "fmt"
    "os"
)

func setupGoogleCredentials() error {
    creds := os.Getenv("GOOGLE_CREDENTIALS")
    if creds != "" {
        tmp := "/tmp/google-creds.json"
        err := os.WriteFile(tmp, []byte(creds), 0600)
        if err != nil {
            return err
        }
        os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", tmp)
        return nil
    }

    local := "./config/gen-lang-client-0235190640-7bd0c00a7ced.json"
    if _, err := os.Stat(local); err == nil {
        os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", local)
        return nil
    }

    return fmt.Errorf("no Google credentials found")
}
