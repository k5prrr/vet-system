## поднимает БД и adminer 
podman-compose -f db.yml up -d
http://localhost:45010/

## Поднятие WEB front для теста
podman-compose -f front.yml up -d
http://localhost:8081/
http://localhost:8081/lk/

## показывает статусы поднятого 
podman ps -a

