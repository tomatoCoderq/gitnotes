# ✅ TODO: `gitnotes` CLI Tool (with TDD Segments)

A personal Git note-taking CLI tool for attaching, viewing, and managing notes on commits or branches.

---


## 🏗️ Project Setup
- [X] Create GitHub repository `gitnotes`
- [ ] Initialize Go module: `go mod init github.com/yourusername/gitnotes`
- [ ] Install Cobra CLI: `go get github.com/spf13/cobra`
- [ ] Scaffold project with: `cobra init --pkg-name github.com/yourusername/gitnotes`
- [ ] Create base structure:
  - [] `cmd/gitnotes/main.go`
  - [ ] `internal/storage/`
  - [ ] `internal/git/`
  - [ ] `internal/commands/`
  - [ ] `testdata/`

---

## 📦 Data Model
- [ ] Define `Note` struct:
  ```go
  type Note struct {
      Ref       string    `json:"ref"`
      Message   string    `json:"message"`
      CreatedAt time.Time `json:"created_at"`
  }
  ```

### 🧪 Tests:
- [ ] Marshall/unmarshal JSON to Note struct
- [ ] Validate empty fields and timestamps

---

## 💾 Storage Layer
- [ ] Create `~/.gitnotes.json` file if it doesn't exist
- [ ] Functions to implement:
  - [ ] `LoadNotes() ([]Note, error)`
  - [ ] `SaveNotes([]Note) error`
  - [ ] `FindNoteByRef(ref string) *Note`
- [ ] Use `os.UserHomeDir()` for path resolution

### 🧪 Tests:
- [ ] Load file with valid JSON
- [ ] Handle missing/empty file gracefully
- [ ] Add/remove note, then reload and verify
- [ ] Handle invalid JSON errors

---

## ⚙️ Git Ref Resolution
- [ ] Use `git rev-parse <ref>` to resolve commits/branches
- [ ] Create `ResolveRef(ref string) (string, error)`

### 🧪 Tests:
- [ ] Valid short SHA to full SHA
- [ ] Branch name to SHA
- [ ] Nonexistent ref returns error

---

## 🔨 Core Commands (TDD First!)

### 1️⃣ `add <ref>`
- [ ] Resolve ref
- [ ] Prompt for note message via terminal
- [ ] Save note to storage
- [ ] Add flag for tags {Tags (`TODO`, `BUG`, `INFO`, `CRITICAL`)}

#### 🧪 Tests:
- [ ] Validate ref resolution
- [ ] Handle input via mocked stdin
- [ ] Check note is persisted to file

---

### 2️⃣ `show <ref>`
- [ ] Resolve and display associated note
- [ ] Display timestamp

#### 🧪 Tests:
- [ ] Show existing note by SHA
- [ ] Handle missing ref gracefully
- [ ] Print expected output

---

### 3️⃣ `list`
- [ ] Load and display all notes
- [ ] Show preview (truncate message)
- [ ] Optional: Sort by date

#### 🧪 Tests:
- [ ] Ensure notes load and render in order
- [ ] Handle empty state
- [ ] Truncation formatting

---

### 4️⃣ `rm <ref>`
- [ ] Resolve ref
- [ ] Confirm deletion unless `--force`
- [ ] Remove from storage
- [ ] Remove by refernce (removes all notes related to the branch)
- [ ] Remove by note title (removes one specific note)

#### 🧪 Tests:
- [ ] Remove existing note
- [ ] Error on missing ref
- [ ] Force flag skips confirmation

---

### 5️⃣ `search <text>`
- [ ] Search note content by keyword (case-insensitive)
- [ ] Display matching entries

#### 🧪 Tests:
- [ ] Match on full and partial terms
- [ ] Handle no-match case
- [ ] Multiple match results

---

## 🎨 Usability Features
- [X] Format timestamps for readability (`Mon Jan 2 15:04`)
- [ ] Use colors or bold text for output (optional)
- [X] Handle paths across OS (Windows/macOS/Linux)

#### 🧪 Tests:
- [ ] Format function returns correct date strings
- [ ] Cross-platform path resolution

---

## 🔍 Integration & CLI Behavior
- [ ] End-to-end: `add → show → list → rm → show`
- [ ] Use temp test directories for integration tests
- [ ] Simulate real CLI behavior with `os.Args`

#### 🧪 Tests:
- [ ] Full flow from empty state to data mutation
- [ ] Test help text and command parsing

---

## 📄 Documentation & README
- [ ] Write `README.md` with:
  - [ ] Installation guide (`go install`)
  - [ ] Command examples
  - [ ] Screenshots or terminal output
- [ ] Add `.gitignore`, `LICENSE`, and Go report badge

#### 🧪 Tests:
- [ ] All example commands work when copy-pasted
- [ ] Markdown renders correctly

---

## 🚀 Stretch Features
- [ ] Markdown support in notes
- [ ] Tags (`TODO`, `BUG`, `INFO`)
- [ ] Encrypt/decrypt note messages
- [ ] Export notes to `notes.md`
- [ ] `fzf` fuzzy search integration

#### 🧪 Tests:
- [ ] Tag-based filtering
- [ ] Markdown render check
- [ ] Encrypt/decrypt roundtrip
- [ ] Export file integrity check