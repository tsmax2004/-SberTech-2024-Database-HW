# Работа CouchDB + PouchDB

1. Поднимаем CouchDB с помощью [docker-compose](hw3_pouchdb/docker-compose.yaml)

![image](imgs/hw_3_1.png)

2. Создаем БД `name_db` с помошью `curl`:

![image](imgs/hw_3_2.png)

3. Указываем в [html-файле](hw3_pouchdb/without_name.html) путь до поднятой БД - `http://admin:password@localhost:5984/name_db`

4. Открываем, видим, что никакое имя не отображается:

![image](imgs/hw_3_3.png)

5. Добавляем в бд документ с полем "name":

![image](imgs/hw_3_4.png)

6. Видим, что имя подгрузилось:

![image](imgs/hw_3_5.png)

7. После остановки бд на локалхосте и перезагрузки страницы имя все равно подгружается.
Файл с прочитанным именем сохранил [здесь](hw3_pouchdb/with_name.html)
