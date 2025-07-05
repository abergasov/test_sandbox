import requests
import json
import os
from dotenv import load_dotenv  # Импортируем dotenv

# Загружаем переменные из .env
load_dotenv()

# Получаем токен и хост из окружения
my_token = os.getenv("AUTH_TOKEN")
my_host = os.getenv("REMOTE_HOST")

# Выводим токен и хост в консоль для проверки (временно)
print("Токен из .env:", my_token)
print("Хост из .env:", my_host)

def test_get_all_messages():
    my_url = f"{my_host}/api/secret/all_messages"

    my_headers = {
        "Authorization": my_token,
        "Content-Type": "application/json"
    }

    response = requests.get(my_url, headers=my_headers)

    print("Статус-код:", response.status_code)
    assert response.status_code == 200, "❌ Статус не 200"

    my_messages = response.json()

    print(f"Всего сообщений: {len(my_messages)}\n")

    print("📨 Полученные сообщения:")
    for msg in my_messages:
        print(json.dumps(msg, ensure_ascii=False, indent=2))

# Запуск теста
test_get_all_messages()
