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
    fileName: "Note-{{.Day.Pad}}-{{.Month.Pad}}-{{.Year.Short}}.{{.FileExt}}"
    directory: "{{.EntryID}}s"
    schedule:
      frequency: daily
  - id: standup
    fileName: "{{.Day.Short}}-{{.Day.Ord}}-{{.Month.Short}}-{{.Year.Num}}.{{.FileExt}}"
    directory: "{{.EntryID}}s/wc-{{.WkCom.Day.Pad}}-{{.WkCom.Month.Pad}}-{{.WkCom.Year.Short}}"
    schedule:
      frequency: daily
paths:
  baseDirectory: "/home/user/journal"
  journalDirectory: "{{.Year.Num}}/{{.Month.Num}}"
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

## Templating

Directories and filenames (with the exception of the base directory) can be templated. At its core, the templating contains:

* **Year**: Information on the current year.
* **Month**: Information on the current month.
* **Day**: Information on the current day.
* **WkCom**: Contains sub-year, month, and day details for the Monday of the current week.
* **EntryID**: ID or name of the target entry.
* **FileExt**: File extension for the target entry.
* **Topic**: Topic specified for the entry.

> At some point in the future, documents will be generated based on templates (at which point all the above will apply to the document templates too)

### Date Hierarchy and Fields

Dates have a hierarchical structure relative to their parent container. This allows dates to have different values based on context. For example, 2nd August 2024 is the 2nd day of August, the 5th day of the week (Friday), and the 215th day of the year (out of 366, as it is a leap year).

This hierarchical structure allows these values to be referenced relative to their parent. In the example above, `Day` would be `2` (2nd August), `Week.Day` would be `5` (5th day of the week), `Month.Day` would be `2`, and `Year.Day` would be `215` (215th day of the year). These values are available in multiple formats, such as names (Monday/Tuesday/etc. or January/February/etc.), where applicable.

Here’s an explanation of the date tags available:

**Day**:
* **Num**: The day of the month (e.g. 1, 2, 3...) `{{.Day.Num}}`
* **Pad**: Zero-padded day of the month (e.g. 01, 02, 03...) `{{.Day.Pad}}`
* **Ord**: Day of the month with ordinal suffix (e.g. 1st, 2nd, 3rd...) `{{.Day.Ord}}`
* **Name**: Name of the day of the week (e.g. Monday, Tuesday...) `{{.Day.Name}}`
* **Short**: Short name of the day of the week (e.g. Mon, Tue...) `{{.Day.Short}}`

The `Day` is relative to its parent container. For example:
- `{{.Day}}`: Day of the month (1-31)
- `{{.Year.Day}}`: Day of the year (1-366)
- `{{.Week.Day}}`: Day of the week (1-7)

**Month**:
* **Num**: The month number (e.g. 1, 2, 3...) `{{.Month.Num}}`
* **Pad**: Zero-padded month number (e.g. 01 for January) `{{.Month.Pad}}`
* **Ord**: Month with ordinal suffix (e.g. 1st) `{{.Month.Ord}}`
* **Name**: Full month name (e.g. January) `{{.Month.Name}}`
* **Short**: Short month name (e.g. Jan) `{{.Month.Short}}`
* **DaysIn**: Number of days in the month (e.g. 31 for January) `{{.Month.DaysIn}}`
* **Day**: Day of the month (1-31) `{{.Month.Day}}`
* **Week**: Week of the month `{{.Month.Week}}`

**Week**:
* **Num**: Week number `{{.Week.Num}}`
* **Pad**: Zero-padded week number `{{.Week.Pad}}`
* **Ord**: Week with ordinal suffix (e.g. 1st) `{{.Week.Ord}}`
* **Day**: Day of the week `{{.Week.Day}}`

The `Week` number is relative to its parent container:
- `{{.Month.Week}}`: Week of the month (1-5)
- `{{.Year.Week}}`: Week of the year (1-53)

**Year**:
* **Num**: Full year number (e.g. 2024) `{{.Year.Num}}`
* **Short**: Short form of the year (e.g. 24) `{{.Year.Short}}`
* **Month**: Details about the current month `{{.Year.Month}}`
* **Week**: Details about the current week of the year `{{.Year.Week}}`
* **Day**: Details about the current day of the year `{{.Year.Day}}`
* **DaysIn**: Number of days in the year `{{.Year.DaysIn}}`

**WkCom**:
* Holds the same date structure as `Year`, `Month`, and `Day` but for the Monday of the current week. e.g.:
  - `{{.WkCom.Year.Num}}`
  - `{{.WkCom.Month.Num}}`
  - `{{.WkCom.Day.Num}}`
  - `{{.WkCom.Year.Day.Num}}`
  - `{{.WkCom.Month.Week.Day.Num}}`

### Common Fields

- **Num**: Full number representation (1, 2, ... 20, 21... 101, 102...)
- **Pad**: Zero-padded number representation (01, 02... 09, 10)
- **Ord**: Number with ordinal suffix (1st, 2nd, 3rd... 10th, 11th... 101st, 102nd...)
- **Name**: Full name (for months or weekdays) (Monday, Tuesday... or January, February...)
- **Short**: Short name (for months or weekdays) (Mon, Tue... or Jan, Feb...)
- **DaysIn**: Number of days (for months or years) (28/29/30/31 or 365/366)

### Example Patterns

Assuming the current date is Friday, 2nd August 2024, with an entry called "note" and a topic called "ProjectA".

| Pattern | Parsed |
|---------|--------|
| `{{.Year.Short}}-{{.Month.Pad}}-{{.Day.Pad}}/` | `24-08-02` |
| `{{.Year.Num}}/{{.Month.Short}}/{{.Day.Name}}-{{.Day.Ord}}-{{.Month.Name}}.md` | `2024/Aug/Friday-2nd-August.md` |
| `{{.WkCom.Year.Short}}/{{.WkCom.Month.Pad}}/{{.WkCom.Day.Pad}}` | `24/07/29` (Given that 29th July 2024 is the Monday of the current week) |
| `{{.EntryID}}_{{.Year.Num}}_{{.Month.Short}}_{{.Day.Pad}}.{{.FileExt}}` | `notes_2024_Aug_02.md` |
| `{{.Topic}}/{{.Year.Num}}/{{.Month.Name}}/{{.Day.Name}}` | `projectA/2024/August/Friday` |
| `{{.EntryID}}s/{{.Year.Day.Num}}-of-{{.Year.DaysIn}}` | `notes/215-of-366` (366 because 2024 is a leap year) |

This hierarchical and structured approach allows for flexible and dynamic generation of directory and file names based on the current date and entry details.

## Logging
`journal` uses `zerolog` for logging. There are two levels that can be specified: `--info` and `--debug`.

Logs are saved by default to the `~/.journal` directory in JSON format at `info` level. Logs are not automatically output to the console unless a log level is specified. For example, if a user specifies `--info` the console will now output `info` level logs, as well as still saving `info` level logs to the log file.

When outputting to the console, the logger will opt for a human-readable format. Users may specify `--logjson` to have the console output in the regular json format instead.

## Contribution

We welcome contributions! Please ensure that all pull requests and commit messages are informative. All code should build using the `make build` command and be appropriately unit tested. For any large changes, please open an issue first, for discussion.

## License

This project is licensed under the MIT License.
