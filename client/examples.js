// This file contains examples of scenarios implementation using
// the SDK for channels management.

const forums = require('./forums/client');
let myUrl = 'http://localhost:8080';

const client = forums.Client(myUrl);

const printForums = res => {
    console.log('Displaying forums:');
    res.forEach((f) => {
        console.log('Forum name: ', f.name);
        console.log('Topic keyword: ', f.topicKeyword);
        console.log('Users: ', f.users.join(', '));
        console.log('_______________________');
    });
}

// Scenario 1: Show existing forums.
client.showForums()
    .then((list) => printForums(list))
    .catch((e) => console.log(e));

// Scenario 2: Create new user.
let newNick = 'torass';
const newInters = ['ukraine-politics','rock-music'];

client.addUser(newNick, newInters)
    .then((resp) => {
        console.log('Adding user:', newNick);
        console.log('Inserting data:', resp);

        //forums now have modified 'users' field
        return client.showForums()
    })
            .then((list) => printForums(list))
    .catch((e) => console.log(e));
