class Main {
    constructor() {

    }

    name: string = "";

    show(text: string): void {
        let $ele: JQuery = $("h1");
        $ele.text(text);
    }
}

let main = new Main();
main.name = "Hello TypeScript33ddd";
main.show(main.name);