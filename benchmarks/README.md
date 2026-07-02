# MegaVerse Benchmarks

## Overview

Performance benchmarking suites and results.

## Structure

```
benchmarks/
├── api/            # API performance benchmarks
├── ai/             # AI model benchmarks
└── database/       # Database query benchmarks
```

## Running

```bash
# API benchmarks
k6 run benchmarks/api/load-test.js

# AI benchmarks
python benchmarks/ai/evaluate.py

# Database benchmarks
go test -bench=. ./benchmarks/database/
```

## Results

Benchmark results are stored in `benchmarks/*/results/` and tracked over time.
