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


## показывает статусы поднятого 
podman ps -a


## INIT
in ROOT
chmod +x make.sh
./make.sh init





## Прочее
Формируем api
по всем лучшим правилам REST API json
напиши под это openapi.yaml

У нас есть роли:
не зарегистрированный пользователь
пользователь
доктор
админ

чтоб системе понять кто они, и проставить им роли, они отправляют
auth id и token в запросе либо в заголовках



кто угодно может создать запись, главное указать фио и номер телефона,
врач и админ могут менять запись

Post /api/record
{
"fio": "Иванов Иван Иванович",
"phone": "+7",
}
Get /api/record/5
Закрыт от неавторизованных

Put /api/record/5
Delete /api/record/5

Get /api/records Показать все
(у пользователя отображается только его)








создавать и менять пользователей может только админ

Создавать менять клиентов могут врачи и админы

Создавать менять животных могут врачи и админы

Создавать удалять timesheet могут только админы

