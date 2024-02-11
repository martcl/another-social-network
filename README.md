# ActivityPub experiment with Golang

Experimenting with ActivityPub and how it can be implemented. What can this protocol enable us to do other than serve as a _social network_ protocol?

## Local development
*"Works on my machine™"*

Start the services
```sh
nix-shell -p skaffold minikube kubectl go --run zsh
minikube start
minikube addons enable ingress
skaffold dev
```

Update `/etc/hosts` with the minikube IP and the domain name
```sh
echo "# Kubernetes social network" >> /etc/hosts
echo "$(minikube ip) social-network.local" >> /etc/hosts
echo "$(minikube ip) couchdb.local" >> /etc/hosts
```

Test the service
```sh
curl http://social-network.local
curl http://social-network.local/.well-known/webfinger
```

Look at the database http://couchdb.local/_utils