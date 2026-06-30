# Pokedex CLI

A command-line application built in Go that uses the PokeAPI to explore
location areas, catch Pokemon, and build your own Pokedex right in the terminal.

This was my first project in Go and helped me understand how to structure
a CLI app, work with REST APIs, and implement caching from scratch.

## What I Learned

- Building a REPL (Read-Eval-Print Loop) in Go
- Making HTTP requests and parsing JSON responses
- Writing an in-memory cache with TTL (time-to-live) to avoid redundant API calls
- Writing unit tests in Go
- Organizing a Go project into packages

## Features

- Interactive REPL with multiple commands
- Explore Pokemon location areas
- Catch Pokemon with randomized success rates
- Inspect caught Pokemon stats (height, weight, types, base stats)
- View your full Pokedex
- Caching layer to improve performance

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or higher

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/thenitverse/pokedex.git
   ```

2. Navigate into the project:
   ```bash
   cd pokedex
   ```

3. Run the app:
   ```bash
   go run .
   ```

## Commands

| Command             | Description                          |
|---------------------|--------------------------------------|
| `help`              | Show available commands              |
| `exit`              | Exit the REPL                        |
| `map`               | Show next page of location areas     |
| `mapb`              | Show previous page of location areas |
| `explore <area>`    | List Pokemon in a location area      |
| `catch <pokemon>`   | Try to catch a Pokemon               |
| `inspect <pokemon>` | View stats of a caught Pokemon       |
| `pokedex`           | List all your caught Pokemon         |

## Example

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area

Pokedex > explore pastoria-city-area
Found Pokemon:
- tentacool
- shellos
- budew

Pokedex > catch shellos
Throwing a Pokeball at shellos...
shellos was caught!

Pokedex > inspect shellos
Name: shellos
Height: 3
Weight: 63
Types:
  - water

Pokedex > pokedex
Your Pokedex:
- shellos
```

## Running Tests

```bash
go test ./...
```

## Acknowledgements

- [PokeAPI](https://pokeapi.co/) for the free Pokemon data API
