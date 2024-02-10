# ActivityPub experiment with Golang

Experimenting with ActivityPub and how it can be implemented. What can this protocol enable us to do other than serve as a _social network_ protocol?

## Local development
*Works on my machine™*

Start the services
```sh
nix-shell -p skaffold minikube kubectl go --run zsh
minikube start
skaffold dev
```

Get the minikube IP
```sh
minikube ip
```

Test the service
```sh
curl http://<minikube-ip>
curl http://<minikube-ip>/.well-known/webfinger
```
