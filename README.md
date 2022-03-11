# nobahar-1401

## set-up postgresql database

```zsh
helm repo add bitnami https://charts.bitnami.com/bitnami

helm install my-release bitnami/postgresql --set primary.persistence.size=1Gi --set readReplicas.persistence.size=1Gi

kubectl port-forward svc/postgresql 5432:5432
```
