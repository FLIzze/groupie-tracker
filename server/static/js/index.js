function enlarge(boxdiv, imgdiv, divtext) {
    var xDiv = document.getElementById(boxdiv);
    var yDiv = document.getElementById(imgdiv);
    var wDiv = document.getElementById(divtext);

    if (xDiv.style.width < '400px')
    {
        wDiv.classList.remove("notHidden")
        wDiv.classList.add("hidden")
        xDiv.style.width = '550px' 
        yDiv.style.borderTopRightRadius = "0px"
        yDiv.style.borderBottomRightRadius = "0px"
        yDiv.classList.add("noHover")
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