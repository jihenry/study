"use strict";
var Main = /** @class */ (function () {
    function Main() {
        this.name = "";
    }
    Main.prototype.show = function (text) {
        var $ele = $("h1");
        $ele.text(text);
    };
    return Main;
}());
var main = new Main();
main.name = "Hello TypeScript33ddd";
main.show(main.name);
//# sourceMappingURL=index.js.map