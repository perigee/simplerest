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

