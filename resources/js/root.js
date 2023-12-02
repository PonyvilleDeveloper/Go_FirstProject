// furniture.js

function Main() {
    this.__proto__ = new Util();
    const util = this.__proto__;

    this.furnitureList = [];

    this.newFurniture = () => {
        console.log("Функция добавления новой мебели");
    };

    const action = () => {
        document.querySelectorAll('.action').forEach(el => {
            el.onclick = () => {
                this[el.dataset.action](el.dataset);
            };
        });
    };

    this.modals = {};

    this.init = function() {
        action();
        util.get('/furniture', info => {
            if (info.error && info.error === "invalid_verifier") {
                this.auth();
                return;
            }
            this.furnitureList = info;
            this.view();
        });

        if (util.emptyToken()) {
            this.auth();
            return;
        }

        this.getFurniture();
    };

    this.auth = () => {
        util.modal("auth", "show");
        util.id("auth").modal("show");
    };

    this.getFurniture = () => {
        util.get('/furniture', info => {
            if (info.error && info.error === "invalid_verifier") {
                this.auth();
                return;
            }
            this.furnitureList = info;
            this.view();
        });
    };

    this.view = () => {
        action();
        if (this.furnitureList.length < 1) return;
        const furniture = this.furnitureList
            .map(f => util.parse(util.tpl.tr, f)).join("");
        util.id("root").innerHTML = util.parse(util.tpl.table, {
            furniture: furniture
        });
    };

    this.authIn = () => {
        util.id("auth").modal("hide");
        util.post("/login", {
            login: util.id("authLogin").value,
            password: util.id("authPassword").value,
        }, resp => {
            if (resp.error) {
                this.auth();
                return;
            }
            util.setUser(resp);
            this.getFurniture();
        });
    };

    this.deleteFurniture = id => {
        util.delete("/furniture/" + id);
    };

    this.init();
}

const main = new Main();
