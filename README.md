# Overview
Personal solutions to projecteuler.net problems, written in go.

# Installation

Install golang, clone this repo, then build the `project-euler` binary:

```bash
go build .
```

# Usage

```bash
[john@acmecorp ~]# ./project-euler
Usage of ./project-euler:
  -p string
        The problem to run/test
  -r    Run a problem
  -t    Test a problem
[john@acmecorp ~]# ./project-euler -p 12 -r
Looking for 500 divisors
Found: 31337
Run time: 592.265162ms
[john@acmecorp ~]# ./project-euler -p 12 -t
Looking for 500 divisors
Found: 31337
Run time: 665.191849ms
Test passed!
```

# Problems/Answers

All problems are encrypted with git-crypt to obfuscate source code.
