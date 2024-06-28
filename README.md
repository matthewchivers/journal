# Journal

`journal` is a simple yet powerful command-line tool designed to streamline the creation of various types of documents within a user-defined directory structure. Whether it's for daily stand-up notes, planning meetings, or general diary entries, `journal` automates the process, allowing you to focus on your content.

## Features

- **Automated Directory Creation**: Generates directories based on a user-specified structure.
- **Document Creation**: Quickly create documents for specific purposes like daily stand-ups, planning meetings, or general notes.
- **Configuration Driven**: Reads settings from a user-specific configuration file for flexibility.
- **Future Enhancements**: Upcoming support for document templating to pre-populate documents entirely from templates.

## Installation

To install `journal`, you can use `go install`:

```sh
go install github.com/matthewchivers/journal@latest
```

## Configuration

The configuration file is stored in the user's home directory at `~/.journal/config.yaml`. Below is an example configuration file:

```yaml
defaultEntry: note
defaultFileExt: md
entries:
  - id: note
    fileName: "Note-{{.Day}}-{{.Month}}-{{.Year}}.{{.FileExt}}"
    directory: "{{.EntryID}}s"
    schedule:
      frequency: daily
  - id: standup
    fileName: "{{.WeekdayNameShort}}-{{.Day}}{{.DayOrdinal}}-{{.MonthNameShort}}-{{.YearShort}}.{{.FileExt}}"
    directory: "{{.EntryID}}s/wc-{{.WeekCommencing}}"
    schedule:
      frequency: daily
paths:
  baseDirectory: "/home/user/journal"
  journalDirectory: "{{.Year}}/{{.Month}}"
```

### Example Usage & Output

Based on the above configuration, here is an example of what `journal` would create:

- For a **Note** created on June 21, 2024:
  - **Command**: `journal create -t note`
    - Or just `journal create` because the config default file type is `note`
  - **Directory**: `/home/user/journal/2024/06/notes/`
  - **File**: `Note-21-06-2024.md`

- For a **Daily Stand-up** created on June 21, 2024:
  - **Command**: `journal create -t standup`
  - **Directory**: `/home/user/journal/2024/06/standups/wc-17-06-24/`
  - **File**: `Fri-21st-Jun-24.md`

## Contribution

We welcome contributions! Please ensure that all pull requests and commit messages are informative. All code should build using the `make build` command and be appropriately unit tested. For any large changes, please open an issue first, for discussion.

## License

This project is licensed under the MIT License.
