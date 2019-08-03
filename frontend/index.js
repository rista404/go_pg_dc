const fetch = require('node-fetch')

module.exports = async (_req, res) => {
	const go = await (await fetch('http://backend-service:8080/')).json()
	console.log(go)
	res.end('From Go app: ' + JSON.stringify(go))
}