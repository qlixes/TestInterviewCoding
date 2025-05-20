const http = require("http");
const url = require("url");

const server = http.createServer((req, res) => {

    switch(req.method) {

    }
});

const port = 3000;
server.listen(port, () => {
    console.log("Server running in ${port}");
});
