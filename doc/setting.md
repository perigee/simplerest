# Design

## Openstack/AWS credential as environment variables

Create the container with environment variables setting

### AWS

- AWS_ACCESS_KEY_ID – AWS access key.
- AWS_SECRET_ACCESS_KEY – AWS secret key.
- AWS_SESSION_TOKEN – session token.
- AWS_DEFAULT_REGION – AWS region.


### Openstack

- OS_TENANT_NAME
- OS_USERNAME
- OS_PASSWORD
- OS_AUTH_URL

## Access keys

### VM default keypair

### Chef keypair



## Dependencies

- github.com/fsouza/go-dockerclient



## Parameters


### From gateway
SpaceID
RequestID


### Self tracking IDs
StackID
ResourceID



## TODO list

1. launch the docker container with all parameters
2. parser yml file - using HCL package to parser the variables (variables file)

### CRUD API
1. resource: GET


## References

- A repo using docker client github.com/controlroom/lincoln/backends/docker/container.go
- Async web service golang http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
- chan chan http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/