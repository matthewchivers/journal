# Journal: Project Overview

> The following is a _prospective_ readme. Or "something to work towards". It does not describe the current capabilities of the tool, but acts as a reference to align work on the tool.

## Current capabilities (03-04-24):
* None

## Overview

Journal is a versatile command-line tool designed for creating and managing journal entries tailored to the unique workflow of professionals. It simplifies the documentation process for various types of meetings and personal notes, allowing for efficient and structured record-keeping directly from the command line.

## Key Features

- **Customizable Document Types**:

Users can define their own types of documents and document structures, so that you always have the right kind of entry for the activity being journalled.

- **Flexible Scheduling**:

Journal supports creation of entries based on their scheduled occurrence. Got a status meeting that happens every day? Or perhaps a planning meeting every week on a Tuesday? Journal places the entry on the correct day for that week. Use flags such as `--previous` or `--next` to look back or plan ahead.

- **Markdown**:

The developer's best friend; Leverage markdown for note-taking, including support for code snippets, structured headings, and metadata.

- **Configuration via YAML**:

Define document schemas and templates in a YAML configuration file, offering a high degree of customization.Configuration GuideSetting Up Your Journal ConfigurationCreate a journal-config.yaml file in your project root. This file allows you to specify the document types you need, their schedules, and the templates for each document type.

Example `journal-config.yaml`
``` yaml
defaultDocType: status
documentTypes:
  - name: status
    schedule:
      frequency: [daily/weekly/monthly/annually]
      interval: [1-99]
      days: [1-7]
      dates: [1-31]
      weeks: [1-4]
      months: [1-12]
    templatePath: status.md
paths:
  templatesDir: ~/.journal/templates/
  journalDir: ~/journal/
  journalPathPattern: "{{.Year}}/{{.Month}}/{{.Day}}/{{.TemplateName}}"
userSettings:
  timezone: 'Europe/London'
  ```


#### Document and Template Definitions

`documents`: A list of your document types along with their schedule and template.
`schedule`: When the document should be created (daily, weekly, monthly, or a specific day).
`template`: The markdown template that will be used for each document type ()

## Using Journal
Journal's command-line interface is intuitive, designed to make creating and managing entries as straightforward as possible.

### Basic Commands
Creating Today's Note: Simply type `journal` to generate today's note based on the default template.Specifying a Document Type: Use `journal scrum` or `journal playback` to create entries of those types.
### Using Flags:
Apply `--previous` or `--next` to create or access entries relative to today, based on the document type's frequency.
`journal --previous` creates an entry for the previous occurrence based on its schedule. Forgot to write notes for yesterday's status meeting? `journal status --previous` takes you there. Planning ahead for next week's playback? `journal playback --next` will create an entry for the next occurrence.

## Future Directions:
### Search Functionality:
Implement search capabilities to quickly find entries by date, tags, or content.
### GUI
Integration into some nice markdown editor, perhaps.

## Conclusion:
Journal aims to be a dynamic and essential tool for anyone looking to streamline their daily documentation process, providing the flexibility to adapt to any workflow while maintaining simplicity and efficiency in note-taking.

## Contribution

Project uses golangci-lint for linting, and pre-commit for the pre-commit hook.
