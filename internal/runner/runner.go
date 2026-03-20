package runner

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ensurePomExists() error {
	pomPath := filepath.Join(".", "pom.xml")

	info, err := os.Stat(pomPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("pom.xml was not found in the current directory. Please run this command from the root of a Maven project")
		}
		return fmt.Errorf("could not check for pom.xml: %w", err)
	}

	if info.IsDir() {
		return fmt.Errorf("a path named pom.xml exists, but it is a directory, not a file")
	}

	return nil
}

func runOutput(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = new(bytes.Buffer)
	return cmd.Output()
}

func runCombinedOutput(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

func firstNonEmptyValue(values ...string) string {
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func cleanMavenValue(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return ""
	}

	lines := strings.Split(v, "\n")
	cleaned := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "[INFO]") || strings.HasPrefix(line, "[WARNING]") || strings.HasPrefix(line, "Download") {
			continue
		}
		cleaned = append(cleaned, line)
	}

	if len(cleaned) == 0 {
		return ""
	}

	result := strings.TrimSpace(cleaned[len(cleaned)-1])
	if strings.Contains(result, "${") {
		return ""
	}

	return result
}

func evaluateMavenExpression(expression string) string {
	out, err := runOutput(
		"mvn",
		"help:evaluate",
		"-Dexpression="+expression,
		"-q",
		"-DforceStdout",
	)
	if err != nil {
		return ""
	}

	return cleanMavenValue(string(out))
}

func parseJavaMajorVersion(raw string) (int, error) {
	value := strings.TrimSpace(raw)
	if value == "" {
		return 0, fmt.Errorf("empty Java version")
	}

	if strings.HasPrefix(value, "1.") {
		parts := strings.Split(value, ".")
		if len(parts) > 1 {
			return strconv.Atoi(parts[1])
		}
	}

	re := regexp.MustCompile(`^\d+`)
	match := re.FindString(value)
	if match == "" {
		return 0, fmt.Errorf("could not parse Java major version from %q", raw)
	}

	major, err := strconv.Atoi(match)
	if err != nil {
		return 0, fmt.Errorf("could not parse Java major version from %q: %w", raw, err)
	}

	return major, nil
}

func getRequiredJavaVersion() (int, string, error) {
	release := evaluateMavenExpression("maven.compiler.release")
	target := evaluateMavenExpression("maven.compiler.target")
	javaVersion := evaluateMavenExpression("java.version")

	raw := firstNonEmptyValue(release, target, javaVersion)
	if raw == "" {
		return 0, "", fmt.Errorf("could not determine the Java version required by the project")
	}

	major, err := parseJavaMajorVersion(raw)
	if err != nil {
		return 0, raw, err
	}

	return major, raw, nil
}

func getCurrentMavenJavaVersion() (int, string, error) {
	out, err := runCombinedOutput("mvn", "-version")
	if err != nil {
		return 0, "", fmt.Errorf("could not get Maven version information: %w", err)
	}

	text := string(out)
	lines := strings.Split(text, "\n")

	re := regexp.MustCompile(`(?i)Java version:\s*([^\s,]+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		matches := re.FindStringSubmatch(line)
		if len(matches) == 2 {
			raw := matches[1]
			major, err := parseJavaMajorVersion(raw)
			if err != nil {
				return 0, raw, err
			}
			return major, raw, nil
		}
	}

	return 0, "", fmt.Errorf("could not detect the Java version used by Maven")
}

func ensureJavaVersionMatches() error {
	requiredMajor, requiredRaw, err := getRequiredJavaVersion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: %s. Skipping Java version validation.\n", err.Error())
		return nil
	}

	currentMajor, currentRaw, err := getCurrentMavenJavaVersion()
	if err != nil {
		return err
	}

	if requiredMajor != currentMajor {
		return fmt.Errorf(
			"project requires Java %d (%s), but Maven is currently running with Java %d (%s). Please switch Java version before running this command. Example: jenv global %d",
			requiredMajor,
			requiredRaw,
			currentMajor,
			currentRaw,
			requiredMajor,
		)
	}

	return nil
}

func Run(name string, args ...string) error {
	if name == "mvn" {
		if err := ensurePomExists(); err != nil {
			return err
		}
		if err := ensureJavaVersionMatches(); err != nil {
			return err
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("Running: %s %v\n\n", name, args)
	return cmd.Run()
}

func RunAndCapture(name string, args ...string) ([]byte, error) {
	if name == "mvn" {
		if err := ensurePomExists(); err != nil {
			return nil, err
		}
		if err := ensureJavaVersionMatches(); err != nil {
			return nil, err
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}
