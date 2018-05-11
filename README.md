# Overview
Personal solutions to projecteuler.net problems, written in go.

# Usage

Install golang, clone this repo, then build the `project-euler` binary:

```bash
go build .
```

```bash
Usage of ./project-euler:
  -p string
        The problem to run/test
  -r    Run a problem
  -t    Test a problem
```

For example:

```bash
/project-euler -p 12 -r
Looking for 500 divisors
Found: 31337
Run time: 592.265162ms
```

# Problems/Answers

All problems are encrypted with git-crypt to obfuscate source code.
