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
git clone https://github.com/yourusername/gitnotes.git
cd gitnotes
go build -o gitnotes
