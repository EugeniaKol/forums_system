// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require('./channels/client');
let myUrl = 'http://localhost:8080';

const client = channels.Client(myUrl);

const printForums(res) => {
    console.log('Displaying forums:');
    res.forEach((f) => {
        console.log('Forum name: ', f.name);
        console.log('Topic keyword: ', f.topicKeyword);
        console.log('Users: ', f.users.join(', '));
        console.lof('_______________________');
    });
}

// Scenario 1: Show existing forums.
client.showForums()
    .then((list) => printForums(list);
    )
    .catch((e) => console.log(e);
    );

// Scenario 2: Create new user.
let newNick = 'zhenshen';
const newInters = ['marvel movies', 'rock music'];

client.addUser(newNick, newInters)
    .then((resp) => {
        console.log('Adding user:', );
        console.log('Adding channel response:', resp);

    //forums now have modified 'users' field
        return client.showForums()
            .then((list) => printForums(list);
    })
    .catch((e) => console.log(e);
    );
