Запуск "docker run -p 5000:5000 image_name"

Api имеет 5 запросов:
POST http://127.0.0.1:8080/control/task/add_new_task/
    -Добавляет задачу
PUT http://127.0.0.1:8080/control/task/update_task/
    -Изменяет состояние задачи
DELETE http://127.0.0.1:8080/control/task/delete_task/
    -Удаляет задачу
GET http://127.0.0.1:8080/control/task/get_info_task/
    -Выдает информацию по конктерной задаче
GET http://127.0.0.1:8080/control/task/get_all_task/
    -Выдает все задачи
