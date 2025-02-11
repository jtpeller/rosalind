# rosalind

[![Go Report Card](https://goreportcard.com/badge/github.com/jtpeller/rosalind)](https://goreportcard.com/report/github.com/jtpeller/rosalind)
[![Release](https://img.shields.io/github/release/jtpeller/rosalind.svg?style=flat-square)](https://github.com/jtpeller/rosalind/releases)
![GitHub License](https://img.shields.io/github/license/jtpeller/rosalind)

## Overview

Coding problems from [Rosalind](https://rosalind.info/problems/list-view/)

From their site:
> Rosalind is a platform for learning bioinformatics and programming through problem solving.

This repo contains solutions to some of these problems; which are algorithms that produce the desired result. Rosalind contains methods of checking whether the algorithm is correct, so the solutions can be verified as correct (or incorrect).

## Content

- `problems` - Folder with provided text files to be used as input.
  - `bsdata` - Folder containing the problem data for "Bioinformatics Stronghold" problems.
  - `bs.go` - Contains the implementations for the "Bioinformatics Stronghold" problems that I have solved.
  - `btt.go` - Contains the implementations for the "Bioinformatics Textbook Track" problems that I have solved.
  - `README.md` - This is the README for the problems folder.
- `utils` - Folder containing common utility functions.
- `go.mod` - Module file for Golang.
- `LICENSE` - GNU GENERAL PUBLIC LICENSE Version 3
- `main.go` - File to run whichever problems desired.
- `README.md` - This is the README that you are reading.

## Usage

Run the program with `go run main.go` and add options if desired.

For more help, use either `go run main.go -h` or `go run main.go --help`.

Options:

- `-id` -- Which problem to run. IDs examples are: `dna` or `hamm`
- `-t` -- whether or not to output the computation time. Example: `go run main.go -t true`
