# Cassini
An HTTP probe built for sanity checking production REST endpoints post-deployment 

### MVP
- Takes in a config file and makes a get request to some dummy server, and makes assertions based on the responses

### Usage
```bash
# Run dummy server (runs on :8000)
cd server && go run dummyServer.go

#Build
make

# Run cassini
./cassini path/to/config
```

### Config Options
We use a .yaml file for configuration

For now we can only assert status codes

```yaml
# Declares a Cassini Service
service:
  name: "test service"
  url: "http://localhost:8000/"
  intervalUnit: "second" # Options: second, hour, minute
  intervalAmount: 5
  tests:
    - name: "test1"
      assertStatus: 200 # this is what gets asserted when we run our scheduled tests in the service
      method: "GET"
    - name: "testing post req"
      assertStatus: 201
      method: "POST"
```

Based on the above config we are targeting localhost:8000 every 5 seconds with 2 tests
