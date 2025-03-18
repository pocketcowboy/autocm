# Automatic Commit Message Generator (autocm)

Generates commit messages based on git diffs using the Gemini API.

## Prerequisites

* Requires a Gemini API key, see https://ai.google.dev/gemini-api/docs/api-key 

## Installation

```bash
go install github.com/pocketcowboy/autocm@latest
```

Make sure that your Go bin directory is in your system's PATH. The Go bin directory is typically located at `$HOME/go/bin`. You can add it to your PATH by adding the following line to your shell's configuration file (e.g., `~/.bashrc`, `~/.zshrc`):

```bash
export PATH=$PATH:$HOME/go/bin
```

## Usage

**Set your Gemini API key as an environment variable:**
```bash
# Add this line to your shell's configuration file (e.g., `~/.bashrc`, `~/.zshrc`).
export GEMINI_API_KEY="YOUR_API_KEY"
```

**Generate a commit message:**
```bash
git diff --staged | autocm
```

**Copy the commit message to the clipboard (macOS):**
```bash
git diff --staged | autocm | pbcopy
```

**Insert the commit message inside Vim:**
```vim
:r!git diff --staged | autocm
```
