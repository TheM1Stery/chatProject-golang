import * as Htmx from "htmx.org";


declare global {
    var htmx: typeof Htmx;
    interface Window {
        htmx: typeof Htmx;
    }
}



window.htmx = Htmx;
