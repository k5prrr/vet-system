
Из самого проекта
тут админка для вет клиники, для управления записями и графиками работ
так же ценный crud на go
не стал использовать дженерики и горм, так как определение типов затратно (да из-за этого есть дубли, но таков выбор в сторону скорости)
Цель: Максимально чистый проект без сторонних зависимостей + максимально быстрый(относительно go)

По архитектуре: гексоганальная
внешние и внутренние адаптеры разделены расположением
ещё есть инверсия зависимостей


## START
В корне проекта
cp .env.example .env

chmod +x make.sh
./make.sh help
./make.sh rm_base
./make.sh base
./make.sh run

Далее
Смотрим http://localhost:8080/lk/
Доступы в .env
Телефон: 10001110011
Второе поле: ps


## Прочее
curl -X 'POST' 'http://localhost:8080/api/create_admin'
podman compose -f app.yml up --build -d

## поднимает БД и adminer
podman-compose down
podman-compose -f db.yml up -d
podman ps -a
http://localhost:45010/

## Поднятие WEB front для теста
podman-compose -f front.yml up -d
http://localhost:8081/
http://localhost:8081/lk/
http://localhost:8081/api/doc/
