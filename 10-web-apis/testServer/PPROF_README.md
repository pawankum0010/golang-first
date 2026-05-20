# PPROF Setup Guide

This file explains what was added in this route folder, why it was added, and how to use it later for checking logs and learning `pprof`.

## What Was Added

The server in this folder now includes:

- Request logging for every HTTP call
- `pprof` endpoints for profiling
- CPU, heap, goroutine, mutex, block, allocs, and trace profiling routes

Main implementation file:

- `10-web-apis/testServer/main.go`

Logger utility used by the server:

- `utils/myutil/logger.go`

## What Changed in Code

### 1. Added `pprof` import

In `main.go`:

```go
import "net/http/pprof"
```

This enables Go's built-in profiling handlers.

### 2. Added logging middleware

A middleware was added to:

- capture request start time
- capture response status code
- log request path and duration

Log format example:

```text
2026/05/20 08:34:53.182146 GET /health | status=200 | remote=[::1]:54021 | duration=0.315 ms
```

### 3. Added `pprof` routes

These routes were registered on the same server:

```text
/debug/pprof/
/debug/pprof/heap
/debug/pprof/goroutine
/debug/pprof/profile
/debug/pprof/allocs
/debug/pprof/block
/debug/pprof/mutex
/debug/pprof/threadcreate
/debug/pprof/trace
```

### 4. Added execution logger

The server uses:

```go
myutil.NewExecutionLogger()
```

This writes logs into:

- project root, if writable
- current `testServer` folder, if root writing is blocked

## Step-by-Step Usage

### Step 1: Open this folder

```powershell
cd D:\Golang\golang-first\10-web-apis\testServer
```

### Step 2: Run the server

```powershell
go run main.go
```

Expected startup output includes:

- server port
- normal API routes
- `pprof` routes

### Step 3: Test normal routes

Open another terminal and run:

```powershell
curl http://localhost:8080/
curl http://localhost:8080/health
curl http://localhost:8080/api/users
curl "http://localhost:8080/api/user?id=1"
```

### Step 4: Check request logs

If log file is created inside this folder:

```powershell
type .\execution-log-YYYY-MM-DD.txt
```

If log file is created in project root:

```powershell
type ..\..\execution-log-YYYY-MM-DD.txt
```

What to check in logs:

- request method
- route path
- response status
- remote client address
- request duration

## Step-by-Step `pprof` Checks

### Step 5: Open pprof index

```powershell
start http://localhost:8080/debug/pprof/
```

You will see available profiles.

### Step 6: Check heap profile

```powershell
go tool pprof http://localhost:8080/debug/pprof/heap
```

Useful commands inside `pprof`:

```text
top
list main.rootHandler
help
quit
```

### Step 7: Check CPU profile

Capture 5 seconds of CPU usage:

```powershell
go tool pprof "http://localhost:8080/debug/pprof/profile?seconds=5"
```

Quick top view:

```powershell
go tool pprof -top "http://localhost:8080/debug/pprof/profile?seconds=5"
```

### Step 8: Open browser UI for heap or CPU

```powershell
go tool pprof -http=:8090 http://localhost:8080/debug/pprof/heap
```

Then open:

```text
http://localhost:8090
```

This helps for future visual analysis.

### Step 9: Check goroutines

```powershell
curl "http://localhost:8080/debug/pprof/goroutine?debug=1"
```

Detailed stack dump:

```powershell
curl "http://localhost:8080/debug/pprof/goroutine?debug=2"
```

Use this when:

- request is hanging
- too many goroutines are running
- you suspect deadlock or blocked work

### Step 10: Check trace

```powershell
curl http://localhost:8080/debug/pprof/trace?seconds=5 -o trace.out
go tool trace trace.out
```

Use trace when you want deeper timing and scheduler analysis.

## How To Read Results

### When checking logs

Look for:

- repeated slow requests
- wrong status codes like `400`, `404`, `500`
- unusual paths
- too many repeated hits on one endpoint

### When checking heap

Look for:

- large memory consumers
- functions allocating too much
- repeated allocations during requests

### When checking CPU

Look for:

- functions spending the most execution time
- unexpected hot paths
- expensive loops or JSON work

### When checking goroutines

Look for:

- blocked goroutines
- too many goroutines created
- repeated waiting patterns

## Future Reference Commands

### Normal route checks

```powershell
curl http://localhost:8080/health
curl http://localhost:8080/api/users
```

### Logs

```powershell
type .\execution-log-YYYY-MM-DD.txt
```

### Heap

```powershell
go tool pprof http://localhost:8080/debug/pprof/heap
```

### CPU

```powershell
go tool pprof -top "http://localhost:8080/debug/pprof/profile?seconds=5"
```

### UI

```powershell
go tool pprof -http=:8090 http://localhost:8080/debug/pprof/heap
```

### Goroutine dump

```powershell
curl "http://localhost:8080/debug/pprof/goroutine?debug=2"
```

## Important Notes

- `pprof` is very useful for study and debugging
- Do not expose `/debug/pprof/` publicly in production without protection
- Use `pprof` only when the server is running
- Generate some traffic before checking profiles, otherwise data will be minimal

## Summary

What was implemented:

- request logging middleware
- `pprof` routes on the same HTTP server
- fallback log-file handling

Where to start next time:

1. Run the server
2. Hit `/health` or `/api/users`
3. Check the log file
4. Open `/debug/pprof/`
5. Run `go tool pprof -top ...`
