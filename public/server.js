/**
 * This is created just for serving this directory absolutely static with nodejs
 * The main go project serve this direcory itself
 */


var connect = require('connect');
var serveStatic = require('serve-static');
connect().use(serveStatic(__dirname)).listen(8080, function(){
    console.log('Server running on 8080...');
});
