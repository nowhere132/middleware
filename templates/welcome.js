fetch("http://localhost:8000/checkuser", {
    mode: "no-cors",
    method: "GET",
    headers: {
        "Content-Type": "application/json",
    }
})
.then(response => response.json())
.then(data => {
    if (data.notice == "no cookie") {
        location.replace("http://localhost:8000/templates/")
    } else {
        console.log(data.notice)
    }
})