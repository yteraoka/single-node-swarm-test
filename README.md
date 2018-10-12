Single node swarm でお手軽 rolling update
=========================================

```
nginx (global service) -> app (service) -> app container
                             `-----------> app container
                             `-----------> app container
```

create docker swarm (cluster)
-----------------------------

Single node swarm の場合は他の node と通信する必要がないので 127.0.0.1 だけ Listen させておく

```
docker swarm init --advertise-addr 127.0.0.1 --listen-addr 127.0.0.1:2377
```

build docker image
------------------

docker stack の deploy には build 済みの image が必要なので build する

container があれば

```
docker-compose build
```

deploy docker stack
-------------------

```
docker stack deploy --compose-file docker-compose.yml myapp
```

update docker image
-------------------

```
docker service update myapp_web --image myapp-app:1.0.3
```

scale container
---------------

```
docker service scale myapp_app=3
```
