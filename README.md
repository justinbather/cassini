# Cassini
An HTTP probe built for sanity checking production REST endpoints post-deployment


# thinking out loud
- It would be nice if a user could easily specify through some sort of config file:
  - What endpoint to hit
  - What method
  - The data
  - And how often
  - They should also be able to have assertions in the responses that if failed, can throw an alert

- Cassini can be deployed as a long running service in a k8s cluster or something similar
- We can create a docker image so that a user could create their own image with their config in it to be deployed 


### MVP
- Takes in a config file and makes a get request to some dummy server, and makes assertions based on the responses

### Usage
```bash
./cassini path/to/config
```

