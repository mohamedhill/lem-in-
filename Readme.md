# Lem-in (Go Implementation)

## Overview

This project is a Go implementation of the classic **lem-in** challenge. The goal is to move a given number of ants from a start room to an end room through a network of rooms and links, using the minimum number of turns

---

## Features

- Parses a custom text input format describing ants, rooms, and links.
- Finds all possible paths from start to end using depth-first search.
- Selects the largest set of non-overlapping paths.
- Distributes ants for minimal total turns.
- Simulates and prints ant movements turn by turn.

---

## Usage

go run . test.txt

**Example:**
```
9
##start
richard 0 6
gilfoyle 6 3
##end
peter 14 6
richard-gilfoyle
gilfoyle-peter
```

---

## projet Structure

- `main.go` — Entry point, handles input and output.
- `funcs/graph.go` — structs for Graph and rooms
- `funcs/parseinput.go` — Input parsing.
- `funcs/pathfinding.go` — Pathfinding,and ant distribution.
- `funcs/simulation.go` — Ant movement simulation and output.

---

## Authors

-boulhaj
-mhilli