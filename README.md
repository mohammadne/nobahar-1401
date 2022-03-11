# nobahar-1401

## how to run

```zsh
go run cmd/root.go server
```

## set-up postgresql database

```zsh
helm repo add bitnami https://charts.bitnami.com/bitnami

helm install my-release bitnami/postgresql --set primary.persistence.size=1Gi --set readReplicas.persistence.size=1Gi

kubectl port-forward svc/postgresql 5432:5432
```

## connect to postgresql database

```zsh
kubectl exec --stdin --tty postgreql -- /bin/bash

psql -d nobahar -U username
```
