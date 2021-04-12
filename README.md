# Conventional Commits Autoversion

### `autoversion`

## Overview

This maintains a go based CLI tool to obtain the semantic version numbers from a git repository based on tags and conventional commits. An `actions.yml` file is included to run this as part of a github actions pipeline.

### Build locally

`go install github.com/localline/autoversion`

### Docker

```
docker build -t autoversion .
docker run -it autoversion
```

### Using the CLI

`autoversion --help`

### Use as a Github Action

```yml
name: Autoversion Action
on: [push]

jobs:
  autoversion-job:
    runs-on: ubuntu-latest
    name: Release
    steps:
      - uses: actions/checkout@master
        with:
          fetch-depth: 0 # Required if using v2 so all tags are included.
      - name: Autoversion
        id: autoversion
        uses: localline/autoversion@master
        with:
          path: "."
      - name: Create Release
        id: create_release
        if: steps.autoversion.outputs.previous_version != steps.autoversion.outputs.next_version
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ steps.autoversion.outputs.next_version }}
          release_name: Release ${{ steps.autoversion.outputs.next_version }}
          body: New Release Test
          draft: false
          prerelease: false
```
