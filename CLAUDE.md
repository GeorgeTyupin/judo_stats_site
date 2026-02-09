# Judo Archive - Tech Stack

## Tech Stack

**Backend:**
- Go 1.21+ - основной язык программирования
- Chi v5 - минималистичный HTTP роутер
- PostgreSQL - база данных

**Frontend:**
- Templ - типизированные HTML шаблоны (компилируются в Go)
- Vanilla CSS - стили в отдельных файлах (static/css/)
- Vanilla JS - скрипты в отдельных файлов (static/js/)
- HTMX - динамические обновления DOM
- Lucide Icons - SVG иконки (https://lucide.dev)

## Правила использования иконок

**ВАЖНО:** В проекте используются только SVG иконки из библиотеки Lucide. Эмоджи НЕ ИСПОЛЬЗУЮТСЯ.

### Почему SVG вместо эмоджи:
- Эмоджи выглядят инфантильно и непрофессионально
- SVG иконки обеспечивают консистентный дизайн
- SVG легко кастомизируются через CSS (размер, цвет)
- SVG иконки выглядят одинаково во всех браузерах

### Как использовать Lucide иконки:

```html
<!-- Базовое использование -->
<i data-lucide="icon-name" class="w-4 h-4"></i>

<!-- С кастомными размерами -->
<i data-lucide="search" class="w-6 h-6"></i>

<!-- В группе с текстом -->
<span class="flex items-center gap-1">
    <i data-lucide="flag" class="w-4 h-4"></i>
    Текст рядом с иконкой
</span>
```

### Часто используемые иконки:
- `search` - поиск
- `user`, `users` - пользователи
- `flag` - страна/флаг
- `medal`, `award`, `trophy` - награды и достижения
- `calendar` - дата
- `map-pin`, `map` - место, локация
- `building` - здание, клуб
- `weight` - весовая категория
- `settings` - настройки

### Инициализация:
Lucide иконки автоматически инициализируются в base.templ через:
```javascript
lucide.createIcons();
```

## Архитектура

- **Server-Side Rendering** - Go генерирует HTML через Templ
- **Partial Updates** - HTMX обновляет части страницы без перезагрузки
- **No SPA** - традиционный серверный рендеринг с интерактивностью

## Структура проекта

```
judo-archive/
├── cmd/server/main.go              # Точка входа
├── internal/
│   ├── app/                        # Конфигурация приложения
│   ├── handlers/                   # HTTP handlers
│   ├── models/                     # Go структуры данных
│   ├── services/                   # Бизнес логика
│   └── repository/                 # Работа с БД
├── templates/
│   ├── layouts/                    # Базовые layouts
│   ├── pages/                      # Страницы
│   └── components/                 # Переиспользуемые компоненты
└── static/
    ├── css/                        # CSS файлы (style.css - общий, остальные - для страниц)
    └── js/                         # JavaScript файлы
```

## Окружение разработки

**ВАЖНО:** Проект разрабатывается в Docker контейнерах. Используй Docker Desktop или docker compose команды.

## Команды разработки (Docker)

### Через Docker Desktop GUI:
- Запуск/остановка контейнеров через интерфейс
- Просмотр логов в реальном времени
- **После изменений кода**: перезапусти контейнер `server` через GUI (правый клик → Restart)

### Через терминал (docker compose):

```bash
# Первый запуск (собрать образы и запустить)
docker compose up -d --build

# Обычный запуск
docker compose up -d

# Остановить все сервисы
docker compose down

# ВАЖНО: Перезапустить контейнер после изменений кода
docker compose restart server

# Посмотреть логи в реальном времени
docker compose logs -f server

# Открыть shell внутри контейнера
docker compose exec server sh

# Открыть psql для работы с базой данных
docker compose exec db psql -U judo -d judo_archive

# Полная очистка (удалить контейнеры и volumes)
docker compose down -v
docker compose up -d --build
```

### Workflow разработки

1. **Редактируешь код** (`.templ`, `.go`, `.css`, `.js` файлы)
2. **Перезапускаешь контейнер**:
   - Через Docker Desktop: правый клик на `server` → Restart
   - Через терминал: `docker compose restart server`
3. **Templ шаблоны генерируются автоматически** при каждом запуске контейнера
4. **Перезапуск занимает ~2-3 секунды**

**Важно:**
- НЕ нужно запускать `templ generate` вручную - это делается автоматически при старте контейнера
- После КАЖДОГО изменения кода нужно перезапускать контейнер `server`
- После изменений в Dockerfile или compose.yaml пересобери: `docker compose up -d --build`

## Принципы анимаций и UI

### Философия анимаций
Все интерактивные элементы должны иметь плавные, приятные анимации. Анимации делают интерфейс живым и современным.

### Стандартные параметры анимаций

**Timing функции:**
- `cubic-bezier(0.4, 0, 0.2, 1)` - основная для большинства переходов
- `ease-out` - для появления элементов
- `ease-in-out` - для двусторонних переходов

**Длительность:**
- Короткие действия (hover, focus): `0.3s`
- Средние действия (раскрытие, переключение): `0.4-0.5s`
- Длинные действия (сложные анимации): `0.6s`

### Типы анимаций

**1. Hover эффекты для кнопок:**
```css
.button {
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
    overflow: hidden;
}

/* Ripple эффект */
.button::before {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    width: 0;
    height: 0;
    border-radius: 50%;
    background: rgba(20, 184, 166, 0.1);
    transform: translate(-50%, -50%);
    transition: width 0.6s ease, height 0.6s ease;
}

.button:hover::before {
    width: 300px;
    height: 300px;
}

.button:hover {
    transform: translateY(-2px) scale(1.02);
}

.button:active {
    transform: scale(0.98);
}
```

**2. Появление элементов с задержкой (staggered):**
```css
.item {
    animation: fadeIn 0.3s ease-out backwards;
}

.item:nth-child(1) { animation-delay: 0.05s; }
.item:nth-child(2) { animation-delay: 0.1s; }
.item:nth-child(3) { animation-delay: 0.15s; }

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: scale(0.95);
    }
    to {
        opacity: 1;
        transform: scale(1);
    }
}
```

**3. Раскрытие/сворачивание контента:**
```css
.wrapper {
    max-height: 1000px;
    opacity: 1;
    transform: translateY(0);
    transition: max-height 0.5s cubic-bezier(0.4, 0, 0.2, 1),
                opacity 0.4s ease,
                transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.wrapper.hidden {
    max-height: 0;
    opacity: 0;
    overflow: hidden;
    transform: translateY(-20px);
}
```

**4. Вращение иконок:**
```css
.icon {
    transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    display: inline-block;
}

.collapsed .icon {
    transform: rotate(180deg);
}
```

**5. Переключение контента (fade-in-up):**
```css
.content {
    animation: fadeInUp 0.4s ease-out;
}

@keyframes fadeInUp {
    from {
        opacity: 0;
        transform: translateY(15px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
```

### Правила применения

1. **Всегда добавляй анимации для:**
   - Кнопок (hover, active состояния)
   - Появления/скрытия элементов
   - Переключения табов/категорий
   - Раскрытия/сворачивания секций
   - Фокуса на input полях

2. **Используй transform вместо абсолютных значений:**
   - ✅ `transform: translateY(-2px)`
   - ❌ `margin-top: -2px`

3. **Группируй связанные элементы:**
   - Если элементы появляются вместе, используй staggered animation
   - Добавляй небольшие задержки (0.05s шаг)

4. **Ripple эффект для интерактивных элементов:**
   - Используй `::before` или `::after` псевдоэлементы
   - Центрируй эффект с помощью `translate(-50%, -50%)`
   - Используй `overflow: hidden` на родителе

## Конвенции коммитов

Используйте следующие префиксы для коммитов:

- `feat:` - новая функциональность
- `fix:` - исправление бага
- `refactor:` - рефакторинг кода без изменения функциональности
- `style:` - изменения стилей, форматирования
- `docs:` - изменения в документации
- `test:` - добавление или изменение тестов
- `chore:` - рутинные задачи, обновление зависимостей
- `perf:` - улучшение производительности

**Примеры:**
```
feat: добавлена страница профиля дзюдоиста с медалями
fix: исправлено форматирование даты турнира
refactor: вынесен компонент поиска
docs: обновлен README с инструкциями по Docker
chore: добавлен .gitignore для Go проекта
```