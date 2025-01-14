
let msg = new SpeechSynthesisUtterance();//create a new utterance object to store the text to be spoken by the screen reader
let voices = speechSynthesis.getVoices();// get all the voices installed on the user's computer
msg.voice = voices[0]; // we want to use the first voice you can use another voice by changing the index
let tags = document.querySelectorAll('p,a,h1,h2,h3'); // add more tags for you project

// in here we loop through all the tags and add an event listener to each one when the user clicks on it we want to speak the text inside the tag
tags.forEach((tag) => {
    tag.addEventListener('click', (e) => {
        
        msg.text = e.target.innerText;
        tag.style.backgroundColor = "yellow";
        speechSynthesis.speak(msg);
        
        let interval = setInterval(() => {
            if(!speechSynthesis.speaking){
                tag.style.removeProperty('background-color');
                clearInterval(interval);
            }
        }, 100);
        
    });
});