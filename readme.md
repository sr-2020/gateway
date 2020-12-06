[![Build Status](https://travis-ci.org/sr-2020/gateway.svg?branch=master)](https://travis-ci.org/sr-2020/gateway)
# Gateway

Helm deploy production
```
helm secrets upgrade gateway helm/gateway/ \
    -f helm/gateway/values.yaml \
    -f helm/gateway/values/production/values.yaml \
    -f helm/gateway/values/production/secrets.yaml
```

Helm deploy rc
```
helm secrets upgrade rc-gateway helm/gateway/ \
    -f helm/gateway/values.yaml \
    -f helm/gateway/values/rc/values.yaml \
    -f helm/gateway/values/rc/secrets.yaml
```

Update SSL
```
certbot certonly --manual -d '*.evarun.ru'  --logs-dir certbot --config-dir certbot --work-dir certbot
```
Create secret for Ingress

```
kubectl create secret tls tls-secret --key=privkey.pem --cert=fullchain.pem

```
