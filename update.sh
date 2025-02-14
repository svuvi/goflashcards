#!/bin/bash

SCRIPT_DIR=$(dirname "$0")

# 1. Загружаем обновление
echo "Загружаю обновление с github.com/svuvi/goflashcards.git..."
git -C "$SCRIPT_DIR" pull origin main

if [ $? -ne 0 ]; then
    echo "Git pull не удался. Отмена."
    exit 1
fi

# 2. Билдим приложение
echo "Создаю билд приложения..."
make

if [ $? -ne 0 ]; then
    echo "Билд не удался. Отмена."
    exit 1
fi

# 3. Останавливаем сервис
echo "Останавливаю сервис goflashcards..."
sudo service goflashcards stop

if [ $? -ne 0 ]; then
    echo "Не удалось остановить сервис goflashcards. Отмена."
    exit 1
fi

# 4. Бэкап базы данных
datetime=$(date +"%Y%m%d-%H%M%S")
backup_path="$HOME/DBbackups/goflashcards/$datetime-prod.db"
echo "Бэкаплю базу данных в $backup_path..."
cp "$SCRIPT_DIR/database.db" "$backup_path"

if [ $? -ne 0 ]; then
    echo "Бэкап базы данных не удался. Восстанавливаю старую версию сервиса."
    sudo service goflashcards start
    if [ $? -ne 0 ]; then
        echo "Рестарт сервиса не удался. Требуется ручное вмешательство."
        exit 1
    fi
    exit 1
fi

# 6. Заменяем старую версию
echo "Заменяем старую версию на новую..."

if mv "$SCRIPT_DIR/goflashcards" "$SCRIPT_DIR/goflashcards-old"; then
    if mv "$SCRIPT_DIR/goflashcards-new" "$SCRIPT_DIR/goflashcards"; then
        echo "Успешно заменена старая версия на новую."
    else
        echo "Не удалось заменить goflashcards-new на goflashcards. Восстанавливаю старую версию."
        if mv "$SCRIPT_DIR/goflashcards-old" "$SCRIPT_DIR/goflashcards"; then
            echo "Старая версия восстановлена."
        else
            echo "Не удалось восстановить старую версию. Требуется ручное вмешательство."
        fi
        exit 1
    fi
else
    echo "Не удалось переименовать goflashcards. Отмена."
    exit 1
fi

# 7. Запускаем сервис
echo "Запускаю сервис goflashcards..."
sudo service goflashcards start

if [ $? -ne 0 ]; then
    echo "Не удалось запустить сервис goflashcards. Проверьте логи."
    exit 1
fi

# 8. Удаляем старую версию
echo "Удаляем старую версию..."
rm -f "$SCRIPT_DIR/goflashcards-old"