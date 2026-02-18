# Vulcan CLI — Go CLI для создания модульных микросервисов

**Vulcan** — лёгкий и понятный CLI-инструмент для быстрого scaffolding API-модулей на Go по паттерну MVC.
Использует [Gorilla MUX](https://github.com/gorilla/mux) для маршрутизации.

С Vulcan можно создавать новые модули, добавлять контроллеры, сервисы, миддлвары и другие компоненты одной командой.

---

## Основные возможности

* Быстро scaffold модульной MVC-архитектуры.
* Автоматическая генерация контроллеров, сервисов, моделей, миддлваров, репозиториев, интерфейсов и конфигов.
* Поддержка кастомных шаблонов и структуры проекта.
* Приведение названий компонентов к единому виду: `BlogController`, `blog_controller`, `blog-controller` → `blog_controller.go` с PascalCase структурами и camelCase переменными.
* Поддержка ключа `--force` для перезаписи файлов.
* Подсветка успешных действий и ошибок через [fatih/color](https://github.com/fatih/color).
* Управление командами, шаблонами и структурой проекта через `settings.yaml`.

---

## Быстрый старт

### 1. Клонируйте Vulcan

```bash
git clone https://github.com/knyazev-ro/vulcan-cli.git
```

Или скопируйте папку `vulcan/` в ваш проект.

---

### 2. Соберите CLI

```bash
cd vulcan
go build -o ../vulcan.exe .
cd ..
```

В корне проекта появится `vulcan.exe`.

---

### 3. Используйте команды Vulcan

```bash
vulcan.exe create:module <module_name>
```

Создаёт новый модуль с базовой структурой и шаблонами.

---

## Доступные команды

```plaintext
vulcan.exe create:module <module_name>                   - Создать новый модуль
vulcan.exe create:controller <name> <module>            - Добавить контроллер
vulcan.exe create:middleware <name> <module>            - Добавить middleware
vulcan.exe create:model <name> <module>                 - Создать модель
vulcan.exe create:repository <name> <module>            - Создать репозиторий
vulcan.exe create:service <name> <module>               - Создать сервис
vulcan.exe create:interface <name> <module>             - Создать интерфейс
vulcan.exe create:config <name> <module>                - Создать конфиг
vulcan.exe create:test <name> <module>                  - Создать тест
vulcan.exe remove:module <module_name>                  - Удалить модуль

vulcan.exe help                                         - Показать справку

Пример использования:

vulcan.exe create:controller user blog --force
```

---

## Настройка команд и структуры проекта

Файл `vulcan-cli/settings.yaml` позволяет управлять доступными командами, шаблонами и структурой проекта:

```yaml
commands:
  create-init: true
  create-model: true
  create-config: true
  create-service: true
  create-interface: true
  create-controller: true
  create-middleware: true
  create-repository: true
  remove-module: true

templates:
  service: "service.tmpl"
  controller: "controller.tmpl"
  model: "model.tmpl"
  interface: "interface.tmpl"
  config: "config.tmpl"
  config-utils: "config-utils.tmpl"
  config-base: "config-base.tmpl"
  config-database: "config-database.tmpl"
  config-server: "config-server.tmpl"
  module: "module.tmpl"
  middleware: "middleware.tmpl"
  repository: "repository.tmpl"
  dockerfile: "dockerfile.tmpl"
  gitignore: "gitignore.tmpl"
  readme: "readme.tmpl"
  route: "route.tmpl"
  env-example: "env-example.tmpl"
  github-workflows: "github-workflows.tmpl"

generated-file-structure:
  docs: "docs"
  scripts: ".scripts"
  configs: "configs"
  config_utils: "configs/utils"
  tests: "tests"
  docker: "docker"
  src: "src"
  utils: "src/utils"
  enums: "src/enums"
  models: "src/models"
  routes: "src/routes"
  services: "src/services"
  interfaces: "src/interfaces"
  middlewares: "src/middlewares"
  controllers: "src/controllers"
  repositories: "src/repositories"
  github-workflows: ".github/workflows"
```

---

## Структура проекта после инициализации

```
your_project/
├── vulcan.exe
├── <module_name>/
│   ├── src/
│   │   ├── controllers/
│   │   │   └── user_controller.go
│   │   ├── middlewares/
│   │   │   └── auth_middleware.go
│   │   ├── routes/
│   │   │   └── routes.go
│   │   ├── models/
│   │   ├── repositories/
│   │   ├── services/
│   │   ├── interfaces/
│   │   ├── utils/
│   │   └── enums/
│   ├── configs/
│   ├── tests/
│   ├── scripts/
│   ├── docker/
│   └── docs/
```

---

## Требования

* Go 1.18 или выше.
* Поддержка Windows (`vulcan.exe`) или сборка для других платформ без ключа `-o`.

---

## Дополнительно

* Все имена нормализуются: snake_case для файлов и папок, PascalCase для структур, camelCase для переменных.
* Ключ `--force` позволяет перезаписывать файлы.
* Подсветка ошибок и успешных сообщений через `github.com/fatih/color`.
* Планируется поддержка генерации моделей, Swagger документации и автокомплита CLI.

---

## Генерация тестов

Для любой команды `create:` можно указать флаг `--with-tests`, чтобы автоматически создать тестовый файл в папке `tests`.

* Имя теста формируется как `<component_name>_test.go`.
* Внутри файла шаблон теста: `func Test<ComponentName>(t *testing.T)`.

Пример:

```bash
vulcan.exe create:service payment_service shop --with-tests
```

Создаст:

```
src/services/payment_service.go
tests/payment_service_test.go
```