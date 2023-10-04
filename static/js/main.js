let imgArray = [
    "static/img/rock_icon.png",
    "static/img/paper_icon.png",
    "static/img/scissors_icon.png",
]

function choose(x){
    fetch("/play?option="+x)
    .then(response => response.json())
    .then(data =>{
        if(x==0){
            document.getElementById("player_choice").innerHTML = "The Player chosen Rock"
        }else if(x==1){            
            document.getElementById("player_choice").innerHTML = "The Player chosen Paper"
        }else {            
            document.getElementById("player_choice").innerHTML = "The Player chosen Scissors"
        }
        document.getElementById("computer_choice").innerHTML = data.computer_choice_name
        document.getElementById("round_result").innerHTML = data.round_result
        document.getElementById("round_message").innerHTML = data.message

        document.getElementById("player_score").innerHTML = data.player_score
        document.getElementById("computer_score").innerHTML = data.computer_score

        var imgElement =  document.getElementById("img_computer")
        imgElement.src = imgArray[data.computer_choice]
        imgElement.width = 188
    })
}