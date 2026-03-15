package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>PodPulse</title>

<style>

body{
	margin:0;
	font-family:"Segoe UI",system-ui,-apple-system,sans-serif;
	background:linear-gradient(135deg,#0f172a,#1e293b,#020617);
	height:100vh;
	display:flex;
	justify-content:center;
	align-items:center;
	color:white;
}

.container{
	backdrop-filter:blur(12px);
	background:rgba(255,255,255,0.05);
	border:1px solid rgba(255,255,255,0.08);
	border-radius:18px;
	padding:45px;
	width:420px;
	text-align:center;
	box-shadow:0 20px 45px rgba(0,0,0,0.45);
}

h1{
	font-size:2.2rem;
	margin-bottom:8px;
	background:linear-gradient(90deg,#38bdf8,#22c55e);
	-webkit-background-clip:text;
	-webkit-text-fill-color:transparent;
}

.tag{
	font-size:0.9rem;
	color:#94a3b8;
	margin-bottom:30px;
}

.label{
	color:#cbd5f5;
	font-size:1rem;
	margin-bottom:8px;
}

.host{
	font-size:1.3rem;
	background:#020617;
	padding:14px;
	border-radius:10px;
	border:1px solid #334155;
	font-weight:600;
	color:#22c55e;
}

.footer{
	margin-top:28px;
	font-size:0.8rem;
	color:#64748b;
}

.refresh{
	margin-top:18px;
	padding:10px 16px;
	border:none;
	border-radius:8px;
	background:#2563eb;
	color:white;
	cursor:pointer;
	font-size:0.9rem;
}

.refresh:hover{
	background:#1d4ed8;
}

</style>

</head>

<body>

<div class="container">

<h1>PodPulse 🚀</h1>

<div class="tag">
Kubernetes Pod Identity Demo
</div>

<div class="label">
Request served by
</div>

<div class="host">
{{.Host}}
</div>

<button class="refresh" onclick="location.reload()">
Refresh
</button>

<div class="footer">
Refresh the page to hit another pod in the cluster
</div>

</div>

</body>
</html>
`))

type PageData struct {
	Host string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	data := PageData{
		Host: host,
	}

	tpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("PodPulse server running on :3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
