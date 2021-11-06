# S/KEY Melnikov Dorofeev 9317{23,24}

## Описание алгоритма

1. Вначале и клиент, и сервер нужно настроить на единую парольную фразу. Далее клиент отправляет серверу пакет инициализации, а сервер в ответ отправляет порядковый номер и случайное число - «зерно» (seed).
2. Клиент вводит секретную парольную фразу. Парольная фраза соединяется с «зерном», полученным от сервера в открытом виде. «Зерно» дает клиенту возможность использовать одну и ту же парольную фразу на множестве машин с разными «зернами».
3. Клиент многократно использует хэш-функцию и получает 64-разрядную итоговую величину. При каждом новом использовании количество хэш-циклов уменьшается на один.
4. Клиент передает одноразовый пароль на сервер, где он и проверяется. На сервере есть файл, в котором хранится одноразовый пароль, использованный в последнем успешном сеансе связи с каждым отдельным пользователем. Для проверки  система однократно пропускает полученный одноразовый пароль через хэш-функцию. Если результат этой операции совпадает с предыдущим паролем, хранящимся в файле, результат аутентификации считается положительным, а новый пароль сохраняется для дальнейшего использования.
5. Пользователь может инициализировать систему с помощью команды keyinit, которая дает возможность изменить секретную парольную фразу, количество циклов итерации и «зерно».

## Меню клиента и сервера
IP адрес и порт передаются, как аргументы при вызове программы

1. Launch	_запустить сервер/клиент для приёма/передачи сообщений_
2. KeyInit	_изменить стандартные параметры_
3. Exit		_выход из программы_

### Key init
Позволяет изменять параметры по умолчанию:

1. Парольная фраза (default: `passphrase`)
2. Количество циклов итерации (default: `1000`)
3. Зерно (default: random seed)	_только для сервера_

## Спецификация протокола

- Используемый язык -- **Go**. Для обмена данными используются сокеты
- Используемая хэш-функция -- **MD5**
- Размер передаваемого блока всегда **128 бит**
- Парольная фраза изначально определена у клиента и сервера
- Число итераций по умолчанию `n = 1000`

### Инициализация общения

- Клиент инициирует общение с сервером, посылая сообщение `!hello`
- Сервер отправляет двумя сообщениями: `id` клиента и `seed`
- Оба вырабатывают пароли для всех будущих идентификаций
- Далее клиент начинает отправлять сообщения

### Передача сообщений

Информация передаётся кусками по 128 бит. Порядок передачи следующий:

1. Пароль `i`-ой итерции
2. Сообщение, разбитое на блоки по 128 бит
3. Пустой кусок, как окончание передачи сообщения
	
### Обновление seed

В случае, если число итерации закончилось, а клиенту всё ещё нужно продолжить общение с сервером, то они меняют `seed` и вырабатывают ещё `n` одноразовых паролей

Как это происходит: как только сервер использовал последний одноразовый ключ, он рандомит новый seed и отправляет его клиенту. Клиент принимает seed и оба вырабатывают новые пароли. Далее клиент продолжает передачу сообщений
