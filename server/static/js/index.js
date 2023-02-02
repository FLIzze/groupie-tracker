function enlarge(boxdiv, imgdiv, divtext) {
    const xDiv = document.getElementById(boxdiv);
    const yDiv = document.getElementById(imgdiv);
    const wDiv = document.getElementById(divtext);

    if (xDiv.style.width < '400px')
    {
        wDiv.classList.remove("notHidden")
        wDiv.classList.add("hidden")
        yDiv.style.borderTopRightRadius = "0px"
        yDiv.style.borderBottomRightRadius = "0px"
        yDiv.classList.add("noHover")
        if (window.innerWidth < "966")
        {
            xDiv.style.width = '500px'
        } 
        else {
            xDiv.style.width = '540px' 
        }
        setTimeout(function() {
            wDiv.style.marginLeft = "245px"
            wDiv.classList.add("notHidden")
            wDiv.classList.remove("hidden")
        }, 700)
        }
    else 
    {
        wDiv.classList.add("hidden")
        wDiv.classList.remove("notHidden")
        yDiv.classList.remove("noHover")
        yDiv.style.borderTopRightRadius = "10px"
        yDiv.style.borderBottomRightRadius = "10px"
        xDiv.style.width = '240px' 
        setTimeout(function() {
            wDiv.classList.add("notHidden")
            wDiv.classList.remove("hidden")
            wDiv.style.marginLeft = "5px"
        }, 700)
    }
}

function slidechange1() {
    const print = document.getElementById("result")
    const change = document.getElementById("change")
    print.innerText = change.value;

    const form = document.getElementById("slide1")
}

function slidechange2() {
    const print = document.getElementById("result2")
    const change = document.getElementById("change2")
    print.innerText = change.value;

    const form = document.getElementById("slide2")
}

function submit(form) {
    form.submit()
}

const searchInput = document.querySelector("[data-search]")

console.log(objApi)

let bandName = []
let membersName = []
let albumDate = []
let creationDate = []

searchInput.addEventListener("input", e => {
    const value = e.target.value.toLowerCase()
    console.log(value)
})

.then(data => {
    user = data.map(user => {
        return { bandName: user.bandName, membersName: user.membersName}
    })
})
