function Util() {
    let token = "";
    const ls = localStorage.getItem("token"); // Добавлено ключевое слово const перед ls
    if (ls !== null) token = ls;

    this.emptyToken = () => {
        return token === "";
    };

    let user = {};

    this.get = (url, callback) => {
        fetch(url, {
            method: "GET",
            headers: {
                "Content-Type": "application/json", // Исправлено значение Content-Type
                "Authorization": token
            }
        }).then(data => data.json())
            .then(callback);
    };

    this.post = (url, data, callback) => {
        fetch(url, {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "Content-Type": "application/json", // Исправлено значение Content-Type
                "Authorization": token
            }
        }).then(data => data.json())
            .then(callback);
    };

    this.delete = (url, callback) => { // Исправлено название метода на delete
        fetch(url, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json", // Исправлено значение Content-Type
                "Authorization": token
            }
        }).then(data => data.json())
            .then(callback);
    };

    this.id = el => document.getElementById(el);

    this.modals = {};

    this.modal = (id, action) => {
        if (!this.modals[id]) {
            this.modals[id] = new bootstrap.Modal('#' + id);
        }
        this.modals[id][action]();
    };

    this.setUser = usr => {
        user = usr;
        token = user.token;
        localStorage.setItem("token", token);
    };


    this.parse = (content, params) => {
        const param = Object.assign({}, params); // Добавлено ключевое слово const перед param
        return content.replace(/{{(\w+)}}/g, (str, key) => {
            if (param[key] === undefined) // Убрано лишнее присваивание
                return '';
            return param[key];
        });
    };
}
