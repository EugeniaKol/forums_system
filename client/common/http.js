const fetch = require('node-fetch');

const Client = (baseUrl) => {

    return {
        get: (url) => {
            var res_forums = fetch(baseUrl + url).then((resp) => resp.json());
            return res_forums
                
        },
        post: (url, nickname, interests) => {
            let user = {
                nickname,
                interests
            };
            var res_user = fetch(baseUrl + url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
                })
                .then((resp) => resp.json());
                return res_user
        }
    };
};

module.exports = { Client };
