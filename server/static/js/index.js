function enlarge(boxdiv, imgdiv, divtext) {
    const xDiv = document.getElementById(boxdiv);
    const yDiv = document.getElementById(imgdiv);
    const wDiv = document.getElementById(divtext);

    if (xDiv.style.width < '400px')
    {
        wDiv.classList.remove("notHidden")
        wDiv.style.marginLeft = "250px"
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
            wDiv.classList.add("notHidden")
            wDiv.classList.remove("hidden")
        }, 700)
        }
    else 
    {
        wDiv.classList.add("hidden")
        wDiv.classList.remove("notHidden")
        yDiv.classList.remove("noHover")
        wDiv.style.marginLeft = "5px"
        yDiv.style.borderTopRightRadius = "10px"
        yDiv.style.borderBottomRightRadius = "10px"
        xDiv.style.width = '240px' 
        setTimeout(function() {
            wDiv.classList.add("notHidden")
            wDiv.classList.remove("hidden")
        }, 700)
    }
}

function slidechange1() {
    const print = document.getElementById("result")
    const change = document.getElementById("change")
    print.innerText = change.value;

    const form = document.getElementById("slide1")

    console.log(change.value)

    for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i+1}`)
        if (albumDate[i].includes(change.value)) {
            li.style.display = "flex";
        } else {
            li.style.display = "none";
        }
    }
}

function slidechange2() {
    const print = document.getElementById("result2")
    const change = document.getElementById("change2")
    print.innerText = change.value;

    const form = document.getElementById("slide2")

    for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i+1}`)
        if (creationDate[i].includes(change.value)) {
            li.style.display = "flex";
        } else {
            li.style.display = "none";
        }
    }
}

const searchInput = document.querySelector("[data-search]")

let bandName = []
let membersName = []
let albumDate = []
let creationDate = []

objApi.forEach(artist => {
    bandName.push(artist.Name.toLowerCase()) //k
    membersName.push(artist.Members) //k
    albumDate.push(artist.FirstAlbum) //k
    creationDate.push(artist.CreationDate.toString())  //k
});

const artistsCards = [];
for (i = 1; i < 53; i++) {
    li = document.getElementById(`box${i}`)
    artistsCards.push(li)
}

searchInput.addEventListener("input", e => {  
    const value = e.target.value.toLowerCase()
        for (i = 0; i < 52; i++) {
        li = document.getElementById(`box${i+1}`)
        if (bandName[i].includes(value)) {
            li.style.display = "flex";
        } else if (albumDate[i].includes(value)) {
            li.style.display = "flex";
        } else if (creationDate[i].includes(value)) {
            li.style.display = "flex";
        } else if ((membersName[i].join(' ').toLowerCase()).includes(value)) {
            li.style.display = "flex";
        } else {
            li.style.display = "none";  
        }
    }
}) 

function validate() {
    displayNothing();
    for (i = 1; i < 8; i++) {
        if (document.getElementById('checkbox'+i).checked) {
            for (j = 0; j < 52; j++) {
                li = document.getElementById(`box${j+1}`)
                if (membersName[j].length == i) {
                        console.log(membersName[j].length);
                        console.log(membersName[j])
                        li.style.display = "flex";
                 }  
            }   
        }
    }
}

function displayNothing() {
    for (i = 1; i < 53; i++) { 
        document.getElementById('box'+i).style.display = "none";
    }
}