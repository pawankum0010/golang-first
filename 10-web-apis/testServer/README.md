# Test Server

A simple HTTP server for testing your Go applications on GCP.

## Usage

### Run Locally
```bash
go run main.go
```

### Build and Run
```bash
go build -o test-server main.go
./test-server
```

### Test Endpoints

1. **Root endpoint**
   ```bash
   curl http://localhost:8080/
   ```

2. **Health check**
   ```bash
   curl http://localhost:8080/health
   ```

3. **Get all users**
   ```bash
   curl http://localhost:8080/api/users
   ```

4. **Get user by ID**
   ```bash
   curl http://localhost:8080/api/user?id=1
   ```

5. **Open pprof index**
   ```bash
   curl http://localhost:8080/debug/pprof/
   ```

6. **Capture CPU profile for 5 seconds**
   ```bash
   go tool pprof http://localhost:8080/debug/pprof/profile?seconds=5
   ```

7. **Inspect heap profile**
   ```bash
   go tool pprof http://localhost:8080/debug/pprof/heap
   ```

8. **Inspect goroutines in browser or terminal**
   ```bash
   curl http://localhost:8080/debug/pprof/goroutine?debug=1
   ```

### Step-by-Step: Learn pprof and Logs

1. Start the server:
   ```bash
   go run main.go
   ```

2. Generate some traffic in another terminal:
   ```bash
   curl http://localhost:8080/health
   curl http://localhost:8080/api/users
   curl http://localhost:8080/api/user?id=1
   ```

3. Check the request log created at the project root.
   If root write permission is blocked, the log is created in the current `testServer` folder:
   ```bash
   type ..\..\execution-log-YYYY-MM-DD.txt
   ```

4. Open the pprof index page:
   ```bash
   start http://localhost:8080/debug/pprof/
   ```

5. Record a CPU profile and inspect top functions:
   ```bash
   go tool pprof -top http://localhost:8080/debug/pprof/profile?seconds=5
   ```

6. Open the interactive pprof UI:
   ```bash
   go tool pprof -http=:8090 http://localhost:8080/debug/pprof/heap
   ```

7. Read goroutine details when debugging stuck requests:
   ```bash
   curl http://localhost:8080/debug/pprof/goroutine?debug=2
   ```

### Test on GCP Server

1. **SSH into your server**
   ```bash
   gcloud compute ssh YOUR_INSTANCE_NAME --zone YOUR_ZONE
   ```

2. **Navigate to directory**
   ```bash
   cd /path/to/golang-first/10-web-apis/testServer
   ```

3. **Run the server**
   ```bash
   go run main.go
   # OR run in background
   nohup go run main.go > server.log 2>&1 &
   ```

4. **Test from server**
   ```bash
   curl http://localhost:8080/health
   ```

5. **Test from external machine**
   ```bash
   curl http://YOUR_GCP_EXTERNAL_IP:8080/health
   ```

## Notes

- Server runs on port 8080 by default
- `pprof` is available on the same server under `/debug/pprof/`
- Every request is logged to `execution-log-YYYY-MM-DD.txt` in the project root, or the current folder if root writing is blocked
- Make sure firewall allows port 8080
- For production, use environment variables for configuration
- In production, do not expose `pprof` publicly without protection
- Add authentication/authorization as needed
