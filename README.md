# gitnotes

**gitnotes** is a command-line tool for attaching personal notes to Git commits or branches. Notes are stored locally in a persistent BoltDB database and are not committed to your Git repository. This allows developers to maintain contextual or task-related information without polluting the codebase or Git history.

---

## Features

- Add notes to specific commits or branches
- View notes for a ref
- List all stored notes
- Remove notes
- Store multiple notes per ref
- Persistent local storage using BoltDB

---

## Installation

### Prerequisites

- Go 1.18 or newer
- Git installed and available in your system PATH

### Build from source

```bash
git clone https://github.com/tomatoCoderq/gitnotes.git
cd gitnotes
go build -o gitnotes
```

You can now run the binary:

```bash
./gitnotes --help
```

---

## Usage

### Add a Note

```bash
gitnotes add HEAD
```

You'll be prompted to enter your note. The note is saved against the full SHA that `HEAD` points to.

### Show Notes for a Ref

```bash
gitnotes show <ref>
```

Example:

```bash
gitnotes show abc1234
```

### List All Notes

```bash
gitnotes list
```

This displays all refs that have associated notes and their contents.

### Remove Notes

```bash
gitnotes rm <ref>
gitnotes rm -p title SomeTitle
```

Removes all notes for a given ref or for a given title (requires specification).

---

## Data Storage

Notes are stored locally in:

```
./gitnotes.db
```

---

## Roadmap

- [ ] Expand tags/categories for notes
- [ ] Export notes to Markdown
- [X] Shell completions (bash, zsh)

---