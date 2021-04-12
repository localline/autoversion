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

## Using the CLI

`autoversion --help`

## Use as a Github Action

```
name: Autoversion Action
on: [push]

jobs:
  autoversion-job:
    runs-on: ubuntu-latest
    name: A job to use autoversion
    steps:
    - uses: actions/checkout@master
    - name: Autoversion
      id: autoversion
      uses: localline/autoversion@master
      env:
        GITHUB_ACCESS_TOKEN: ${{ secrets.GITHUB_ACCESS_TOKEN }}
      with:
        path: "/path/to/repo/"
```
