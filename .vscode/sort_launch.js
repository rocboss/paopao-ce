let launch = require('./launch.json');
launch.configurations.sort((a, b) => a.name.localeCompare(b.name));
let fs = require('fs');
fs.writeFileSync('launch.json', JSON.stringify(launch, null, 4));
