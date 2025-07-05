import requests
import json

def test_get_all_messages():
    my_token = "QUReZgN4xAE3YOdrNVc5RKoaShLYmLlkN"
    my_url = "http://127.0.0.1:8000/api/secret/all_messages"

    my_headers = {
        "Authorization": my_token,
        "Content-Type": "application/json"
    }

    response = requests.get(my_url, headers=my_headers)

    print("Статус-код:", response.status_code)

    assert response.status_code == 200, "❌ Статус не 200"

    my_messages = response.json()

    print(f"Всего сообщений: {len(my_messages)}")

    print("\n📨 Полученные сообщения:")
    for msg in my_messages:
        print(json.dumps(msg, ensure_ascii=False, indent=2))  # красиво, построчно

# Не забудь вызвать функцию:
test_get_all_messages()