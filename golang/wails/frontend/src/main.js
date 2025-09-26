import {GetURL} from '../wailsjs/go/main/App';

document.addEventListener('DOMContentLoaded', () => {
    try {
        GetURL()
            .then((result) => {
                window.location.href = result
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
});
