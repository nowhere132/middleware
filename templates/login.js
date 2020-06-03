async function Login(event) {
    event.preventDefault()
    nameBlock = document.getElementById("name")  
    passBlock = document.getElementById("pass")
    console.log(nameBlock.value)
    await fetch("http://localhost:8000/checklogin", {
        mode: "no-cors",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            name: nameBlock.value,
            pass: passBlock.value
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data)
        if (data.notice == "manhdeptrai") {
            location.replace("http://localhost:8000/templates/welcome.html")
        } else {
            document.getElementById("noti").innerHTML = "Wrong in4"
            nameBlock.value = ""
            passBlock.value = ""
        }
    })
}

