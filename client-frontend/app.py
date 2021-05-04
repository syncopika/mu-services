from flask import Flask, render_template, url_for
import requests

app = Flask(__name__)

@app.route('/')
def hello_world():
    return render_template('app.html')
	
@app.route('/somewhere', methods=['GET'])
def go_somewhere():
	return render_template('somewhere.html')

# this is kinda dumb but just for fun
# eventually can turn to POST to customize the request? i.e. get diff timezone times, etc.
@app.route('/getTime', methods=['GET'])
def get_time():
	# make a call to the timekeeper service and return the info
	# note we're currently assuming it's on port 6060, i.e. the service was started like:
	# docker run --publish 6060:8080 --name test timekeeper
	time = "sorry, there appears to be an issue with the timekeeper right now. please try again later."
	
	# https://docs.docker.com/docker-for-windows/networking/
	try:
		req = requests.get('http://host.docker.internal:6060/getTime')
	
		if req.status_code == 200:
			time = req.text
			
	except Exception as err:
		print(err)
		
	return {'time': time}