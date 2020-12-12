const fetch = require('node-fetch');

const Client = (baseUrl) => {

    return {
        get: (url) => {
            fetch(url)
                .then((resp) => resp.json());
        },
        post: (url, nickName, interests) => {
            let user = {
                nickName,
                interests
            };
            fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
                })
                .then((resp) => resp.json);
        }
    };
};

module.exports = { Client };
