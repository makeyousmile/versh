// Токен бота и ID чата
const token = '468353575:AAFBaLyb20M3iMmjqBnm3js-qNrRjL81bFk';
const chatId = '-4514269523';

// Функция для отправки сообщения
function sendMessage() {
    const message = document.getElementById('message').value;

    // URL для запроса к Telegram API
    const url = `https://api.telegram.org/bot${token}/sendMessage`;

    // Данные запроса
    const data = {
        chat_id: chatId,
        text: message
    };

    // Отправка POST запроса на сервер Телеграм
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(result => {
            if (result.ok) {
                document.getElementById('status').textContent = 'Сообщение успешно отправлено!';
            } else {
                document.getElementById('status').textContent = 'Ошибка отправки сообщения!';
            }
        })
        .catch(error => {
            document.getElementById('status').textContent = 'Ошибка подключения!';
            console.error('Error:', error);
        });
}
