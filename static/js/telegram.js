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

function sendForm() {
    const form = document.getElementById('form');
    const formData = new FormData(form);
    const entries = Array.from(formData.entries());


    const url = `https://api.telegram.org/bot${token}/sendMessage`;

    const dataString = entries.map(([key, value]) => `${key}=${value}`).join(' ');
    const maxLength = 4096;

    const sendMessage = async (text) => {
        const data = {
            chat_id: chatId,
            text: text,
        };

        try {
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();
            return result.ok;
        } catch (error) {
            console.error('Ошибка подключения:', error);
            return false;
        }
    };

    const sendInChunks = async (text) => {
        for (let i = 0; i < text.length; i += maxLength) {
            const chunk = text.substring(i, i + maxLength);
            const success = await sendMessage(chunk);
            if (!success) {
                document.getElementById('status').textContent = 'Ошибка отправки сообщения!';
                return;
            }
        }
        document.getElementById('status').textContent = 'Сообщение успешно отправлено!';
    };

    sendInChunks(dataString);
}

// Функция для отправки сообщения
function sendForm2() {
    const form = document.getElementById('form');
    const formData = new FormData(form);
    const entries = Array.from(formData.entries());

    // URL для запроса к Telegram API
    const url = `https://api.telegram.org/bot${token}/sendMessage`;
    const dataString = entries.map(([key, value]) => `${key}=${value}`).join('&');
    alert(dataString);
    // Данные запроса
    const data = {
        chat_id: chatId,
        text: dataString.slice(0, 70),
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