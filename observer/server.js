const express = require('express');
const path = require('path')
const app = express();
const port = 3000;
const currServices = {'services': {}};

app.use(express.json());
app.use(express.static((path.join(__dirname , ""))));

// for when the client polls this server for updates on current services
app.get('/getCurrServices', (req, res) => {
	res.send(currServices);
});

// for when services update this server on their current status
app.post('/serviceUpdate', (req, res) => {
	//console.log(req.body);
	const serviceName = req.body.serviceName;
	const serviceStatus = req.body.serviceStatus;
	const timestamp = req.body.timestamp;
	
	currServices.services[serviceName] = {
		'status': serviceStatus,
		'timestamp': timestamp,
	}
	
	res.send("ok thanks");
});

app.listen(port, () => console.log("listening on port: " + port));