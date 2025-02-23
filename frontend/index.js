document.getElementById("loginForm").addEventListener("submit", async function (event) {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const message = document.getElementById("message");

    console.log("Отправляем:", {email, password});

    try {
        const response = await fetch("http://localhost:8080/signup", {
            method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify({email, password})
        });

        const data = await response.json();
        console.log("Ответ сервера:", data);

        if (response.ok) {
            message.style.color = "green";
            message.textContent = "Успешный вход!";
        } else {
            message.textContent = data.message || "Ошибка входа";
        }
    } catch (error) {
        console.error("Ошибка запроса:", error);
        message.textContent = "Ошибка сети";
    }
});
