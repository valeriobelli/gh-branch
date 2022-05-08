# gh-branch

GitHub CLI extension for managing the current branch

## Usage

Following you can find some of the available commands

### Create a branch

```bash
# Default branch
gh branch create 1

# Specify a branch type
gh branch create 1 --type feat
```

### Rebase the current Pull Request's branch

```bash
# Default behaviour
gh branch rebase

# Specified rebase strategy
gh branch rebase --type interactive
```
