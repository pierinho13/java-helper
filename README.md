# java-helper

`java-helper` provides shortcuts for common Maven commands used in Java projects.

## What it looks like

```text
java-helper
java-helper provides shortcuts for common Maven commands used in Java projects.

Usage:
  java-helper [command]

Available Commands:
  completion          Generate the autocompletion script for the specified shell
  fmt                 Run spotless apply
  help                Help about any command
  instructions        Show Java and Maven installation instructions
  java                Inspect Java-related Maven config from effective-pom
  manual-instructions Show the manual Maven commands
  menu                Show an interactive menu
  tree                Run mvn dependency:tree
  verify              Run mvn verify with the required local flags

Flags:
  -h, --help   help for java-helper

Use "java-helper [command] --help" for more information about a command.
```

Example:

```text
$ java-helper menu

java-helper
-----------
1) Java version hints
2) Spotless apply
3) Verify
4) Dependency tree
5) Instructions about Java installation
6) Manual instructions
7) Exit
Choose an option:
```

## Installation

```bash
brew install pierinho13/java-helper/java-helper
```

Or:

```bash
brew tap pierinho13/java-helper
brew install java-helper
```

Or, in a `Brewfile`:

```ruby
tap "pierinho13/java-helper"
brew "java-helper"
```
