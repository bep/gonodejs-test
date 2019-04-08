const http = require('http');
const babel = require('@babel/core');
const port = 8182

http.createServer(function (req, res) {
	if (req.method === 'POST') {
		let body = '';
		req.on('data', chunk => {
			body += chunk.toString();
		});
		req.on('end', () => {
			const { code } = babel.transformSync(body, {code: true})
			console.log("Got %s", code)
			res.end();
		});
	}
	res.write("Node!");
	res.end();
}).listen(port);

console.log('Node server started at port: %d', port);