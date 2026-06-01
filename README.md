# Deal or No Deal

A clean, console-based implementation of the classic game show **Deal or No Deal**, written in Go.

## Features

- Fisher-Yates shuffle for fair randomization of case values
- Persistent player briefcase selection
- Dynamic bank offers based on remaining case maximums with controlled randomness


## Tech Stack

- **Go** 1.21+
- Standard library only (zero dependencies)

## Quick Start

```bash
git clone https://github.com/BrodieFriday/Deal-or-No-Deal.git
cd Deal-or-No-Deal

# Run directly
go run DealOrNoDeal.go

# Or build executable
go build -o dealornodeal DealOrNoDeal.go
./dealornodeal
