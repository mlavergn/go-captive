package captive

// PortalJS text
var PortalJS = `function send() {
	const url = document.getElementById('url');
	const code = document.getElementById('code');
	const message = document.getElementById('message');
	const xhr = new XMLHttpRequest();
    xhr.open('POST', '/captive');
    xhr.onreadystatechange = () => {
        if (xhr.status === 200 && xhr.readyState === XMLHttpRequest.DONE) {
			const data = JSON.parse(xhr.response);
			message.innerHTML = data.message;
		}	
	}
	xhr.send(JSON.stringify({
		op: 'code',
		data: {
			url: url.value,
			code: code.value
		}
	}));
}`

// PortalHTML text
var PortalHTML = `<!DOCTYPE html>
<html lang="en">

<head>
	<title>Captive Portal Handler</title>
	<style>
		body {
			background-color: black;
		}

		label {
			color: white;
			font-family: sans-serif;
		}

		input {
			font-family: monospace;
		}
	</style>
	<script>` +
	PortalJS +
	`</script>
	</head>

<body>
	<form action="/captive" method="get">
		<table>
			<tr>
				<td><label>URL</label></td>
				<td><input type="text" id="url" value="http://10.21.1.1:8002/" /></td>
				<td></td>
			</tr>
			<tr>
				<td><label>Code</label></td>
				<td><input type="text" id="code" value="" /></td>
				<td></td>
			</tr>
			<tr>
				<td></td>
				<td><input type="button" value="Send" onclick="send();"/></td>
				<td></td>
			</tr>
			<tr>
				<td></td>
				<td><label id="message"></label></td>
				<td></td>
			</tr>
		</table>
	</form>
</body>

</html>`
