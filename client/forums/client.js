const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        showForums: () => client.get('/forums'),
        addUser: (nickName, interests) => client.post('/forums', { nickName, interests })
    }

};

module.exports = { Client };
